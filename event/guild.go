package event

// GuildUpdatedEvent 表示服务器信息更新事件的数据。
type GuildUpdatedEvent struct {
	// GuildID 是被更新服务器的 ID
	GuildID string `json:"id"`
	// Name 是更新后的服务器名称
	Name string `json:"name"`
	// Icon 是更新后的服务器图标 URL
	Icon string `json:"icon"`
	// NotifyType 是通知类型
	NotifyType int `json:"notify_type"`
	// Region 是服务器所在区域
	Region string `json:"region"`
	// EnableOpen 表示是否开启公开服务器
	EnableOpen bool `json:"enable_open"`
	// OpenID 是公开服务器的 ID
	OpenID string `json:"open_id"`
	// DefaultChannelID 是默认频道的 ID
	DefaultChannelID string `json:"default_channel_id"`
	// WelcomeChannelID 是欢迎频道的 ID
	WelcomeChannelID string `json:"welcome_channel_id"`
}

// GuildDeletedEvent 表示服务器被删除事件的数据。
type GuildDeletedEvent struct {
	// GuildID 是被删除服务器的 ID
	GuildID string `json:"id"`
}

// BlacklistAddedEvent 表示用户被加入黑名单事件的数据。
type BlacklistAddedEvent struct {
	// OperatorID 是执行操作的用户 ID
	OperatorID string `json:"operator_id"`
	// Remark 是拉黑备注
	Remark string `json:"remark"`
	// UserID 是被拉黑的用户 ID
	UserID string `json:"user_id"`
}

// BlacklistDeletedEvent 表示用户被移出黑名单事件的数据。
type BlacklistDeletedEvent struct {
	// OperatorID 是执行操作的用户 ID
	OperatorID string `json:"operator_id"`
	// UserID 是被移出黑名单的用户 ID
	UserID string `json:"user_id"`
}
