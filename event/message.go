package event

// MessageUpdatedEvent 表示消息更新事件的数据。
type MessageUpdatedEvent struct {
	// MsgID 是被更新消息的 ID
	MsgID string `json:"msg_id"`
	// Content 是更新后的消息内容
	Content string `json:"content"`
	// ChannelID 是消息所属频道的 ID
	ChannelID string `json:"channel_id"`
	// Mention 是 @ 的用户 ID 列表
	Mention []string `json:"mention"`
	// MentionAll 表示是否 @ 所有人
	MentionAll bool `json:"mention_all"`
	// MentionRoles 是 @ 的角色 ID 列表
	MentionRoles []int `json:"mention_roles"`
	// MentionHere 表示是否 @ 在线成员
	MentionHere bool `json:"mention_here"`
	// UpdatedAt 是消息更新的时间戳（毫秒）
	UpdatedAt int64 `json:"updated_at"`
}

// MessageDeletedEvent 表示消息删除事件的数据。
type MessageDeletedEvent struct {
	// MsgID 是被删除消息的 ID
	MsgID string `json:"msg_id"`
	// ChannelID 是消息所属频道的 ID
	ChannelID string `json:"channel_id"`
	// DeletedAt 是消息删除的时间戳（毫秒）
	DeletedAt int64 `json:"deleted_at"`
}

// MessagePinnedEvent 表示消息置顶事件的数据。
type MessagePinnedEvent struct {
	// ChannelID 是置顶消息所属频道的 ID
	ChannelID string `json:"channel_id"`
	// OperatorID 是执行置顶操作的用户 ID
	OperatorID string `json:"operator_id"`
	// MsgID 是被置顶的消息 ID
	MsgID string `json:"msg_id"`
}

// MessageUnpinnedEvent 表示消息取消置顶事件的数据。
type MessageUnpinnedEvent struct {
	// ChannelID 是取消置顶消息所属频道的 ID
	ChannelID string `json:"channel_id"`
	// OperatorID 是执行取消置顶操作的用户 ID
	OperatorID string `json:"operator_id"`
	// MsgID 是被取消置顶的消息 ID
	MsgID string `json:"msg_id"`
}
