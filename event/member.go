package event

// GuildMemberJoinedEvent 表示成员加入服务器事件的数据。
type GuildMemberJoinedEvent struct {
	// UserID 是加入成员的用户 ID
	UserID string `json:"user_id"`
	// JoinedAt 是加入时间戳（毫秒）
	JoinedAt int64 `json:"joined_at"`
}

// GuildMemberExitedEvent 表示成员退出服务器事件的数据。
type GuildMemberExitedEvent struct {
	// UserID 是退出成员的用户 ID
	UserID string `json:"user_id"`
	// ExitedAt 是退出时间戳（毫秒）
	ExitedAt int64 `json:"exited_at"`
}

// GuildMemberUpdatedEvent 表示服务器成员信息更新事件的数据。
type GuildMemberUpdatedEvent struct {
	// UserID 是更新成员的用户 ID
	UserID string `json:"user_id"`
	// Nickname 是更新后的昵称
	Nickname string `json:"nickname"`
}

// GuildMemberOnlineEvent 表示服务器成员上线事件的数据。
type GuildMemberOnlineEvent struct {
	// UserID 是上线成员的用户 ID
	UserID string `json:"user_id"`
	// EventTime 是事件发生的时间戳（毫秒）
	EventTime int64 `json:"event_time"`
	// Guilds 是成员所在的服务器列表
	Guilds []GuildMemberStatus `json:"guilds"`
}

// GuildMemberOfflineEvent 表示服务器成员离线事件的数据。
type GuildMemberOfflineEvent struct {
	// UserID 是离线成员的用户 ID
	UserID string `json:"user_id"`
	// EventTime 是事件发生的时间戳（毫秒）
	EventTime int64 `json:"event_time"`
	// Guilds 是成员所在的服务器列表
	Guilds []GuildMemberStatus `json:"guilds"`
}

// GuildMemberStatus 表示成员在某个服务器中的状态信息。
type GuildMemberStatus struct {
	// ID 是服务器的 ID
	ID string `json:"id"`
	// Name 是服务器的名称
	Name string `json:"name"`
}
