package api

import (
	"context"

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

// ListMessages 获取频道消息列表，支持分页和消息筛选。
// GET /message/list
//
// 参数说明：
//   - targetID: 频道 ID（必填）
//   - msgID: 参考消息 ID，用于翻页定位
//   - pin: 是否只获取置顶消息
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
func ListMessages(ctx context.Context, client Doer, targetID string, msgID string, pin bool, page int, pageSize int) (*model.PageResult[model.Message], error) {
	params := map[string]interface{}{
		"target_id": targetID,
	}
	if msgID != "" {
		params["msg_id"] = msgID
	}
	if pin {
		params["pin"] = true
	}
	if page > 0 {
		params["page"] = page
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}

	var result model.PageResult[model.Message]
	err := client.Do(ctx, "GET", "/message/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMessage 获取指定频道消息的详细信息。
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
// POST /message/create
//
// content 支持纯文本、KMarkdown（type=9）、卡片 JSON（type=10）。
// 当 content 以 "[" 开头时，自动识别为卡片消息并设置 type=10。
//
// 参数说明：
//   - targetID: 目标频道 ID（必填）
//   - content: 消息内容（必填）
//   - quote: 引用消息 ID，为空表示不引用
//   - nonce: 去重标识，为空则不启用去重
//   - tempTargetID: 临时消息目标用户 ID，为空则发送为普通消息
func CreateMessage(ctx context.Context, client Doer, targetID string, content string, quote string, nonce string, tempTargetID string) (*CreateMessageResponse, error) {
	body := map[string]interface{}{
		"target_id": targetID,
		"content":   content,
	}
	// 自动检测卡片消息：以 "[" 开头的 JSON 数组视为卡片消息
	if len(content) > 0 && content[0] == '[' {
		body["type"] = 10
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

	var result CreateMessageResponse
	err := client.Do(ctx, "POST", "/message/create", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateMessage 更新已发送的频道消息。
// POST /message/update
//
// 参数说明：
//   - msgID: 消息 ID（必填）
//   - content: 新的消息内容（必填）
//   - quote: 引用消息 ID，为空表示不修改
//   - tempTargetID: 临时消息目标用户 ID，为空表示不修改
func UpdateMessage(ctx context.Context, client Doer, msgID string, content string, quote string, tempTargetID string) (*model.Message, error) {
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

	var result model.Message
	err := client.Do(ctx, "POST", "/message/update", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteMessage 删除指定的频道消息。
// POST /message/delete
func DeleteMessage(ctx context.Context, client Doer, msgID string) error {
	return client.Do(ctx, "POST", "/message/delete", map[string]interface{}{
		"msg_id": msgID,
	}, nil)
}

// ListReactions 获取消息回应的用户列表。
// GET /message/reaction-list?msg_id={msgID}&emoji={emoji}
//
// emoji 参数为表情的数字标识，完整对照表参见 https://img.kookapp.cn/assets/emoji.json
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
// POST /message/add-reaction
//
// emoji 参数为表情的数字标识，完整对照表参见 https://img.kookapp.cn/assets/emoji.json
func AddReaction(ctx context.Context, client Doer, msgID string, emoji string) error {
	return client.Do(ctx, "POST", "/message/add-reaction", map[string]interface{}{
		"msg_id": msgID,
		"emoji":  emoji,
	}, nil)
}

// DeleteReaction 删除消息的回应（表情）。
// POST /message/delete-reaction
//
// 参数说明：
//   - msgID: 消息 ID（必填）
//   - emoji: 表情的数字标识（必填），完整对照表参见 https://img.kookapp.cn/assets/emoji.json
//   - userID: 要删除回应的用户 ID，为空则删除当前用户的回应
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

// SendPipeMessage 发送管道消息，用于机器人之间的通信。
// POST /message/send-pipemsg
func SendPipeMessage(ctx context.Context, client Doer, targetID string, content string) error {
	return client.Do(ctx, "POST", "/message/send-pipemsg", map[string]interface{}{
		"target_id": targetID,
		"content":   content,
	}, nil)
}

// PinMessage 置顶指定消息。
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

// UnpinMessage 取消置顶指定消息。
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
