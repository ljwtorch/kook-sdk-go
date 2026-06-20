package api

import (
	"context"
)

// GuildMuteListResult 表示服务器静音/闭麦列表的响应结果。
type GuildMuteListResult struct {
	// Mic 是麦克风闭麦信息（type=1）
	Mic *GuildMuteGroup `json:"mic"`
	// Headset 是耳机静音信息（type=2）
	Headset *GuildMuteGroup `json:"headset"`
}

// GuildMuteGroup 表示一组静音/闭麦用户。
type GuildMuteGroup struct {
	// Type 表示静音类型：1=麦克风闭麦，2=耳机静音
	Type int `json:"type"`
	// UserIDs 是被静音/闭麦的用户 ID 列表
	UserIDs []string `json:"user_ids"`
}

// ListGuildMutes 获取服务器静音/闭麦用户列表。
// 参考文档：https://developer.kookapp.cn/doc/http/guild#服务器静音闭麦列表
// GET /guild-mute/list
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - returnType: 返回格式，建议传 "detail"，其他值仅作兼容
func ListGuildMutes(ctx context.Context, client Doer, guildID string, returnType string) (*GuildMuteListResult, error) {
	params := map[string]interface{}{
		"guild_id": guildID,
	}
	if returnType != "" {
		params["return_type"] = returnType
	}

	var result GuildMuteListResult
	err := client.Do(ctx, "GET", "/guild-mute/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateGuildMute 对服务器中的用户添加静音或闭麦。
// 参考文档：https://developer.kookapp.cn/doc/http/guild#添加服务器静音或闭麦
// POST /guild-mute/create
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - userID: 目标用户 ID（必填）
//   - muteType: 静音类型，1=麦克风闭麦，2=耳机静音
func CreateGuildMute(ctx context.Context, client Doer, guildID string, userID string, muteType int) error {
	return client.Do(ctx, "POST", "/guild-mute/create", map[string]interface{}{
		"guild_id": guildID,
		"user_id":  userID,
		"type":     muteType,
	}, nil)
}

// DeleteGuildMute 取消服务器中用户的静音或闭麦状态。
// 参考文档：https://developer.kookapp.cn/doc/http/guild#删除服务器静音或闭麦
// POST /guild-mute/delete
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - userID: 目标用户 ID（必填）
//   - muteType: 静音类型，1=麦克风闭麦，2=耳机静音
func DeleteGuildMute(ctx context.Context, client Doer, guildID string, userID string, muteType int) error {
	return client.Do(ctx, "POST", "/guild-mute/delete", map[string]interface{}{
		"guild_id": guildID,
		"user_id":  userID,
		"type":     muteType,
	}, nil)
}
