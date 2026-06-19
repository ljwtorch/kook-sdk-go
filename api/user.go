package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// GetCurrentUser 获取当前认证用户（机器人）的信息。
// GET /user/me
func GetCurrentUser(ctx context.Context, client Doer) (*model.User, error) {
	var result model.User
	err := client.Do(ctx, "GET", "/user/me", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUser 获取目标用户的详细信息。
// GET /user/view?user_id={userID}&guild_id={guildID}
//
// 参数说明：
//   - userID: 目标用户 ID（必填）
//   - guildID: 服务器 ID，传入可获取该用户在服务器中的额外信息
func GetUser(ctx context.Context, client Doer, userID string, guildID string) (*model.User, error) {
	params := map[string]interface{}{
		"user_id": userID,
	}
	if guildID != "" {
		params["guild_id"] = guildID
	}

	var result model.User
	err := client.Do(ctx, "GET", "/user/view", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// OfflineBot 下线机器人，使其停止接收消息。
// POST /user/offline
func OfflineBot(ctx context.Context, client Doer) error {
	return client.Do(ctx, "POST", "/user/offline", nil, nil)
}

// OnlineBot 上线机器人，恢复接收消息。
// POST /user/online
func OnlineBot(ctx context.Context, client Doer) error {
	return client.Do(ctx, "POST", "/user/online", nil, nil)
}

// botOnlineStatus 用于解析机器人在线状态接口响应。
type botOnlineStatus struct {
	Online bool `json:"online"`
}

// GetBotOnlineStatus 获取机器人当前在线状态。
// GET /user/get-online-status
//
// 返回 true 表示机器人在线，false 表示离线。
func GetBotOnlineStatus(ctx context.Context, client Doer) (bool, error) {
	var result botOnlineStatus
	err := client.Do(ctx, "GET", "/user/get-online-status", nil, &result)
	if err != nil {
		return false, err
	}
	return result.Online, nil
}
