// Package kook 提供了 KOOK（原开黑啦）开放平台的 Go SDK。
//
// SDK 提供完整的 HTTP API 封装、WebSocket 事件监听和卡片消息构建能力。
//
// # 基本使用
//
// 创建客户端后，可以通过 api 包调用 HTTP API：
//
//	client := kook.NewClient("your-bot-token")
//	guilds, err := api.GetGuildList(ctx, client)
//
// # WebSocket 使用
//
// 注册事件处理器后连接 WebSocket：
//
//	client := kook.NewClient("your-bot-token")
//	client.OnMessage(func(ctx context.Context, evt *kook.EventData) {
//	    // 处理消息
//	})
//	client.Connect(ctx)
//
// # 便捷方法
//
// Client 提供了一系列便捷方法，可直接调用而无需导入 api 包：
//
//	me, err := client.Me(ctx)
//	msg, err := client.SendMessage(ctx, "channel-id", "Hello!")
package kook

import "errors"

// Version 是 SDK 版本号。
const Version = "0.1.0-dev.1"

// DefaultBaseURL 是 KOOK API 默认基础地址。
const DefaultBaseURL = "https://www.kookapp.cn"

// DefaultAPIVersion 是 KOOK API 默认版本。
const DefaultAPIVersion = "v3"

// ErrConnected 表示 WebSocket 已连接，不可重复连接。
var ErrConnected = errors.New("kook: already connected")

// MessageHandler 是消息事件处理函数类型的别名，等同于 EventHandler。
// 保留此类型以向后兼容。
type MessageHandler = EventHandler

// SystemHandler 是系统事件处理函数类型的别名，等同于 EventHandler。
// 保留此类型以向后兼容。
type SystemHandler = EventHandler

// ClientOption 是客户端配置选项类型的别名，等同于 Option。
// 保留此类型以向后兼容。
type ClientOption = Option

// HTTPRequester 是 HTTP 请求接口，由 Client 实现。
// api 包中的函数可通过此接口发起 HTTP 请求，便于测试时 Mock。
type HTTPRequester interface {
	// Token 返回 Bot Token。
	Token() string
	// BaseURL 返回 API 基础 URL。
	BaseURL() string
}
