package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// GetGuildList 获取当前用户加入的服务器列表。
// GET /guild/list
func GetGuildList(ctx context.Context, client Doer) (*model.PageResult[model.Guild], error) {
	var result model.PageResult[model.Guild]
	err := client.Do(ctx, "GET", "/guild/list", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetGuild 获取指定服务器的详细信息。
// GET /guild/view?guild_id={guildID}
func GetGuild(ctx context.Context, client Doer, guildID string) (*model.Guild, error) {
	var result model.Guild
	err := client.Do(ctx, "GET", "/guild/view", map[string]interface{}{
		"guild_id": guildID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetGuildUserList 获取服务器成员列表，支持分页和多条件筛选。
// GET /guild/user-list
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
//   - channelID: 频道 ID，筛选在该频道的用户
//   - search: 搜索关键字
//   - roleID: 角色 ID，筛选拥有指定角色的用户
//   - mobileVerified: 是否只筛选已验证手机号的用户
//   - activeTime: 活跃时间筛选（天）
//   - joinedAt: 加入时间筛选（时间戳）
func GetGuildUserList(ctx context.Context, client Doer, guildID string, page int, pageSize int, channelID string, search string, roleID int64, mobileVerified bool, activeTime int, joinedAt int) (*model.PageResult[model.GuildUser], error) {
	params := map[string]interface{}{
		"guild_id": guildID,
	}
	if page > 0 {
		params["page"] = page
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}
	if channelID != "" {
		params["channel_id"] = channelID
	}
	if search != "" {
		params["search"] = search
	}
	if roleID > 0 {
		params["role_id"] = roleID
	}
	if mobileVerified {
		params["mobile_verified"] = true
	}
	if activeTime > 0 {
		params["active_time"] = activeTime
	}
	if joinedAt > 0 {
		params["joined_at"] = joinedAt
	}

	var result model.PageResult[model.GuildUser]
	err := client.Do(ctx, "GET", "/guild/user-list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SetGuildNickname 修改用户在服务器中的昵称。
// POST /guild/nickname
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - nickname: 新昵称，空字符串表示删除昵称
//   - userID: 目标用户 ID，为空则修改当前用户
func SetGuildNickname(ctx context.Context, client Doer, guildID string, nickname string, userID string) error {
	body := map[string]interface{}{
		"guild_id": guildID,
		"nickname": nickname,
	}
	if userID != "" {
		body["user_id"] = userID
	}
	return client.Do(ctx, "POST", "/guild/nickname", body, nil)
}

// LeaveGuild 离开指定服务器。
// POST /guild/leave
func LeaveGuild(ctx context.Context, client Doer, guildID string) error {
	return client.Do(ctx, "POST", "/guild/leave", map[string]interface{}{
		"guild_id": guildID,
	}, nil)
}

// KickoutGuildMember 将指定用户踢出服务器。
// POST /guild/kickout
func KickoutGuildMember(ctx context.Context, client Doer, guildID string, targetID string) error {
	return client.Do(ctx, "POST", "/guild/kickout", map[string]interface{}{
		"guild_id":  guildID,
		"target_id": targetID,
	}, nil)
}
