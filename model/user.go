package model

// User 表示 KOOK 用户的信息。
type User struct {
	// ID 是用户的唯一标识
	ID string `json:"id"`
	// Username 是用户名
	Username string `json:"username"`
	// IdentifyNum 是用户名后的识别数字（如 #1234）
	IdentifyNum string `json:"identify_num"`
	// Online 表示用户是否在线
	Online bool `json:"online"`
	// Status 是用户状态，0 表示正常，10 表示被封禁
	Status int `json:"status"`
	// Avatar 是用户头像 URL
	Avatar string `json:"avatar"`
	// Bot 表示用户是否为机器人
	Bot bool `json:"bot"`
	// MobileVerified 表示用户是否已验证手机号
	MobileVerified bool `json:"mobile_verified"`
	// System 表示是否为系统用户
	System bool `json:"system"`
	// MobilePrefix 是手机号前缀（国际区号）
	MobilePrefix string `json:"mobile_prefix"`
	// Mobile 是用户手机号（脱敏）
	Mobile string `json:"mobile"`
	// InvitedCount 是用户邀请的人数
	InvitedCount int `json:"invited_count"`
}

// UserChat 表示用户私信会话的信息。
type UserChat struct {
	// Code 是私信会话的标识
	Code string `json:"code"`
	// LastReadTime 是最后阅读时间戳（毫秒）
	LastReadTime int64 `json:"last_read_time"`
	// LatestMsgTime 是最新消息时间戳（毫秒）
	LatestMsgTime int64 `json:"latest_msg_time"`
	// UnreadCount 是未读消息数量
	UnreadCount int `json:"unread_count"`
	// TargetInfo 是私信对象的用户信息
	TargetInfo *User `json:"target_info"`
}
