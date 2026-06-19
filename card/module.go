package card

// Module 卡片模块
type Module struct {
	Type      string    `json:"type"`
	Text      *Element  `json:"text,omitempty"`
	Elements  []Element `json:"elements,omitempty"`
	Accessory *Element  `json:"accessory,omitempty"`
	Mode      string    `json:"mode,omitempty"`
	Src       string    `json:"src,omitempty"`
	Title     string    `json:"title,omitempty"`
	Cover     string    `json:"cover,omitempty"`
	Size      int       `json:"size,omitempty"`
	FileType  string    `json:"fileType,omitempty"`
	StartTime int64     `json:"startTime,omitempty"`
	EndTime   int64     `json:"endTime,omitempty"`
	Code      string    `json:"code,omitempty"`
	Theme     string    `json:"theme,omitempty"`
	Value     string    `json:"value,omitempty"`
}

// 模块类型常量
const (
	ModuleTypeHeader      = "header"
	ModuleTypeSection     = "section"
	ModuleTypeImageGroup  = "image-group"
	ModuleTypeContainer   = "container"
	ModuleTypeActionGroup = "action-group"
	ModuleTypeContext     = "context"
	ModuleTypeDivider     = "divider"
	ModuleTypeFile        = "file"
	ModuleTypeAudio       = "audio"
	ModuleTypeVideo       = "video"
	ModuleTypeCountdown   = "countdown"
	ModuleTypeInvite      = "invite"
)

// Section Mode 常量
const (
	SectionModeLeft  = "left"
	SectionModeRight = "right"
)

// ActionGroup Mode 常量
const (
	ActionGroupModeLeft  = "left"
	ActionGroupModeRight = "right"
)

// Countdown Mode 常量
const (
	CountdownModeDay    = "day"
	CountdownModeHour   = "hour"
	CountdownModeSecond = "second"
)
