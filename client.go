package kook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"

	"github.com/ljwtorch/kook-sdk-go/internal"
)

const (
	// defaultUserAgent 是默认的 User-Agent 头。
	defaultUserAgent = "KookBotGo/" + Version
)

// Client 是 KOOK SDK 的核心客户端。
// 它封装了 HTTP 请求、速率限制、WebSocket 连接等功能，
// 是调用 KOOK API 和接收事件的入口。
// Client 是并发安全的，可以被多个 goroutine 同时使用。
type Client struct {
	// token 是 Bot 的认证令牌。
	token string
	// baseURL 是 API 的基础 URL。
	baseURL string
	// apiVersion 是 API 版本。
	apiVersion string
	// userAgent 是 HTTP 请求的 User-Agent 头。
	userAgent string
	// debug 控制是否输出调试日志。
	debug bool
	// httpClient 是底层的 HTTP 客户端。
	httpClient *http.Client
	// rateLimiter 管理 API 请求的速率限制。
	rateLimiter *RateLimiter
	// gateway 管理 WebSocket 连接（在 gateway.go 中定义完整逻辑）。
	gateway *Gateway
	// mu 保护 connected 等字段的并发读写。
	mu sync.RWMutex
	// connected 标记 WebSocket 是否已连接。
	connected bool
}

// NewClient 创建一个新的 KOOK SDK 客户端。
// token 参数为 Bot 的认证令牌，opts 为可选的配置项。
//
// 使用示例:
//
//	client := kook.NewClient("your-bot-token")
//	client := kook.NewClient("your-bot-token", kook.WithDebug(true))
func NewClient(token string, opts ...Option) *Client {
	c := &Client{
		token:       token,
		baseURL:     DefaultBaseURL,
		apiVersion:  DefaultAPIVersion,
		userAgent:   defaultUserAgent,
		httpClient:  http.DefaultClient,
		rateLimiter: NewRateLimiter(),
	}

	for _, opt := range opts {
		opt(c)
	}

	c.gateway = &Gateway{
		client:   c,
		handlers: make(map[string][]EventHandler),
		done:     make(chan struct{}),
	}

	return c
}

// Get 发送 GET 请求到指定的 API 路径。
// body 参数如果非 nil，会被转换为 URL 查询参数。
// result 参数用于接收响应中的 data 字段，如果为 nil 则忽略响应数据。
func (c *Client) Get(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Do(ctx, http.MethodGet, path, body, result)
}

// Post 发送 POST 请求到指定的 API 路径。
// body 参数会被序列化为 JSON 作为请求体。
// result 参数用于接收响应中的 data 字段，如果为 nil 则忽略响应数据。
func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Do(ctx, http.MethodPost, path, body, result)
}

// Delete 发送 DELETE 请求到指定的 API 路径。
// body 参数如果非 nil，会被转换为 URL 查询参数。
// result 参数用于接收响应中的 data 字段，如果为 nil 则忽略响应数据。
func (c *Client) Delete(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Do(ctx, http.MethodDelete, path, body, result)
}

// Do 执行一个 API 请求，这是 Client 的核心方法。
// 所有 API 调用最终都通过此方法发送。
//
// 处理流程:
//  1. GET/DELETE 请求：将 body 转为查询参数附加到 URL
//  2. POST/PUT 请求：将 body 序列化为 JSON 请求体
//  3. 自动添加 Authorization: Bot {token} 认证头
//  4. 发送请求并处理速率限制（429 时自动等待重试）
//  5. 解析响应 JSON，检查 code 是否为 0
//  6. 如果 result 不为 nil，将 data 反序列化到 result
func (c *Client) Do(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	// 构建查询参数（GET/DELETE）或请求体（POST/PUT）
	var (
		queryParams url.Values
		reqBody     io.Reader
	)

	method = strings.ToUpper(method)

	if body != nil {
		switch method {
		case http.MethodGet, http.MethodDelete:
			queryParams = bodyToQuery(body)
		default:
			data, err := internal.MarshalBody(body)
			if err != nil {
				return fmt.Errorf("kook: failed to marshal request body: %w", err)
			}
			reqBody = bytes.NewReader(data)
		}
	}

	fullURL := internal.BuildURL(c.baseURL, c.apiVersion, path, queryParams)

	if c.debug {
		log.Printf("[KOOK DEBUG] %s %s", method, fullURL)
	}

	// 速率限制等待
	rateKey := method + ":" + path
	if err := c.rateLimiter.Wait(ctx, rateKey); err != nil {
		return fmt.Errorf("kook: rate limit wait: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequestWithContext(ctx, method, fullURL, reqBody)
	if err != nil {
		return fmt.Errorf("kook: failed to create request: %w", err)
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bot "+c.token)
	req.Header.Set("User-Agent", c.userAgent)
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("kook: request failed: %w", err)
	}
	defer resp.Body.Close()

	// 更新速率限制信息
	c.rateLimiter.Update(rateKey, resp)

	// 处理 429 Too Many Requests
	if resp.StatusCode == http.StatusTooManyRequests {
		return c.handleRateLimit(ctx, method, path, body, result, rateKey)
	}

	// 读取响应体
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("kook: failed to read response: %w", err)
	}

	if c.debug {
		log.Printf("[KOOK DEBUG] Response status=%d body=%s", resp.StatusCode, string(respData))
	}

	// 检查 HTTP 状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Body:       string(respData),
		}
	}

	// 解析响应
	code, message, rawData, err := internal.ParseResponse(respData)
	if err != nil {
		return err
	}

	// 检查业务错误码
	if code != 0 {
		return &APIError{Code: code, Message: message}
	}

	// 反序列化 data 到 result
	if result != nil && rawData != nil {
		if err := json.Unmarshal(rawData, result); err != nil {
			return fmt.Errorf("kook: failed to unmarshal response data: %w", err)
		}
	}

	return nil
}

// handleRateLimit 处理 HTTP 429 速率限制响应。
// 它会等待响应头中指定的重试时间后重新发起请求。
func (c *Client) handleRateLimit(ctx context.Context, method, path string, body interface{}, result interface{}, rateKey string) error {
	// 等待一段时间后重试，使用 rate limiter 中记录的重置时间
	if err := c.rateLimiter.Wait(ctx, rateKey); err != nil {
		return fmt.Errorf("kook: rate limit retry wait: %w", err)
	}

	if c.debug {
		log.Printf("[KOOK DEBUG] Retrying after rate limit: %s %s", method, path)
	}

	return c.Do(ctx, method, path, body, result)
}

// DoMultipart 执行一个 multipart/form-data 文件上传请求。
// fields 为额外的表单字段，fieldName 为文件字段名，
// fileName 为文件名，file 为文件内容的 Reader。
// result 参数用于接收响应中的 data 字段。
func (c *Client) DoMultipart(ctx context.Context, path string, fields map[string]string, fieldName string, fileName string, file io.Reader, result interface{}) error {
	fullURL := internal.BuildURL(c.baseURL, c.apiVersion, path, nil)

	// 速率限制等待
	rateKey := "POST:" + path
	if err := c.rateLimiter.Wait(ctx, rateKey); err != nil {
		return fmt.Errorf("kook: rate limit wait: %w", err)
	}

	// 构建 multipart 请求体
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// 添加文件字段
	part, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		return fmt.Errorf("kook: failed to create form file: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return fmt.Errorf("kook: failed to copy file content: %w", err)
	}

	// 添加额外的表单字段
	for k, v := range fields {
		if err := writer.WriteField(k, v); err != nil {
			return fmt.Errorf("kook: failed to write field %q: %w", k, err)
		}
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("kook: failed to close multipart writer: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, &buf)
	if err != nil {
		return fmt.Errorf("kook: failed to create request: %w", err)
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bot "+c.token)
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("kook: request failed: %w", err)
	}
	defer resp.Body.Close()

	// 更新速率限制信息
	c.rateLimiter.Update(rateKey, resp)

	// 处理 429
	if resp.StatusCode == http.StatusTooManyRequests {
		if err := c.rateLimiter.Wait(ctx, rateKey); err != nil {
			return fmt.Errorf("kook: rate limit retry wait: %w", err)
		}
		// 注意：multipart 请求体无法重新读取，因此对于文件上传的 429 不自动重试
		return &APIError{Code: 42900, Message: "rate limited on multipart request, please retry manually"}
	}

	// 读取响应体
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("kook: failed to read response: %w", err)
	}

	// 检查 HTTP 状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Body:       string(respData),
		}
	}

	// 解析响应
	code, message, rawData, err := internal.ParseResponse(respData)
	if err != nil {
		return err
	}

	if code != 0 {
		return &APIError{Code: code, Message: message}
	}

	if result != nil && rawData != nil {
		if err := json.Unmarshal(rawData, result); err != nil {
			return fmt.Errorf("kook: failed to unmarshal response data: %w", err)
		}
	}

	return nil
}

// bodyToQuery 将 struct 或 map 类型的 body 转换为 URL 查询参数。
// 对于 struct 类型，使用 JSON tag 作为参数名；
// 对于 map 类型，直接将键值对转为查询参数。
// 零值字段会被忽略。
func bodyToQuery(body interface{}) url.Values {
	params := url.Values{}
	if body == nil {
		return params
	}

	v := reflect.ValueOf(body)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return params
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			if !val.IsZero() {
				params.Set(fmt.Sprintf("%v", key.Interface()), fmt.Sprintf("%v", val.Interface()))
			}
		}
	case reflect.Struct:
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			if !field.IsExported() {
				continue
			}
			fieldVal := v.Field(i)
			if fieldVal.IsZero() {
				continue
			}
			// 使用 json tag 作为参数名
			name := field.Tag.Get("json")
			if name == "" || name == "-" {
				name = field.Name
			}
			// 去除 omitempty 等选项
			if idx := strings.Index(name, ","); idx != -1 {
				name = name[:idx]
			}
			if name == "" {
				continue
			}
			params.Set(name, fmt.Sprintf("%v", fieldVal.Interface()))
		}
	}

	return params
}

// Token 返回 Bot Token，实现 HTTPRequester 接口。
// api 包中的函数通过此方法获取 Token 用于认证。
func (c *Client) Token() string {
	return c.token
}

// BaseURL 返回 API 基础 URL，实现 HTTPRequester 接口。
func (c *Client) BaseURL() string {
	return c.baseURL
}

// Close 关闭 WebSocket 连接并释放相关资源。
// 如果当前未连接，则不执行任何操作。
func (c *Client) Close() error {
	c.mu.Lock()
	c.connected = false
	c.mu.Unlock()
	return c.Disconnect()
}
