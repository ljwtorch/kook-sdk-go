package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// ListUserChats 获取私聊会话列表，支持分页。
// GET /user-chat/list
//
// 参数说明：
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
func ListUserChats(ctx context.Context, client Doer, page int, pageSize int) (*model.PageResult[model.UserChat], error) {
	params := map[string]interface{}{}
	if page > 0 {
		params["page"] = page
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}

	var result model.PageResult[model.UserChat]
	err := client.Do(ctx, "GET", "/user-chat/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserChat 获取指定私聊会话的详细信息。
// GET /user-chat/view?chat_code={chatCode}
func GetUserChat(ctx context.Context, client Doer, chatCode string) (*model.UserChat, error) {
	var result model.UserChat
	err := client.Do(ctx, "GET", "/user-chat/view", map[string]interface{}{
		"chat_code": chatCode,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateUserChat 创建与指定用户的私聊会话。
// POST /user-chat/create
func CreateUserChat(ctx context.Context, client Doer, targetID string) (*model.UserChat, error) {
	var result model.UserChat
	err := client.Do(ctx, "POST", "/user-chat/create", map[string]interface{}{
		"target_id": targetID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteUserChat 删除指定的私聊会话。
// POST /user-chat/delete
func DeleteUserChat(ctx context.Context, client Doer, chatCode string) error {
	return client.Do(ctx, "POST", "/user-chat/delete", map[string]interface{}{
		"chat_code": chatCode,
	}, nil)
}
