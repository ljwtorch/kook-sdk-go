package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// ListChannels 获取指定服务器下的频道列表。
// GET /channel/list?guild_id={guildID}
func ListChannels(ctx context.Context, client Doer, guildID string) (*model.ListResult[model.Channel], error) {
	var result model.ListResult[model.Channel]
	err := client.Do(ctx, "GET", "/channel/list", map[string]interface{}{
		"guild_id": guildID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetChannel 获取指定频道的详细信息。
// GET /channel/view?target_id={targetID}
func GetChannel(ctx context.Context, client Doer, targetID string) (*model.Channel, error) {
	var result model.Channel
	err := client.Do(ctx, "GET", "/channel/view", map[string]interface{}{
		"target_id": targetID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateChannel 在指定服务器中创建新频道。
// POST /channel/create
//
// 当 isCategory 传 1 时，只接收 guildID、name、isCategory 三个字段。
//
// 参考文档：https://developer.kookapp.cn/doc/http/channel#创建频道
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
// POST /channel/update
//
// 参数说明：
//   - channelID: 频道 ID（必填）
//   - name: 频道名称
//   - topic: 频道主题
//   - slowMode: 慢速模式间隔（毫秒），传 0 表示关闭慢速模式，传 -1 表示不修改
func UpdateChannel(ctx context.Context, client Doer, channelID string, name string, topic string, slowMode int) (*model.Channel, error) {
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

	var result model.Channel
	err := client.Do(ctx, "POST", "/channel/update", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteChannel 删除指定频道。
// POST /channel/delete
func DeleteChannel(ctx context.Context, client Doer, channelID string) error {
	return client.Do(ctx, "POST", "/channel/delete", map[string]interface{}{
		"channel_id": channelID,
	}, nil)
}

// ListVoiceChannelUsers 获取语音频道中的用户列表。
// GET /channel/user-list?channel_id={channelID}
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
// POST /channel/move-user
//
// 参数说明：
//   - targetID: 目标语音频道 ID
//   - userIDs: 要移动的用户 ID 列表
func MoveVoiceUser(ctx context.Context, client Doer, targetID string, userIDs []string) error {
	return client.Do(ctx, "POST", "/channel/move-user", map[string]interface{}{
		"target_id": targetID,
		"user_ids":  userIDs,
	}, nil)
}

// KickoutVoiceUser 将用户踢出语音频道。
// POST /channel/kickout
func KickoutVoiceUser(ctx context.Context, client Doer, channelID string, userID string) error {
	return client.Do(ctx, "POST", "/channel/kickout", map[string]interface{}{
		"channel_id": channelID,
		"user_id":    userID,
	}, nil)
}
