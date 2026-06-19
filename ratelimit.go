package kook

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// RateLimiter 管理 KOOK API 的速率限制。
// 它通过解析 HTTP 响应头中的 Rate Limit 信息，
// 在请求发送前自动判断是否需要等待。
// RateLimiter 是并发安全的。
type RateLimiter struct {
	mu      sync.Mutex
	buckets map[string]*bucket
}

// bucket 存储单个路由的速率限制状态。
type bucket struct {
	mu        sync.Mutex
	remaining int
	resetAt   time.Time
	limit     int
}

// NewRateLimiter 创建一个新的速率限制器实例。
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		buckets: make(map[string]*bucket),
	}
}

// Wait 等待直到可以发送指定路由的请求。
// 如果当前路由的剩余请求数为 0 且重置时间未到，则阻塞等待。
// 如果 ctx 被取消，立即返回 ctx.Err()。
func (rl *RateLimiter) Wait(ctx context.Context, key string) error {
	rl.mu.Lock()
	b, ok := rl.buckets[key]
	if !ok {
		b = &bucket{remaining: 1}
		rl.buckets[key] = b
	}
	rl.mu.Unlock()

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.remaining > 0 {
		return nil
	}

	now := time.Now()
	if b.resetAt.Before(now) {
		return nil
	}

	wait := b.resetAt.Sub(now)
	timer := time.NewTimer(wait)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// Update 根据 HTTP 响应头中的 Rate Limit 信息更新对应路由的限流状态。
// 如果路由对应的 bucket 不存在，则创建一个新的。
//
// KOOK API 的 Rate Limit 响应头:
//   - X-Rate-Limit-Limit: 窗口内允许的最大请求数
//   - X-Rate-Limit-Remaining: 窗口内剩余的请求数
//   - X-Rate-Limit-Reset: 窗口重置的 Unix 时间戳（秒）
//   - X-Rate-Limit-Bucket: 限流桶标识（可选）
func (rl *RateLimiter) Update(key string, resp *http.Response) {
	if resp == nil {
		return
	}

	limit := parseIntHeader(resp.Header, "X-Rate-Limit-Limit")
	remaining := parseIntHeader(resp.Header, "X-Rate-Limit-Remaining")
	reset := parseIntHeader(resp.Header, "X-Rate-Limit-Reset")

	rl.mu.Lock()
	b, ok := rl.buckets[key]
	if !ok {
		b = &bucket{}
		rl.buckets[key] = b
	}
	rl.mu.Unlock()

	b.mu.Lock()
	defer b.mu.Unlock()

	if limit > 0 {
		b.limit = limit
	}
	b.remaining = remaining
	if reset > 0 {
		b.resetAt = time.Unix(int64(reset), 0)
	}
}

// parseIntHeader 从 HTTP 头中解析整数值，解析失败返回 0。
func parseIntHeader(h http.Header, key string) int {
	v := h.Get(key)
	if v == "" {
		return 0
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return n
}
