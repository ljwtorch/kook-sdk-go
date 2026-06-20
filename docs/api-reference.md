# API 参考

本文档介绍 KOOK SDK for Go 提供的所有 HTTP API 模块。

## 目录

- [用户 API](#用户-api)
- [服务器 API](#服务器-api)
- [频道 API](#频道-api)
- [消息 API](#消息-api)
- [私聊 API](#私聊-api)
- [角色 API](#角色-api)
- [频道权限 API](#频道权限-api)
- [邀请 API](#邀请-api)
- [亲密度 API](#亲密度-api)
- [黑名单 API](#黑名单-api)
- [静音 API](#静音-api)
- [助力 API](#助力-api)
- [媒体资源 API](#媒体资源-api)

---

## 用户 API

获取和管理用户信息。

### 获取当前用户

```go
me, err := client.Me(ctx)
// 或
me, err := api.GetCurrentUser(ctx, client)
```

**返回：** `*model.User`

**字段说明：**
- `ID` - 用户 ID
- `Username` - 用户名
- `IdententifyNum` - 用户识别号
- `Online` - 是否在线
- `Status` - 状态（0=正常，1=被封禁）
- `Avatar` - 头像 URL
- `Bot` - 是否为机器人

### 获取目标用户

```go
user, err := client.GetUser(ctx, "user-id", "guild-id")
// 或
user, err := api.GetUser(ctx, client, "user-id", "guild-id")
```

**参数：**
- `userID` - 目标用户 ID
- `guildID` - 服务器 ID（可选，用于获取用户在该服务器的信息）

### 上线/下线机器人

```go
// 上线
err := client.OnlineBot(ctx)

// 下线
err := client.OfflineBot(ctx)
```

### 获取机器人在线状态

```go
status, err := client.GetBotOnlineStatus(ctx)
// 返回 *model.BotOnlineStatus
if status.Online {
    fmt.Println("机器人在线")
    fmt.Printf("在线平台: %v\n", status.OnlineOS)
}
```

---

## 服务器 API

管理服务器（Guild）相关操作。

### 获取服务器列表

```go
guilds, err := client.GetGuildList(ctx)
// 或
guilds, err := api.GetGuildList(ctx, client)
```

**返回：** `*model.PageResult[model.Guild]`

**遍历：**
```go
for _, guild := range guilds.Items {
    fmt.Printf("服务器: %s (ID: %s)\n", guild.Name, guild.ID)
}
```

### 获取服务器详情

```go
guild, err := client.GetGuild(ctx, "guild-id")
// 或
guild, err := api.GetGuild(ctx, client, "guild-id")
```

### 获取服务器成员列表

```go
members, err := client.GetGuildUserList(ctx, "guild-id", 1, 20)
// 或使用完整参数
members, err := api.GetGuildUserList(ctx, client, "guild-id", page, pageSize,
    "channel-id", "search-keyword", 0, false, 7, 0)
```

**参数说明：**
- `page` - 页码（从 1 开始）
- `pageSize` - 每页数量
- `channelID` - 频道 ID（可选，筛选在该频道的用户）
- `search` - 搜索关键词（可选）
- `roleID` - 角色 ID（可选，筛选拥有该角色的用户）
- `joinedAt` - 加入时间筛选（可选）
- `offset` - 偏移量（可选）
- `filterUserIDs` - 用户 ID 列表（可选）

### 修改服务器昵称

```go
err := client.SetGuildNickname(ctx, "guild-id", "新昵称", "")
```

**参数：**
- `guildID` - 服务器 ID
- `nickname` - 新昵称（空字符串表示删除昵称）
- `userID` - 用户 ID（空字符串表示修改自己）

### 离开服务器

```go
err := client.LeaveGuild(ctx, "guild-id")
```

### 踢出服务器成员

```go
err := client.KickoutGuildMember(ctx, "guild-id", "user-id")
```

---

## 频道 API

管理频道（Channel）相关操作。

### 获取频道列表

```go
channels, err := client.ListChannels(ctx, "guild-id")
// 或
channels, err := api.ListChannels(ctx, client, "guild-id")
```

### 获取频道详情

```go
channel, err := client.GetChannel(ctx, "channel-id")
// 或
channel, err := api.GetChannel(ctx, client, "channel-id")
```

### 创建频道

```go
channel, err := client.CreateChannel(ctx, "guild-id", "频道名称", 1, "", 0, "", 0)
```

**参数：**
- `guildID` - 服务器 ID
- `name` - 频道名称
- `typ` - 频道类型（1=文字，2=语音）
- `parentID` - 父分类 ID（可选）
- `limitAmount` - 语音频道人数限制（0=不限制）
- `voiceQuality` - 语音质量（可选）
- `isCategory` - 是否为分类（0=否，1=是）

### 编辑频道

```go
channel, err := client.UpdateChannel(ctx, "channel-id", "新名称", "新主题", 0)
```

**参数：**
- `channelID` - 频道 ID
- `name` - 新名称（空字符串表示不修改）
- `topic` - 新主题（空字符串表示不修改）
- `slowMode` - 慢速模式（0=关闭，-1=不修改，其他值=秒数）

### 删除频道

```go
err := client.DeleteChannel(ctx, "channel-id")
```

---

## 消息 API

管理频道消息相关操作。

### 获取消息列表

```go
messages, err := client.ListMessages(ctx, "channel-id", 1, 20)
// 或
messages, err := api.ListMessages(ctx, client, "channel-id", "", false, 1, 20)
```

**参数：**
- `targetID` - 频道 ID
- `msgID` - 消息 ID（可选，从此消息开始获取）
- `pin` - 是否只获取置顶消息
- `page` - 页码
- `pageSize` - 每页数量

### 获取消息详情

```go
msg, err := client.GetMessage(ctx, "msg-id")
// 或
msg, err := api.GetMessage(ctx, client, "msg-id")
```

### 发送消息

```go
// 简单发送
msg, err := client.SendMessage(ctx, "channel-id", "Hello!")

// 带引用发送
msg, err := client.SendReplyMessage(ctx, "channel-id", "Hello!", "quote-msg-id")

// 完整参数发送
msg, err := client.SendMessageEx(ctx, "channel-id", "Hello!", "", "", "")
```

**返回：** `*api.CreateMessageResponse`
- `MsgID` - 消息 ID

### 更新消息

```go
msg, err := client.UpdateMessage(ctx, "msg-id", "新内容")
// 或
msg, err := api.UpdateMessage(ctx, client, "msg-id", "新内容", "", "")
```

### 删除消息

```go
err := client.DeleteMessage(ctx, "msg-id")
```

### 添加回应

```go
// 使用 emoji ID
err := client.AddReaction(ctx, "msg-id", "128077")

// 使用 Unicode 字符
err := client.AddReactionWithEmoji(ctx, "msg-id", '👍')
```

### 删除回应

```go
// 删除自己的回应
err := client.DeleteReaction(ctx, "msg-id", "128077", "")

// 删除指定用户的回应
err := client.DeleteReaction(ctx, "msg-id", "128077", "user-id")

// 使用 Unicode 字符
err := client.DeleteReactionWithEmoji(ctx, "msg-id", '👍', "")
```

### 置顶/取消置顶消息

```go
// 置顶
err := client.PinMessage(ctx, "msg-id", "channel-id")

// 取消置顶
err := client.UnpinMessage(ctx, "msg-id", "channel-id")
```

---

## 私聊 API

管理私聊消息和会话。

### 获取私聊会话列表

```go
chats, err := client.ListUserChats(ctx, 1, 20)
// 或
chats, err := api.ListUserChats(ctx, client, 1, 20)
```

### 获取私聊会话详情

```go
chat, err := client.GetUserChat(ctx, "chat-code")
// 或
chat, err := api.GetUserChat(ctx, client, "chat-code")
```

### 创建私聊会话

```go
chat, err := client.CreateUserChat(ctx, "user-id")
// 或
chat, err := api.CreateUserChat(ctx, client, "user-id")
```

### 删除私聊会话

```go
err := client.DeleteUserChat(ctx, "chat-code")
```

### 发送私聊消息

```go
// 简单发送
dm, err := client.SendDirectMessage(ctx, "user-id", "Hello!")

// 完整参数发送
dm, err := client.SendDirectMessageEx(ctx, "user-id", "", "Hello!", "", "")
```

### 更新私聊消息

```go
dm, err := client.UpdateDirectMessage(ctx, "msg-id", "新内容")
```

### 删除私聊消息

```go
err := client.DeleteDirectMessage(ctx, "msg-id")
```

---

## 角色 API

管理服务器角色。

### 获取角色列表

```go
roles, err := client.ListGuildRoles(ctx, "guild-id", 1, 20)
// 或
roles, err := api.ListGuildRoles(ctx, client, "guild-id", 1, 20)
```

### 创建角色

```go
role, err := client.CreateGuildRole(ctx, "guild-id", "角色名称")
```

### 更新角色

```go
role, err := client.UpdateGuildRole(ctx, "guild-id", roleID, "新名称", 
    16776960, true, true, model.PermAdmin)
```

**参数：**
- `guildID` - 服务器 ID
- `roleID` - 角色 ID
- `name` - 角色名称
- `color` - 颜色（24 位 RGB）
- `hoist` - 是否在成员列表中单独显示
- `mentionable` - 是否允许被 @提及
- `permissions` - 权限位掩码

### 删除角色

```go
err := client.DeleteGuildRole(ctx, "guild-id", roleID)
```

### 赋予用户角色

```go
role, err := client.GrantRole(ctx, "guild-id", "user-id", roleID)
```

### 移除用户角色

```go
err := client.RevokeRole(ctx, "guild-id", "user-id", roleID)
```

---

## 频道权限 API

管理频道角色权限。

### 获取频道角色权限列表

```go
roles, err := client.GetChannelRoles(ctx, "channel-id")
// 或
roles, err := api.GetChannelRoles(ctx, client, "channel-id")
```

### 创建频道角色权限

```go
role, err := client.CreateChannelRole(ctx, "channel-id", "role_id", "123")
```

**参数：**
- `channelID` - 频道 ID
- `typ` - 类型（`role_id` 或 `user_id`）
- `value` - 角色 ID 或用户 ID

### 更新频道角色权限

```go
role, err := client.UpdateChannelRole(ctx, "channel-id", "role_id", "123", allow, deny)
```

**参数：**
- `allow` - 允许的权限位掩码
- `deny` - 拒绝的权限位掩码

### 同步频道角色权限

```go
err := client.SyncChannelRole(ctx, "channel-id")
```

### 删除频道角色权限

```go
err := client.DeleteChannelRole(ctx, "channel-id", "role_id", "123")
```

---

## 邀请 API

管理邀请链接。

### 获取邀请列表

```go
invites, err := client.ListInvites(ctx, "guild-id", "channel-id", 1, 20)
// 或
invites, err := api.ListInvites(ctx, client, "guild-id", "channel-id", 1, 20)
```

### 创建邀请链接

```go
url, err := client.CreateInvite(ctx, "guild-id", "channel-id", 3600, 1)
```

**参数：**
- `guildID` - 服务器 ID
- `channelID` - 频道 ID
- `duration` - 有效期（秒）
- `settingTimes` - 使用次数限制

### 删除邀请链接

```go
err := client.DeleteInvite(ctx, "url-code", api.WithDeleteInviteGuildID("guild-id"))
```

### 获取被邀请用户列表

```go
invitees, err := api.ListInvitees(ctx, client, "guild-id", 1, 10,
    api.WithInviteesID("url-code"))
```

---

## 亲密度 API

管理用户亲密度。

### 获取亲密度

```go
intimacy, err := client.GetIntimacy(ctx, "user-id")
// 或
intimacy, err := api.GetIntimacy(ctx, client, "user-id")
```

**返回字段：**
- `ImgURL` - 形象图片 URL
- `Score` - 亲密度分数
- `SocialInfo` - 社交信息
- `LastRead` - 上次查看时间
- `LastModify` - 最后修改时间
- `ImgList` - 形象图片列表

### 更新亲密度

```go
err := client.UpdateIntimacy(ctx, "user-id", 10, "社交信息", "")
```

**参数：**
- `userID` - 用户 ID
- `score` - 亲密度分数（0-2200）
- `socialInfo` - 社交信息（500字以内）
- `imgID` - 形象图片 ID

---

## 黑名单 API

管理服务器黑名单。

### 获取黑名单列表

```go
blacklist, err := client.ListBlacklist(ctx, "guild-id", 1, 20)
// 或
blacklist, err := api.ListBlacklist(ctx, client, "guild-id", 1, 20)
```

### 加入黑名单

```go
err := client.AddBlacklist(ctx, "guild-id", "user-id", "违规", 7)
```

**参数：**
- `guildID` - 服务器 ID
- `targetID` - 目标用户 ID
- `remark` - 备注
- `delMsgDays` - 删除消息天数（0=不删除，7=删除最近 7 天）

### 移除黑名单

```go
err := client.RemoveBlacklist(ctx, "guild-id", "user-id")
```

---

## 静音 API

管理服务器静音和闭麦。

### 获取静音列表

```go
mutes, err := client.ListGuildMutes(ctx, "guild-id", "", 1, 20)
// 或
mutes, err := api.ListGuildMutes(ctx, client, "guild-id", "", 1, 20)
```

**参数：**
- `returnType` - 返回类型（空=全部，`1`=静音，`2`=闭麦）

### 添加静音/闭麦

```go
err := client.CreateGuildMute(ctx, "guild-id", "user-id", "1")
```

**参数：**
- `muteType` - 类型（`1`=静音，`2`=闭麦）

### 取消静音/闭麦

```go
err := client.DeleteGuildMute(ctx, "guild-id", "user-id", "1")
```

---

## 助力 API

获取服务器助力历史。

### 获取助力历史

```go
boosts, err := client.GetBoostHistory(ctx, "guild-id", 0, 0)
// 或
boosts, err := api.GetBoostHistory(ctx, client, "guild-id", startTime, endTime)
```

**参数：**
- `startTime` - 开始时间戳（0=不限制）
- `endTime` - 结束时间戳（0=不限制）

---

## 媒体资源 API

上传文件和图片。

### 上传资源

```go
file, _ := os.Open("image.png")
defer file.Close()

asset, err := client.UploadAsset(ctx, file, "image.png")
// 或
asset, err := api.UploadAsset(ctx, client, file, "image.png")
```

**返回字段：**
- `URL` - 资源 URL

---

## 分页响应

所有分页接口返回 `model.PageResult[T]` 或 `model.ListResult[T]`：

```go
type PageResult[T any] struct {
	Items  []T   `json:"items"`
	Total  int   `json:"total"`
	Page   int   `json:"page"`
	Size   int   `json:"page_size"`
	Sort   string `json:"sort"`
}

type ListResult[T any] struct {
	Items []T `json:"items"`
}
```

**使用示例：**
```go
guilds, err := client.GetGuildList(ctx)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("总数: %d\n", guilds.Total)
for _, guild := range guilds.Items {
    fmt.Printf("- %s\n", guild.Name)
}
```

---

## 错误处理

### APIError

API 返回的业务错误：

```go
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
```

### HTTPError

HTTP 请求错误：

```go
type HTTPError struct {
	StatusCode int
	Body       string
}
```

### 错误检查

```go
me, err := client.Me(ctx)
if err != nil {
	var apiErr *kook.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("API 错误 %d: %s\n", apiErr.Code, apiErr.Message)
	}

	var httpErr *kook.HTTPError
	if errors.As(err, &httpErr) {
		fmt.Printf("HTTP 错误 %d\n", httpErr.StatusCode)
	}
}
```
