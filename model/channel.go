package model

import (
	"encoding/json"
	"fmt"
)

// 频道类型常量。
// KOOK API 返回的 type 字段为数字。
const (
	ChannelTypeCategory = 0 // 分组（分类）频道
	ChannelTypeText     = 1 // 文字频道
	ChannelTypeVoice    = 2 // 语音频道
)

// SlowMode 支持的值常量（毫秒）。
// 用于设置频道的慢速模式，仅对文字频道有效。
const (
	SlowModeOff   = 0        // 关闭慢速模式
	SlowMode5s    = 5000     // 5秒
	SlowMode10s   = 10000    // 10秒
	SlowMode15s   = 15000    // 15秒
	SlowMode30s   = 30000    // 30秒
	SlowMode1min  = 60000    // 1分钟
	SlowMode2min  = 120000   // 2分钟
	SlowMode5min  = 300000   // 5分钟
	SlowMode10min = 600000   // 10分钟
	SlowMode15min = 900000   // 15分钟
	SlowMode30min = 1800000  // 30分钟
	SlowMode1hour = 3600000  // 1小时
	SlowMode2hour = 7200000  // 2小时
	SlowMode6hour = 21600000 // 6小时
)

// FlexibleBool 是一个可以同时解析 JSON bool 和 int 类型的自定义类型。
// KOOK API 在不同接口中对 is_category 字段返回类型不一致：
//   - /channel/update 返回 int (0 或 1)
//   - /channel/view 返回 bool (true 或 false)
//
// FlexibleBool 统一处理这两种情况，内部存储为 int（0 或 1）。
type FlexibleBool int

// UnmarshalJSON 实现自定义 JSON 反序列化，支持 bool 和 int 两种格式。
func (fb *FlexibleBool) UnmarshalJSON(data []byte) error {
	// 尝试解析为 bool
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		if b {
			*fb = 1
		} else {
			*fb = 0
		}
		return nil
	}
	// 尝试解析为 int
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*fb = FlexibleBool(i)
		return nil
	}
	return fmt.Errorf("cannot unmarshal %s into FlexibleBool", string(data))
}

// Bool 返回 FlexibleBool 对应的 bool 值。
func (fb FlexibleBool) Bool() bool {
	return fb != 0
}

// Channel 表示 KOOK 频道的信息。
type Channel struct {
	// ID 是频道的唯一标识
	ID string `json:"id"`
	// Name 是频道名称
	Name string `json:"name"`
	// UserID 是频道创建者的用户 ID
	UserID string `json:"user_id"`
	// GuildID 是频道所属服务器的 ID
	GuildID string `json:"guild_id"`
	// Topic 是频道主题
	Topic string `json:"topic"`
	// IsCategory 表示是否为分组（分类）频道，1 是，0 否
	// KOOK API 在不同接口返回类型不一致（bool 或 int），使用 FlexibleBool 统一处理
	IsCategory FlexibleBool `json:"is_category"`
	// ParentID 是父频道的 ID（用于分组）
	ParentID string `json:"parent_id"`
	// Level 是频道排序层级
	Level int `json:"level"`
	// SlowMode 是慢速模式时间（毫秒），0 表示关闭慢速模式
	// 支持的值：0, 5000, 10000, 15000, 30000, 60000, 120000, 300000, 600000, 900000, 1800000, 3600000, 7200000, 21600000
	SlowMode int `json:"slow_mode"`
	// Type 是频道类型：0=分组, 1=文字频道, 2=语音频道
	Type int `json:"type"`
	// MasterID 是频道管理员的用户 ID
	MasterID string `json:"master_id"`
	// LastMsgContent 是频道最后一条消息的内容
	LastMsgContent string `json:"last_msg_content"`
	// LastMsgID 是频道最后一条消息的 ID
	LastMsgID string `json:"last_msg_id"`
	// HasPassword 表示频道是否设有密码
	HasPassword bool `json:"has_password"`
	// LimitAmount 是频道人数限制，0 表示不限制
	LimitAmount int `json:"limit_amount"`
	// PermissionSync 表示是否同步分组权限，1 表示同步
	PermissionSync int `json:"permission_sync"`
	// PermissionOverwrites 是频道的角色权限覆写列表
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
	// PermissionUsers 是频道的用户权限覆写列表
	PermissionUsers []PermissionUser `json:"permission_users"`
}

// PermissionOverwrite 表示频道的角色权限覆写规则。
type PermissionOverwrite struct {
	// RoleID 是角色 ID
	RoleID int64 `json:"role_id"`
	// Allow 是允许的权限位
	Allow int64 `json:"allow"`
	// Deny 是拒绝的权限位
	Deny int64 `json:"deny"`
}

// PermissionUser 表示频道的用户权限覆写规则。
type PermissionUser struct {
	// UserID 是用户 ID
	UserID string `json:"user_id"`
	// Allow 是允许的权限位
	Allow int64 `json:"allow"`
	// Deny 是拒绝的权限位
	Deny int64 `json:"deny"`
}

// ChannelRole 表示频道中的角色权限信息。
type ChannelRole struct {
	// RoleID 是角色 ID
	RoleID int64 `json:"role_id"`
	// Allow 是允许的权限位
	Allow int64 `json:"allow"`
	// Deny 是拒绝的权限位
	Deny int64 `json:"deny"`
}
