// 私聊 API 示例：展示私聊相关接口的使用方法
//
// 环境变量：
//   - KOOK_BOT_TOKEN:        Bot Token（必填）
//   - KOOK_TARGET_USER_ID:   目标用户 ID（必填）
//   - KOOK_CHAT_CODE:        私聊会话 Code（获取消息列表时需要）
//   - KOOK_DM_MSG_ID:        私聊消息 ID（编辑/删除消息时需要）
//
// 运行方式：
//
//	go run examples/direct_message/main.go
//	go run examples/direct_message/main.go --help
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

	fmt.Println("========== 私聊 API 测试 ==========")
	runTests(ctx, client)
	fmt.Println("\n========== 测试完成 ==========")
}

func printUsage() {
	fmt.Println("私聊 API 示例")
	fmt.Println()
	fmt.Println("用法: go run examples/direct_message/main.go [options]")
	fmt.Println()
	fmt.Println("选项:")
	fmt.Println("  --help    显示帮助信息")
	fmt.Println()
	fmt.Println("环境变量:")
	fmt.Println("  KOOK_BOT_TOKEN        Bot Token（必填）")
	fmt.Println("  KOOK_TARGET_USER_ID   目标用户 ID（必填）")
	fmt.Println("  KOOK_CHAT_CODE        私聊会话 Code（获取消息列表时需要）")
	fmt.Println("  KOOK_DM_MSG_ID        私聊消息 ID（编辑/删除消息时需要）")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  export KOOK_BOT_TOKEN=\"your-bot-token\"")
	fmt.Println("  export KOOK_TARGET_USER_ID=\"target-user-id\"")
	fmt.Println("  go run examples/direct_message/main.go")
}

func runTests(ctx context.Context, client *kook.Client) {
	// 获取私聊会话列表
	testListUserChats(ctx, client)

	// 创建私聊会话
	testCreateUserChat(ctx, client)

	// 发送私聊消息
	testSendDirectMessage(ctx, client)

	// 编辑私聊消息
	testUpdateDirectMessage(ctx, client)

	// 删除私聊消息
	testDeleteDirectMessage(ctx, client)

	// 删除私聊会话
	testDeleteUserChat(ctx, client)
}

func testListUserChats(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取私聊会话列表 ---")

	chats, err := client.ListUserChats(ctx, 1, 10)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 私聊会话数量: %d\n", len(chats.Items))
	for i, chat := range chats.Items {
		fmt.Printf("     [%d] Code: %s\n", i+1, chat.Code)
	}
}

func testCreateUserChat(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 创建私聊会话 ---")

	targetUserID := os.Getenv("KOOK_TARGET_USER_ID")
	if targetUserID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_TARGET_USER_ID")
		return
	}

	chat, err := client.CreateUserChat(ctx, targetUserID)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 私聊会话已创建: Code=%s\n", chat.Code)
	fmt.Printf("     请记录此 chatCode 用于后续测试: %s\n", chat.Code)
}

func testSendDirectMessage(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 发送私聊消息 ---")

	targetUserID := os.Getenv("KOOK_TARGET_USER_ID")
	if targetUserID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_TARGET_USER_ID")
		return
	}

	content := fmt.Sprintf("KOOK Go SDK 私聊测试消息 - %d", time.Now().Unix())
	dm, err := client.SendDirectMessage(ctx, targetUserID, content)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 私聊消息已发送 (ID: %s)\n", dm.MsgID)
	fmt.Printf("     请记录此 dmMsgID 用于后续测试: %s\n", dm.MsgID)
}

func testUpdateDirectMessage(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 编辑私聊消息 ---")

	msgID := os.Getenv("KOOK_DM_MSG_ID")
	if msgID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_DM_MSG_ID")
		return
	}

	content := fmt.Sprintf("KOOK Go SDK 私聊测试消息（已编辑）- %d", time.Now().Unix())
	dm, err := client.UpdateDirectMessage(ctx, msgID, content)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 私聊消息已编辑: %s\n", dm.Content)
}

func testDeleteDirectMessage(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 删除私聊消息 ---")

	msgID := os.Getenv("KOOK_DM_MSG_ID")
	if msgID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_DM_MSG_ID")
		return
	}

	fmt.Println("[SKIP] 此操作会删除私聊消息，请谨慎使用")
	fmt.Println("       如需测试，请取消注释代码")
	// err := client.DeleteDirectMessage(ctx, msgID)
	// if err != nil {
	//     fmt.Printf("[FAIL] %v\n", err)
	//     return
	// }
	// fmt.Println("[OK] 私聊消息已删除")
}

func testDeleteUserChat(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 删除私聊会话 ---")

	chatCode := os.Getenv("KOOK_CHAT_CODE")
	if chatCode == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_CHAT_CODE")
		return
	}

	fmt.Println("[SKIP] 此操作会删除私聊会话，请谨慎使用")
	fmt.Println("       如需测试，请取消注释代码")
	// err := client.DeleteUserChat(ctx, chatCode)
	// if err != nil {
	//     fmt.Printf("[FAIL] %v\n", err)
	//     return
	// }
	// fmt.Println("[OK] 私聊会话已删除")
}
