package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// ListGuildRoles 获取服务器角色列表，支持分页。
// GET /guild-role/list
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
func ListGuildRoles(ctx context.Context, client Doer, guildID string, page int, pageSize int) (*model.PageResult[model.Role], error) {
	params := map[string]interface{}{
		"guild_id": guildID,
	}
	if page > 0 {
		params["page"] = page
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}

	var result model.PageResult[model.Role]
	err := client.Do(ctx, "GET", "/guild-role/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateGuildRole 在服务器中创建新角色。
// POST /guild-role/create
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - name: 角色名称，为空则使用默认名称
func CreateGuildRole(ctx context.Context, client Doer, guildID string, name string) (*model.Role, error) {
	body := map[string]interface{}{
		"guild_id": guildID,
	}
	if name != "" {
		body["name"] = name
	}

	var result model.Role
	err := client.Do(ctx, "POST", "/guild-role/create", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateGuildRole 更新服务器角色的属性。
// POST /guild-role/update
//
// 注意：正常角色由上向下排序，这个先后顺序是角色的优先级（position 字段）。
// 如果你有管理员权限，你只能管理优先级比自己低的用户，不能管理优先级等于或比自己高的用户。
// 因此，在使用授予权限、更新等接口时，要注意一下，可能机器人虽然有管理权限，
// 但是也不是什么角色都可以授予，也不是什么人都可以操作。
// 参考文档：https://developer.kookapp.cn/doc/http/guild-role#更新服务器角色
//
// color 参数使用 24 位 RGB 颜色整数，参见 model.Role.Color。
// permissions 参数为权限位掩码，权限常量定义在 model 包中（如 model.PermAdmin）。
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - roleID: 角色 ID（必填）
//   - name: 角色名称，为空表示不修改
//   - color: 角色颜色（24 位 RGB 十进制整数），传 0 表示不修改。计算公式：(R << 16) | (G << 8) | B
//   - hoist: 是否在成员列表中单独显示
//   - mentionable: 是否允许被 @提及
//   - permissions: 权限位掩码，传 0 表示不修改。权限常量定义在 model 包中（如 model.PermAdmin）
func UpdateGuildRole(ctx context.Context, client Doer, guildID string, roleID int64, name string, color int, hoist bool, mentionable bool, permissions int64) (*model.Role, error) {
	body := map[string]interface{}{
		"guild_id": guildID,
		"role_id":  roleID,
	}
	if name != "" {
		body["name"] = name
	}
	if color > 0 {
		body["color"] = color
	}
	if hoist {
		body["hoist"] = 1
	}
	if mentionable {
		body["mentionable"] = 1
	}
	if permissions > 0 {
		body["permissions"] = permissions
	}

	var result model.Role
	err := client.Do(ctx, "POST", "/guild-role/update", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteGuildRole 删除服务器中的指定角色。
// POST /guild-role/delete
func DeleteGuildRole(ctx context.Context, client Doer, guildID string, roleID int64) error {
	return client.Do(ctx, "POST", "/guild-role/delete", map[string]interface{}{
		"guild_id": guildID,
		"role_id":  roleID,
	}, nil)
}

// GrantRole 为服务器中的用户赋予指定角色。
// POST /guild-role/grant
func GrantRole(ctx context.Context, client Doer, guildID string, userID string, roleID int64) (*model.Role, error) {
	var result model.Role
	err := client.Do(ctx, "POST", "/guild-role/grant", map[string]interface{}{
		"guild_id": guildID,
		"user_id":  userID,
		"role_id":  roleID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RevokeRole 移除服务器中用户的指定角色。
// POST /guild-role/revoke
func RevokeRole(ctx context.Context, client Doer, guildID string, userID string, roleID int64) error {
	return client.Do(ctx, "POST", "/guild-role/revoke", map[string]interface{}{
		"guild_id": guildID,
		"user_id":  userID,
		"role_id":  roleID,
	}, nil)
}
