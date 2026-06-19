package kook

import (
	"context"
	"io"

	"github.com/ljwtorch/kook-sdk-go/api"
	"github.com/ljwtorch/kook-sdk-go/emoji"
	"github.com/ljwtorch/kook-sdk-go/model"
)

// --- 用户 ---

// Me 获取当前认证用户（机器人）的信息。
// 等同于 api.GetCurrentUser(ctx, c)。
func (c *Client) Me(ctx context.Context) (*model.User, error) {
	return api.GetCurrentUser(ctx, c)
}

// GetUser 获取目标用户的详细信息。
// 等同于 api.GetUser(ctx, c, userID, guildID)。
func (c *Client) GetUser(ctx context.Context, userID string, guildID string) (*model.User, error) {
	return api.GetUser(ctx, c, userID, guildID)
}

// OnlineBot 上线机器人，恢复接收消息。
func (c *Client) OnlineBot(ctx context.Context) error {
	return api.OnlineBot(ctx, c)
}

// OfflineBot 下线机器人，停止接收消息。
func (c *Client) OfflineBot(ctx context.Context) error {
	return api.OfflineBot(ctx, c)
}

// GetBotOnlineStatus 获取机器人当前在线状态。
func (c *Client) GetBotOnlineStatus(ctx context.Context) (bool, error) {
	return api.GetBotOnlineStatus(ctx, c)
}

// --- 服务器 ---

// GetGuildList 获取当前用户加入的服务器列表。
func (c *Client) GetGuildList(ctx context.Context) (*model.PageResult[model.Guild], error) {
	return api.GetGuildList(ctx, c)
}

// GetGuild 获取指定服务器的详细信息。
func (c *Client) GetGuild(ctx context.Context, guildID string) (*model.Guild, error) {
	return api.GetGuild(ctx, c, guildID)
}

// GetGuildUserList 获取服务器成员列表。
func (c *Client) GetGuildUserList(ctx context.Context, guildID string, page int, pageSize int) (*model.PageResult[model.GuildUser], error) {
	return api.GetGuildUserList(ctx, c, guildID, page, pageSize, "", "", 0, false, 0, 0)
}

// SetGuildNickname 修改用户在服务器中的昵称。
func (c *Client) SetGuildNickname(ctx context.Context, guildID string, nickname string, userID string) error {
	return api.SetGuildNickname(ctx, c, guildID, nickname, userID)
}

// LeaveGuild 离开指定服务器。
func (c *Client) LeaveGuild(ctx context.Context, guildID string) error {
	return api.LeaveGuild(ctx, c, guildID)
}

// KickoutGuildMember 将指定用户踢出服务器。
func (c *Client) KickoutGuildMember(ctx context.Context, guildID string, targetID string) error {
	return api.KickoutGuildMember(ctx, c, guildID, targetID)
}

// --- 频道 ---

// ListChannels 获取指定服务器下的频道列表。
func (c *Client) ListChannels(ctx context.Context, guildID string) (*model.ListResult[model.Channel], error) {
	return api.ListChannels(ctx, c, guildID)
}

// GetChannel 获取指定频道的详细信息。
func (c *Client) GetChannel(ctx context.Context, targetID string) (*model.Channel, error) {
	return api.GetChannel(ctx, c, targetID)
}

// CreateChannel 在指定服务器中创建新频道。
//
// 参考文档：https://developer.kookapp.cn/doc/http/channel#创建频道
func (c *Client) CreateChannel(ctx context.Context, guildID string, name string, typ int, parentID string, limitAmount int, voiceQuality string, isCategory int) (*model.Channel, error) {
	return api.CreateChannel(ctx, c, guildID, name, typ, parentID, limitAmount, voiceQuality, isCategory)
}

// UpdateChannel 编辑指定频道的信息。
// slowMode 参数：慢速模式间隔（毫秒），传 0 表示关闭慢速模式，传 -1 表示不修改
func (c *Client) UpdateChannel(ctx context.Context, channelID string, name string, topic string, slowMode int) (*model.Channel, error) {
	return api.UpdateChannel(ctx, c, channelID, name, topic, slowMode)
}

// DeleteChannel 删除指定频道。
func (c *Client) DeleteChannel(ctx context.Context, channelID string) error {
	return api.DeleteChannel(ctx, c, channelID)
}

// --- 频道消息 ---

// SendMessage 发送频道消息的快捷方法。
// 等同于 api.CreateMessage(ctx, c, targetID, content, "", "", "")。
func (c *Client) SendMessage(ctx context.Context, targetID string, content string) (*api.CreateMessageResponse, error) {
	return api.CreateMessage(ctx, c, targetID, content, "", "", "")
}

// SendReplyMessage 发送带引用的频道消息。
func (c *Client) SendReplyMessage(ctx context.Context, targetID string, content string, quoteMsgID string) (*api.CreateMessageResponse, error) {
	return api.CreateMessage(ctx, c, targetID, content, quoteMsgID, "", "")
}

// SendMessageEx 发送频道消息（完整参数）。
func (c *Client) SendMessageEx(ctx context.Context, targetID string, content string, quote string, nonce string, tempTargetID string) (*api.CreateMessageResponse, error) {
	return api.CreateMessage(ctx, c, targetID, content, quote, nonce, tempTargetID)
}

// UpdateMessage 更新已发送的频道消息。
func (c *Client) UpdateMessage(ctx context.Context, msgID string, content string) (*model.Message, error) {
	return api.UpdateMessage(ctx, c, msgID, content, "", "")
}

// DeleteMessage 删除指定的频道消息。
func (c *Client) DeleteMessage(ctx context.Context, msgID string) error {
	return api.DeleteMessage(ctx, c, msgID)
}

// ListMessages 获取频道消息列表。
func (c *Client) ListMessages(ctx context.Context, targetID string, page int, pageSize int) (*model.PageResult[model.Message], error) {
	return api.ListMessages(ctx, c, targetID, "", false, page, pageSize)
}

// GetMessage 获取频道消息详情。
func (c *Client) GetMessage(ctx context.Context, msgID string) (*model.Message, error) {
	return api.GetMessage(ctx, c, msgID)
}

// AddReaction 为消息添加回应（表情）。
// emoji 参数为表情的数字标识，完整对照表参见 https://img.kookapp.cn/assets/emoji.json
func (c *Client) AddReaction(ctx context.Context, msgID string, emoji string) error {
	return api.AddReaction(ctx, c, msgID, emoji)
}

// DeleteReaction 删除消息的回应。
// emoji 参数为表情的数字标识，完整对照表参见 https://img.kookapp.cn/assets/emoji.json
func (c *Client) DeleteReaction(ctx context.Context, msgID string, emoji string, userID string) error {
	return api.DeleteReaction(ctx, c, msgID, emoji, userID)
}

// AddReactionWithEmoji 为消息添加回应（表情）。
// emojiChar 为 Unicode 字符（如 '😆'），会自动转换为 KOOK emoji 格式。
//
// KOOK 支持的完整表情对照表请参见：
// https://img.kookapp.cn/assets/emoji.json
func (c *Client) AddReactionWithEmoji(ctx context.Context, msgID string, emojiChar rune) error {
	return api.AddReaction(ctx, c, msgID, emoji.ID(emojiChar))
}

// DeleteReactionWithEmoji 删除消息的回应（表情）。
// emojiChar 为 Unicode 字符（如 '😆'），会自动转换为 KOOK emoji 格式。
// userID 为空表示删除自己的回应。
//
// KOOK 支持的完整表情对照表请参见：
// https://img.kookapp.cn/assets/emoji.json
func (c *Client) DeleteReactionWithEmoji(ctx context.Context, msgID string, emojiChar rune, userID string) error {
	return api.DeleteReaction(ctx, c, msgID, emoji.ID(emojiChar), userID)
}

// PinMessage 置顶指定消息。
func (c *Client) PinMessage(ctx context.Context, msgID string, targetID string) error {
	return api.PinMessage(ctx, c, msgID, targetID)
}

// UnpinMessage 取消置顶指定消息。
func (c *Client) UnpinMessage(ctx context.Context, msgID string, targetID string) error {
	return api.UnpinMessage(ctx, c, msgID, targetID)
}

// --- 私聊消息 ---

// SendDirectMessage 发送私聊消息的快捷方法。
// 通过 targetID（用户 ID）发送私聊。
func (c *Client) SendDirectMessage(ctx context.Context, targetID string, content string) (*api.CreateMessageResponse, error) {
	return api.CreateDirectMessage(ctx, c, targetID, "", content, "", "")
}

// SendDirectMessageEx 发送私聊消息（完整参数）。
func (c *Client) SendDirectMessageEx(ctx context.Context, targetID string, chatCode string, content string, quote string, nonce string) (*api.CreateMessageResponse, error) {
	return api.CreateDirectMessage(ctx, c, targetID, chatCode, content, quote, nonce)
}

// UpdateDirectMessage 更新已发送的私聊消息。
func (c *Client) UpdateDirectMessage(ctx context.Context, msgID string, content string) (*model.Message, error) {
	return api.UpdateDirectMessage(ctx, c, msgID, content, "")
}

// DeleteDirectMessage 删除指定的私聊消息。
func (c *Client) DeleteDirectMessage(ctx context.Context, msgID string) error {
	return api.DeleteDirectMessage(ctx, c, msgID)
}

// --- 私聊会话 ---

// ListUserChats 获取私聊会话列表。
func (c *Client) ListUserChats(ctx context.Context, page int, pageSize int) (*model.PageResult[model.UserChat], error) {
	return api.ListUserChats(ctx, c, page, pageSize)
}

// GetUserChat 获取指定私聊会话的详细信息。
func (c *Client) GetUserChat(ctx context.Context, chatCode string) (*model.UserChat, error) {
	return api.GetUserChat(ctx, c, chatCode)
}

// CreateUserChat 创建与指定用户的私聊会话。
func (c *Client) CreateUserChat(ctx context.Context, targetID string) (*model.UserChat, error) {
	return api.CreateUserChat(ctx, c, targetID)
}

// DeleteUserChat 删除指定的私聊会话。
func (c *Client) DeleteUserChat(ctx context.Context, chatCode string) error {
	return api.DeleteUserChat(ctx, c, chatCode)
}

// --- 服务器角色 ---

// ListGuildRoles 获取服务器角色列表。
func (c *Client) ListGuildRoles(ctx context.Context, guildID string, page int, pageSize int) (*model.PageResult[model.Role], error) {
	return api.ListGuildRoles(ctx, c, guildID, page, pageSize)
}

// CreateGuildRole 在服务器中创建新角色。
func (c *Client) CreateGuildRole(ctx context.Context, guildID string, name string) (*model.Role, error) {
	return api.CreateGuildRole(ctx, c, guildID, name)
}

// UpdateGuildRole 更新服务器角色的属性。
//
// 注意：正常角色由上向下排序，这个先后顺序是角色的优先级（position 字段）。
// 如果你有管理员权限，你只能管理优先级比自己低的用户，不能管理优先级等于或比自己高的用户。
//
// 参考文档：https://developer.kookapp.cn/doc/http/guild-role#更新服务器角色
//
// color 参数使用 24 位 RGB 颜色整数，参见 model.Role.Color。
// permissions 参数为权限位掩码，权限常量定义在 model 包中（如 model.PermAdmin）。
func (c *Client) UpdateGuildRole(ctx context.Context, guildID string, roleID int64, name string, color int, hoist bool, mentionable bool, permissions int64) (*model.Role, error) {
	return api.UpdateGuildRole(ctx, c, guildID, roleID, name, color, hoist, mentionable, permissions)
}

// DeleteGuildRole 删除服务器中的指定角色。
func (c *Client) DeleteGuildRole(ctx context.Context, guildID string, roleID int64) error {
	return api.DeleteGuildRole(ctx, c, guildID, roleID)
}

// GrantRole 为服务器中的用户赋予指定角色。
func (c *Client) GrantRole(ctx context.Context, guildID string, userID string, roleID int64) (*model.Role, error) {
	return api.GrantRole(ctx, c, guildID, userID, roleID)
}

// RevokeRole 移除服务器中用户的指定角色。
func (c *Client) RevokeRole(ctx context.Context, guildID string, userID string, roleID int64) error {
	return api.RevokeRole(ctx, c, guildID, userID, roleID)
}

// --- 频道权限 ---

// GetChannelRoles 获取频道角色权限列表。
func (c *Client) GetChannelRoles(ctx context.Context, channelID string) (*model.ListResult[model.ChannelRole], error) {
	return api.GetChannelRoles(ctx, c, channelID)
}

// CreateChannelRole 为频道创建角色权限。
func (c *Client) CreateChannelRole(ctx context.Context, channelID string, typ string, value string) (*model.ChannelRole, error) {
	return api.CreateChannelRole(ctx, c, channelID, typ, value)
}

// UpdateChannelRole 更新频道角色权限设置。
func (c *Client) UpdateChannelRole(ctx context.Context, channelID string, typ string, value string, allow int64, deny int64) (*model.ChannelRole, error) {
	return api.UpdateChannelRole(ctx, c, channelID, typ, value, allow, deny)
}

// SyncChannelRole 将频道角色权限同步到服务器默认设置。
func (c *Client) SyncChannelRole(ctx context.Context, channelID string) error {
	return api.SyncChannelRole(ctx, c, channelID)
}

// DeleteChannelRole 删除频道的角色权限设置。
func (c *Client) DeleteChannelRole(ctx context.Context, channelID string, typ string, value string) error {
	return api.DeleteChannelRole(ctx, c, channelID, typ, value)
}

// --- 邀请 ---

// ListInvites 获取邀请列表。
func (c *Client) ListInvites(ctx context.Context, guildID string, channelID string, page int, pageSize int) (*model.PageResult[api.InviteInfo], error) {
	return api.ListInvites(ctx, c, guildID, channelID, page, pageSize)
}

// CreateInvite 创建邀请链接。
func (c *Client) CreateInvite(ctx context.Context, guildID string, channelID string, duration int, settingTimes int) (string, error) {
	return api.CreateInvite(ctx, c, guildID, channelID, duration, settingTimes)
}

// DeleteInvite 删除邀请链接。
// urlCode 为必填参数（邀请码），opts 为可选参数（如 guildID、channelID）。
func (c *Client) DeleteInvite(ctx context.Context, urlCode string, opts ...api.DeleteInviteOption) error {
	return api.DeleteInvite(ctx, c, urlCode, opts...)
}

// --- 媒体资源 ---

// UploadAsset 上传文件/图片到 KOOK 平台。
func (c *Client) UploadAsset(ctx context.Context, file io.Reader, fileName string) (*api.AssetResult, error) {
	return api.UploadAsset(ctx, c, file, fileName)
}

// --- 亲密度 ---

// GetIntimacy 获取指定用户的亲密度信息。
func (c *Client) GetIntimacy(ctx context.Context, userID string) (*api.IntimacyInfo, error) {
	return api.GetIntimacy(ctx, c, userID)
}

// UpdateIntimacy 更新指定用户的亲密度信息。
func (c *Client) UpdateIntimacy(ctx context.Context, userID string, score int, socialInfo string, imgID int) error {
	return api.UpdateIntimacy(ctx, c, userID, score, socialInfo, imgID)
}

// --- 黑名单 ---

// ListBlacklist 获取服务器黑名单列表。
func (c *Client) ListBlacklist(ctx context.Context, guildID string, page int, pageSize int) (*model.PageResult[api.BlacklistUser], error) {
	return api.ListBlacklist(ctx, c, guildID, page, pageSize)
}

// AddBlacklist 将用户加入服务器黑名单。
func (c *Client) AddBlacklist(ctx context.Context, guildID string, targetID string, remark string, delMsgDays int) error {
	return api.AddBlacklist(ctx, c, guildID, targetID, remark, delMsgDays)
}

// RemoveBlacklist 将用户从服务器黑名单中移除。
func (c *Client) RemoveBlacklist(ctx context.Context, guildID string, targetID string) error {
	return api.RemoveBlacklist(ctx, c, guildID, targetID)
}

// --- 服务器静音 ---

// ListGuildMutes 获取服务器静音/闭麦用户列表。
func (c *Client) ListGuildMutes(ctx context.Context, guildID string, returnType string, page int, pageSize int) (*model.PageResult[api.GuildMuteUser], error) {
	return api.ListGuildMutes(ctx, c, guildID, returnType, page, pageSize)
}

// CreateGuildMute 对服务器中的用户添加静音或闭麦。
func (c *Client) CreateGuildMute(ctx context.Context, guildID string, userID string, muteType string) error {
	return api.CreateGuildMute(ctx, c, guildID, userID, muteType)
}

// DeleteGuildMute 取消服务器中用户的静音或闭麦状态。
func (c *Client) DeleteGuildMute(ctx context.Context, guildID string, userID string, muteType string) error {
	return api.DeleteGuildMute(ctx, c, guildID, userID, muteType)
}

// --- 服务器助力 ---

// GetBoostHistory 获取服务器助力历史记录。
func (c *Client) GetBoostHistory(ctx context.Context, guildID string, startTime int64, endTime int64) (*model.ListResult[api.BoostRecord], error) {
	return api.GetBoostHistory(ctx, c, guildID, startTime, endTime)
}
