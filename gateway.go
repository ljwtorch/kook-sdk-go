package kook

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ljwtorch/kook-sdk-go/internal"

	"github.com/gorilla/websocket"
)

const (
	// pingInterval 是客户端心跳（PING）的发送间隔。
	// KOOK 官方建议为 30 秒，容差 ±5 秒。
	pingInterval = 30 * time.Second
	// pongTimeout 是等待 PONG 响应的超时时间。
	pongTimeout = 6 * time.Second
)

// EventHandler 是事件处理函数的类型。
// 当 Gateway 接收到匹配的事件时，会调用注册的 EventHandler。
// ctx 为事件上下文，evt 为解析后的事件数据。
type EventHandler func(ctx context.Context, evt *EventData)

// Gateway 管理 KOOK WebSocket 连接。
// 它负责建立连接、维护心跳、接收事件、处理重连等功能。
// Gateway 的读写操作通过互斥锁保护，是并发安全的。
type Gateway struct {
	// client 是关联的 SDK 客户端。
	client *Client
	// conn 是 WebSocket 连接实例。
	conn *websocket.Conn
	// sessionID 是当前 WebSocket 会话的标识，用于断线恢复。
	sessionID string
	// lastSN 是收到的最后一个 EVENT 的序列号，用于 RESUME。
	lastSN int64
	// handlers 存储按事件类型注册的事件处理器。
	// key 为事件类型字符串（如 "message"、"system"），value 为处理器列表。
	handlers map[string][]EventHandler
	// msgHandlers 存储消息类型（type 1~10）的处理器列表。
	msgHandlers []EventHandler
	// sysHandlers 存储系统消息（type 255）的处理器列表。
	sysHandlers []EventHandler
	// mu 保护 conn、sessionID、lastSN 等字段的并发读写。
	mu sync.RWMutex
	// writeMu 保护 WebSocket 写入操作，确保同一时间只有一个 goroutine 写入。
	writeMu sync.Mutex
	// done 用于通知所有后台 goroutine 停止运行。
	done chan struct{}
	// compress 控制是否启用 zlib 压缩。
	compress bool
}

// gatewayResponse 是 /gateway/index API 的响应数据结构。
type gatewayResponse struct {
	URL string `json:"url"`
}

// Connect 获取 Gateway URL 并建立 WebSocket 连接。
// 连接成功后会自动启动心跳和事件接收循环。
// 该方法会阻塞直到 ctx 被取消或调用 Disconnect()。
// 如果已处于连接状态，返回 ErrConnected。
func (c *Client) Connect(ctx context.Context) error {
	c.mu.Lock()
	if c.connected {
		c.mu.Unlock()
		return ErrConnected
	}
	c.connected = true
	c.mu.Unlock()

	defer func() {
		c.mu.Lock()
		c.connected = false
		c.mu.Unlock()
	}()

	// 获取 Gateway URL
	var gw gatewayResponse
	if err := c.Get(ctx, "/gateway/index", map[string]string{"compress": "0"}, &gw); err != nil {
		return fmt.Errorf("kook: failed to get gateway URL: %w", err)
	}

	if gw.URL == "" {
		return fmt.Errorf("kook: gateway URL is empty")
	}

	if err := c.gateway.connect(ctx, gw.URL); err != nil {
		return err
	}

	// 阻塞直到 ctx 被取消或 disconnect 被调用
	select {
	case <-ctx.Done():
		_ = c.gateway.disconnect()
		return ctx.Err()
	case <-c.gateway.done:
		return nil
	}
}

// ConnectWithCompress 获取 Gateway URL 并建立带压缩的 WebSocket 连接。
// 启用压缩后，服务端会对 WebSocket 消息进行 zlib 压缩。
func (c *Client) ConnectWithCompress(ctx context.Context) error {
	var gw gatewayResponse
	if err := c.Get(ctx, "/gateway/index", map[string]string{"compress": "1"}, &gw); err != nil {
		return fmt.Errorf("kook: failed to get gateway URL: %w", err)
	}

	if gw.URL == "" {
		return fmt.Errorf("kook: gateway URL is empty")
	}

	c.gateway.compress = true
	return c.gateway.connect(ctx, gw.URL)
}

// Disconnect 断开 WebSocket 连接并停止所有后台 goroutine。
func (c *Client) Disconnect() error {
	return c.gateway.disconnect()
}

// OnEvent 注册一个事件处理器。
// eventType 为事件类型字符串，如 "message"、"system" 等。
// 当接收到匹配的事件时，handler 会被调用。
//
// 使用示例:
//
//	client.OnEvent("message", func(ctx context.Context, evt *kook.EventData) {
//	    fmt.Println("Received:", evt.Content)
//	})
func (c *Client) OnEvent(eventType string, handler EventHandler) {
	c.gateway.mu.Lock()
	defer c.gateway.mu.Unlock()
	c.gateway.handlers[eventType] = append(c.gateway.handlers[eventType], handler)
}

// OnMessage 注册消息处理器。
// 消息类型包括文字(1)、图片(2)、视频(3)、文件(4)、音频(8)、KMarkdown(9)、卡片(10)。
// 此处理器仅处理 type >= 1 且 type <= 10 的事件。
//
// 使用示例:
//
//	client.OnMessage(func(ctx context.Context, evt *kook.EventData) {
//	    fmt.Printf("Message from %s: %s\n", evt.AuthorID, evt.Content)
//	})
func (c *Client) OnMessage(handler EventHandler) {
	c.gateway.mu.Lock()
	defer c.gateway.mu.Unlock()
	c.gateway.msgHandlers = append(c.gateway.msgHandlers, handler)
}

// OnSystem 注册系统消息处理器。
// 系统消息的 type 值为 255，包括频道变更、成员加入/退出、角色变更等事件。
//
// 使用示例:
//
//	client.OnSystem(func(ctx context.Context, evt *kook.EventData) {
//	    fmt.Printf("System event type=%d target=%s\n", evt.Type, evt.TargetID)
//	})
func (c *Client) OnSystem(handler EventHandler) {
	c.gateway.mu.Lock()
	defer c.gateway.mu.Unlock()
	c.gateway.sysHandlers = append(c.gateway.sysHandlers, handler)
}

// connect 建立 WebSocket 连接并启动后台 goroutine。
func (g *Gateway) connect(ctx context.Context, wsURL string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	// 关闭之前的连接
	if g.conn != nil {
		_ = g.conn.Close()
	}

	// 重置 done channel
	g.done = make(chan struct{})

	// 建立 WebSocket 连接
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}
	header := http.Header{}
	header.Set("Authorization", "Bot "+g.client.token)

	conn, _, err := dialer.DialContext(ctx, wsURL, header)
	if err != nil {
		return fmt.Errorf("kook: failed to connect to gateway: %w", err)
	}
	g.conn = conn

	if g.client.debug {
		log.Printf("[KOOK DEBUG] WebSocket connected to %s", wsURL)
	}

	// 等待 HELLO 信令
	if err := g.waitForHello(); err != nil {
		_ = conn.Close()
		g.conn = nil
		return err
	}

	// 启动心跳和读取循环
	go g.pingLoop(ctx)
	go g.readLoop(ctx)

	return nil
}

// disconnect 断开 WebSocket 连接并通知所有后台 goroutine 停止。
func (g *Gateway) disconnect() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	// 通知所有 goroutine 停止
	select {
	case <-g.done:
		// 已经关闭
	default:
		close(g.done)
	}

	if g.conn != nil {
		err := g.conn.Close()
		g.conn = nil
		return err
	}

	return nil
}

// waitForHello 等待服务端发送的 HELLO (s=1) 信令。
// HELLO 信令包含 session_id，需要保存以便后续 RESUME 使用。
func (g *Gateway) waitForHello() error {
	// 设置读取超时
	_ = g.conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	_, msg, err := g.conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("kook: failed to read HELLO: %w", err)
	}

	// 如果启用压缩，先解压
	if g.compress {
		msg, err = internal.Decompress(msg)
		if err != nil {
			return fmt.Errorf("kook: failed to decompress HELLO: %w", err)
		}
	}

	var signal Signal
	if err := json.Unmarshal(msg, &signal); err != nil {
		return fmt.Errorf("kook: failed to parse HELLO signal: %w", err)
	}

	if signal.S != SignalHello {
		return fmt.Errorf("kook: expected HELLO (s=1), got s=%d", signal.S)
	}

	var hello HelloData
	if err := json.Unmarshal(signal.D, &hello); err != nil {
		return fmt.Errorf("kook: failed to parse HELLO data: %w", err)
	}

	if hello.Code != 0 {
		return &APIError{Code: hello.Code, Message: "gateway hello failed"}
	}

	g.sessionID = hello.SessionID

	if g.client.debug {
		log.Printf("[KOOK DEBUG] HELLO received, session_id=%s", g.sessionID)
	}

	// 清除读取超时
	_ = g.conn.SetReadDeadline(time.Time{})

	return nil
}

// pingLoop 定期发送 PING (s=2) 心跳信令。
// 每 30 秒发送一次，容差 ±5 秒。
func (g *Gateway) pingLoop(ctx context.Context) {
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-g.done:
			return
		case <-ticker.C:
			if err := g.sendPing(); err != nil {
				if g.client.debug {
					log.Printf("[KOOK DEBUG] Failed to send PING: %v", err)
				}
				// 尝试重连
				g.reconnect(ctx)
				return
			}
		}
	}
}

// sendPing 发送 PING (s=2) 心跳信令。
func (g *Gateway) sendPing() error {
	signal := Signal{S: SignalPing}
	data, err := json.Marshal(signal)
	if err != nil {
		return err
	}

	g.writeMu.Lock()
	defer g.writeMu.Unlock()

	g.mu.RLock()
	conn := g.conn
	g.mu.RUnlock()

	if conn == nil {
		return fmt.Errorf("kook: connection is nil")
	}

	if g.client.debug {
		log.Printf("[KOOK DEBUG] Sending PING (sn=%d)", g.lastSN)
	}

	return conn.WriteMessage(websocket.TextMessage, data)
}

// readLoop 持续读取 WebSocket 消息并分发处理。
func (g *Gateway) readLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-g.done:
			return
		default:
		}

		_, msg, err := g.conn.ReadMessage()
		if err != nil {
			select {
			case <-g.done:
				return
			case <-ctx.Done():
				return
			default:
			}

			if g.client.debug {
				log.Printf("[KOOK DEBUG] WebSocket read error: %v", err)
			}

			// 尝试重连
			g.reconnect(ctx)
			return
		}

		// 如果启用压缩，先解压
		if g.compress {
			msg, err = internal.Decompress(msg)
			if err != nil {
				if g.client.debug {
					log.Printf("[KOOK DEBUG] Failed to decompress message: %v", err)
				}
				continue
			}
		}

		var signal Signal
		if err := json.Unmarshal(msg, &signal); err != nil {
			if g.client.debug {
				log.Printf("[KOOK DEBUG] Failed to parse signal: %v", err)
			}
			continue
		}

		g.handleSignal(ctx, &signal)
	}
}

// handleSignal 根据信令类型分发处理。
func (g *Gateway) handleSignal(ctx context.Context, signal *Signal) {
	switch signal.S {
	case SignalEvent:
		g.handleEvent(ctx, signal)
	case SignalPong:
		if g.client.debug {
			log.Printf("[KOOK DEBUG] Received PONG")
		}
	case SignalReconnect:
		if g.client.debug {
			log.Printf("[KOOK DEBUG] Received RECONNECT, reconnecting...")
		}
		g.reconnect(ctx)
	case SignalResumeAck:
		if g.client.debug {
			log.Printf("[KOOK DEBUG] Received RESUME ACK")
		}
	default:
		if g.client.debug {
			log.Printf("[KOOK DEBUG] Received unknown signal s=%d", signal.S)
		}
	}
}

// handleEvent 处理 EVENT (s=0) 信令，解析事件数据并分发给注册的处理器。
func (g *Gateway) handleEvent(ctx context.Context, signal *Signal) {
	// 更新 lastSN
	if signal.SN > 0 {
		g.mu.Lock()
		if signal.SN > g.lastSN {
			g.lastSN = signal.SN
		}
		g.mu.Unlock()
	}

	var evt EventData
	if err := json.Unmarshal(signal.D, &evt); err != nil {
		if g.client.debug {
			log.Printf("[KOOK DEBUG] Failed to parse event data: %v", err)
		}
		return
	}

	if g.client.debug {
		log.Printf("[KOOK DEBUG] Event: type=%d channel=%s sn=%d", evt.Type, evt.ChannelType, signal.SN)
	}

	// 分发消息处理器 (type 1~10)
	if evt.Type >= 1 && evt.Type <= 10 {
		g.mu.RLock()
		msgHandlers := make([]EventHandler, len(g.msgHandlers))
		copy(msgHandlers, g.msgHandlers)
		g.mu.RUnlock()

		for _, handler := range msgHandlers {
			go handler(ctx, &evt)
		}
	}

	// 分发系统消息处理器 (type 255)
	if evt.Type == 255 {
		g.mu.RLock()
		sysHandlers := make([]EventHandler, len(g.sysHandlers))
		copy(sysHandlers, g.sysHandlers)
		g.mu.RUnlock()

		for _, handler := range sysHandlers {
			go handler(ctx, &evt)
		}
	}

	// 分发给按类型注册的处理器
	g.mu.RLock()
	eventHandlers := make([]EventHandler, len(g.handlers["*"]))
	copy(eventHandlers, g.handlers["*"])
	g.mu.RUnlock()

	for _, handler := range eventHandlers {
		go handler(ctx, &evt)
	}
}

// reconnect 使用指数退避策略进行重连。
// 重连时会尝试使用 RESUME 信令恢复之前的会话。
func (g *Gateway) reconnect(ctx context.Context) {
	backoff := internal.NewBackoff()

	// 先断开当前连接
	g.mu.Lock()
	if g.conn != nil {
		_ = g.conn.Close()
		g.conn = nil
	}
	g.mu.Unlock()

	for {
		select {
		case <-ctx.Done():
			return
		case <-g.done:
			return
		default:
		}

		wait := backoff.Next()
		if g.client.debug {
			log.Printf("[KOOK DEBUG] Reconnecting in %v...", wait)
		}

		timer := time.NewTimer(wait)
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-g.done:
			timer.Stop()
			return
		case <-timer.C:
		}

		// 重新获取 Gateway URL
		var gw gatewayResponse
		if err := g.client.Get(ctx, "/gateway/index", map[string]string{
			"compress": compressFlag(g.compress),
		}, &gw); err != nil {
			if g.client.debug {
				log.Printf("[KOOK DEBUG] Failed to get gateway URL: %v", err)
			}
			continue
		}

		if gw.URL == "" {
			continue
		}

		// 建立新连接
		dialer := websocket.Dialer{
			HandshakeTimeout: 10 * time.Second,
		}
		header := http.Header{}
		header.Set("Authorization", "Bot "+g.client.token)

		conn, _, err := dialer.DialContext(ctx, gw.URL, header)
		if err != nil {
			if g.client.debug {
				log.Printf("[KOOK DEBUG] Failed to reconnect: %v", err)
			}
			continue
		}

		g.mu.Lock()
		g.conn = conn
		g.mu.Unlock()

		// 等待 HELLO
		if err := g.waitForHello(); err != nil {
			if g.client.debug {
				log.Printf("[KOOK DEBUG] HELLO failed after reconnect: %v", err)
			}
			_ = conn.Close()
			g.mu.Lock()
			g.conn = nil
			g.mu.Unlock()
			continue
		}

		// 尝试 RESUME
		if g.sessionID != "" && g.lastSN > 0 {
			if err := g.sendResume(); err != nil {
				if g.client.debug {
					log.Printf("[KOOK DEBUG] RESUME failed: %v", err)
				}
				// RESUME 失败不阻断，继续正常工作（事件可能丢失）
			}
		}

		if g.client.debug {
			log.Printf("[KOOK DEBUG] Reconnected successfully")
		}

		// 重新启动心跳和读取循环
		go g.pingLoop(ctx)
		go g.readLoop(ctx)
		return
	}
}

// sendResume 发送 RESUME (s=4) 信令以恢复之前的会话。
// RESUME 需要携带 session_id 和最后收到的 SN。
func (g *Gateway) sendResume() error {
	g.mu.RLock()
	resumeData := ResumeData{
		SN:        g.lastSN,
		SessionID: g.sessionID,
	}
	g.mu.RUnlock()

	data, err := json.Marshal(resumeData)
	if err != nil {
		return err
	}

	signal := Signal{
		S: SignalResume,
		D: data,
	}
	signalBytes, err := json.Marshal(signal)
	if err != nil {
		return err
	}

	g.writeMu.Lock()
	defer g.writeMu.Unlock()

	g.mu.RLock()
	conn := g.conn
	g.mu.RUnlock()

	if conn == nil {
		return fmt.Errorf("kook: connection is nil")
	}

	if g.client.debug {
		log.Printf("[KOOK DEBUG] Sending RESUME (sn=%d, session_id=%s)", resumeData.SN, resumeData.SessionID)
	}

	return conn.WriteMessage(websocket.TextMessage, signalBytes)
}

// compressFlag 根据 compress 设置返回查询参数值。
func compressFlag(compress bool) string {
	if compress {
		return "1"
	}
	return "0"
}
