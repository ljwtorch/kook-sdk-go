package kook

import "encoding/json"

// WebSocket 信令类型常量。
// 这些常量定义了 KOOK WebSocket Gateway 的所有信令类型。
const (
	// SignalEvent 表示服务端推送的事件信令 (EVENT)。
	SignalEvent = 0
	// SignalHello 表示连接成功的欢迎信令 (HELLO)。
	SignalHello = 1
	// SignalPing 表示客户端发送的心跳信令 (PING)。
	SignalPing = 2
	// SignalPong 表示服务端回复的心跳响应信令 (PONG)。
	SignalPong = 3
	// SignalResume 表示客户端请求恢复连接的信令 (RESUME)。
	SignalResume = 4
	// SignalReconnect 表示服务端要求客户端重连的信令 (RECONNECT)。
	SignalReconnect = 5
	// SignalResumeAck 表示恢复连接成功的确认信令 (RESUME ACK)。
	SignalResumeAck = 6
)

// Signal 是 KOOK WebSocket Gateway 的信令结构。
// 每个 WebSocket 消息都遵循此格式。
type Signal struct {
	// S 是信令类型，取值范围参见 Signal* 常量。
	S int `json:"s"`
	// D 是信令携带的数据，具体内容取决于信令类型。
	D json.RawMessage `json:"d,omitempty"`
	// SN 是事件序列号，仅 EVENT (s=0) 信令包含此字段。
	// 客户端需要记录最新的 SN 以便在重连时恢复会话。
	SN int64 `json:"sn,omitempty"`
}

// HelloData 是 HELLO (s=1) 信令的数据结构。
// 连接成功后服务端发送此信令，客户端据此判断连接状态。
type HelloData struct {
	// Code 为 0 表示连接成功，非 0 表示连接失败。
	Code int `json:"code"`
	// SessionID 是当前 WebSocket 会话的唯一标识，用于断线恢复。
	SessionID string `json:"session_id"`
}

// ResumeData 是 RESUME (s=4) 信令的数据结构。
// 客户端在重连时发送此信令以恢复之前的会话。
type ResumeData struct {
	// SN 是客户端收到的最后一个 EVENT 的序列号。
	SN int64 `json:"sn"`
	// SessionID 是之前 HELLO 信令返回的会话标识。
	SessionID string `json:"session_id"`
}

// EventData 是 EVENT (s=0) 信令的数据结构。
// 包含了所有服务端推送事件的基础信息。
type EventData struct {
	// ChannelType 表示消息来源的频道类型。
	// "GROUP" 表示群聊频道，"PERSON" 表示私聊。
	ChannelType string `json:"channel_type"`
	// Type 表示消息类型：
	// 1=文字, 2=图片, 3=视频, 4=文件, 8=音频, 9=KMarkdown, 10=卡片, 255=系统消息。
	Type int `json:"type"`
	// TargetID 是消息目标，频道消息中为频道 ID，私聊消息中为用户 ID。
	TargetID string `json:"target_id"`
	// AuthorID 是消息发送者的用户 ID。
	AuthorID string `json:"author_id"`
	// Content 是消息内容（文字消息时为文本，其他类型时可能为 URL 或 JSON）。
	Content string `json:"content"`
	// MsgID 是消息的唯一标识。
	MsgID string `json:"msg_id"`
	// MsgTimestamp 是消息发送时间的 Unix 毫秒时间戳。
	MsgTimestamp int64 `json:"msg_timestamp"`
	// Nonce 是消息的随机字符串，用于去重。
	Nonce string `json:"nonce"`
	// Extra 包含消息的额外数据，根据 Type 不同内容不同。
	// 文字消息包含发送者信息，系统消息包含事件详细数据等。
	Extra json.RawMessage `json:"extra"`
}
