# 卡片消息

本文档介绍如何使用 KOOK SDK for Go 构建卡片消息。

## 目录

- [基本使用](#基本使用)
- [模块类型](#模块类型)
- [元素类型](#元素类型)
- [主题和尺寸](#主题和尺寸)
- [高级用法](#高级用法)
- [限制](#限制)

---

## 基本使用

```go
package main

import (
	"context"
	"log"

	kook "github.com/ljwtorch/kook-sdk-go"
	"github.com/ljwtorch/kook-sdk-go/card"
)

func main() {
	client := kook.NewClient("your-bot-token")
	ctx := context.Background()

	// 创建卡片构建器
	builder := card.New()

	// 构建卡片
	builder.Card(card.ThemePrimary, card.SizeLarge).
		Header("通知标题").
		Section("这是一条通知消息内容").
		Divider().
		Section("**重要**: 请查看详情").
		End()

	// 获取 JSON
	cardJSON := builder.Build()

	// 发送卡片消息
	msg, err := client.SendMessage(ctx, "channel-id", cardJSON)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("卡片已发送: %s", msg.MsgID)
}
```

---

## 模块类型

### Header - 标题模块

```go
builder.Header("标题文本")
```

### Section - 内容模块

```go
// 简单文本
builder.Section("普通文本")

// KMarkdown 文本
builder.Section("**粗体** *斜体* ~~删除线~~")

// 带附属元素的内容
builder.SectionWithAccessories("左侧文本",
	card.Button(card.ButtonClickReturnVal, "value", card.PlainText("按钮"), card.ThemePrimary),
)
```

### Divider - 分割线

```go
builder.Divider()
```

### ImageGroup - 图片组

```go
builder.ImageGroup(
	card.Image("https://example.com/img1.png", "图片1", card.ImageSizeLarge),
	card.Image("https://example.com/img2.png", "图片2", card.ImageSizeLarge),
)
```

### Container - 容器

```go
builder.Container(
	card.Image("https://example.com/img.png", "图片", card.ImageSizeLarge),
)
```

### ActionGroup - 交互组件组

```go
builder.ActionGroup(card.ActionGroupModeLeft,
	card.Button(card.ButtonClickReturnVal, "value1", card.PlainText("按钮1"), card.ThemePrimary),
	card.Button(card.ButtonClickReturnVal, "value2", card.PlainText("按钮2"), card.ThemeDanger),
)
```

### Context - 备注模块

```go
builder.Context(
	card.PlainText("备注信息"),
	card.Image("https://example.com/icon.png", "图标", card.ImageSizeSmall),
)
```

### File - 文件模块

```go
builder.File("https://example.com/file.pdf", "文件名", "file", 1024)
```

### Audio - 音频模块

```go
builder.Audio("https://example.com/audio.mp3", "音频标题", "https://example.com/cover.png")
```

### Video - 视频模块

```go
builder.Video("https://example.com/video.mp4", "视频标题", "https://example.com/cover.png")
```

### Countdown - 倒计时模块

```go
builder.Countdown(card.CountdownModeDay, startTime, endTime)
```

### Invite - 邀请模块

```go
builder.Invite("invite-code")
```

---

## 元素类型

### PlainText - 纯文本

```go
card.PlainText("普通文本")
```

### KMarkdown - KMarkdown 文本

```go
card.KMarkdown("**粗体** *斜体*")
```

### Image - 图片

```go
card.Image("https://example.com/img.png", "图片描述", card.ImageSizeLarge)
```

**图片尺寸：**
- `card.ImageSizeSmall` - 小图
- `card.ImageSizeLarge` - 大图
- `card.ImageSizeMedium` - 中图

### Button - 按钮

```go
card.Button(
	card.ButtonClickReturnVal,  // 点击类型
	"button-value",              // 值
	card.PlainText("按钮文本"),   // 文本
	card.ThemePrimary,           // 主题
)
```

**点击类型：**
- `card.ButtonClickReturnVal` - 返回值
- `card.ButtonClickLink` - 打开链接

**按钮主题：**
- `card.ThemePrimary` - 主要（蓝色）
- `card.ThemeSuccess` - 成功（绿色）
- `card.ThemeDanger` - 危险（红色）
- `card.ThemeWarning` - 警告（黄色）
- `card.ThemeInfo` - 信息（青色）
- `card.ThemeNone` - 无主题（灰色）
- `card.ThemeSecondary` - 次要

---

## 主题和尺寸

### 卡片主题

```go
builder.Card(card.ThemePrimary, card.SizeLarge)  // 蓝色主题
builder.Card(card.ThemeSuccess, card.SizeLarge)  // 绿色主题
builder.Card(card.ThemeDanger, card.SizeLarge)   // 红色主题
builder.Card(card.ThemeWarning, card.SizeLarge)  // 黄色主题
builder.Card(card.ThemeInfo, card.SizeLarge)     // 青色主题
builder.Card(card.ThemeNone, card.SizeLarge)     // 无主题
builder.Card(card.ThemeSecondary, card.SizeLarge) // 次要主题
```

### 卡片尺寸

```go
builder.Card(card.ThemePrimary, card.SizeLarge)  // 大尺寸
builder.Card(card.ThemePrimary, card.SizeMedium) // 中尺寸
builder.Card(card.ThemePrimary, card.SizeSmall)  // 小尺寸
```

---

## 高级用法

### 多张卡片

```go
builder := card.New()

// 第一张卡片
builder.Card(card.ThemePrimary, card.SizeLarge).
	Header("卡片 1").
	Section("内容 1").
	End()

// 第二张卡片
builder.Card(card.ThemeSuccess, card.SizeLarge).
	Header("卡片 2").
	Section("内容 2").
	End()

cardJSON := builder.Build()
```

### 复杂布局示例

```go
builder := card.New()

builder.Card(card.ThemeInfo, card.SizeLarge).
	// 标题
	Header("服务器状态报告").
	// 内容
	Section("**CPU**: 45%  \n**内存**: 62%  \n**磁盘**: 78%").
	// 分割线
	Divider().
	// 图片组
	ImageGroup(
		card.Image("https://example.com/cpu.png", "CPU 图表", card.ImageSizeLarge),
		card.Image("https://example.com/mem.png", "内存图表", card.ImageSizeLarge),
	).
	// 分割线
	Divider().
	// 交互按钮
	ActionGroup(card.ActionGroupModeLeft,
		card.Button(card.ButtonClickReturnVal, "refresh", card.PlainText("刷新"), card.ThemePrimary),
		card.Button(card.ButtonClickLink, "https://panel.example.com", card.PlainText("管理面板"), card.ThemeNone),
	).
	// 备注
	Context(
		card.PlainText("更新时间: 2024-01-01 12:00:00"),
	).
	End()

cardJSON := builder.Build()
```

### 带按钮的卡片

```go
builder := card.New()

builder.Card(card.ThemeWarning, card.SizeLarge).
	Header("快捷操作").
	Section("请选择要执行的操作：").
	ActionGroup(card.ActionGroupModeLeft,
		card.Button(card.ButtonClickReturnVal, "restart", card.PlainText("重启服务"), card.ThemeDanger),
		card.Button(card.ButtonClickReturnVal, "refresh", card.PlainText("刷新缓存"), card.ThemePrimary),
		card.Button(card.ButtonClickLink, "https://panel.example.com", card.PlainText("管理面板"), card.ThemeNone),
	).
	End()

cardJSON := builder.Build()
```

---

## 限制

- 单条消息最多 **5** 张卡片
- 单张卡片最多 **50** 个模块
- 按钮点击回调需要在 KOOK 开放平台配置

---

## 完整示例

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	kook "github.com/ljwtorch/kook-sdk-go"
	"github.com/ljwtorch/kook-sdk-go/card"
)

func main() {
	token := os.Getenv("KOOK_BOT_TOKEN")
	client := kook.NewClient(token)
	ctx := context.Background()

	// 构建卡片
	builder := card.New()

	// 卡片 1：信息展示
	builder.Card(card.ThemeInfo, card.SizeLarge).
		Header("系统通知").
		Section("这是一条系统通知消息，包含以下信息：").
		Divider().
		Section("**类型**: 系统更新\n**时间**: 2024-01-01\n**状态**: 已完成").
		End()

	// 卡片 2：操作按钮
	builder.Card(card.ThemePrimary, card.SizeSmall).
		Header("快捷操作").
		ActionGroup(card.ActionGroupModeLeft,
			card.Button(card.ButtonClickReturnVal, "confirm", card.PlainText("确认"), card.ThemeSuccess),
			card.Button(card.ButtonClickReturnVal, "cancel", card.PlainText("取消"), card.ThemeDanger),
		).
		End()

	// 发送
	cardJSON := builder.Build()
	msg, err := client.SendMessage(ctx, "channel-id", cardJSON)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("卡片已发送: %s\n", msg.MsgID)
}
```

更多示例请参考 `examples/card_message/` 目录。
