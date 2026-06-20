package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// GetGuildList 获取当前用户加入的服务器列表。
// 参考文档：https://developer.kookapp.cn/doc/http/guild#获取当前用户加入的服务器列表
// GET /guild/list
//
// 参数说明：
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
//   - sort: 排序字段，如 "-id" 表示按 ID 降序，"id" 表示按 ID 升序
func GetGuildList(ctx context.Context, client Doer, page int, pageSize int, sort string) (*model.PageResult[model.Guild], error) {
	params := map[string]interface{}{}
	if page > 0 {
		params["page"] = page
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}
	if sort != "" {
		params["sort"] = sort
	}

	var result model.PageResult[model.Guild]
	err := client.Do(ctx, "GET", "/guild/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetGuild 获取指定服务器的详细信息，包含角色和频道列表。
// 参考文档：https://developer.kookapp.cn/doc/http/guild#获取服务器详情
// GET /guild/view
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
func GetGuild(ctx context.Context, client Doer, guildID string) (*model.GuildDetail, error) {
	var result model.GuildDetail
	err := client.Do(ctx, "GET", "/guild/view", map[string]interface{}{
		"guild_id": guildID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetGuildUserList 获取服务器成员列表，支持分页和多条件筛选。
// 参考文档：https://developer.kookapp.cn/doc/http/guild#获取服务器中的用户列表
// GET /guild/user-list
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
//   - channelID: 频道 ID，筛选在该频道的用户
//   - search: 搜索关键字，在用户名或昵称中搜索
//   - roleID: 角色 ID，筛选拥有指定角色的用户
//   - mobileVerified: 手机认证筛选，0=未认证，1=已认证，传 -1 不筛选
//   - activeTime: 根据活跃时间排序，0=顺序，1=倒序，传 -1 不排序
//   - joinedAt: 根据加入时间排序，0=顺序，1=倒序，传 -1 不排序
//   - filterUserID: 获取指定 ID 所属用户的信息
func GetGuildUserList(ctx context.Context, client Doer, guildID string, page int, pageSize int, channelID string, search string, roleID int64, mobileVerified int, activeTime int, joinedAt int, filterUserID string) (*model.PageResult[model.GuildUser], error) {
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
	if mobileVerified >= 0 {
		params["mobile_verified"] = mobileVerified
	}
	if activeTime >= 0 {
		params["active_time"] = activeTime
	}
	if joinedAt >= 0 {
		params["joined_at"] = joinedAt
	}
	if filterUserID != "" {
		params["filter_user_id"] = filterUserID
	}

	var result model.PageResult[model.GuildUser]
	err := client.Do(ctx, "GET", "/guild/user-list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SetGuildNickname 修改用户在服务器中的昵称。
// 参考文档：https://developer.kookapp.cn/doc/http/guild#修改服务器中用户的昵称
// POST /guild/nickname
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - nickname: 新昵称，2-64 长度，空字符串表示清空昵称
//   - userID: 目标用户 ID，为空则修改当前登录用户的昵称
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
// 参考文档：https://developer.kookapp.cn/doc/http/guild#离开服务器
// POST /guild/leave
func LeaveGuild(ctx context.Context, client Doer, guildID string) error {
	return client.Do(ctx, "POST", "/guild/leave", map[string]interface{}{
		"guild_id": guildID,
	}, nil)
}

// KickoutGuildMember 将指定用户踢出服务器。
// 参考文档：https://developer.kookapp.cn/doc/http/guild#踢出服务器
// POST /guild/kickout
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - targetID: 目标用户 ID（必填）
func KickoutGuildMember(ctx context.Context, client Doer, guildID string, targetID string) error {
	return client.Do(ctx, "POST", "/guild/kickout", map[string]interface{}{
		"guild_id":  guildID,
		"target_id": targetID,
	}, nil)
}
