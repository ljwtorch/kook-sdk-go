// Package event 定义了 KOOK 系统事件（type=255）的 extra 字段结构。
package event

import "github.com/ljwtorch/kook-sdk-go/model"

// ChannelAddedEvent 表示频道新增事件的数据。
type ChannelAddedEvent struct {
	// ID 是新增频道的唯一标识
	ID string `json:"id"`
	// GuildID 是频道所属服务器的 ID
	GuildID string `json:"guild_id"`
	// Name 是频道名称
	Name string `json:"name"`
	// UserID 是频道创建者的用户 ID
	UserID string `json:"user_id"`
	// Type 是频道类型
	Type string `json:"type"`
	// ParentID 是父频道的 ID
	ParentID string `json:"parent_id"`
	// Level 是频道排序层级
	Level int `json:"level"`
	// LimitAmount 是频道人数限制
	LimitAmount int `json:"limit_amount"`
	// IsCategory 表示是否为分组频道
	IsCategory model.FlexibleBool `json:"is_category"`
}

// ChannelUpdatedEvent 表示频道更新事件的数据。
type ChannelUpdatedEvent struct {
	// ID 是更新频道的唯一标识
	ID string `json:"id"`
	// GuildID 是频道所属服务器的 ID
	GuildID string `json:"guild_id"`
	// Name 是频道名称
	Name string `json:"name"`
	// UserID 是频道创建者的用户 ID
	UserID string `json:"user_id"`
	// Type 是频道类型
	Type string `json:"type"`
	// ParentID 是父频道的 ID
	ParentID string `json:"parent_id"`
	// Level 是频道排序层级
	Level int `json:"level"`
	// LimitAmount 是频道人数限制
	LimitAmount int `json:"limit_amount"`
	// IsCategory 表示是否为分组频道
	IsCategory model.FlexibleBool `json:"is_category"`
}

// ChannelDeletedEvent 表示频道删除事件的数据。
type ChannelDeletedEvent struct {
	// ID 是被删除频道的唯一标识
	ID string `json:"id"`
	// DeletedAt 是频道删除的时间戳（毫秒）
	DeletedAt int64 `json:"deleted_at"`
}
