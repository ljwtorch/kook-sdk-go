// 服务器 API 示例：展示服务器相关接口的使用方法
//
// 环境变量：
//   - KOOK_BOT_TOKEN:        Bot Token（必填）
//   - KOOK_GUILD_ID:         服务器 ID（部分功能需要）
//   - KOOK_TARGET_USER_ID:   目标用户 ID（踢出成员时需要）
//
// 运行方式：
//
//	go run examples/guild/main.go
//	go run examples/guild/main.go --help
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

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

	fmt.Println("========== 服务器 API 测试 ==========")
	runTests(ctx, client)
	fmt.Println("\n========== 测试完成 ==========")
}

func printUsage() {
	fmt.Println("服务器 API 示例")
	fmt.Println()
	fmt.Println("用法: go run examples/guild/main.go [options]")
	fmt.Println()
	fmt.Println("选项:")
	fmt.Println("  --help    显示帮助信息")
	fmt.Println()
	fmt.Println("环境变量:")
	fmt.Println("  KOOK_BOT_TOKEN        Bot Token（必填）")
	fmt.Println("  KOOK_GUILD_ID         服务器 ID（部分功能需要）")
	fmt.Println("  KOOK_TARGET_USER_ID   目标用户 ID（踢出成员时需要）")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  export KOOK_BOT_TOKEN=\"your-bot-token\"")
	fmt.Println("  export KOOK_GUILD_ID=\"your-guild-id\"")
	fmt.Println("  go run examples/guild/main.go")
}

func runTests(ctx context.Context, client *kook.Client) {
	// 获取服务器列表
	testGetGuildList(ctx, client)

	// 获取服务器详情
	testGetGuild(ctx, client)

	// 获取服务器成员列表
	testGetGuildUserList(ctx, client)

	// 修改服务器昵称
	testSetGuildNickname(ctx, client)

	// 离开服务器
	testLeaveGuild(ctx, client)

	// 踢出服务器成员
	testKickoutGuildMember(ctx, client)
}

func testGetGuildList(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取服务器列表 ---")
	guilds, err := client.GetGuildList(ctx)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 服务器数量: %d\n", len(guilds.Items))
	for i, guild := range guilds.Items {
		fmt.Printf("     [%d] %s (ID: %s)\n", i+1, guild.Name, guild.ID)
	}
}

func testGetGuild(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取服务器详情 ---")

	guildID := os.Getenv("KOOK_GUILD_ID")
	if guildID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_GUILD_ID")
		return
	}

	guild, err := client.GetGuild(ctx, guildID)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 服务器名称: %s\n", guild.Name)
	fmt.Printf("     ID: %s\n", guild.ID)
	fmt.Printf("     等级: %d\n", guild.Level)
	fmt.Printf("     所有者 ID: %s\n", guild.MasterID)
	fmt.Printf("     图标: %s\n", guild.Icon)
}

func testGetGuildUserList(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取服务器成员列表 ---")

	guildID := os.Getenv("KOOK_GUILD_ID")
	if guildID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_GUILD_ID")
		return
	}

	members, err := client.GetGuildUserList(ctx, guildID, 1, 10)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 成员数量: %d\n", len(members.Items))
	for i, member := range members.Items {
		fmt.Printf("     [%d] %s (昵称: %s, ID: %s)\n", i+1, member.Username, member.Nickname, member.ID)
	}
}

func testSetGuildNickname(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 修改服务器昵称 ---")

	guildID := os.Getenv("KOOK_GUILD_ID")
	if guildID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_GUILD_ID")
		return
	}

	// 设置昵称
	err := client.SetGuildNickname(ctx, guildID, "SDK测试昵称", "")
	if err != nil {
		fmt.Printf("[FAIL] 设置昵称: %v\n", err)
		return
	}
	fmt.Println("[OK] 昵称已设置为 'SDK测试昵称'")

	// 恢复昵称
	err = client.SetGuildNickname(ctx, guildID, "", "")
	if err != nil {
		fmt.Printf("[FAIL] 恢复昵称: %v\n", err)
		return
	}
	fmt.Println("[OK] 昵称已恢复")
}

func testLeaveGuild(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 离开服务器 ---")
	fmt.Println("[SKIP] 此操作会导致机器人离开服务器，请谨慎使用")
	fmt.Println("       如需测试，请取消注释代码")
	// err := client.LeaveGuild(ctx, "guild-id")
	// if err != nil {
	//     fmt.Printf("[FAIL] %v\n", err)
	//     return
	// }
	// fmt.Println("[OK] 已离开服务器")
}

func testKickoutGuildMember(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 踢出服务器成员 ---")

	guildID := os.Getenv("KOOK_GUILD_ID")
	targetUserID := os.Getenv("KOOK_TARGET_USER_ID")

	if guildID == "" || targetUserID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_GUILD_ID 和 KOOK_TARGET_USER_ID")
		return
	}

	fmt.Println("[SKIP] 此操作会将用户踢出服务器，请谨慎使用")
	fmt.Println("       如需测试，请取消注释代码")
	// err := client.KickoutGuildMember(ctx, guildID, targetUserID)
	// if err != nil {
	//     fmt.Printf("[FAIL] %v\n", err)
	//     return
	// }
	// fmt.Println("[OK] 已踢出成员")
}
