// WebSocket 事件监听示例：展示如何接收 KOOK 的实时事件
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
	// 从环境变量读取 Bot Token
	token := os.Getenv("KOOK_BOT_TOKEN")
	if token == "" {
		log.Fatal("请设置 KOOK_BOT_TOKEN 环境变量")
	}

	// 创建客户端
	client := kook.NewClient(token)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// --- 注册消息事件处理器 ---
	client.OnMessage(func(ctx context.Context, evt *kook.EventData) {
		fmt.Printf("[%s] 消息来自用户 %s: %s\n",
			evt.ChannelType, evt.AuthorID, evt.Content)

		// 示例：自动回复
		if evt.Content == "ping" && evt.ChannelType == "GROUP" {
			_, err := client.SendMessage(ctx, evt.TargetID, "pong!")
			if err != nil {
				log.Printf("回复消息失败: %v", err)
			}
		}
	})

	// --- 注册系统事件处理器 ---
	client.OnSystem(func(ctx context.Context, evt *kook.EventData) {
		fmt.Printf("[系统事件] 类型=%d, 内容=%s\n", evt.Type, evt.Content)

		// 解析不同的系统事件
		switch evt.Content {
		case "added_reaction":
			var re event.ReactionAddedEvent
			if err := json.Unmarshal(evt.Extra, &re); err == nil {
				fmt.Printf("  用户 %s 对消息 %s 添加了回应 %s\n",
					re.UserID, re.MsgID, re.Emoji.Name)
			}
		case "deleted_reaction":
			var re event.ReactionDeletedEvent
			if err := json.Unmarshal(evt.Extra, &re); err == nil {
				fmt.Printf("  用户 %s 对消息 %s 删除了回应 %s\n",
					re.UserID, re.MsgID, re.Emoji.Name)
			}
		case "updated_message":
			var mu event.MessageUpdatedEvent
			if err := json.Unmarshal(evt.Extra, &mu); err == nil {
				fmt.Printf("  消息 %s 已更新，新内容: %s\n", mu.MsgID, mu.Content)
			}
		case "deleted_message":
			var md event.MessageDeletedEvent
			if err := json.Unmarshal(evt.Extra, &md); err == nil {
				fmt.Printf("  消息 %s 已被删除\n", md.MsgID)
			}
		case "added_channel":
			var ca event.ChannelAddedEvent
			if err := json.Unmarshal(evt.Extra, &ca); err == nil {
				fmt.Printf("  新频道: %s (ID: %s)\n", ca.Name, ca.ID)
			}
		case "joined_guild":
			var mj event.GuildMemberJoinedEvent
			if err := json.Unmarshal(evt.Extra, &mj); err == nil {
				fmt.Printf("  用户 %s 加入了服务器\n", mj.UserID)
			}
		case "exited_guild":
			var me event.GuildMemberExitedEvent
			if err := json.Unmarshal(evt.Extra, &me); err == nil {
				fmt.Printf("  用户 %s 退出了服务器\n", me.UserID)
			}
		}
	})

	// --- 连接 WebSocket ---
	fmt.Println("正在连接 WebSocket...")
	go func() {
		if err := client.Connect(ctx); err != nil {
			log.Printf("WebSocket 连接错误: %v", err)
		}
	}()

	// --- 等待退出信号 ---
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("\n正在断开连接...")
	cancel()
	if err := client.Disconnect(); err != nil {
		log.Printf("断开连接时出错: %v", err)
	}
	fmt.Println("已断开连接")
}
