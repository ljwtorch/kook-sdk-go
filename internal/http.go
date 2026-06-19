package internal

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// BuildURL 构建完整的请求 URL。
// baseURL 为基础地址（如 https://www.kookapp.cn），
// apiVersion 为 API 版本（如 v3），
// path 为 API 路径（如 /guild/list），
// params 为可选的查询参数（仅用于 GET/DELETE 请求）。
func BuildURL(baseURL, apiVersion, path string, params url.Values) string {
	baseURL = strings.TrimRight(baseURL, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	// 自动拼接API版本前缀
	if !strings.HasPrefix(path, "/api/") {
		path = "/api/" + apiVersion + path
	}
	u := baseURL + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	return u
}

// MarshalBody 将请求体序列化为 JSON 字节切片。
// 如果 body 为 nil，返回 nil 且无错误。
// 如果 body 已经是 []byte 类型，直接返回。
func MarshalBody(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	if b, ok := body.([]byte); ok {
		return b, nil
	}
	return json.Marshal(body)
}

// ParseResponse 解析 KOOK API 的标准 JSON 响应。
// KOOK API 的响应格式为: {"code": 0, "message": "success", "data": {...}}
// 返回 code、message、原始 data JSON 和可能的解析错误。
func ParseResponse(data []byte) (code int, message string, rawData json.RawMessage, err error) {
	var resp struct {
		Code    int             `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
	}
	if err = json.Unmarshal(data, &resp); err != nil {
		return 0, "", nil, fmt.Errorf("kook: failed to parse response: %w", err)
	}
	return resp.Code, resp.Message, resp.Data, nil
}
