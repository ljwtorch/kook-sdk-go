# KOOK SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/ljwtorch/kook-sdk-go)](https://pkg.go.dev/github.com/ljwtorch/kook-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ljwtorch/kook-sdk-go)](https://goreportcard.com/report/github.com/ljwtorch/kook-sdk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> **⚠️ Dev 版本须知**
>
> 本项目目前处于 **Dev 阶段**，API 可能会发生变化，**不建议用于生产环境**。
>
> 版本号遵循 [语义化版本](https://semver.org/lang/zh-CN/) 规范，当前版本迭代规则：`dev.1` → `dev.2` → ... → `beta.1` → `rc.1` → `0.1.0`
> 
> **备注：** 
> 
> 1. 项目中部分文档是由AI读取项目自生成的，后续会逐个进行核对，如和实际情况有出入可提Issue来协助解决；
>
> 2. 目前只测试了服务器/频道/私聊会话/用户/消息(卡片)等增删改查接口，websocket相关未测试；

KOOK（原开黑啦）开放平台的 Go 语言 SDK，提供完整的 HTTP API 封装、WebSocket 事件监听和卡片消息构建能力。

## 安装

```bash
go get github.com/ljwtorch/kook-sdk-go
```

## 快速开始

### 基本使用

```go
package main

import (
	"context"
	"fmt"
	"log"

	kook "github.com/ljwtorch/kook-sdk-go"
)

func main() {
	// 创建客户端
	client := kook.NewClient("your-bot-token")
	defer client.Close()

	ctx := context.Background()

	// 获取当前用户信息
	me, err := client.Me(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello, %s!\n", me.Username)

	// 获取服务器列表
	guilds, err := client.GetGuildList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, guild := range guilds.Items {
		fmt.Printf("Guild: %s\n", guild.Name)
	}
}
```

### API 版本配置

SDK 支持配置 API 版本，默认使用 v3。当 KOOK 官方发布新版本 API 时，可以通过以下方式切换：

```go
// 使用默认版本 (v3)
client := kook.NewClient("your-bot-token")

// 指定使用 v4 版本
client := kook.NewClient("your-bot-token",
	kook.WithAPIVersion("v4"),
)
```

**注意：** API 版本配置会影响所有 HTTP API 调用和 WebSocket Gateway 连接。

### 自定义配置

```go
client := kook.NewClient("your-bot-token",
	kook.WithBaseURL("https://custom-api.example.com"),
	kook.WithDebug(true),
	kook.WithUserAgent("MyBot/1.0"),
)
```

## Client 便捷方法

`Client` 提供了一系列便捷方法，可直接在 Client 实例上调用，无需额外导入 `api` 包。

```go
// 获取当前用户
me, err := client.Me(ctx)

// 发送频道消息
msg, err := client.SendMessage(ctx, "channel-id", "Hello!")

// 发送私聊消息
dm, err := client.SendDirectMessage(ctx, "user-id", "Hi!")

// 获取服务器列表
guilds, err := client.GetGuildList(ctx)
```

完整的便捷方法包括：用户、服务器、频道、消息、私聊、角色、邀请、黑名单等 50+ 个方法。详见 [API 参考文档](docs/api-reference.md)。

## API 模块

| 模块 | 说明 | 接口数 | 文档 |
|------|------|--------|------|
| 服务器 (Guild) | 服务器管理 | 6 | [API](https://developer.kookapp.cn/doc/http/guild) |
| 频道 (Channel) | 频道管理 | 8 | [API](https://developer.kookapp.cn/doc/http/channel) |
| 频道权限 (Channel Role) | 频道角色权限 | 5 | [API](https://developer.kookapp.cn/doc/http/channel) |
| 消息 (Message) | 频道消息 | 11 | [API](https://developer.kookapp.cn/doc/http/message) |
| 私聊消息 (Direct Message) | 私聊消息 | 8 | [API](https://developer.kookapp.cn/doc/http/direct-message) |
| 私聊会话 (User Chat) | 私聊会话管理 | 4 | [API](https://developer.kookapp.cn/doc/http/user-chat) |
| 用户 (User) | 用户信息 | 5 | [API](https://developer.kookapp.cn/doc/http/user) |
| 服务器角色 (Guild Role) | 角色管理 | 6 | [API](https://developer.kookapp.cn/doc/http/guild-role) |
| 邀请 (Invite) | 邀请管理 | 4 | [API](https://developer.kookapp.cn/doc/http/invite) |
| 亲密度 (Intimacy) | 亲密度管理 | 2 | [API](https://developer.kookapp.cn/doc/http/intimacy) |
| 黑名单 (Blacklist) | 黑名单管理 | 3 | [API](https://developer.kookapp.cn/doc/http/blacklist) |
| 服务器静音 (Guild Mute) | 静音管理 | 3 | [API](https://developer.kookapp.cn/doc/http/guild-mute) |
| 服务器助力 (Guild Boost) | 助力历史 | 1 | [API](https://developer.kookapp.cn/doc/http/guild) |
| 媒体资源 (Asset) | 文件上传 | 1 | [API](https://developer.kookapp.cn/doc/http/asset) |

## 示例代码

`examples/` 目录包含了可独立运行的示例代码，展示各模块的使用方法。

### 示例列表

| 示例 | 说明 | 运行命令 |
|------|------|----------|
| [user](examples/user/) | 用户 API（获取用户、上下线） | `go run examples/user/main.go` |
| [guild](examples/guild/) | 服务器 API（列表、详情、成员管理） | `go run examples/guild/main.go` |
| [channel](examples/channel/) | 频道 API（创建、编辑、删除） | `go run examples/channel/main.go` |
| [message](examples/message/) | 消息 API（发送、编辑、删除、回应） | `go run examples/message/main.go` |
| [card](examples/card/) | 卡片消息（构建和发送） | `go run examples/card/main.go` |
| [direct_message](examples/direct_message/) | 私聊 API（会话、消息） | `go run examples/direct_message/main.go` |
| [websocket](examples/websocket/) | WebSocket 事件监听 | `go run examples/websocket/main.go` |

### 运行示例

每个示例都需要设置环境变量 `KOOK_BOT_TOKEN`（从 [KOOK 开放平台](https://developer.kookapp.cn/) 获取）：

```bash
# 设置 Token
export KOOK_BOT_TOKEN="your-bot-token"

# 运行用户示例
go run examples/user/main.go

# 查看帮助
go run examples/user/main.go --help
```

部分示例需要额外的环境变量，详见各示例的 `--help` 输出，完整示例说明请参考 [examples/README.md](examples/README.md)。

## 项目结构

```
kook-sdk-go/
├── client.go          # HTTP Client 核心（Do, DoMultipart, Get, Post, Delete）
├── client_options.go  # Functional Options 配置
├── convenience.go     # Client 便捷方法（50+ 个）
├── errors.go          # 错误定义（APIError, HTTPError）
├── ratelimit.go       # 速率限制（自动等待重试）
├── signal.go          # WebSocket 信令定义（Signal 0-6）
├── gateway.go         # WebSocket Gateway（心跳、重连、事件分发）
├── kook.go            # 包入口（版本号、类型别名）
├── api/               # HTTP API 模块（15 个文件，67 个接口）
├── model/             # 数据模型（Guild, Channel, User, Message, Role, Card 等）
├── event/             # 事件类型定义（频道、消息、成员、角色、回应、用户、服务器）
├── card/              # 卡片消息构建器（Builder 链式调用）
├── internal/          # 内部实现（HTTP 封装、压缩、退避策略）
├── docs/              # 详细文档
└── examples/          # 使用示例（可独立运行）
    ├── user/          # 用户 API 示例
    ├── guild/         # 服务器 API 示例
    ├── channel/       # 频道 API 示例
    ├── message/       # 消息 API 示例
    ├── card/          # 卡片消息示例
    ├── direct_message/# 私聊 API 示例
    └── websocket/     # WebSocket 事件监听示例
```

## 详细文档

更详细的使用说明请参考 [docs](docs/) 目录：

- [快速开始](docs/getting-started.md) - 安装、配置、第一个程序
- [API 参考](docs/api-reference.md) - 所有 API 模块的详细使用说明
- [WebSocket 事件](docs/websocket.md) - 实时事件监听和处理
- [卡片消息](docs/card-message.md) - 卡片消息构建器使用指南

## Kook官方参考文档

- [KOOK 官方文档](https://developer.kookapp.cn/doc/intro)
- [KOOK HTTP API](https://developer.kookapp.cn/doc/reference)
- [KOOK WebSocket](https://developer.kookapp.cn/doc/websocket)
- [KOOK 卡片消息](https://developer.kookapp.cn/doc/cardmessage)

## 参与贡献

- **提交 Issue** - 发现 Bug 或有功能建议？[提交 Issue](https://github.com/ljwtorch/kook-sdk-go/issues/new)
- **提交 PR** - 想要贡献代码？请先阅读 [贡献指南](docs/contributing.md)
- **Star** - 如果这个项目对你有帮助，请给我们一个 Star！

## License

MIT
