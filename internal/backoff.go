// Package internal 提供 KOOK SDK 的内部实现工具，不对外导出。
package internal

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

// Backoff 实现指数退避策略，用于 WebSocket 重连等场景。
type Backoff struct {
	// Min 是最小退避间隔。
	Min time.Duration
	// Max 是最大退避间隔。
	Max time.Duration
	// Factor 是退避倍数。
	Factor float64
	// Jitter 控制是否在退避时间上添加随机抖动，避免惊群效应。
	Jitter bool

	mu      sync.Mutex
	attempt int
}

// NewBackoff 创建一个默认的退避策略实例。
// 默认值: Min=1s, Max=30s, Factor=2, Jitter=true。
func NewBackoff() *Backoff {
	return &Backoff{
		Min:    time.Second,
		Max:    30 * time.Second,
		Factor: 2,
		Jitter: true,
	}
}

// Next 返回下一次退避的等待时间，并递增重试计数。
// 计算公式: min(Min * Factor^attempt, Max)，可选添加随机抖动。
func (b *Backoff) Next() time.Duration {
	b.mu.Lock()
	defer b.mu.Unlock()

	d := time.Duration(float64(b.Min) * math.Pow(b.Factor, float64(b.attempt)))
	b.attempt++

	if d > b.Max {
		d = b.Max
	}

	if b.Jitter {
		d = time.Duration(rand.Int63n(int64(d)))
		if d < b.Min {
			d = b.Min
		}
	}

	return d
}

// Reset 将重试计数归零，重新开始退避周期。
func (b *Backoff) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.attempt = 0
}
