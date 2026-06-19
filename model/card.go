package model

// Card 表示 KOOK 卡片消息。
// 卡片消息由多个模块（Module）组成，支持丰富的消息展示形式。
type Card struct {
	// Type 是卡片类型，固定为 "card"
	Type string `json:"type"`
	// Theme 是卡片主题，可选值：primary、success、danger、warning、info、secondary、none
	Theme string `json:"theme"`
	// Size 是卡片大小，可选值：lg（大）、sm（小）
	Size string `json:"size"`
	// Modules 是卡片包含的模块列表
	Modules []Module `json:"modules"`
	// Color 是卡片边框颜色（十六进制），当 Theme 为 none 时生效
	Color string `json:"color,omitempty"`
}

// Module 表示卡片中的模块。
// 不同模块类型（type）决定展示形式，如 section、image-group、action-group 等。
type Module struct {
	// Type 是模块类型，如 section、text、image-group、action-group、divider、note、context、file、audio、video、countdown
	Type string `json:"type"`
	// Text 是模块的文本内容（部分模块使用）
	Text *Element `json:"text,omitempty"`
	// Elements 是模块包含的元素列表（部分模块使用）
	Elements []Element `json:"elements,omitempty"`
	// Mode 是文本展示模式，如 left、right（用于 section 模块）
	Mode string `json:"mode,omitempty"`
	// Value 是模块的值（如 countdown 模块的目标时间戳）
	Value string `json:"value,omitempty"`
	// Src 是媒体资源的 URL（如 image、video、audio 模块）
	Src string `json:"src,omitempty"`
}

// Element 表示卡片模块中的元素。
// 元素是卡片的最小组成单元，如文本、图片、按钮等。
type Element struct {
	// Type 是元素类型，如 plain-text、kmarkdown、image、button、paragraph
	Type string `json:"type"`
	// Content 是元素的文本内容（用于 plain-text、kmarkdown）
	Content string `json:"content,omitempty"`
	// Src 是图片/媒体资源的 URL（用于 image）
	Src string `json:"src,omitempty"`
	// Text 是按钮等元素上的文字内容
	Text *Element `json:"text,omitempty"`
	// Size 是元素大小，如 lg、sm
	Size string `json:"size,omitempty"`
	// Click 是点击行为，如 link（跳转链接）、return-val（返回值）
	Click string `json:"click,omitempty"`
	// Value 是点击行为对应的值（链接 URL 或返回值）
	Value string `json:"value,omitempty"`
	// Target 是链接打开方式，如 _blank（新窗口）
	Target string `json:"target,omitempty"`
	// Emoji 是表情元素
	Emoji *Element `json:"emoji,omitempty"`
	// Href 是链接地址
	Href string `json:"href,omitempty"`
}
