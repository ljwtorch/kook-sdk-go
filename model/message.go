package model

// 消息类型常量定义了 KOOK 支持的消息类型。
const (
	// MessageTypeText 表示文字消息
	MessageTypeText = 1
	// MessageTypeImage 表示图片消息
	MessageTypeImage = 2
	// MessageTypeVideo 表示视频消息
	MessageTypeVideo = 3
	// MessageTypeFile 表示文件消息
	MessageTypeFile = 4
	// MessageTypeVoice 表示音频消息
	MessageTypeVoice = 8
	// MessageTypeKMarkdown 表示 KMarkdown 格式消息
	MessageTypeKMarkdown = 9
	// MessageTypeCard 表示卡片消息
	MessageTypeCard = 10
	// MessageTypeSystem 表示系统消息
	MessageTypeSystem = 255
)

// Message 表示 KOOK 中的消息。
type Message struct {
	// ID 是消息的唯一标识
	ID string `json:"id"`
	// Type 是消息类型，参见 MessageType 常量
	Type int `json:"type"`
	// Content 是消息内容
	Content string `json:"content"`
	// Mention 是 @ 的用户 ID 列表
	Mention []string `json:"mention"`
	// MentionAll 表示是否 @ 所有人
	MentionAll bool `json:"mention_all"`
	// MentionRoles 是 @ 的角色 ID 列表
	MentionRoles []int `json:"mention_roles"`
	// MentionHere 表示是否 @ 在线成员
	MentionHere bool `json:"mention_here"`
	// Embeds 是消息中的嵌入内容列表
	Embeds []Embed `json:"embeds"`
	// Attachments 是消息的附件列表
	Attachments []Attachment `json:"attachments"`
	// CreateAt 是消息创建时间戳（毫秒）
	CreateAt int64 `json:"create_at"`
	// UpdateAt 是消息更新时间戳（毫秒）
	UpdateAt int64 `json:"update_at"`
	// EditedAt 是消息编辑时间戳（毫秒）
	EditedAt int64 `json:"edited_at"`
	// ChannelID 是消息所属频道的 ID
	ChannelID string `json:"channel_id"`
	// ParentID 是引用消息的 ID
	ParentID string `json:"parent_id"`
	// Quote 是引用的消息信息
	Quote *Quote `json:"quote"`
	// Author 是消息作者的用户信息
	Author *User `json:"author"`
	// Reactions 是消息的回应列表
	Reactions []Reaction `json:"reactions"`
	// MentionInfo 是 @特定用户或特定角色的详细信息
	MentionInfo *MentionInfo `json:"mention_info"`
}

// MentionInfo 表示消息中 @特定用户或 @特定角色 的详细信息。
type MentionInfo struct {
	// MentionPart 是 @特定用户 的详情列表
	MentionPart []MentionUser `json:"mention_part"`
	// MentionRolePart 是 @特定角色 的详情列表
	MentionRolePart []MentionRolePart `json:"mention_role_part"`
}

// MentionUser 表示被 @的用户的详细信息。
type MentionUser struct {
	// ID 是用户的唯一标识
	ID string `json:"id"`
	// Username 是用户名
	Username string `json:"username"`
	// FullName 是用户的完整名称（用户名#识别号）
	FullName string `json:"full_name"`
	// Avatar 是用户头像 URL
	Avatar string `json:"avatar"`
}

// MentionRolePart 表示被 @的角色的详细信息。
type MentionRolePart struct {
	// RoleID 是角色 ID
	RoleID int `json:"role_id"`
	// Name 是角色名称
	Name string `json:"name"`
	// Color 是角色颜色
	Color int `json:"color"`
	// Position 是角色位置
	Position int `json:"position"`
	// Hoist 是否在成员列表中单独展示
	Hoist int `json:"hoist"`
	// Mentionable 是否可被 @
	Mentionable int `json:"mentionable"`
	// Permissions 是角色权限
	Permissions int `json:"permissions"`
}

// Embed 表示消息中的嵌入内容。
type Embed struct {
	// Type 是嵌入类型，如 video、file、image
	Type string `json:"type"`
	// URL 是嵌入内容的 URL
	URL string `json:"url"`
	// OriginURL 是原始 URL
	OriginURL string `json:"origin_url"`
	// Size 是文件大小（字节）
	Size int `json:"size"`
	// Width 是宽度（像素）
	Width int `json:"width"`
	// Height 是高度（像素）
	Height int `json:"height"`
}

// Attachment 表示消息的附件信息。
type Attachment struct {
	// Type 是附件类型
	Type string `json:"type"`
	// Name 是附件文件名
	Name string `json:"name"`
	// URL 是附件的下载 URL
	URL string `json:"url"`
	// FileType 是文件类型（MIME 类型）
	FileType string `json:"file_type"`
	// Size 是文件大小（字节）
	Size int `json:"size"`
	// Duration 是音频时长（秒），仅音频文件有效
	Duration int `json:"duration"`
	// Width 是图片/视频宽度（像素）
	Width int `json:"width"`
	// Height 是图片/视频高度（像素）
	Height int `json:"height"`
}

// Quote 表示消息引用的信息。
type Quote struct {
	// ID 是被引用消息的 ID
	ID string `json:"id"`
	// Type 是被引用消息的类型
	Type int `json:"type"`
	// RongID 是被引用消息的融云 ID
	RongID string `json:"rong_id"`
	// Content 是被引用消息的内容
	Content string `json:"content"`
	// CreateAt 是被引用消息的创建时间戳（毫秒）
	CreateAt int64 `json:"create_at"`
	// Author 是被引用消息的作者
	Author *User `json:"author"`
}

// Reaction 表示消息的回应（表情回复）信息。
type Reaction struct {
	// Emoji 是回应的表情
	Emoji Emoji `json:"emoji"`
	// Count 是回应数量
	Count int `json:"count"`
	// Me 表示当前用户是否已回应
	Me bool `json:"me"`
}

// Emoji 表示一个表情。
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

// Pin 表示频道中的置顶消息。
type Pin struct {
	// ChannelID 是置顶消息所属频道的 ID
	ChannelID string `json:"channel_id"`
	// UserID 是执行置顶操作的用户 ID
	UserID string `json:"user_id"`
	// MsgID 是被置顶的消息 ID
	MsgID string `json:"msg_id"`
	// Msg 是被置顶的消息内容
	Msg *Message `json:"msg"`
}
