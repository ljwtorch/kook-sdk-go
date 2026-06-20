package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// ChannelRoleIndexResult 表示获取频道角色权限详情的响应结果。
type ChannelRoleIndexResult struct {
	// PermissionOverwrites 是频道权限覆写的角色列表
	PermissionOverwrites []model.PermissionOverwrite `json:"permission_overwrites"`
	// PermissionUsers 是频道权限覆写的用户列表
	PermissionUsers []model.PermissionUser `json:"permission_users"`
	// PermissionSync 表示是否与分组频道同步权限，1 表示同步
	PermissionSync int `json:"permission_sync"`
}

// GetChannelRoles 获取频道角色权限详情。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#频道角色权限详情
// GET /channel-role/index
//
// 参数说明：
//   - channelID: 频道 ID（必填）
func GetChannelRoles(ctx context.Context, client Doer, channelID string) (*ChannelRoleIndexResult, error) {
	var result ChannelRoleIndexResult
	err := client.Do(ctx, "GET", "/channel-role/index", map[string]interface{}{
		"channel_id": channelID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateChannelRole 为频道创建角色权限。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#创建频道角色权限
// POST /channel-role/create
//
// 如果频道是分组的 ID，会同步给所有 sync=1 的子频道。
//
// 参数说明：
//   - channelID: 频道 ID（必填）
//   - typ: 权限对象类型，"role_id" 表示角色，"user_id" 表示用户，不传默认为 "user_id"
//   - value: 权限对象的值（角色 ID 或用户 ID）
func CreateChannelRole(ctx context.Context, client Doer, channelID string, typ string, value string) (*model.ChannelRole, error) {
	body := map[string]interface{}{
		"channel_id": channelID,
	}
	if typ != "" {
		body["type"] = typ
	}
	if value != "" {
		body["value"] = value
	}

	var result model.ChannelRole
	err := client.Do(ctx, "POST", "/channel-role/create", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateChannelRole 更新频道角色权限设置。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#更新频道角色权限
// POST /channel-role/update
//
// 如果频道是分组的 ID，会同步给所有 sync=1 的子频道。
//
// 参数说明：
//   - channelID: 频道 ID（必填）
//   - typ: 权限对象类型，"role_id" 或 "user_id"，不传默认为 "user_id"
//   - value: 权限对象的值（角色 ID 或用户 ID）
//   - allow: 允许的权限位掩码，传 0 表示不修改
//   - deny: 拒绝的权限位掩码，传 0 表示不修改
func UpdateChannelRole(ctx context.Context, client Doer, channelID string, typ string, value string, allow int64, deny int64) (*model.ChannelRole, error) {
	body := map[string]interface{}{
		"channel_id": channelID,
		"type":       typ,
		"value":      value,
		"allow":      allow,
		"deny":       deny,
	}

	var result model.ChannelRole
	err := client.Do(ctx, "POST", "/channel-role/update", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SyncChannelRole 将频道角色权限同步到分组默认设置。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#同步频道角色权限
// POST /channel-role/sync
//
// 参数说明：
//   - channelID: 频道 ID（必填）
func SyncChannelRole(ctx context.Context, client Doer, channelID string) error {
	return client.Do(ctx, "POST", "/channel-role/sync", map[string]interface{}{
		"channel_id": channelID,
	}, nil)
}

// DeleteChannelRole 删除频道的角色权限设置。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#删除频道角色权限
// POST /channel-role/delete
//
// 如果频道是分组的 ID，会同步给所有 sync=1 的子频道。
//
// 参数说明：
//   - channelID: 频道 ID（必填）
//   - typ: 权限对象类型，"role_id" 或 "user_id"，不传默认为 "user_id"
//   - value: 权限对象的值（角色 ID 或用户 ID）
func DeleteChannelRole(ctx context.Context, client Doer, channelID string, typ string, value string) error {
	return client.Do(ctx, "POST", "/channel-role/delete", map[string]interface{}{
		"channel_id": channelID,
		"type":       typ,
		"value":      value,
	}, nil)
}
