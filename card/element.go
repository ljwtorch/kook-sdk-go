package card

// Element 卡片元素
type Element struct {
	Type      string   `json:"type"`
	Content   string   `json:"content,omitempty"`
	Src       string   `json:"src,omitempty"`
	Text      *Element `json:"text,omitempty"`
	Size      string   `json:"size,omitempty"`
	Click     string   `json:"click,omitempty"`
	Value     string   `json:"value,omitempty"`
	Target    string   `json:"target,omitempty"`
	Emoji     *Element `json:"emoji,omitempty"`
	Alt       string   `json:"alt,omitempty"`
	Href      string   `json:"href,omitempty"`
	Mode      string   `json:"mode,omitempty"`
	StartTime int64    `json:"startTime,omitempty"`
	EndTime   int64    `json:"endTime,omitempty"`
	Theme     string   `json:"theme,omitempty"`
}

// 元素类型常量
const (
	ElementTypePlainText = "plain-text"
	ElementTypeKMarkdown = "kmarkdown"
	ElementTypeImage     = "image"
	ElementTypeButton    = "button"
)

// Image Size 常量
const (
	ImageSizeSmall = "sm"
	ImageSizeLarge = "lg"
)

// Button Click 常量
const (
	ButtonClickReturnVal = "return-val"
	ButtonClickLink      = "link"
)

// PlainText 创建纯文本元素
func PlainText(content string) Element {
	return Element{
		Type:    ElementTypePlainText,
		Content: content,
	}
}

// KMarkdown 创建 KMarkdown 文本元素
func KMarkdown(content string) Element {
	return Element{
		Type:    ElementTypeKMarkdown,
		Content: content,
	}
}

// Image 创建图片元素
//
// src 为图片地址，alt 为替代文本，size 可选 ImageSizeSmall 或 ImageSizeLarge。
func Image(src string, alt string, size string) Element {
	return Element{
		Type: ElementTypeImage,
		Src:  src,
		Alt:  alt,
		Size: size,
	}
}

// Button 创建按钮元素
//
// click 决定点击行为（ButtonClickReturnVal 或 ButtonClickLink），
// value 为返回值或链接地址，text 为按钮文字元素，theme 为按钮主题色。
func Button(click string, value string, text Element, theme string) Element {
	return Element{
		Type:  ElementTypeButton,
		Click: click,
		Value: value,
		Text:  &text,
		Theme: theme,
	}
}
