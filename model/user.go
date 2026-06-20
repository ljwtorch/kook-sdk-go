package model

// User 表示 KOOK 用户的信息。
//
// 注意：部分字段仅在特定接口返回：
//   - /user/me 独有: BotStatus, TagInfo, ClientID, Verified, Intent
//   - /user/view 独有: Nickname(服务器昵称), JoinedAt, ActiveTime, KpmVip, WealthLevel
type User struct {
	// ID 是用户的唯一标识
	ID string `json:"id"`
	// Username 是用户名
	Username string `json:"username"`
	// IdentifyNum 是用户名后的识别数字（如 #1234）
	IdentifyNum string `json:"identify_num"`
	// Online 表示用户是否在线
	Online bool `json:"online"`
	// OS 是当前连接方式（如 "Websocket"）
	OS string `json:"os"`
	// Status 是用户状态，0 和 1 代表正常，10 代表被封禁
	Status int `json:"status"`
	// Avatar 是用户头像 URL
	Avatar string `json:"avatar"`
	// VipAvatar 是 VIP 用户头像 URL
	VipAvatar string `json:"vip_avatar"`
	// Banner 是用户横幅 URL
	Banner string `json:"banner"`
	// Nickname 是用户昵称（/user/view 返回的是服务器昵称）
	Nickname string `json:"nickname"`
	// Roles 是用户角色列表（/user/view 返回角色 ID 列表，/user/me 返回角色信息）
	Roles []int `json:"roles"`
	// IsVip 表示用户是否为 VIP
	IsVip bool `json:"is_vip"`
	// VipAmp 表示用户是否为年度 VIP
	VipAmp bool `json:"vip_amp"`
	// Bot 表示用户是否为机器人
	Bot bool `json:"bot"`
	// BotStatus 是机器人状态（仅 /user/me 返回）
	BotStatus bool `json:"bot_status"`
	// TagInfo 是用户标签信息（仅 /user/me 返回）
	TagInfo *TagInfo `json:"tag_info"`
	// MobileVerified 表示用户是否已验证手机号
	MobileVerified bool `json:"mobile_verified"`
	// IsSys 表示是否为系统账号
	IsSys bool `json:"is_sys"`
	// ClientID 是客户端 ID（仅 /user/me 返回）
	ClientID string `json:"client_id"`
	// Verified 表示是否已验证（仅 /user/me 返回）
	Verified bool `json:"verified"`
	// MobilePrefix 是手机号前缀（国际区号）
	MobilePrefix string `json:"mobile_prefix"`
	// Mobile 是用户手机号（脱敏）
	Mobile string `json:"mobile"`
	// InvitedCount 是用户邀请的人数
	InvitedCount int `json:"invited_count"`
	// Intent 是事件订阅掩码，默认为 -1 接收所有（仅 /user/me 返回）
	Intent int `json:"intent"`
	// JoinedAt 是加入服务器时间（仅 /user/view 返回，需传 guild_id）
	JoinedAt int64 `json:"joined_at"`
	// ActiveTime 是活跃时间（仅 /user/view 返回，需传 guild_id）
	ActiveTime int64 `json:"active_time"`
	// KpmVip 是 KPM 服务器 VIP 信息（仅 /user/view 返回）
	KpmVip *KpmVip `json:"kpm_vip"`
	// WealthLevel 是语音房财富等级（仅 /user/view 返回）
	WealthLevel int `json:"wealth_level"`
}

// TagInfo 表示用户标签信息。
type TagInfo struct {
	// Color 是标签颜色
	Color string `json:"color"`
	// BgColor 是标签背景颜色
	BgColor string `json:"bg_color"`
	// Text 是标签文本
	Text string `json:"text"`
}

// KpmVip 表示 KPM 服务器的 VIP 信息。
type KpmVip struct {
	// Exp 是经验值
	Exp int `json:"exp"`
	// Discount 是折扣
	Discount int `json:"discount"`
	// Icon 是图标 URL
	Icon string `json:"icon"`
	// Text 是描述文本
	Text string `json:"text"`
	// Level 是 VIP 等级
	Level int `json:"level"`
}

// BotOnlineStatus 表示机器人在线状态。
type BotOnlineStatus struct {
	// Online 表示是否在线
	Online bool `json:"online"`
	// OnlineOS 是在线的平台列表
	OnlineOS []string `json:"online_os"`
}

// UserChat 表示用户私信会话的信息。
//
// 注意：ListUserChats 接口不返回 IsFriend、IsBlocked、IsTargetBlocked 字段，
// 这些字段仅在 GetUserChat 和 CreateUserChat 接口中返回。
type UserChat struct {
	// Code 是私信会话的标识
	Code string `json:"code"`
	// LastReadTime 是最后阅读时间戳（毫秒）
	LastReadTime int64 `json:"last_read_time"`
	// LatestMsgTime 是最新消息时间戳（毫秒）
	LatestMsgTime int64 `json:"latest_msg_time"`
	// UnreadCount 是未读消息数量
	UnreadCount int `json:"unread_count"`
	// IsFriend 表示是否是好友（仅 GetUserChat/CreateUserChat 返回）
	IsFriend bool `json:"is_friend"`
	// IsBlocked 表示是否已屏蔽对方（仅 GetUserChat/CreateUserChat 返回）
	IsBlocked bool `json:"is_blocked"`
	// IsTargetBlocked 表示是否已被对方屏蔽（仅 GetUserChat/CreateUserChat 返回）
	IsTargetBlocked bool `json:"is_target_blocked"`
	// TargetInfo 是私信对象的用户信息
	TargetInfo *User `json:"target_info"`
}
