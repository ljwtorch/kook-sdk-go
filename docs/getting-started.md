# 快速开始

本指南将帮助你从零开始使用 KOOK SDK for Go 创建一个 KOOK Bot。

## 前置条件

- Go 1.21 或更高版本
- 一个 KOOK Bot Token（从 [KOOK 开放平台](https://developer.kookapp.cn/) 获取）

## 安装

```bash
go get github.com/ljwtorch/kook-sdk-go
```

## 创建客户端

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
	fmt.Printf("Bot 已登录: %s\n", me.Username)
}
```

## 配置选项

SDK 使用 Functional Options 模式进行配置：

```go
client := kook.NewClient("your-bot-token",
	// 启用调试模式，输出请求日志
	kook.WithDebug(true),

	// 自定义 User-Agent
	kook.WithUserAgent("MyBot/1.0"),

	// 自定义 API 基础地址
	kook.WithBaseURL("https://custom-api.example.com"),

	// 指定 API 版本（默认 v3）
	kook.WithAPIVersion("v3"),
)
```

### 可用配置项

| 选项 | 说明 | 默认值 |
|------|------|--------|
| `WithDebug(bool)` | 启用调试日志 | `false` |
| `WithUserAgent(string)` | 自定义 User-Agent | `KookBotGo/{version}` |
| `WithBaseURL(string)` | API 基础地址 | `https://www.kookapp.cn` |
| `WithAPIVersion(string)` | API 版本 | `v3` |

## 使用环境变量管理 Token

建议使用环境变量管理 Bot Token，避免硬编码：

```go
package main

import (
	"context"
	"log"
	"os"

	kook "github.com/ljwtorch/kook-sdk-go"
)

func main() {
	token := os.Getenv("KOOK_BOT_TOKEN")
	if token == "" {
		log.Fatal("请设置 KOOK_BOT_TOKEN 环境变量")
	}

	client := kook.NewClient(token)
	defer client.Close()

	// ...
}
```

运行时设置环境变量：

```bash
export KOOK_BOT_TOKEN="your-bot-token"
go run main.go
```

## 两种调用方式

SDK 提供两种方式调用 API：

### 1. Client 便捷方法

直接在 Client 实例上调用，无需导入 `api` 包：

```go
// 获取当前用户
me, err := client.Me(ctx)

// 发送消息
msg, err := client.SendMessage(ctx, "channel-id", "Hello!")

// 获取服务器列表
guilds, err := client.GetGuildList(ctx)
```

### 2. api 包直接调用

对于需要更精细控制的场景，使用 `api` 包：

```go
import "github.com/ljwtorch/kook-sdk-go/api"

// 获取带筛选条件的成员列表
members, err := api.GetGuildUserList(ctx, client, guildID, 1, 20,
    "channel-id", "search", 0, false, 7, 0)
```

## 错误处理

SDK 返回标准的 Go 错误类型：

```go
import (
	"errors"
	kook "github.com/ljwtorch/kook-sdk-go"
)

me, err := client.Me(ctx)
if err != nil {
	// 检查是否为 API 错误
	var apiErr *kook.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("API 错误: code=%d, message=%s\n", apiErr.Code, apiErr.Message)
		return
	}

	// 检查是否为 HTTP 错误
	var httpErr *kook.HTTPError
	if errors.As(err, &httpErr) {
		fmt.Printf("HTTP 错误: status=%d\n", httpErr.StatusCode)
		return
	}

	// 其他错误
	log.Fatal(err)
}
```

## 完整示例

更多完整示例请参考 `examples/` 目录：

- [基础使用示例](../examples/basic/main.go) - API 调用示例
- [WebSocket 示例例](../examples/websocket/main.go) - 事件监听示例
- [卡片消息示例](../examples/card_message/main.go) - 卡片消息构建示例

## 下一步

- [API 参考](api-reference.md) - 了解所有可用的 API 接口
- [WebSocket 事件](websocket.md) - 学习如何监听实时事件
- [卡片消息](card-message.md) - 构建丰富的卡片消息
