package retry

import (
	"math"
	"math/rand"
	"time"
)

// Strategy 接口定义了一个方法 Duration，用于计算重试操作之间的等待时间。
type Strategy interface {
	Duration(attempt int) time.Duration
}

// ExponentialStrategy 结构体定义了指数回退策略的参数。
type ExponentialStrategy struct {
	Min       time.Duration // 最小等待时间
	Max       time.Duration // 最大等待时间
	MaxJitter time.Duration // 最大抖动时间
}

// Duration 方法根据重试次数计算等待时间，使用指数回退算法，并添加抖动时间。
func (e *ExponentialStrategy) Duration(attempt int) time.Duration {
	var jitter time.Duration // 非负抖动时间
	if e.MaxJitter > 0 {
		// 生成一个随机的抖动时间，范围在 [0, MaxJitter) 之间
		jitter = time.Duration(rand.Int63n(e.MaxJitter.Nanoseconds()))
	}
	if attempt < 0 {
		// 如果重试次数小于 0，返回最小等待时间加上抖动时间
		return e.Min + jitter
	}
	// 计算指数回退时间
	durFloat := float64(e.Min)
	durFloat += math.Pow(2, float64(attempt)) * float64(time.Second)
	dur := time.Duration(durFloat)
	// 确保等待时间不超过最大等待时间
	if durFloat > float64(e.Max) {
		dur = e.Max
	}
	// 添加抖动时间
	dur += jitter

	return dur
}

// Exponential 函数返回一个默认配置的 ExponentialStrategy 实例。
func Exponential() Strategy {
	return &ExponentialStrategy{
		Min:       0,
		Max:       10 * time.Second,
		MaxJitter: 250 * time.Millisecond,
	}
}

// FixedStrategy 结构体定义了固定等待时间策略的参数。
type FixedStrategy struct {
	Dur time.Duration // 固定等待时间
}

// Duration 方法返回固定的等待时间。
func (f *FixedStrategy) Duration(attempt int) time.Duration {
	return f.Dur
}

// Fixed 函数返回一个配置了固定等待时间的 FixedStrategy 实例。
func Fixed(dur time.Duration) Strategy {
	return &FixedStrategy{
		Dur: dur,
	}
}
