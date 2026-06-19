package model

// 角色颜色常量。
// 颜色值使用 24 位 RGB 编码的十进制整数。
// 取值范围：0x000000 ~ 0xFFFFFF（十进制 0 ~ 16777215）。
//
// 计算公式：color = (R << 16) | (G << 8) | B
//
// 例如红色 (R=255, G=0, B=0)：(255 << 16) | (0 << 8) | 0 = 16711680
const (
	ColorRed    int = 16711680 // 红色: 0xFF0000
	ColorGreen  int = 65280    // 绿色: 0x00FF00
	ColorBlue   int = 255      // 蓝色: 0x0000FF
	ColorYellow int = 16776960 // 黄色: 0xFFFF00
	ColorWhite  int = 16777215 // 白色: 0xFFFFFF
	ColorBlack  int = 0        // 黑色: 0x000000
)

// Role 表示服务器中的角色。
type Role struct {
	// RoleID 是角色的唯一标识
	RoleID int64 `json:"role_id"`
	// Name 是角色名称
	Name string `json:"name"`
	// Color 是角色颜色值，使用 24 位 RGB 编码的十进制整数。
	// 取值范围：0x000000 ~ 0xFFFFFF（十进制 0 ~ 16777215）。
	// 计算公式：color = (R << 16) | (G << 8) | B
	//
	// 可使用预定义常量：ColorRed、ColorGreen、ColorBlue、ColorYellow、ColorWhite、ColorBlack
	Color int `json:"color"`
	// Position 是角色在层级中的位置
	Position int `json:"position"`
	// Permissions 是角色的权限位
	Permissions int64 `json:"permissions"`
}
