package event

// RoleAddedEvent 表示角色新增事件的数据。
type RoleAddedEvent struct {
	// RoleID 是新增角色的 ID
	RoleID int64 `json:"role_id"`
	// Name 是角色名称
	Name string `json:"name"`
}

// RoleDeletedEvent 表示角色删除事件的数据。
type RoleDeletedEvent struct {
	// RoleID 是被删除角色的 ID
	RoleID int64 `json:"role_id"`
	// Name 是角色名称
	Name string `json:"name"`
}

// RoleUpdatedEvent 表示角色更新事件的数据。
type RoleUpdatedEvent struct {
	// RoleID 是被更新角色的 ID
	RoleID int64 `json:"role_id"`
	// Name 是更新后的角色名称
	Name string `json:"name"`
}
