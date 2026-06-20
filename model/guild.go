package model

// Guild 表示 KOOK 服务器（频道）的信息。
type Guild struct {
	// ID 是服务器的唯一标识
	ID string `json:"id"`
	// Name 是服务器名称
	Name string `json:"name"`
	// Topic 是服务器主题
	Topic string `json:"topic"`
	// MasterID 是服务器主的用户 ID
	MasterID string `json:"master_id"`
	// UserID 是服务器主的用户 ID（同 MasterID）
	UserID string `json:"user_id"`
	// IsMaster 表示当前用户是否为服务器主
	IsMaster bool `json:"is_master"`
	// Icon 是服务器图标 URL
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
	// BoostNum 是服务器助力数量
	BoostNum int `json:"boost_num"`
	// Level 是服务器等级
	Level int `json:"level"`
	// AutoDeleteTime 是消息自动删除天数，0 表示不自动删除
	AutoDeleteTime FlexInt `json:"auto_delete_time"`
}

// GuildDetail 表示服务器的详细信息，包含角色和频道列表。
// 由 GetGuild（guild/view）接口返回。
type GuildDetail struct {
	Guild
	// Roles 是服务器的角色列表
	Roles []Role `json:"roles"`
	// Channels 是服务器的频道列表
	Channels []Channel `json:"channels"`
}

// GuildUser 表示服务器中的用户信息。
type GuildUser struct {
	// ID 是用户的唯一标识
	ID string `json:"id"`
	// Username 是用户名
	Username string `json:"username"`
	// Nickname 是用户在服务器中的昵称
	Nickname string `json:"nickname"`
	// Avatar 是用户头像 URL
	Avatar string `json:"avatar"`
	// Online 表示用户是否在线
	Online bool `json:"online"`
	// Bot 表示用户是否为机器人
	Bot bool `json:"bot"`
	// Roles 是用户在服务器中的角色 ID 列表
	Roles []int `json:"roles"`
	// JoinedAt 是用户加入服务器的时间戳（毫秒）
	JoinedAt int64 `json:"joined_at"`
	// ActiveTime 是用户最后活跃的时间戳（毫秒）
	ActiveTime int64 `json:"active_time"`
	// IsOwner 表示用户是否为服务器所有者
	IsOwner bool `json:"is_owner"`
	// MobileVerified 表示用户是否已验证手机号
	MobileVerified bool `json:"mobile_verified"`
	// Color 是用户角色颜色，使用 24 位 RGB 编码的十进制整数（参见 model.Role.Color）。
	Color int `json:"color"`
}
