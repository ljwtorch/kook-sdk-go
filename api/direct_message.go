package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// ListDirectMessages 获取私聊消息列表，基于参考消息进行游标分页。
// 参考文档：https://developer.kookapp.cn/doc/http/direct-message#获取私信聊天消息列表
// GET /direct-message/list
//
// chatCode 和 targetID 二选一：提供 chatCode 时使用已有会话，提供 targetID 时自动创建会话。
//
// 参数说明：
//   - chatCode: 私聊会话 Code，与 targetID 二选一
//   - targetID: 目标用户 ID，与 chatCode 二选一
//   - msgID: 参考消息 ID，用于翻页定位，不传则查询最新消息
//   - flag: 查询模式，"before"（参考消息之前）、"around"（参考消息前后）、"after"（参考消息之后），不传则查询最新消息
//   - page: 目标页数，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值（50）
func ListDirectMessages(ctx context.Context, client Doer, chatCode string, targetID string, msgID string, flag string, page int, pageSize int) (*model.ListResult[model.Message], error) {
	params := map[string]interface{}{}
	if chatCode != "" {
		params["chat_code"] = chatCode
	}
	if targetID != "" {
		params["target_id"] = targetID
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

	var result model.ListResult[model.Message]
	err := client.Do(ctx, "GET", "/direct-message/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDirectMessage 获取指定私聊消息的详细信息。
// 参考文档：https://developer.kookapp.cn/doc/http/direct-message#获取私信聊天消息详情
// GET /direct-message/view
//
// 参数说明：
//   - chatCode: 私聊会话 Code（必填）
//   - msgID: 私聊消息 ID（必填）
func GetDirectMessage(ctx context.Context, client Doer, chatCode string, msgID string) (*model.Message, error) {
	var result model.Message
	err := client.Do(ctx, "GET", "/direct-message/view", map[string]interface{}{
		"chat_code": chatCode,
		"msg_id":    msgID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDirectMessage 发送私聊消息。
// 参考文档：https://developer.kookapp.cn/doc/http/direct-message#发送私信聊天消息
// POST /direct-message/create
//
// targetID 和 chatCode 二选一：提供 chatCode 时使用已有会话，提供 targetID 时自动创建会话。
//
// 参数说明：
//   - targetID: 目标用户 ID，与 chatCode 二选一
//   - chatCode: 私聊会话 Code，与 targetID 二选一
//   - content: 消息内容（必填）
//   - msgType: 消息类型，0 表示自动检测，1 表示文本（默认），9 表示 KMarkdown，10 表示卡片消息
//   - quote: 引用消息 ID，为空表示不引用
//   - nonce: 去重标识，为空则不启用去重
//   - templateID: 模板消息 ID，使用后 content 作为模板消息的 input
//   - replyMsgID: 当前消息回复的用户 5 分钟内发送给当前机器人消息的 ID，用于配额折扣
func CreateDirectMessage(ctx context.Context, client Doer, targetID string, chatCode string, content string, msgType int, quote string, nonce string, templateID string, replyMsgID string) (*CreateMessageResponse, error) {
	body := map[string]interface{}{
		"content": content,
	}
	if msgType != 0 {
		body["type"] = msgType
	} else if len(content) > 0 && content[0] == '[' {
		body["type"] = 10
	} else {
		body["type"] = 1
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
	if templateID != "" {
		body["template_id"] = templateID
	}
	if replyMsgID != "" {
		body["reply_msg_id"] = replyMsgID
	}

	var result CreateMessageResponse
	err := client.Do(ctx, "POST", "/direct-message/create", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDirectMessage 更新已发送的私聊消息。
// 参考文档：https://developer.kookapp.cn/doc/http/direct-message#更新私信聊天消息
// POST /direct-message/update
//
// 目前支持消息 type 为 9、10 的修改，即 KMarkdown 和 CardMessage。
//
// 参数说明：
//   - msgID: 消息 ID（必填）
//   - content: 新的消息内容（必填）
//   - quote: 引用消息 ID，为空表示删除引用，不传则无影响
//   - templateID: 模板消息 ID，使用后 content 作为模板消息的 input
//   - replyMsgID: 当前消息回复的用户 5 分钟内发送给当前机器人消息的 ID，用于配额折扣
func UpdateDirectMessage(ctx context.Context, client Doer, msgID string, content string, quote string, templateID string, replyMsgID string) error {
	body := map[string]interface{}{
		"msg_id":  msgID,
		"content": content,
	}
	if quote != "" {
		body["quote"] = quote
	}
	if templateID != "" {
		body["template_id"] = templateID
	}
	if replyMsgID != "" {
		body["reply_msg_id"] = replyMsgID
	}

	return client.Do(ctx, "POST", "/direct-message/update", body, nil)
}

// DeleteDirectMessage 删除指定的私聊消息，只能删除自己的消息。
// 参考文档：https://developer.kookapp.cn/doc/http/direct-message#删除私信聊天消息
// POST /direct-message/delete
func DeleteDirectMessage(ctx context.Context, client Doer, msgID string) error {
	return client.Do(ctx, "POST", "/direct-message/delete", map[string]interface{}{
		"msg_id": msgID,
	}, nil)
}

// ListDirectReactions 获取私聊消息回应的用户列表。
// 参考文档：https://developer.kookapp.cn/doc/http/direct-message#获取频道消息某回应的用户列表
// GET /direct-message/reaction-list
//
// 参数说明：
//   - msgID: 消息的 ID（必填）
//   - emoji: 表情的数字标识（必填），可为 GuildEmoji 或 Emoji
//
// KOOK 支持的完整表情对照表请参见：https://img.kookapp.cn/assets/emoji.json
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
// 参考文档：https://developer.kookapp.cn/doc/http/direct-message#给某个消息添加回应
// POST /direct-message/add-reaction
//
// 参数说明：
//   - msgID: 消息的 ID（必填）
//   - emoji: 表情的数字标识（必填），可为 GuildEmoji 或 Emoji
//
// KOOK 支持的完整表情对照表请参见：https://img.kookapp.cn/assets/emoji.json
func AddDirectReaction(ctx context.Context, client Doer, msgID string, emoji string) error {
	return client.Do(ctx, "POST", "/direct-message/add-reaction", map[string]interface{}{
		"msg_id": msgID,
		"emoji":  emoji,
	}, nil)
}

// DeleteDirectReaction 删除私聊消息的回应（表情）。
// 参考文档：https://developer.kookapp.cn/doc/http/direct-message#删除消息的某个回应
// POST /direct-message/delete-reaction
//
// 参数说明：
//   - msgID: 消息的 ID（必填）
//   - emoji: 表情的数字标识（必填），完整对照表参见 https://img.kookapp.cn/assets/emoji.json
//   - userID: 要删除回应的用户 ID，为空则删除当前用户的回应。删除他人的回应需要有管理频道消息的权限
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
