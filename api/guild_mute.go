package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// GuildMuteUser 表示服务器静音列表中的用户信息。
type GuildMuteUser struct {
	UserID      string `json:"user_id"`
	CreatedTime int64  `json:"created_time"`
	// Type 表示静音类型："1" 为静音（耳机），"2" 为闭麦（麦克风）。
	Type string `json:"type"`
}

// ListGuildMutes 获取服务器静音/闭麦用户列表，支持分页。
// GET /guild-mute/list
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - returnType: 返回类型筛选，如 "1" 静音 / "2" 闭麦
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
func ListGuildMutes(ctx context.Context, client Doer, guildID string, returnType string, page int, pageSize int) (*model.PageResult[GuildMuteUser], error) {
	params := map[string]interface{}{
		"guild_id": guildID,
	}
	if returnType != "" {
		params["return_type"] = returnType
	}
	if page > 0 {
		params["page"] = page
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}

	var result model.PageResult[GuildMuteUser]
	err := client.Do(ctx, "GET", "/guild-mute/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateGuildMute 对服务器中的用户添加静音或闭麦。
// POST /guild-mute/create
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - userID: 目标用户 ID（必填）
//   - muteType: 静音类型，"1" 为静音（耳机），"2" 为闭麦（麦克风）
func CreateGuildMute(ctx context.Context, client Doer, guildID string, userID string, muteType string) error {
	return client.Do(ctx, "POST", "/guild-mute/create", map[string]interface{}{
		"guild_id": guildID,
		"user_id":  userID,
		"type":     muteType,
	}, nil)
}

// DeleteGuildMute 取消服务器中用户的静音或闭麦状态。
// POST /guild-mute/delete
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - userID: 目标用户 ID（必填）
//   - muteType: 静音类型，"1" 为静音（耳机），"2" 为闭麦（麦克风）
func DeleteGuildMute(ctx context.Context, client Doer, guildID string, userID string, muteType string) error {
	return client.Do(ctx, "POST", "/guild-mute/delete", map[string]interface{}{
		"guild_id": guildID,
		"user_id":  userID,
		"type":     muteType,
	}, nil)
}
