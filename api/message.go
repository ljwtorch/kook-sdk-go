package api

import (
	"context"
	"net/url"
	"strconv"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// CreateMessageResponse 表示发送消息接口的响应。
// KOOK API 发送消息（频道/私聊）返回的是 msg_id 而非完整消息对象。
type CreateMessageResponse struct {
	// MsgID 是服务端生成的消息 ID
	MsgID string `json:"msg_id"`
	// MsgTimestamp 是消息发送时间（服务器时间戳，毫秒）
	MsgTimestamp int64 `json:"msg_timestamp"`
	// Nonce 是客户端提供的去重标识，服务端原样返回
	Nonce string `json:"nonce"`
}

// ListMessages 获取频道消息列表，基于参考消息进行游标分页。
// 参考文档：https://developer.kookapp.cn/doc/http/message#获取频道聊天消息列表
// GET /message/list
//
// 参数说明：
//   - targetID: 频道 ID（必填）
//   - msgID: 参考消息 ID，用于翻页定位，不传则查询最新消息
//   - pin: 是否只获取置顶消息，0 或 1
//   - flag: 查询模式，"before"（参考消息之前）、"around"（参考消息前后）、"after"（参考消息之后），不传则查询最新消息
//   - pageSize: 每页数量，传 0 使用默认值（50）
func ListMessages(ctx context.Context, client Doer, targetID string, msgID string, pin int, flag string, pageSize int) (*model.ListResult[model.Message], error) {
	params := map[string]interface{}{
		"target_id": targetID,
	}
	if msgID != "" {
		params["msg_id"] = msgID
	}
	if pin != 0 {
		params["pin"] = pin
	}
	if flag != "" {
		params["flag"] = flag
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}

	var result model.ListResult[model.Message]
	err := client.Do(ctx, "GET", "/message/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMessage 获取指定频道消息的详细信息。
// 参考文档：https://developer.kookapp.cn/doc/http/message#获取频道聊天消息详情
// GET /message/view?msg_id={msgID}
func GetMessage(ctx context.Context, client Doer, msgID string) (*model.Message, error) {
	var result model.Message
	err := client.Do(ctx, "GET", "/message/view", map[string]interface{}{
		"msg_id": msgID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateMessage 发送频道消息。
// 参考文档：https://developer.kookapp.cn/doc/http/message#发送频道聊天消息
// POST /message/create
//
// content 支持纯文本、KMarkdown（type=9）、卡片 JSON（type=10）。
// 当 msgType 为 0 时，自动检测：content 以 "[" 开头视为卡片消息（type=10），否则默认 KMarkdown（type=9）。
//
// 参数说明：
//   - targetID: 目标频道 ID（必填）
//   - content: 消息内容（必填）
//   - msgType: 消息类型，0 表示自动检测，9 表示 KMarkdown，10 表示卡片消息
//   - quote: 引用消息 ID，为空表示不引用
//   - nonce: 去重标识，为空则不启用去重
//   - tempTargetID: 临时消息目标用户 ID，为空则发送为普通消息
//   - replyMsgID: 当前消息回复的用户 5 分钟内发送到相同频道的消息 ID，用于配额折扣
//   - templateID: 模板消息 ID，使用后 content 作为模板消息的 input
func CreateMessage(ctx context.Context, client Doer, targetID string, content string, msgType int, quote string, nonce string, tempTargetID string, replyMsgID string, templateID string) (*CreateMessageResponse, error) {
	body := map[string]interface{}{
		"target_id": targetID,
		"content":   content,
	}
	if msgType != 0 {
		body["type"] = msgType
	} else if len(content) > 0 && content[0] == '[' {
		body["type"] = 10
	} else {
		body["type"] = 9
	}
	if quote != "" {
		body["quote"] = quote
	}
	if nonce != "" {
		body["nonce"] = nonce
	}
	if tempTargetID != "" {
		body["temp_target_id"] = tempTargetID
	}
	if replyMsgID != "" {
		body["reply_msg_id"] = replyMsgID
	}
	if templateID != "" {
		body["template_id"] = templateID
	}

	var result CreateMessageResponse
	err := client.Do(ctx, "POST", "/message/create", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateMessage 更新已发送的频道消息。
// 参考文档：https://developer.kookapp.cn/doc/http/message#更新频道聊天消息
// POST /message/update
//
// 目前支持消息 type 为 9、10 的修改，即 KMarkdown 和 CardMessage。
//
// 参数说明：
//   - msgID: 消息 ID（必填）
//   - content: 新的消息内容（必填）
//   - quote: 引用消息 ID，为空表示删除引用，不传则无影响
//   - tempTargetID: 临时消息目标用户 ID，为空表示不修改
//   - replyMsgID: 当前消息回复的用户 5 分钟内发送到相同频道的消息 ID，用于配额折扣
//   - templateID: 模板消息 ID，使用后 content 作为模板消息的 input
func UpdateMessage(ctx context.Context, client Doer, msgID string, content string, quote string, tempTargetID string, replyMsgID string, templateID string) error {
	body := map[string]interface{}{
		"msg_id":  msgID,
		"content": content,
	}
	if quote != "" {
		body["quote"] = quote
	}
	if tempTargetID != "" {
		body["temp_target_id"] = tempTargetID
	}
	if replyMsgID != "" {
		body["reply_msg_id"] = replyMsgID
	}
	if templateID != "" {
		body["template_id"] = templateID
	}

	return client.Do(ctx, "POST", "/message/update", body, nil)
}

// DeleteMessage 删除指定的频道消息。
// 参考文档：https://developer.kookapp.cn/doc/http/message#删除频道聊天消息
// POST /message/delete
func DeleteMessage(ctx context.Context, client Doer, msgID string) error {
	return client.Do(ctx, "POST", "/message/delete", map[string]interface{}{
		"msg_id": msgID,
	}, nil)
}

// ListReactions 获取消息回应的用户列表。
// 参考文档：https://developer.kookapp.cn/doc/http/message#获取频道消息某个回应的用户列表
// GET /message/reaction-list
//
// 参数说明：
//   - msgID: 频道消息的 ID（必填）
//   - emoji: 表情的数字标识（必填），可为 GuildEmoji 或 Emoji
//
// KOOK 支持的完整表情对照表请参见：https://img.kookapp.cn/assets/emoji.json
func ListReactions(ctx context.Context, client Doer, msgID string, emoji string) (*model.ListResult[model.User], error) {
	var result model.ListResult[model.User]
	err := client.Do(ctx, "GET", "/message/reaction-list", map[string]interface{}{
		"msg_id": msgID,
		"emoji":  emoji,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AddReaction 为消息添加回应（表情）。
// 参考文档：https://developer.kookapp.cn/doc/http/message#给某个消息添加回应
// POST /message/add-reaction
//
// 参数说明：
//   - msgID: 频道消息的 ID（必填）
//   - emoji: 表情的数字标识（必填），可为 GuildEmoji 或 Emoji
//
// KOOK 支持的完整表情对照表请参见：https://img.kookapp.cn/assets/emoji.json
func AddReaction(ctx context.Context, client Doer, msgID string, emoji string) error {
	return client.Do(ctx, "POST", "/message/add-reaction", map[string]interface{}{
		"msg_id": msgID,
		"emoji":  emoji,
	}, nil)
}

// DeleteReaction 删除消息的回应（表情）。
// 参考文档：https://developer.kookapp.cn/doc/http/message#删除消息的某个回应
// POST /message/delete-reaction
//
// 参数说明：
//   - msgID: 频道消息的 ID（必填）
//   - emoji: 表情的数字标识（必填），完整对照表参见 https://img.kookapp.cn/assets/emoji.json
//   - userID: 要删除回应的用户 ID，为空则删除当前用户的回应。删除他人的回应需要有管理频道消息的权限
func DeleteReaction(ctx context.Context, client Doer, msgID string, emoji string, userID string) error {
	body := map[string]interface{}{
		"msg_id": msgID,
		"emoji":  emoji,
	}
	if userID != "" {
		body["user_id"] = userID
	}
	return client.Do(ctx, "POST", "/message/delete-reaction", body, nil)
}

// SendPipeMessage 发送管道消息，需要在开发者后台先创建管道。
// 参考文档：https://developer.kookapp.cn/doc/http/message#发送管道消息
// POST /message/send-pipemsg
//
// 参数说明：
//   - targetID: 频道 ID，不填则以消息管道的设置为准
//   - msgType: 消息类型，不填则以模板为准，无模板则为 KMarkdown
//   - content: 消息内容（必填），格式参见 CreateMessage
func SendPipeMessage(ctx context.Context, client Doer, targetID string, msgType int, content string) error {
	// target_id 和 type 作为 GET 查询参数，content 作为 POST body
	path := "/message/send-pipemsg"
	query := url.Values{}
	if targetID != "" {
		query.Set("target_id", targetID)
	}
	if msgType != 0 {
		query.Set("type", strconv.Itoa(msgType))
	}
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	body := map[string]interface{}{
		"content": content,
	}

	return client.Do(ctx, "POST", path, body, nil)
}

// PinMessage 置顶指定消息，需要管理消息权限。
// 参考文档：https://developer.kookapp.cn/doc/http/message#置顶频道消息
// POST /message/pin
//
// 参数说明：
//   - msgID: 消息 ID（必填）
//   - targetID: 频道 ID（必填）
func PinMessage(ctx context.Context, client Doer, msgID string, targetID string) error {
	return client.Do(ctx, "POST", "/message/pin", map[string]interface{}{
		"msg_id":    msgID,
		"target_id": targetID,
	}, nil)
}

// UnpinMessage 取消置顶指定消息，需要管理消息权限。
// 参考文档：https://developer.kookapp.cn/doc/http/message#取消置顶频道消息
// POST /message/unpin
//
// 参数说明：
//   - msgID: 消息 ID（必填）
//   - targetID: 频道 ID（必填）
func UnpinMessage(ctx context.Context, client Doer, msgID string, targetID string) error {
	return client.Do(ctx, "POST", "/message/unpin", map[string]interface{}{
		"msg_id":    msgID,
		"target_id": targetID,
	}, nil)
}
