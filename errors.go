package kook

import (
	"errors"
	"fmt"
)

// APIError 表示 KOOK API 返回的业务错误。
// 当 API 响应的 code 字段非 0 时，表示业务层面出错，
// 此时 Do 方法会将此错误返回给调用方。
type APIError struct {
	// Code 是 KOOK API 返回的错误码。
	Code int `json:"code"`
	// Message 是 KOOK API 返回的错误描述。
	Message string `json:"message"`
}

// Error 实现 error 接口，返回格式化的错误信息。
func (e *APIError) Error() string {
	return fmt.Sprintf("kook: API error (code=%d): %s", e.Code, e.Message)
}

// Is 支持 errors.Is 比较。当目标错误也是 *APIError 且 Code 相同时返回 true。
func (e *APIError) Is(target error) bool {
	var apiErr *APIError
	if errors.As(target, &apiErr) {
		return e.Code == apiErr.Code
	}
	return false
}

// 预定义的常见 API 错误。
// 可通过 errors.Is 进行判断：
//
//	if errors.Is(err, ErrUnauthorized) { ... }
var (
	// ErrUnauthorized 表示认证失败（token 无效或过期）。
	ErrUnauthorized = &APIError{Code: 40100, Message: "unauthorized"}
	// ErrForbidden 表示无权限执行该操作。
	ErrForbidden = &APIError{Code: 40300, Message: "forbidden"}
	// ErrNotFound 表示请求的资源不存在。
	ErrNotFound = &APIError{Code: 40400, Message: "not found"}
	// ErrRateLimited 表示请求频率超过限制。
	ErrRateLimited = &APIError{Code: 42900, Message: "rate limited"}
)

// IsAPIError 检查 error 是否为特定 code 的 APIError。
// 用法示例:
//
//	if IsAPIError(err, 40100) {
//	    // 处理认证错误
//	}
func IsAPIError(err error, code int) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code == code
	}
	return false
}

// HTTPError 表示 HTTP 层面的错误（非 API 业务错误）。
// 当 HTTP 响应状态码非 2xx，且无法解析为 APIError 时使用此类型。
type HTTPError struct {
	// StatusCode 是 HTTP 响应状态码。
	StatusCode int
	// Body 是响应体的原始内容。
	Body string
}

// Error 实现 error 接口，返回格式化的 HTTP 错误信息。
func (e *HTTPError) Error() string {
	return fmt.Sprintf("kook: HTTP error (status=%d): %s", e.StatusCode, e.Body)
}
