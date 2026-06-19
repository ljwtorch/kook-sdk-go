package kook

import "net/http"

// Option 是配置 Client 的函数类型（Functional Options 模式）。
// 通过向 NewClient 传递 Option 函数，可以灵活地定制客户端行为。
type Option func(*Client)

// WithHTTPClient 设置自定义的 HTTP 客户端。
// 如果未设置，默认使用 http.DefaultClient。
//
// 使用示例:
//
//	client := kook.NewClient("your-token", kook.WithHTTPClient(&http.Client{
//	    Timeout: 30 * time.Second,
//	}))
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

// WithBaseURL 设置 API 的基础 URL。
// 默认值为 "https://www.kookapp.cn/api/v3"。
// 此选项主要用于测试或对接自建网关。
//
// 使用示例:
//
//	client := kook.NewClient("your-token", kook.WithBaseURL("http://localhost:8080/api/v3"))
func WithBaseURL(url string) Option {
	return func(c *Client) {
		if url != "" {
			c.baseURL = url
		}
	}
}

// WithUserAgent 设置 HTTP 请求的 User-Agent 头。
// 默认值为 "KookBotGo/0.1.0"。
//
// 使用示例:
//
//	client := kook.NewClient("your-token", kook.WithUserAgent("MyBot/1.0"))
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		if ua != "" {
			c.userAgent = ua
		}
	}
}

// WithDebug 启用或禁用调试模式。
// 启用后，客户端会将请求和响应的详细信息输出到标准日志。
//
// 使用示例:
//
//	client := kook.NewClient("your-token", kook.WithDebug(true))
func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.debug = debug
	}
}

// WithRateLimiter 设置自定义的速率限制器。
// 如果未设置，默认使用 NewRateLimiter() 创建一个新的限制器。
//
// 使用示例:
//
//	rl := kook.NewRateLimiter()
//	client := kook.NewClient("your-token", kook.WithRateLimiter(rl))
func WithRateLimiter(rl *RateLimiter) Option {
	return func(c *Client) {
		if rl != nil {
			c.rateLimiter = rl
		}
	}
}

// WithAPIVersion 设置 API 版本。
// 默认值为 "v3"。
// 此选项用于指定使用哪个版本的 KOOK API。
//
// 使用示例:
//
//	client := kook.NewClient("your-token", kook.WithAPIVersion("v4"))
func WithAPIVersion(version string) Option {
	return func(c *Client) {
		if version != "" {
			c.apiVersion = version
		}
	}
}
