package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// BlacklistUser 表示黑名单中的用户信息。
type BlacklistUser struct {
	UserID      string      `json:"user_id"`
	CreatedTime int64       `json:"created_time"`
	Remark      string      `json:"remark"`
	User        *model.User `json:"user"`
}

// ListBlacklist 获取服务器黑名单列表，支持分页。
// GET /blacklist/list
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
func ListBlacklist(ctx context.Context, client Doer, guildID string, page int, pageSize int) (*model.PageResult[BlacklistUser], error) {
	params := map[string]interface{}{
		"guild_id": guildID,
	}
	if page > 0 {
		params["page"] = page
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}

	var result model.PageResult[BlacklistUser]
	err := client.Do(ctx, "GET", "/blacklist/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AddBlacklist 将用户加入服务器黑名单。
// POST /blacklist/create
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - targetID: 目标用户 ID（必填）
//   - remark: 加入黑名单的备注原因
//   - delMsgDays: 删除最近 N 天的消息（0-7），默认为 0
func AddBlacklist(ctx context.Context, client Doer, guildID string, targetID string, remark string, delMsgDays int) error {
	body := map[string]interface{}{
		"guild_id":  guildID,
		"target_id": targetID,
	}
	if remark != "" {
		body["remark"] = remark
	}
	if delMsgDays > 0 {
		body["del_msg_days"] = delMsgDays
	}
	return client.Do(ctx, "POST", "/blacklist/create", body, nil)
}

// RemoveBlacklist 将用户从服务器黑名单中移除。
// POST /blacklist/delete
func RemoveBlacklist(ctx context.Context, client Doer, guildID string, targetID string) error {
	return client.Do(ctx, "POST", "/blacklist/delete", map[string]interface{}{
		"guild_id":  guildID,
		"target_id": targetID,
	}, nil)
}
