// 消息 API 示例：展示消息相关接口的使用方法
//
// 环境变量：
//   - KOOK_BOT_TOKEN:    Bot Token（必填）
//   - KOOK_CHANNEL_ID:   频道 ID（必填）
//   - KOOK_MSG_ID:       消息 ID（获取详情时需要）
//
// 运行方式：
//
//	go run examples/message/main.go
//	go run examples/message/main.go --help
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	kook "github.com/ljwtorch/kook-sdk-go"
)

func main() {
	help := flag.Bool("help", false, "显示帮助信息")
	flag.Parse()

	if *help {
		printUsage()
		return
	}

	token := os.Getenv("KOOK_BOT_TOKEN")
	if token == "" {
		log.Fatal("请设置 KOOK_BOT_TOKEN 环境变量")
	}

	client := kook.NewClient(token, kook.WithDebug(false))
	defer client.Close()

	ctx := context.Background()

	fmt.Println("========== 消息 API 测试 ==========")
	runTests(ctx, client)
	fmt.Println("\n========== 测试完成 ==========")
}

func printUsage() {
	fmt.Println("消息 API 示例")
	fmt.Println()
	fmt.Println("用法: go run examples/message/main.go [options]")
	fmt.Println()
	fmt.Println("选项:")
	fmt.Println("  --help    显示帮助信息")
	fmt.Println()
	fmt.Println("环境变量:")
	fmt.Println("  KOOK_BOT_TOKEN    Bot Token（必填）")
	fmt.Println("  KOOK_CHANNEL_ID   频道 ID（必填）")
	fmt.Println("  KOOK_MSG_ID       消息 ID（获取详情时需要）")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  export KOOK_BOT_TOKEN=\"your-bot-token\"")
	fmt.Println("  export KOOK_CHANNEL_ID=\"your-channel-id\"")
	fmt.Println("  go run examples/message/main.go")
}

func runTests(ctx context.Context, client *kook.Client) {
	// 获取消息列表
	testListMessages(ctx, client)

	// 获取消息详情
	testGetMessage(ctx, client)

	// 发送消息
	testSendMessage(ctx, client)

	// 编辑消息
	testUpdateMessage(ctx, client)

	// 添加回应
	testAddReaction(ctx, client)

	// 删除回应
	testDeleteReaction(ctx, client)

	// 置顶消息
	testPinMessage(ctx, client)

	// 取消置顶
	testUnpinMessage(ctx, client)

	// 删除消息
	testDeleteMessage(ctx, client)
}

func testListMessages(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取消息列表 ---")

	channelID := os.Getenv("KOOK_CHANNEL_ID")
	if channelID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_CHANNEL_ID")
		return
	}

	messages, err := client.ListMessages(ctx, channelID, 1, 10)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 消息数量: %d\n", len(messages.Items))
	for i, msg := range messages.Items {
		fmt.Printf("     [%d] [%d] %s (ID: %s)\n", i+1, msg.Type, msg.Content, msg.ID)
	}
}

func testGetMessage(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取消息详情 ---")

	msgID := os.Getenv("KOOK_MSG_ID")
	if msgID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_MSG_ID")
		return
	}

	msg, err := client.GetMessage(ctx, msgID)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 消息内容: %s\n", msg.Content)
	fmt.Printf("     ID: %s\n", msg.ID)
	fmt.Printf("     类型: %d\n", msg.Type)
	if msg.Author != nil {
		fmt.Printf("     作者: %s (ID: %s)\n", msg.Author.Username, msg.Author.ID)
	}
	fmt.Printf("     频道 ID: %s\n", msg.ChannelID)
}

func testSendMessage(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 发送消息 ---")

	channelID := os.Getenv("KOOK_CHANNEL_ID")
	if channelID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_CHANNEL_ID")
		return
	}

	content := fmt.Sprintf("KOOK Go SDK 测试消息 - %d", time.Now().Unix())
	msg, err := client.SendMessage(ctx, channelID, content)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 消息已发送 (ID: %s)\n", msg.MsgID)
	fmt.Printf("     请记录此 msgID 用于后续测试: %s\n", msg.MsgID)
}

func testUpdateMessage(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 编辑消息 ---")

	msgID := os.Getenv("KOOK_MSG_ID")
	if msgID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_MSG_ID")
		return
	}

	content := fmt.Sprintf("KOOK Go SDK 测试消息（已编辑）- %d", time.Now().Unix())
	msg, err := client.UpdateMessage(ctx, msgID, content)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 消息已编辑: %s\n", msg.Content)
}

func testAddReaction(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 添加回应 ---")

	msgID := os.Getenv("KOOK_MSG_ID")
	if msgID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_MSG_ID")
		return
	}

	// 使用 Unicode 字符
	err := client.AddReactionWithEmoji(ctx, msgID, '👍')
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Println("[OK] 已添加回应 👍")
}

func testDeleteReaction(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 删除回应 ---")

	msgID := os.Getenv("KOOK_MSG_ID")
	if msgID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_MSG_ID")
		return
	}

	// 删除自己的回应
	err := client.DeleteReactionWithEmoji(ctx, msgID, '👍', "")
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Println("[OK] 已删除回应 👍")
}

func testPinMessage(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 置顶消息 ---")

	msgID := os.Getenv("KOOK_MSG_ID")
	channelID := os.Getenv("KOOK_CHANNEL_ID")

	if msgID == "" || channelID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_MSG_ID 和 KOOK_CHANNEL_ID")
		return
	}

	err := client.PinMessage(ctx, msgID, channelID)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Println("[OK] 消息已置顶")
}

func testUnpinMessage(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 取消置顶 ---")

	msgID := os.Getenv("KOOK_MSG_ID")
	channelID := os.Getenv("KOOK_CHANNEL_ID")

	if msgID == "" || channelID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_MSG_ID 和 KOOK_CHANNEL_ID")
		return
	}

	err := client.UnpinMessage(ctx, msgID, channelID)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Println("[OK] 已取消置顶")
}

func testDeleteMessage(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 删除消息 ---")

	msgID := os.Getenv("KOOK_MSG_ID")
	if msgID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_MSG_ID")
		return
	}

	fmt.Println("[SKIP] 此操作会删除消息，请谨慎使用")
	fmt.Println("       如需测试，请取消注释代码")
	// err := client.DeleteMessage(ctx, msgID)
	// if err != nil {
	//     fmt.Printf("[FAIL] %v\n", err)
	//     return
	// }
	// fmt.Println("[OK] 消息已删除")
}
