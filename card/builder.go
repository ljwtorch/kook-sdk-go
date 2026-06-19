package card

import "encoding/json"

// Theme 常量
const (
	ThemePrimary   = "primary"
	ThemeSuccess   = "success"
	ThemeDanger    = "danger"
	ThemeWarning   = "warning"
	ThemeInfo      = "info"
	ThemeSecondary = "secondary"
	ThemeNone      = "none"
)

// Size 常量
const (
	SizeLarge = "lg"
	SizeSmall = "sm"
)

// 卡片消息限制
const (
	MaxCardsPerMessage = 5  // 单条消息最多卡片数
	MaxModulesPerCard  = 50 // 单张卡片最多模块数
)

// Card 卡片
type Card struct {
	Type    string   `json:"type"`
	Theme   string   `json:"theme"`
	Size    string   `json:"size"`
	Modules []Module `json:"modules"`
	Color   string   `json:"color,omitempty"`
}

// Builder 提供链式调用来构建卡片消息。
//
// 使用示例:
//
//	builder := card.New()
//	builder.Card(card.ThemePrimary, card.SizeLarge).
//	    Header("标题").
//	    Section("内容").
//	    End()
//	jsonStr := builder.Build()
type Builder struct {
	cards []*Card
}

// New 创建新的卡片构建器
func New() *Builder {
	return &Builder{
		cards: make([]*Card, 0),
	}
}

// Card 开始一张新卡片，返回 CardBuilder 用于链式添加模块。
//
// theme 可选值: ThemePrimary, ThemeSuccess, ThemeDanger, ThemeWarning,
// ThemeInfo, ThemeSecondary, ThemeNone。
// size 可选值: SizeLarge, SizeSmall。
func (b *Builder) Card(theme string, size string) *CardBuilder {
	c := &Card{
		Type:    "card",
		Theme:   theme,
		Size:    size,
		Modules: make([]Module, 0),
	}
	b.cards = append(b.cards, c)
	return &CardBuilder{
		card:    c,
		builder: b,
	}
}

// Build 构建所有卡片，返回 JSON 字符串。
//
// 即使未调用 CardBuilder.End()，所有已添加的模块也会正确包含在输出中。
// 如果 JSON 序列化失败则返回空字符串。
func (b *Builder) Build() string {
	data, err := json.Marshal(b.cards)
	if err != nil {
		return ""
	}
	return string(data)
}

// BuildCards 构建所有卡片，返回 []Card 值切片。
func (b *Builder) BuildCards() []Card {
	result := make([]Card, len(b.cards))
	for i, c := range b.cards {
		result[i] = *c
	}
	return result
}

// CardBuilder 单张卡片构建器，通过链式调用添加模块。
type CardBuilder struct {
	card    *Card
	builder *Builder
}

// addModule 内部方法：将模块添加到当前卡片。
// 如果已达 MaxModulesPerCard 上限则忽略。
func (cb *CardBuilder) addModule(m Module) {
	if len(cb.card.Modules) >= MaxModulesPerCard {
		return
	}
	cb.card.Modules = append(cb.card.Modules, m)
}

// Header 添加标题模块。文本以 KMarkdown 渲染。
func (cb *CardBuilder) Header(text string) *CardBuilder {
	elem := PlainText(text)
	cb.addModule(Module{
		Type: ModuleTypeHeader,
		Text: &elem,
	})
	return cb
}

// Section 添加内容模块。文本以 KMarkdown 渲染。
func (cb *CardBuilder) Section(text string) *CardBuilder {
	elem := KMarkdown(text)
	cb.addModule(Module{
		Type: ModuleTypeSection,
		Text: &elem,
	})
	return cb
}

// SectionWithAccessory 添加带附属元素的内容模块。
//
// text 为正文内容（KMarkdown），accessory 为附属元素（只能为 image 或 button）。
// 注意：button 不能放置在左侧（mode 为 left 时无效）。
func (cb *CardBuilder) SectionWithAccessory(text string, mode string, accessory Element) *CardBuilder {
	elem := KMarkdown(text)
	cb.addModule(Module{
		Type:      ModuleTypeSection,
		Text:      &elem,
		Accessory: &accessory,
		Mode:      mode,
	})
	return cb
}

// SectionWithAccessories 添加带附属元素的内容模块（已废弃，请使用 SectionWithAccessory）。
//
// Deprecated: 使用 SectionWithAccessory 替代。
func (cb *CardBuilder) SectionWithAccessories(text string, elements ...Element) *CardBuilder {
	elem := KMarkdown(text)
	var accessory *Element
	if len(elements) > 0 {
		accessory = &elements[0]
	}
	cb.addModule(Module{
		Type:      ModuleTypeSection,
		Text:      &elem,
		Accessory: accessory,
	})
	return cb
}

// Divider 添加分割线模块
func (cb *CardBuilder) Divider() *CardBuilder {
	cb.addModule(Module{
		Type: ModuleTypeDivider,
	})
	return cb
}

// ImageGroup 添加图片组模块。
//
// elements 应为 Image 元素，最多 9 张图片。
func (cb *CardBuilder) ImageGroup(elements ...Element) *CardBuilder {
	cb.addModule(Module{
		Type:     ModuleTypeImageGroup,
		Elements: elements,
	})
	return cb
}

// Container 添加容器模块。
//
// 仅支持 Image 元素，图片会纵向排列。
func (cb *CardBuilder) Container(elements ...Element) *CardBuilder {
	cb.addModule(Module{
		Type:     ModuleTypeContainer,
		Elements: elements,
	})
	return cb
}

// ActionGroup 添加交互组件组模块。
//
// mode 可选 ActionGroupModeLeft 或 ActionGroupModeRight。
// elements 通常为 Button 元素。
func (cb *CardBuilder) ActionGroup(mode string, elements ...Element) *CardBuilder {
	cb.addModule(Module{
		Type:     ModuleTypeActionGroup,
		Mode:     mode,
		Elements: elements,
	})
	return cb
}

// Context 添加备注模块。
//
// 支持 plain-text、kmarkdown 和 image 元素。
func (cb *CardBuilder) Context(elements ...Element) *CardBuilder {
	cb.addModule(Module{
		Type:     ModuleTypeContext,
		Elements: elements,
	})
	return cb
}

// File 添加文件模块
//
// src 为文件地址，title 为文件名，type_ 为文件类型（如 "xls"），size 为文件大小（字节）。
func (cb *CardBuilder) File(src string, title string, type_ string, size int) *CardBuilder {
	cb.addModule(Module{
		Type:     ModuleTypeFile,
		Src:      src,
		Title:    title,
		FileType: type_,
		Size:     size,
	})
	return cb
}

// Audio 添加音频模块
//
// src 为音频地址，title 为音频标题，cover 为封面图地址。
func (cb *CardBuilder) Audio(src string, title string, cover string) *CardBuilder {
	cb.addModule(Module{
		Type:  ModuleTypeAudio,
		Src:   src,
		Title: title,
		Cover: cover,
	})
	return cb
}

// Video 添加视频模块
//
// src 为视频地址，title 为视频标题，cover 为封面图地址。
func (cb *CardBuilder) Video(src string, title string, cover string) *CardBuilder {
	cb.addModule(Module{
		Type:  ModuleTypeVideo,
		Src:   src,
		Title: title,
		Cover: cover,
	})
	return cb
}

// Countdown 添加倒计时模块（已废弃，请使用 CountdownDay/CountdownHour/CountdownSecond）
//
// Deprecated: 使用 CountdownDay、CountdownHour 或 CountdownSecond 替代。
func (cb *CardBuilder) Countdown(mode string, startTime int64, endTime int64) *CardBuilder {
	cb.addModule(Module{
		Type:      ModuleTypeCountdown,
		Mode:      mode,
		StartTime: startTime,
		EndTime:   endTime,
	})
	return cb
}

// CountdownDay 添加按天倒计时模块。
//
// endTime 为到期的毫秒时间戳（Unix milliseconds）。
// 注意：endTime 不能小于服务器当前时间戳。
func (cb *CardBuilder) CountdownDay(endTime int64) *CardBuilder {
	cb.addModule(Module{
		Type:    ModuleTypeCountdown,
		Mode:    CountdownModeDay,
		EndTime: endTime,
	})
	return cb
}

// CountdownHour 添加按小时倒计时模块。
//
// endTime 为到期的毫秒时间戳（Unix milliseconds）。
// 注意：endTime 不能小于服务器当前时间戳。
func (cb *CardBuilder) CountdownHour(endTime int64) *CardBuilder {
	cb.addModule(Module{
		Type:    ModuleTypeCountdown,
		Mode:    CountdownModeHour,
		EndTime: endTime,
	})
	return cb
}

// CountdownSecond 添加按秒倒计时模块。
//
// startTime 为起始的毫秒时间戳，endTime 为到期的毫秒时间戳（Unix milliseconds）。
// 注意：startTime 和 endTime 不能小于服务器当前时间戳。
func (cb *CardBuilder) CountdownSecond(startTime int64, endTime int64) *CardBuilder {
	cb.addModule(Module{
		Type:      ModuleTypeCountdown,
		Mode:      CountdownModeSecond,
		StartTime: startTime,
		EndTime:   endTime,
	})
	return cb
}

// Invite 添加邀请模块
//
// code 为邀请码。
func (cb *CardBuilder) Invite(code string) *CardBuilder {
	cb.addModule(Module{
		Type: ModuleTypeInvite,
		Code: code,
	})
	return cb
}

// AddModule 添加自定义模块
func (cb *CardBuilder) AddModule(m Module) *CardBuilder {
	cb.addModule(m)
	return cb
}

// End 结束当前卡片，返回 Builder 可以继续添加新卡片。
//
// 由于 Builder 内部使用指针引用，CardBuilder 上的所有修改会立即生效，
// 即使不调用 End() 也不会丢失数据。调用 End() 是一种良好的风格习惯，
// 使代码结构更清晰。
func (cb *CardBuilder) End() *Builder {
	return cb.builder
}
