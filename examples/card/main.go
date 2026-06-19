// 卡片消息示例：展示卡片消息构建器的使用方法
//
// 环境变量：
//   - KOOK_BOT_TOKEN:    Bot Token（必填）
//   - KOOK_CHANNEL_ID:   频道 ID（必填）
//
// 运行方式：
//
//	go run examples/card/main.go
//	go run examples/card/main.go --help
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	kook "github.com/ljwtorch/kook-sdk-go"
	"github.com/ljwtorch/kook-sdk-go/card"
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

	channelID := os.Getenv("KOOK_CHANNEL_ID")
	if channelID == "" {
		log.Fatal("请设置 KOOK_CHANNEL_ID 环境变量")
	}

	client := kook.NewClient(token, kook.WithDebug(false))
	defer client.Close()

	ctx := context.Background()

	fmt.Println("========== 卡片消息测试 ==========")
	runTests(ctx, client, channelID)
	fmt.Println("\n========== 测试完成 ==========")
}

func printUsage() {
	fmt.Println("卡片消息示例")
	fmt.Println()
	fmt.Println("用法: go run examples/card/main.go [options]")
	fmt.Println()
	fmt.Println("选项:")
	fmt.Println("  --help    显示帮助信息")
	fmt.Println()
	fmt.Println("环境变量:")
	fmt.Println("  KOOK_BOT_TOKEN    Bot Token（必填）")
	fmt.Println("  KOOK_CHANNEL_ID   频道 ID（必填）")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  export KOOK_BOT_TOKEN=\"your-bot-token\"")
	fmt.Println("  export KOOK_CHANNEL_ID=\"your-channel-id\"")
	fmt.Println("  go run examples/card/main.go")
}

func runTests(ctx context.Context, client *kook.Client, channelID string) {
	// 基础卡片
	testBasicCard(ctx, client, channelID)
	time.Sleep(time.Second)

	// 带按钮的卡片
	testButtonCard(ctx, client, channelID)
	time.Sleep(time.Second)

	// 多卡片消息
	testMultiCard(ctx, client, channelID)
	time.Sleep(time.Second)

	// 所有主题展示
	testAllThemes(ctx, client, channelID)
	time.Sleep(time.Second)

	// 带图片的卡片
	testImageCard(ctx, client, channelID)
}

func testBasicCard(ctx context.Context, client *kook.Client, channelID string) {
	fmt.Println("\n--- 基础卡片 ---")

	builder := card.New()
	builder.Card(card.ThemePrimary, card.SizeLarge).
		Header("通知标题").
		Section("这是一条通知消息内容").
		Divider().
		Section("**重要**: 请查看详情").
		End()

	sendCard(ctx, client, channelID, "基础卡片", builder)
}

func testButtonCard(ctx context.Context, client *kook.Client, channelID string) {
	fmt.Println("\n--- 带按钮的卡片 ---")

	builder := card.New()
	builder.Card(card.ThemeWarning, card.SizeLarge).
		Header("操作面板").
		Section("请选择要执行的操作：").
		ActionGroup(card.ActionGroupModeLeft,
			card.Button(card.ButtonClickReturnVal, "deploy", card.PlainText("部署应用"), card.ThemeSuccess),
			card.Button(card.ButtonClickReturnVal, "restart", card.PlainText("重启服务"), card.ThemeDanger),
			card.Button(card.ButtonClickReturnVal, "status", card.PlainText("查看状态"), card.ThemeInfo),
		).
		Divider().
		ActionGroup(card.ActionGroupModeLeft,
			card.Button(card.ButtonClickLink, "https://kookapp.cn", card.PlainText("访问官网"), card.ThemePrimary),
		).
		End()

	sendCard(ctx, client, channelID, "按钮卡片", builder)
}

func testMultiCard(ctx context.Context, client *kook.Client, channelID string) {
	fmt.Println("\n--- 多卡片消息 ---")

	builder := card.New()

	// 卡片 1：状态报告
	builder.Card(card.ThemeInfo, card.SizeLarge).
		Header("服务器状态报告").
		Section("**CPU**: ████████░░ 82%\n**内存**: ██████░░░░ 64%\n**磁盘**: █████████░ 91%").
		Divider().
		Section("**在线人数**: 128\n**活跃频道**: 12\n**消息/分钟**: 45").
		End()

	// 卡片 2：告警信息
	builder.Card(card.ThemeDanger, card.SizeSmall).
		Header("告警信息").
		Section("**磁盘使用率超过 90%**，建议尽快清理。").
		ActionGroup(card.ActionGroupModeLeft,
			card.Button(card.ButtonClickReturnVal, "cleanup", card.PlainText("一键清理"), card.ThemeWarning),
			card.Button(card.ButtonClickReturnVal, "ignore", card.PlainText("忽略告警"), card.ThemeSecondary),
		).
		End()

	sendCard(ctx, client, channelID, "多卡片", builder)
}

func testAllThemes(ctx context.Context, client *kook.Client, channelID string) {
	fmt.Println("\n--- 所有主题展示 ---")

	themes := []struct {
		theme string
		name  string
	}{
		{card.ThemePrimary, "Primary"},
		{card.ThemeSuccess, "Success"},
		{card.ThemeDanger, "Danger"},
		{card.ThemeWarning, "Warning"},
		{card.ThemeInfo, "Info"},
		{card.ThemeSecondary, "Secondary"},
		{card.ThemeNone, "None"},
	}

	for _, t := range themes {
		builder := card.New()
		builder.Card(t.theme, card.SizeLarge).
			Header(fmt.Sprintf("%s 主题", t.name)).
			Section(fmt.Sprintf("这是 %s 主题的卡片消息", t.name)).
			End()

		sendCard(ctx, client, channelID, "Theme_"+t.name, builder)
		time.Sleep(time.Second)
	}
}

func testImageCard(ctx context.Context, client *kook.Client, channelID string) {
	fmt.Println("\n--- 带图片的卡片 ---")

	// 使用 KOOK 官方示例图片
	imageURL := "https://img.kaiheila.cn/assets/2021-01/pWsmcLsPJq08c08c.jpeg"

	builder := card.New()
	builder.Card(card.ThemeSecondary, card.SizeLarge).
		Header("图片展示").
		Section("以下是一张示例图片：").
		Divider().
		ImageGroup(
			card.Image(imageURL, "示例图片", card.ImageSizeLarge),
		).
		End()

	sendCard(ctx, client, channelID, "图片卡片", builder)
}

func sendCard(ctx context.Context, client *kook.Client, channelID, name string, builder *card.Builder) {
	cardJSON := builder.Build()
	msg, err := client.SendMessage(ctx, channelID, cardJSON)
	if err != nil {
		fmt.Printf("[%s] 发送失败: %v\n", name, err)
		return
	}
	fmt.Printf("[%s] 发送成功 (ID: %s)\n", name, msg.MsgID)
}
