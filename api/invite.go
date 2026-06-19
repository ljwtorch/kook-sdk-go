package api

import (
	"context"

	"github.com/ljwtorch/kook-sdk-go/model"
)

// InviteInfo 表示邀请链接的详细信息。
type InviteInfo struct {
	URL       string      `json:"url"`
	URLCode   string      `json:"url_code"`
	GuildID   string      `json:"guild_id"`
	ChannelID string      `json:"channel_id"`
	Status    int         `json:"status"`
	ExpireAt  int64       `json:"expire_at"`
	Creator   *model.User `json:"creator"`
}

// InviteeInfo 表示被邀请用户的信息。
type InviteeInfo struct {
	// Status 表示用户状态：0=未退出，254=已退出
	Status int `json:"status"`
	// JoinedTime 是用户加入服务器的时间戳（毫秒）
	JoinedTime int64 `json:"joined_time"`
	// ActiveTime 是用户最后活跃时间戳（毫秒）
	ActiveTime int64 `json:"active_time"`
	// ShowName 是用户的显示名称
	ShowName string `json:"show_name"`
}

// ListInviteesResult 表示被邀请用户列表的响应结果。
type ListInviteesResult struct {
	// Items 是被邀请用户列表
	Items []InviteeInfo `json:"items"`
	// Meta 是分页信息
	Meta model.PageMeta `json:"meta"`
	// Count 是总用户数
	Count int `json:"count"`
	// KeepCount 是还在服务器内的用户数
	KeepCount int `json:"keep_count"`
	// LossCount 是已经离开的用户数
	LossCount int `json:"loss_count"`
}

// inviteURLResult 用于解析创建邀请链接接口响应。
type inviteURLResult struct {
	URL string `json:"url"`
}

// ListInvites 获取服务器或频道的邀请列表，支持分页。
// GET /invite/list
//
// 参数说明：
//   - guildID: 服务器 ID
//   - channelID: 频道 ID
//   - page: 页码，传 0 使用默认值
//   - pageSize: 每页数量，传 0 使用默认值
func ListInvites(ctx context.Context, client Doer, guildID string, channelID string, page int, pageSize int) (*model.PageResult[InviteInfo], error) {
	params := map[string]interface{}{}
	if guildID != "" {
		params["guild_id"] = guildID
	}
	if channelID != "" {
		params["channel_id"] = channelID
	}
	if page > 0 {
		params["page"] = page
	}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}

	var result model.PageResult[InviteInfo]
	err := client.Do(ctx, "GET", "/invite/list", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateInvite 创建服务器或频道的邀请链接。
// POST /invite/create
//
// 参数说明：
//   - guildID: 服务器 ID
//   - channelID: 频道 ID
//   - duration: 邀请有效期（秒），传 0 表示永久
//   - settingTimes: 邀请使用次数限制，传 0 表示不限
//
// 返回创建的邀请链接 URL。
func CreateInvite(ctx context.Context, client Doer, guildID string, channelID string, duration int, settingTimes int) (string, error) {
	body := map[string]interface{}{}
	if guildID != "" {
		body["guild_id"] = guildID
	}
	if channelID != "" {
		body["channel_id"] = channelID
	}
	if duration > 0 {
		body["duration"] = duration
	}
	if settingTimes > 0 {
		body["setting_times"] = settingTimes
	}

	var result inviteURLResult
	err := client.Do(ctx, "POST", "/invite/create", body, &result)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}

// DeleteInviteOption 配置删除邀请链接的可选参数。
type DeleteInviteOption func(*deleteInviteOptions)

type deleteInviteOptions struct {
	guildID   string
	channelID string
}

// WithDeleteInviteGuildID 设置删除邀请链接时的服务器 ID。
func WithDeleteInviteGuildID(guildID string) DeleteInviteOption {
	return func(o *deleteInviteOptions) { o.guildID = guildID }
}

// WithDeleteInviteChannelID 设置删除邀请链接时的频道 ID。
func WithDeleteInviteChannelID(channelID string) DeleteInviteOption {
	return func(o *deleteInviteOptions) { o.channelID = channelID }
}

// DeleteInvite 删除指定的邀请链接。
// POST /invite/delete
//
// 参数说明：
//   - urlCode: 邀请码（必填）
//   - opts: 可选参数，包括 guildID（服务器 ID）和 channelID（频道 ID）
func DeleteInvite(ctx context.Context, client Doer, urlCode string, opts ...DeleteInviteOption) error {
	o := &deleteInviteOptions{}
	for _, opt := range opts {
		opt(o)
	}
	body := map[string]interface{}{
		"url_code": urlCode,
	}
	if o.guildID != "" {
		body["guild_id"] = o.guildID
	}
	if o.channelID != "" {
		body["channel_id"] = o.channelID
	}
	return client.Do(ctx, "POST", "/invite/delete", body, nil)
}

// ListInviteesOption 配置获取被邀请用户列表的可选参数。
type ListInviteesOption func(*listInviteesOptions)

type listInviteesOptions struct {
	id        string
	inviteURL string
	status    int
	startTime string
	endTime   string
}

// WithInviteesID 设置查询的邀请码。
func WithInviteesID(id string) ListInviteesOption {
	return func(o *listInviteesOptions) { o.id = id }
}

// WithInviteesURL 设置查询的邀请码链接。
func WithInviteesURL(url string) ListInviteesOption {
	return func(o *listInviteesOptions) { o.inviteURL = url }
}

// WithInviteesStatus 设置用户状态筛选：0=未退出，254=已退出，-1=全部。
func WithInviteesStatus(status int) ListInviteesOption {
	return func(o *listInviteesOptions) { o.status = status }
}

// WithInviteesStartTime 设置加入的开始时间，格式如 "2026-06-01 12:00:00"。
func WithInviteesStartTime(t string) ListInviteesOption {
	return func(o *listInviteesOptions) { o.startTime = t }
}

// WithInviteesEndTime 设置加入的结束时间，格式如 "2026-06-02 12:00:00"。
func WithInviteesEndTime(t string) ListInviteesOption {
	return func(o *listInviteesOptions) { o.endTime = t }
}

// ListInvitees 获取被邀请用户列表，支持分页。
// GET /invite/invitees
//
// 参数说明：
//   - guildID: 服务器 ID（必填）
//   - page: 页码（必填）
//   - pageSize: 每页数量（必填）
//   - opts: 可选参数，包括 id（邀请码）、inviteURL（邀请链接）、status（状态筛选）、startTime、endTime
func ListInvitees(ctx context.Context, client Doer, guildID string, page int, pageSize int, opts ...ListInviteesOption) (*ListInviteesResult, error) {
	o := &listInviteesOptions{}
	for _, opt := range opts {
		opt(o)
	}

	params := map[string]interface{}{
		"page":      page,
		"page_size": pageSize,
	}
	if guildID != "" {
		params["guild_id"] = guildID
	}
	if o.id != "" {
		params["id"] = o.id
	}
	if o.inviteURL != "" {
		params["invite_url"] = o.inviteURL
	}
	if o.status != 0 {
		params["status"] = o.status
	}
	if o.startTime != "" {
		params["start_time"] = o.startTime
	}
	if o.endTime != "" {
		params["end_time"] = o.endTime
	}

	var result ListInviteesResult
	err := client.Do(ctx, "GET", "/invite/invitees", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
