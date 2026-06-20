package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// ListChannels 获取指定服务器下的频道列表。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#获取频道列表
// GET /channel/list
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
//   - typ: 频道类型，1=文字频道，2=语音频道，传 0 不筛选
//   - parentID: 父分组频道 ID，为空则不筛选
func ListChannels(ctx context.Context, client Doer, guildID string, page int, pageSize int, typ int, parentID string) (*model.ListResult[model.Channel], error) {
	params := map[string]interface{}{
		"guild_id": guildID,
	}
	if page > 0 {
		params["page"] = page
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}
	if typ > 0 {
		params["type"] = typ
	}
	if parentID != "" {
		params["parent_id"] = parentID
	}

	var result model.ListResult[model.Channel]
	err := client.Do(ctx, "GET", "/channel/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetChannel 获取指定频道的详细信息。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#获取频道详情
// GET /channel/view
//
// 参数说明：
//   - targetID: 频道 ID（必填）
//   - needChildren: 是否需要获取子频道列表
func GetChannel(ctx context.Context, client Doer, targetID string, needChildren bool) (*model.Channel, error) {
	params := map[string]interface{}{
		"target_id": targetID,
	}
	if needChildren {
		params["need_children"] = true
	}

	var result model.Channel
	err := client.Do(ctx, "GET", "/channel/view", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateChannel 在指定服务器中创建新频道。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#创建频道
// POST /channel/create
//
// 当 isCategory 传 1 时，只接收 guildID、name、isCategory 三个字段。
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - name: 频道名称（必填）
//   - typ: 频道类型，1=文字频道, 2=语音频道，默认为 1
//   - parentID: 父分组频道 ID，为空则不分组
//   - limitAmount: 语音频道人数限制，最大 99
//   - voiceQuality: 语音音质，1=流畅, 2=正常, 3=高质量，默认为 2
//   - isCategory: 是否是分组，1=是, 0=否，默认为 0。当该值传 1 时，只接收 guildID、name、isCategory 三个字段
func CreateChannel(ctx context.Context, client Doer, guildID string, name string, typ int, parentID string, limitAmount int, voiceQuality string, isCategory int) (*model.Channel, error) {
	body := map[string]interface{}{
		"guild_id": guildID,
		"name":     name,
	}
	if typ > 0 {
		body["type"] = typ
	}
	if parentID != "" {
		body["parent_id"] = parentID
	}
	if limitAmount > 0 {
		body["limit_amount"] = limitAmount
	}
	if voiceQuality != "" {
		body["voice_quality"] = voiceQuality
	}
	if isCategory > 0 {
		body["is_category"] = isCategory
	}

	var result model.Channel
	err := client.Do(ctx, "POST", "/channel/create", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateChannel 编辑指定频道的信息。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#编辑频道
// POST /channel/update
//
// 参数说明：
//   - channelID: 频道 ID（必填）
//   - name: 频道名称，为空则不修改
//   - topic: 频道简介，文字频道有效，为空则不修改
//   - slowMode: 慢速模式间隔（毫秒），传 -1 表示不修改。支持的值：0, 5000, 10000, 15000, 30000, 60000, 120000, 300000, 600000, 900000, 1800000, 3600000, 7200000, 21600000
//   - level: 频道排序，传 -1 表示不修改
//   - parentID: 分组频道 ID，设置为 "0" 则移出分组，为空则不修改
//   - limitAmount: 语音频道人数限制，最大 99，传 -1 表示不修改
//   - voiceQuality: 语音音质，1=流畅, 2=正常, 3=高质量，语音频道有效，为空则不修改
//   - password: 语音频道密码，为空则不修改
func UpdateChannel(ctx context.Context, client Doer, channelID string, name string, topic string, slowMode int, level int, parentID string, limitAmount int, voiceQuality string, password string) (*model.Channel, error) {
	body := map[string]interface{}{
		"channel_id": channelID,
	}
	if name != "" {
		body["name"] = name
	}
	if topic != "" {
		body["topic"] = topic
	}
	if slowMode >= 0 {
		body["slow_mode"] = slowMode
	}
	if level >= 0 {
		body["level"] = level
	}
	if parentID != "" {
		body["parent_id"] = parentID
	}
	if limitAmount >= 0 {
		body["limit_amount"] = limitAmount
	}
	if voiceQuality != "" {
		body["voice_quality"] = voiceQuality
	}
	if password != "" {
		body["password"] = password
	}

	var result model.Channel
	err := client.Do(ctx, "POST", "/channel/update", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteChannel 删除指定频道。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#删除频道
// POST /channel/delete
func DeleteChannel(ctx context.Context, client Doer, channelID string) error {
	return client.Do(ctx, "POST", "/channel/delete", map[string]interface{}{
		"channel_id": channelID,
	}, nil)
}

// ListVoiceChannelUsers 获取语音频道中的用户列表。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#语音频道用户列表
// GET /channel/user-list
//
// 参数说明：
//   - channelID: 频道 ID（必填）
func ListVoiceChannelUsers(ctx context.Context, client Doer, channelID string) ([]model.GuildUser, error) {
	var result []model.GuildUser
	err := client.Do(ctx, "GET", "/channel/user-list", map[string]interface{}{
		"channel_id": channelID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MoveVoiceUser 将用户移动到指定的语音频道。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#语音频道之间移动用户
// POST /channel/move-user
//
// 只能在语音频道之间移动，用户也必须在其他语音频道在线才能够移动到目标频道。
//
// 参数说明：
//   - targetID: 目标语音频道 ID（必填）
//   - userIDs: 要移动的用户 ID 列表（必填）
func MoveVoiceUser(ctx context.Context, client Doer, targetID string, userIDs []string) error {
	return client.Do(ctx, "POST", "/channel/move-user", map[string]interface{}{
		"target_id": targetID,
		"user_ids":  userIDs,
	}, nil)
}

// KickoutVoiceUser 将用户踢出语音频道。
// 参考文档：https://developer.kookapp.cn/doc/http/channel#踢出语音频道中的用户
// POST /channel/kickout
//
// 只能踢出在语音频道中的用户。
//
// 参数说明：
//   - channelID: 目标频道 ID（必填），需要是语音频道
//   - userID: 用户 ID（必填）
func KickoutVoiceUser(ctx context.Context, client Doer, channelID string, userID string) error {
	return client.Do(ctx, "POST", "/channel/kickout", map[string]interface{}{
		"channel_id": channelID,
		"user_id":    userID,
	}, nil)
}
