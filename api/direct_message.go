package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// ListDirectMessages 获取私聊消息列表，支持分页和筛选。
// GET /direct-message/list
//
// 参数说明：
//   - chatCode: 私聊会话标识（必填）
//   - msgID: 参考消息 ID，用于翻页定位
//   - flag: 筛选标记
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
func ListDirectMessages(ctx context.Context, client Doer, chatCode string, msgID string, flag string, page int, pageSize int) (*model.PageResult[model.Message], error) {
	params := map[string]interface{}{
		"chat_code": chatCode,
	}
	if msgID != "" {
		params["msg_id"] = msgID
	}
	if flag != "" {
		params["flag"] = flag
	}
	if page > 0 {
		params["page"] = page
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}

	var result model.PageResult[model.Message]
	err := client.Do(ctx, "GET", "/direct-message/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDirectMessage 获取指定私聊消息的详细信息。
// GET /direct-message/view?msg_id={msgID}
func GetDirectMessage(ctx context.Context, client Doer, msgID string) (*model.Message, error) {
	var result model.Message
	err := client.Do(ctx, "GET", "/direct-message/view", map[string]interface{}{
		"msg_id": msgID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDirectMessage 发送私聊消息。
// POST /direct-message/create
//
// targetID 和 chatCode 二选一：提供 chatCode 时使用已有会话，提供 targetID 时自动创建会话。
//
// 参数说明：
//   - targetID: 目标用户 ID
//   - chatCode: 私聊会话标识
//   - content: 消息内容（必填）
//   - quote: 引用消息 ID，为空表示不引用
//   - nonce: 去重标识，为空则不启用去重
func CreateDirectMessage(ctx context.Context, client Doer, targetID string, chatCode string, content string, quote string, nonce string) (*CreateMessageResponse, error) {
	body := map[string]interface{}{
		"content": content,
	}
	if targetID != "" {
		body["target_id"] = targetID
	}
	if chatCode != "" {
		body["chat_code"] = chatCode
	}
	if quote != "" {
		body["quote"] = quote
	}
	if nonce != "" {
		body["nonce"] = nonce
	}

	var result CreateMessageResponse
	err := client.Do(ctx, "POST", "/direct-message/create", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDirectMessage 更新已发送的私聊消息。
// POST /direct-message/update
//
// 参数说明：
//   - msgID: 消息 ID（必填）
//   - content: 新的消息内容（必填）
//   - quote: 引用消息 ID，为空表示不修改
func UpdateDirectMessage(ctx context.Context, client Doer, msgID string, content string, quote string) (*model.Message, error) {
	body := map[string]interface{}{
		"msg_id":  msgID,
		"content": content,
	}
	if quote != "" {
		body["quote"] = quote
	}

	var result model.Message
	err := client.Do(ctx, "POST", "/direct-message/update", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDirectMessage 删除指定的私聊消息。
// POST /direct-message/delete
func DeleteDirectMessage(ctx context.Context, client Doer, msgID string) error {
	return client.Do(ctx, "POST", "/direct-message/delete", map[string]interface{}{
		"msg_id": msgID,
	}, nil)
}

// ListDirectReactions 获取私聊消息回应的用户列表。
// GET /direct-message/reaction-list?msg_id={msgID}&emoji={emoji}
//
// emoji 参数为表情的数字标识，完整对照表参见 https://img.kookapp.cn/assets/emoji.json
func ListDirectReactions(ctx context.Context, client Doer, msgID string, emoji string) (*model.ListResult[model.User], error) {
	var result model.ListResult[model.User]
	err := client.Do(ctx, "GET", "/direct-message/reaction-list", map[string]interface{}{
		"msg_id": msgID,
		"emoji":  emoji,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AddDirectReaction 为私聊消息添加回应（表情）。
// POST /direct-message/add-reaction
//
// emoji 参数为表情的数字标识，完整对照表参见 https://img.kookapp.cn/assets/emoji.json
func AddDirectReaction(ctx context.Context, client Doer, msgID string, emoji string) error {
	return client.Do(ctx, "POST", "/direct-message/add-reaction", map[string]interface{}{
		"msg_id": msgID,
		"emoji":  emoji,
	}, nil)
}

// DeleteDirectReaction 删除私聊消息的回应（表情）。
// POST /direct-message/delete-reaction
//
// 参数说明：
//   - msgID: 消息 ID（必填）
//   - emoji: 表情的数字标识（必填），完整对照表参见 https://img.kookapp.cn/assets/emoji.json
//   - userID: 要删除回应的用户 ID，为空则删除当前用户的回应
func DeleteDirectReaction(ctx context.Context, client Doer, msgID string, emoji string, userID string) error {
	body := map[string]interface{}{
		"msg_id": msgID,
		"emoji":  emoji,
	}
	if userID != "" {
		body["user_id"] = userID
	}
	return client.Do(ctx, "POST", "/direct-message/delete-reaction", body, nil)
}
