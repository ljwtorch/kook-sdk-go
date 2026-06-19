package event

// ReactionAddedEvent 表示消息回应（表情回复）新增事件的数据。
type ReactionAddedEvent struct {
	// MsgID 是被回应的消息 ID
	MsgID string `json:"msg_id"`
	// UserID 是添加回应的用户 ID
	UserID string `json:"user_id"`
	// ChannelID 是消息所属频道的 ID
	ChannelID string `json:"channel_id"`
	// Emoji 是回应的表情信息
	Emoji Emoji `json:"emoji"`
}

// ReactionDeletedEvent 表示消息回应（表情回复）删除事件的数据。
type ReactionDeletedEvent struct {
	// MsgID 是被取消回应的消息 ID
	MsgID string `json:"msg_id"`
	// UserID 是取消回应的用户 ID
	UserID string `json:"user_id"`
	// ChannelID 是消息所属频道的 ID
	ChannelID string `json:"channel_id"`
	// Emoji 是取消回应的表情信息
	Emoji Emoji `json:"emoji"`
}

// Emoji 表示一个表情信息。
//
// ID 是表情的数字标识（如 "128077" 代表 👍），
// Name 是表情的英文名称（如 "thumbsup"）。
//
// KOOK 支持的完整表情对照表请参见：
// https://img.kookapp.cn/assets/emoji.json
type Emoji struct {
	// ID 是表情的数字标识
	ID string `json:"id"`
	// Name 是表情的英文名称
	Name string `json:"name"`
}
