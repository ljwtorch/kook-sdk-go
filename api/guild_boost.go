package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// BoostRecord 表示服务器助力记录。
type BoostRecord struct {
	UserID    string      `json:"user_id"`
	GuildID   string      `json:"guild_id"`
	StartTime int64       `json:"start_time"`
	EndTime   int64       `json:"end_time"`
	User      *model.User `json:"user"`
}

// GetBoostHistory 获取服务器助力历史记录。
// GET /guild-boost/history
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - startTime: 查询起始时间（Unix 时间戳，秒）
//   - endTime: 查询结束时间（Unix 时间戳，秒）
func GetBoostHistory(ctx context.Context, client Doer, guildID string, startTime int64, endTime int64) (*model.ListResult[BoostRecord], error) {
	params := map[string]interface{}{
		"guild_id": guildID,
	}
	if startTime > 0 {
		params["start_time"] = startTime
	}
	if endTime > 0 {
		params["end_time"] = endTime
	}

	var result model.ListResult[BoostRecord]
	err := client.Do(ctx, "GET", "/guild-boost/history", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
