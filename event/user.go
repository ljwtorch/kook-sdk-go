package event

// UserUpdatedEvent 表示用户信息更新事件的数据。
type UserUpdatedEvent struct {
	// UserID 是被更新用户的 ID
	UserID string `json:"user_id"`
	// Username 是更新后的用户名
	Username string `json:"username"`
	// Avatar 是更新后的头像 URL
	Avatar string `json:"avatar"`
}

// SelfJoinedGuildEvent 表示当前机器人加入服务器事件的数据。
type SelfJoinedGuildEvent struct {
	// GuildID 是机器人加入的服务器 ID
	GuildID string `json:"guild_id"`
}

// SelfExitedGuildEvent 表示当前机器人退出服务器事件的数据。
type SelfExitedGuildEvent struct {
	// GuildID 是机器人退出的服务器 ID
	GuildID string `json:"guild_id"`
}
