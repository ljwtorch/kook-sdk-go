// 频道 API 示例：展示频道相关接口的使用方法
//
// 环境变量：
//   - KOOK_BOT_TOKEN:    Bot Token（必填）
//   - KOOK_GUILD_ID:     服务器 ID（创建/删除频道时需要）
//   - KOOK_CHANNEL_ID:   频道 ID（编辑/获取详情时需要）
//
// 运行方式：
//
//	go run examples/channel/main.go
//	go run examples/channel/main.go --help
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	kook "github.com/ljwtorch/kook-sdk-go"
	"github.com/ljwtorch/kook-sdk-go/model"
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

	fmt.Println("========== 频道 API 测试 ==========")
	runTests(ctx, client)
	fmt.Println("\n========== 测试完成 ==========")
}

func printUsage() {
	fmt.Println("频道 API 示例")
	fmt.Println()
	fmt.Println("用法: go run examples/channel/main.go [options]")
	fmt.Println()
	fmt.Println("选项:")
	fmt.Println("  --help    显示帮助信息")
	fmt.Println()
	fmt.Println("环境变量:")
	fmt.Println("  KOOK_BOT_TOKEN    Bot Token（必填）")
	fmt.Println("  KOOK_GUILD_ID     服务器 ID（创建/删除频道时需要）")
	fmt.Println("  KOOK_CHANNEL_ID   频道 ID（编辑/获取详情时需要）")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  export KOOK_BOT_TOKEN=\"your-bot-token\"")
	fmt.Println("  export KOOK_GUILD_ID=\"your-guild-id\"")
	fmt.Println("  go run examples/channel/main.go")
}

func runTests(ctx context.Context, client *kook.Client) {
	// 获取频道列表
	testListChannels(ctx, client)

	// 获取频道详情
	testGetChannel(ctx, client)

	// 创建频道
	testCreateChannel(ctx, client)

	// 编辑频道
	testUpdateChannel(ctx, client)

	// 删除频道
	testDeleteChannel(ctx, client)
}

func testListChannels(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取频道列表 ---")

	guildID := os.Getenv("KOOK_GUILD_ID")
	if guildID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_GUILD_ID")
		return
	}

	channels, err := client.ListChannels(ctx, guildID)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 频道数量: %d\n", len(channels.Items))
	for i, ch := range channels.Items {
		typeName := "文字"
		if ch.Type == 2 {
			typeName = "语音"
		}
		fmt.Printf("     [%d] %s (ID: %s, 类型: %s)\n", i+1, ch.Name, ch.ID, typeName)
	}
}

func testGetChannel(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 获取频道详情 ---")

	channelID := os.Getenv("KOOK_CHANNEL_ID")
	if channelID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_CHANNEL_ID")
		return
	}

	channel, err := client.GetChannel(ctx, channelID)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 频道名称: %s\n", channel.Name)
	fmt.Printf("     ID: %s\n", channel.ID)
	fmt.Printf("     类型: %d\n", channel.Type)
	fmt.Printf("     主题: %s\n", channel.Topic)
	fmt.Printf("     服务器 ID: %s\n", channel.GuildID)
}

func testCreateChannel(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 创建频道 ---")

	guildID := os.Getenv("KOOK_GUILD_ID")
	if guildID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_GUILD_ID")
		return
	}

	// 创建文字频道
	channel, err := client.CreateChannel(ctx, guildID, "sdk-test-channel", 1, "", 0, "", 0)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 频道已创建: %s (ID: %s)\n", channel.Name, channel.ID)
	fmt.Printf("     请记录此 channelID 用于后续测试: %s\n", channel.ID)
}

func testUpdateChannel(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 编辑频道 ---")

	channelID := os.Getenv("KOOK_CHANNEL_ID")
	if channelID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_CHANNEL_ID")
		return
	}

	channel, err := client.UpdateChannel(ctx, channelID, "sdk-test-channel-updated", "SDK 测试主题", model.SlowMode5s)
	if err != nil {
		fmt.Printf("[FAIL] %v\n", err)
		return
	}
	fmt.Printf("[OK] 频道已更新: %s\n", channel.Name)
	fmt.Printf("     主题: %s\n", channel.Topic)
}

func testDeleteChannel(ctx context.Context, client *kook.Client) {
	fmt.Println("\n--- 删除频道 ---")

	channelID := os.Getenv("KOOK_CHANNEL_ID")
	if channelID == "" {
		fmt.Println("[SKIP] 需要设置 KOOK_CHANNEL_ID")
		return
	}

	fmt.Println("[SKIP] 此操作会删除频道，请谨慎使用")
	fmt.Println("       如需测试，请取消注释代码")
	// err := client.DeleteChannel(ctx, channelID)
	// if err != nil {
	//     fmt.Printf("[FAIL] %v\n", err)
	//     return
	// }
	// fmt.Println("[OK] 频道已删除")
}
