package emoji

import (
	"fmt"
	"strconv"
)

// ID 将 Unicode 字符转换为 KOOK emoji 数字标识。
// 例如：😆 -> "128518"
//
// KOOK 支持的完整表情对照表请参见：
// https://img.kookapp.cn/assets/emoji.json
func ID(char rune) string {
	return strconv.Itoa(int(char))
}

// Format 将 Unicode 字符转换为 KMarkdown emoji 格式。
// 例如：😆 -> "[#128518;]"
//
// KOOK 支持的完整表情对照表请参见：
// https://img.kookapp.cn/assets/emoji.json
func Format(char rune) string {
	return fmt.Sprintf("[#%d;]", char)
}

// ParseID 将 KOOK emoji 数字标识解析为 Unicode 字符。
// 例如："128518" -> 😆
//
// KOOK 支持的完整表情对照表请参见：
// https://img.kookapp.cn/assets/emoji.json
func ParseID(id string) (rune, error) {
	code, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("invalid emoji id %q: %w", id, err)
	}
	return rune(code), nil
}
