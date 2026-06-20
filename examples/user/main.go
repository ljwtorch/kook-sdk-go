// 用户 API 示例：展示用户相关接口的使用方法
//
// 环境变量：
//   - KOOK_BOT_TOKEN:        Bot Token（必填）
//   - KOOK_GUILD_ID:         服务器 ID（获取目标用户信息时需要）
//   - KOOK_TARGET_USER_ID:   目标用户 ID（获取目标用户信息时需要）
//
// 运行方式：
//
//	go run examples/user/main.go
//	go run examples/user/main.go --help
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	kook "github.com/ljwtorch/kook-sdk-go"
	"github.com/ljwtorch/kook-sdk-go/api"
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

	fmt.Println("========== 用户 API 测试 ==========")
	runTests(ctx, client)
	fmt.Println("\n========== 测试完成 ==========")
}

func printUsage() {
	fmt.Println("用户 API 示例")
	fmt.Println()
	fmt.Println("用法: go run examples/user/main.go [options]")
	fmt.Println()
	fmt.Println("选项:")
	fmt.Println("  --help    显示帮助信息")
	fmt.Println()
	fmt.Println("环境变量:")
	fmt.Println("  KOOK_BOT_TOKEN        Bot Token（必填）")
	fmt.Println("  KOOK_GUILD_ID         服务器 ID（获取目标用户信息时需要）")
	fmt.Println("  KOOK_TARGET_USER_ID   目标用户 ID（获取目标用户信息时需要）")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  export KOOK_BOT_TOKEN=\"your-bot-token\"")
	fmt.Println("  go run examples/user/main.go")
}

func runTests(ctx context.Context, client *kook.Client) {
	// 获取当前用户
	testGetCurrentUser(ctx, client)

	// 获取当前用户（api 包方式）
	testGetCurrentUserAPI(ctx, client)

	// 获取目标用户
	testGetUser(ctx, client)

	// 获取机器人在线状态
	testGetBotOnlineStatus(ctx, client)

	// 上线机器人
	testOnlineBot(ctx, client)

	// 下线机器人
	testOfflineBot(ctx, client)
}

func testGetCurrentUser(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取当前用户 ---")
	me, err := client.Me(ctx)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 用户名: %s\n", me.Username)
	fmt.Printf("     ID: %s\n", me.ID)
	fmt.Printf("     识别号: #%s\n", me.IdentifyNum)
	fmt.Printf("     在线: %v\n", me.Online)
	fmt.Printf("     状态: %d\n", me.Status)
	fmt.Printf("     头像: %s\n", me.Avatar)
	fmt.Printf("     是否机器人: %v\n", me.Bot)
}

func testGetCurrentUserAPI(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取当前用户（api 包方式）---")
	me, err := api.GetCurrentUser(ctx, client)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 用户名: %s (ID: %s)\n", me.Username, me.ID)
}

func testGetUser(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取目标用户 ---")

	guildID := os.Getenv("KOOK_GUILD_ID")
	targetUserID := os.Getenv("KOOK_TARGET_USER_ID")

	if guildID == "" || targetUserID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_GUILD_ID 和 KOOK_TARGET_USER_ID")
		return
	}

	user, err := client.GetUser(ctx, targetUserID, guildID)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 用户名: %s\n", user.Username)
	fmt.Printf("     ID: %s\n", user.ID)
	fmt.Printf("     识别号: #%s\n", user.IdentifyNum)
	fmt.Printf("     在线: %v\n", user.Online)
	fmt.Printf("     状态: %d\n", user.Status)
	fmt.Printf("     头像: %s\n", user.Avatar)
	fmt.Printf("     是否机器人: %v\n", user.Bot)
	fmt.Printf("     手机已验证: %v\n", user.MobileVerified)
	fmt.Printf("     邀请人数: %d\n", user.InvitedCount)
}

func testGetBotOnlineStatus(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取机器人在线状态 ---")
	status, err := client.GetBotOnlineStatus(ctx)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	onlineStatus := "离线"
	if status.Online {
		onlineStatus = "在线"
	}
	fmt.Printf("[OK] 机器人状态: %s\n", onlineStatus)
	if len(status.OnlineOS) > 0 {
		fmt.Printf("     在线平台: %v\n", status.OnlineOS)
	}
}

func testOnlineBot(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 上线机器人 ---")
	err := client.OnlineBot(ctx)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Println("[OK] 机器人已上线")
}

func testOfflineBot(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 下线机器人 ---")
	err := client.OfflineBot(ctx)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Println("[OK] 机器人已下线")
}
