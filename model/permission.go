// Package model 提供 KOOK 平台的数据模型定义。
//
// # 权限说明
//
// 权限是一个 unsigned int 值，由比特位代表是否拥有对应的权限。
// 权限值与对应比特位进行按位与操作，判断是否拥有该权限。
//
//	permissions & (1 << bitValue) == (1 << bitValue)
//
// 注意：正常角色由上向下排序，这个先后顺序是角色的优先级（position 字段）。
// 如果你有管理员权限，你只能管理优先级比自己低的用户，不能管理优先级等于或比自己高的用户。
// 这个地方的逻辑举例来说是这样的：对于一个公司的 HR 来说，他是有招员工的权利也有开除员工的权利
// （类比于管理权限），但是他不能开掉老板，也不是招自己的 boss。
// 因此，在使用授予权限、更新等接口时，要注意一下，可能机器人虽然有管理权限，
// 但是也不是什么角色都可以授予，也不是什么人都可以操作。
//
// 参考文档：https://developer.kookapp.cn/doc/http/guild-role#权限说明
package model

// KOOK 服务器权限常量。
// 权限值为比特位掩码，可通过按位或（|）组合多个权限。
//
// 使用示例：
//
//	// 组合多个权限
//	perms := model.PermSendMessages | model.PermAddReactions
//
//	// 使用辅助函数组合
//	perms := model.CombinePermissions(model.PermSendMessages, model.PermAddReactions)
//
//	// 检查是否拥有某权限
//	if model.HasPermission(role.Permissions, model.PermAdmin) {
//	    // 拥有管理员权限
//	}
//
// 参考文档：https://developer.kookapp.cn/doc/http/guild-role#权限说明
const (
	// PermAdmin 管理员。
	// 拥有此权限会获得完整的管理权，包括绕开所有其他权限（包括频道权限）限制，属于危险权限。
	PermAdmin int64 = 1 << 0

	// PermManageGuild 管理服务器。
	// 拥有此权限的成员可以修改服务器名称和更换区域。
	PermManageGuild int64 = 1 << 1

	// PermViewAuditLog 查看管理日志。
	// 拥有此权限的成员可以查看服务器的管理日志。
	PermViewAuditLog int64 = 1 << 2

	// PermCreateInvite 创建服务器邀请。
	// 能否创建服务器邀请链接。
	PermCreateInvite int64 = 1 << 3

	// PermManageInvite 管理邀请。
	// 拥有该权限可以管理服务器的邀请。
	PermManageInvite int64 = 1 << 4

	// PermManageChannels 频道管理。
	// 拥有此权限的成员可以创建新的频道以及编辑或删除已存在的频道。
	PermManageChannels int64 = 1 << 5

	// PermKickMembers 踢出用户。
	PermKickMembers int64 = 1 << 6

	// PermBanMembers 封禁用户。
	PermBanMembers int64 = 1 << 7

	// PermManageEmoji 管理自定义表情。
	PermManageEmoji int64 = 1 << 8

	// PermChangeNickname 修改服务器昵称。
	// 拥有此权限的用户可以更改他们的昵称。
	PermChangeNickname int64 = 1 << 9

	// PermManageRoles 管理角色权限。
	// 拥有此权限成员可以创建新的角色和编辑删除低于该角色的身份。
	PermManageRoles int64 = 1 << 10

	// PermViewChannels 查看文字、语音频道。
	PermViewChannels int64 = 1 << 11

	// PermSendMessages 发布消息。
	PermSendMessages int64 = 1 << 12

	// PermManageMessages 管理消息。
	// 拥有此权限的成员可以删除其他成员发出的消息和置顶消息。
	PermManageMessages int64 = 1 << 13

	// PermUploadFiles 上传文件。
	PermUploadFiles int64 = 1 << 14

	// PermVoiceConnect 语音链接。
	PermVoiceConnect int64 = 1 << 15

	// PermVoiceManage 语音管理。
	// 拥有此权限的成员可以把其他成员移动和踢出频道；
	// 但此类移动仅限于在该成员和被移动成员均有权限的频道之间进行。
	PermVoiceManage int64 = 1 << 16

	// PermMentionEveryone 提及@全体成员。
	// 拥有此权限的成员可使用@全体成员以提及该频道中所有成员。
	PermMentionEveryone int64 = 1 << 17

	// PermAddReactions 添加反应。
	// 拥有此权限的成员可以对消息添加新的反应。
	PermAddReactions int64 = 1 << 18

	// PermFollowReactions 跟随添加反应。
	// 拥有此权限的成员可以跟随使用已经添加的反应。
	PermFollowReactions int64 = 1 << 19

	// PermPassiveVoice 被动连接语音频道。
	// 拥有此限制的成员无法主动连接语音频道，只能在被动邀请或被人移动时，才可以进入语音频道。
	PermPassiveVoice int64 = 1 << 20

	// PermPTTOnly 仅使用按键说话。
	// 拥有此限制的成员加入语音频道后，只能使用按键说话。
	PermPTTOnly int64 = 1 << 21

	// PermFreeMic 使用自由麦。
	// 没有此权限的成员，必须在频道内使用按键说话。
	PermFreeMic int64 = 1 << 22

	// PermSpeak 说话。
	PermSpeak int64 = 1 << 23

	// PermServerMute 服务器静音。
	PermServerMute int64 = 1 << 24

	// PermServerDeafen 服务器闭麦。
	PermServerDeafen int64 = 1 << 25

	// PermChangeOthersNick 修改他人昵称。
	// 拥有此权限的用户可以更改他人的昵称。
	PermChangeOthersNick int64 = 1 << 26

	// PermPlayMusic 播放伴奏。
	// 拥有此权限的成员可在语音频道中播放音乐伴奏。
	PermPlayMusic int64 = 1 << 27

	// PermScreenShare 屏幕分享。
	// 拥有此权限的成员可在频道中向别人分享自己的屏幕。
	PermScreenShare int64 = 1 << 28

	// PermReplyPost 回复帖子。
	// 拥有此权限的成员可以在此贴子频道回复帖子。
	PermReplyPost int64 = 1 << 29

	// PermStartRecording 开启录音。
	// 拥有此权限的成员可在频道中开启录音。
	PermStartRecording int64 = 1 << 30
)

// CombinePermissions 组合多个权限值为一个权限掩码。
//
// 示例：
//
//	// 组合发布消息和添加反应权限
//	perms := model.CombinePermissions(model.PermSendMessages, model.PermAddReactions)
func CombinePermissions(perms ...int64) int64 {
	var result int64
	for _, p := range perms {
		result |= p
	}
	return result
}

// HasPermission 检查权限掩码中是否包含指定权限。
//
// 示例：
//
//	if model.HasPermission(role.Permissions, model.PermAdmin) {
//	    fmt.Println("拥有管理员权限")
//	}
func HasPermission(permissions int64, perm int64) bool {
	return permissions&perm == perm
}
