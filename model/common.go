// Package model 定义了 KOOK API 的数据模型。
package model

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// FlexInt 是一个灵活的整数类型，可以同时处理 JSON 中的数字和字符串格式。
// KOOK API 的某些字段会根据情况返回数字或字符串，此类型用于兼容两种格式。
type FlexInt int

// UnmarshalJSON 实现 json.Unmarshaler 接口，支持从数字或字符串反序列化。
func (f *FlexInt) UnmarshalJSON(data []byte) error {
	// 尝试作为数字解析
	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		*f = FlexInt(num)
		return nil
	}
	// 尝试作为字符串解析
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if str == "" {
			*f = 0
			return nil
		}
		n, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		*f = FlexInt(n)
		return nil
	}
	return &json.UnmarshalTypeError{Value: string(data), Type: reflect.TypeOf(f)}
}

// APIResponse 是 KOOK API 的标准响应。
// 所有 API 请求都返回此结构，其中 Data 字段包含实际的响应数据。
type APIResponse struct {
	// Code 是响应状态码，0 表示成功
	Code int `json:"code"`
	// Message 是响应消息
	Message string `json:"message"`
	// Data 是实际的响应数据，使用 json.RawMessage 延迟解析
	Data json.RawMessage `json:"data"`
}

// PageMeta 是分页查询的元数据。
type PageMeta struct {
	// Page 是当前页码
	Page int `json:"page"`
	// PageTotal 是总页数
	PageTotal int `json:"page_total"`
	// PageSize 是每页大小
	PageSize int `json:"page_size"`
	// Total 是总记录数
	Total int `json:"total"`
}

// PageResult 是分页查询的结果。
type PageResult[T any] struct {
	// Items 是当前页的数据列表
	Items []T `json:"items"`
	// Meta 是分页元数据
	Meta PageMeta `json:"meta"`
	// Sort 是排序信息
	Sort map[string]string `json:"sort,omitempty"`
}

// ListResult 是简单列表查询的结果。
type ListResult[T any] struct {
	// Items 是数据列表
	Items []T `json:"items"`
	// Meta 是分页元数据
	Meta PageMeta `json:"meta"`
	// Sort 是排序信息
	Sort map[string]string `json:"sort,omitempty"`
}
