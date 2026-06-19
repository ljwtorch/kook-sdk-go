package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// GetChannelRoles 获取频道角色权限列表。
// GET /channel-role/index?channel_id={channelID}
func GetChannelRoles(ctx context.Context, client Doer, channelID string) (*model.ListResult[model.ChannelRole], error) {
	var result model.ListResult[model.ChannelRole]
	err := client.Do(ctx, "GET", "/channel-role/index", map[string]interface{}{
		"channel_id": channelID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateChannelRole 为频道创建角色权限。
// POST /channel-role/create
//
// 参数说明：
//   - channelID: 频道 ID（必填）
//   - typ: 权限对象类型（"role_id" 表示角色，"user_id" 表示用户）
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
// POST /channel-role/update
//
// 参数说明：
//   - channelID: 频道 ID（必填）
//   - typ: 权限对象类型（"role_id" 或 "user_id"）
//   - value: 权限对象的值
//   - allow: 允许的权限位掩码
//   - deny: 拒绝的权限位掩码
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

// SyncChannelRole 将频道角色权限同步到服务器默认设置。
// POST /channel-role/sync
func SyncChannelRole(ctx context.Context, client Doer, channelID string) error {
	return client.Do(ctx, "POST", "/channel-role/sync", map[string]interface{}{
		"channel_id": channelID,
	}, nil)
}

// DeleteChannelRole 删除频道的角色权限设置。
// POST /channel-role/delete
//
// 参数说明：
//   - channelID: 频道 ID（必填）
//   - typ: 权限对象类型（"role_id" 或 "user_id"）
//   - value: 权限对象的值
func DeleteChannelRole(ctx context.Context, client Doer, channelID string, typ string, value string) error {
	return client.Do(ctx, "POST", "/channel-role/delete", map[string]interface{}{
		"channel_id": channelID,
		"type":       typ,
		"value":      value,
	}, nil)
}
