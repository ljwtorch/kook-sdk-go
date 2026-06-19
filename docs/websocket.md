# WebSocket 事件

本文档介绍如何使用 KOOK SDK for Go 监听实时事件。

## 目录

- [基本使用](#基本使用)
- [事件处理器](#事件处理器)
- [事件数据结构](#事件数据结构)
- [系统事件类型](#系统事件类型)
- [消息类型](#消息类型)
- [连接管理](#连接管理)
- [压缩模式](#压缩模式)
- [完整示例](#完整示例)

---

## 基本使用

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	kook "github.com/ljwtorch/kook-sdk-go"
)

func main() {
	token := os.Getenv("KOOK_BOT_TOKEN")
	client := kook.NewClient(token)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 注册消息处理器
	client.OnMessage(func(ctx context.Context, evt *kook.EventData) {
		fmt.Printf("收到消息: %s\n", evt.Content)
	})

	// 连接 WebSocket
	go func() {
		if err := client.Connect(ctx); err != nil {
			log.Printf("连接错误: %v", err)
		}
	}()

	// 等待退出信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	cancel()
	client.Disconnect()
}
```

---

## 事件处理器

SDK 提供三种注册事件处理器的方式：

### OnMessage - 消息事件

处理用户发送的消息（type 1-10）：

```go
client.OnMessage(func(ctx context.Context, evt *kook.EventData) {
    fmt.Printf("频道类型: %s\n", evt.ChannelType) // GROUP 或 PERSON
    fmt.Printf("消息内容: %s\n", evt.Content)
    fmt.Printf("发送者: %s\n", evt.AuthorID)
    fmt.Printf("频道 ID: %s\n", evt.TargetID)
    fmt.Printf("消息 ID: %s\n", evt.MsgID)
})
```

### OnSystem - 系统事件

处理系统事件（type 255）：

```go
client.OnSystem(func(ctx context.Context, evt *kook.EventData) {
    fmt.Printf("系统事件: %s\n", evt.Content)

    // 根据事件类型解析 extra 数据
    switch evt.Content {
    case "added_channel":
        // 频道创建
    case "deleted_channel":
        // 频道删除
    case "joined_guild":
        // 成员加入
    case "exited_guild":
        // 成员退出
    // ... 更多事件类型
    }
})
```

### OnEvent - 自定义事件

注册特定类型的事件处理器：

```go
client.OnEvent("message", func(ctx context.Context, evt *kook.EventData) {
    // 处理消息事件
})

client.OnEvent("*", func(ctx context.Context, evt *kook.EventData) {
    // 处理所有事件
})
```

---

## 事件数据结构

```go
type EventData struct {
	// 频道类型：GROUP=频道消息，PERSON=私聊消息
	ChannelType string `json:"channel_type"`

	// 消息类型
	Type int `json:"type"`

	// 目标 ID（频道 ID 或用户 ID）
	TargetID string `json:"target_id"`

	// 作者 ID
	AuthorID string `json:"author_id"`

	// 消息内容
	Content string `json:"content"`

	// 消息 ID
	MsgID string `json:"msg_id"`

	// 消息时间戳
	MsgTimestamp int64 `json:"msg_timestamp"`

	// 随机数
	Nonce string `json:"nonce"`

	// 额外数据（JSON）
	Extra json.RawMessage `json:"extra"`
}
```

---

## 系统事件类型

系统事件通过 `evt.Content` 字段区分：

### 频道事件

| 事件 | 说明 |
|------|------|
| `added_channel` | 频道创建 |
| `updated_channel` | 频道更新 |
| `deleted_channel` | 频道删除 |

### 消息事件

| 事件 | 说明 |
|------|------|
| `updated_message` | 消息更新 |
| `deleted_message` | 消息删除 |
| `pinned_message` | 消息置顶 |
| `unpinned_message` | 取消置顶 |

### 回应事件

| 事件 | 说明 |
|------|------|
| `added_reaction` | 添加回应 |
| `deleted_reaction` | 删除回应 |

### 成员事件

| 事件 | 说明 |
|------|------|
| `joined_guild` | 成员加入服务器 |
| `exited_guild` | 成员退出服务器 |
| `updated_guild_member` | 成员信息更新 |
| `guild_member_online` | 成员上线 |
| `guild_member_offline` | 成员下线 |

### 角色事件

| 事件 | 说明 |
|------|------|
| `added_role` | 角色创建 |
| `deleted_role` | 角色删除 |
| `updated_role` | 角色更新 |

### 用户事件

| 事件 | 说明 |
|------|------|
| `updated_user` | 用户信息更新 |
| `self_joined_guild` | 机器人加入服务器 |
| `self_exited_guild` | 机器人退出服务器 |

### 服务器事件

| 事件 | 说明 |
|------|------|
| `updated_guild` | 服务器更新 |
| `deleted_guild` | 服务器删除 |
| `added_block_list` | 黑名单添加 |
| `deleted_block_list` | 黑名单移除 |

---

## 消息类型

消息类型通过 `evt.Type` 字段区分：

| 类型 | 说明 |
|------|------|
| 1 | 文字消息 |
| 2 | 图片消息 |
| 3 | 视频消息 |
| 4 | 文件消息 |
| 8 | 音频消息 |
| 9 | KMarkdown 消息 |
| 10 | 卡片消息 |
| 255 | 系统消息 |

---

## 连接管理

### 建立连接

```go
// 普通连接
err := client.Connect(ctx)

// 带压缩的连接
err := client.ConnectWithCompress(ctx)
```

`Connect` 方法会阻塞直到：
- `ctx` 被取消
- 调用 `Disconnect()`
- 发生不可恢复的错误

### 断开连接

```go
err := client.Disconnect()
```

### 检查连接状态

```go
// 使用 Close 方法（会断开连接并清理资源）
err := client.Close()
```

---

## 压缩模式

启用 zlib 压缩可以减少网络传输数据量：

```go
err := client.ConnectWithCompress(ctx)
```

**注意：** 压缩模式下，服务端会对 WebSocket 消息进行 zlib 压缩，SDK 会自动解压。

---

## 完整示例

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	kook "github.com/ljwtorch/kook-sdk-go"
	"github.com/ljwtorch/kook-sdk-go/event"
)

func main() {
	token := os.Getenv("KOOK_BOT_TOKEN")
	if token == "" {
		log.Fatal("请设置 KOOK_BOT_TOKEN 环境变量")
	}

	client := kook.NewClient(token, kook.WithDebug(true))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 注册消息处理器
	client.OnMessage(func(ctx context.Context, evt *kook.EventData) {
		fmt.Printf("[%s] 消息来自 %s: %s\n", evt.ChannelType, evt.AuthorID, evt.Content)

		// 自动回复 ping
		if evt.Content == "ping" && evt.ChannelType == "GROUP" {
			_, err := client.SendMessage(ctx, evt.TargetID, "pong!")
			if err != nil {
				log.Printf("回复失败: %v", err)
			}
		}
	})

	// 注册系统事件处理器
	client.OnSystem(func(ctx context.Context, evt *kook.EventData) {
		fmt.Printf("[系统] %s\n", evt.Content)

		switch evt.Content {
		case "added_reaction":
			var re event.ReactionAddedEvent
			if err := json.Unmarshal(evt.Extra, &re); err == nil {
				fmt.Printf("  用户 %s 添加了回应 %s\n", re.UserID, re.Emoji.Name)
			}

		case "joined_guild":
			var mj event.GuildMemberJoinedEvent
			if err := json.Unmarshal(evt.Extra, &mj); err == nil {
				fmt.Printf("  欢迎用户 %s 加入服务器！\n", mj.UserID)
			}

		case "deleted_message":
			var md event.MessageDeletedEvent
			if err := json.Unmarshal(evt.Extra, &md); err == nil {
				fmt.Printf("  消息 %s 已被删除\n", md.MsgID)
			}
		}
	})

	// 连接 WebSocket
	fmt.Println("正在连接 WebSocket...")
	go func() {
		if err := client.Connect(ctx); err != nil {
			log.Printf("WebSocket 错误: %v", err)
		}
	}()

	// 等待退出信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("\n正在断开连接...")
	cancel()
	if err := client.Disconnect(); err != nil {
		log.Printf("断开连接错误: %v", err)
	}
	fmt.Println("已断开连接")
}
```

---

## 事件类型定义

SDK 在 `event` 包中预定义了常用事件的结构体：

```go
// 频道事件
type ChannelAddedEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 消息事件
type MessageUpdatedEvent struct {
	MsgID   string `json:"msg_id"`
	Content string `json:"content"`
}

type MessageDeletedEvent struct {
	MsgID string `json:"msg_id"`
}

// 回应事件
type ReactionAddedEvent struct {
	MsgID  string `json:"msg_id"`
	UserID string `json:"user_id"`
	Emoji  Emoji  `json:"emoji"`
}

type ReactionDeletedEvent struct {
	MsgID  string `json:"msg_id"`
	UserID string `json:"user_id"`
	Emoji  Emoji  `json:"emoji"`
}

// 成员事件
type GuildMemberJoinedEvent struct {
	UserID  string `json:"user_id"`
	GuildID string `json:"guild_id"`
}

type GuildMemberExitedEvent struct {
	UserID  string `json:"user_id"`
	GuildID string `json:"guild_id"`
}

// Emoji 结构
type Emoji struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
```

更多事件类型请参考 `event/` 包。
