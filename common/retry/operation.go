package retry

import (
	"context"
	"fmt"
	"time"
)

// ErrFailedPermanently 结构体表示操作在多次重试后仍然失败的错误。
type ErrFailedPermanently struct {
	attempts int   // 重试的次数
	LastErr  error // 最后一次错误
}

// Error 方法返回错误的描述信息。
func (e *ErrFailedPermanently) Error() string {
	return fmt.Sprintf("operation failed permanently after %d attempts: %v", e.attempts, e.LastErr)
}

// Unwrap 方法返回最后一次错误，支持错误链的解包。
func (e *ErrFailedPermanently) Unwrap() error {
	return e.LastErr
}

// pair 泛型结构体用于存储两个不同类型的值。
type pair[T, U any] struct {
	a T // 第一个值
	b U // 第二个值
}

// Do2 函数是一个泛型函数，用于处理返回两个值的操作。
// 它将操作封装为返回 pair[T, U] 的函数，并调用 Do 函数进行重试。
// 返回操作的结果和错误。
func Do2[T, U any](ctx context.Context, maxAttempts int, strategy Strategy, op func() (T, U, error)) (T, U, error) {
	// 将操作封装为返回 pair[T, U] 的函数
	f := func() (pair[T, U], error) {
		a, b, err := op()
		return pair[T, U]{a, b}, err
	}
	// 调用 Do 函数进行重试
	res, err := Do(ctx, maxAttempts, strategy, f)
	// 返回操作的结果和错误
	return res.a, res.b, err
}

// Do 函数是一个泛型函数，用于处理返回单个值的操作。
// 它接受一个上下文 ctx、最大重试次数 maxAttempts、重试策略 strategy 和操作 op。
// 如果 maxAttempts 小于 1，返回错误。
// 在重试循环中，如果上下文已取消，返回上下文错误。
// 调用操作 op，如果成功，返回结果；如果失败，根据策略等待一段时间后重试。
// 如果达到最大重试次数仍然失败，返回 ErrFailedPermanently 错误。
func Do[T any](ctx context.Context, maxAttempts int, strategy Strategy, op func() (T, error)) (T, error) {
	var empty, ret T // 定义空值和返回值
	var err error    // 定义错误变量
	if maxAttempts < 1 {
		// 如果 maxAttempts 小于 1，返回错误
		return empty, fmt.Errorf("need at least 1 attempt to run op, but have %d max attempts", maxAttempts)
	}

	// 重试循环
	for i := 0; i < maxAttempts; i++ {
		if ctx.Err() != nil {
			// 如果上下文已取消，返回上下文错误
			return empty, ctx.Err()
		}
		// 调用操作 op
		ret, err = op()
		if err == nil {
			// 如果操作成功，返回结果
			return ret, nil
		}
		if i != maxAttempts-1 {
			// 如果操作失败，根据策略等待一段时间后重试
			time.Sleep(strategy.Duration(i))
		}
	}
	// 如果达到最大重试次数仍然失败，返回 ErrFailedPermanently 错误
	return empty, &ErrFailedPermanently{
		attempts: maxAttempts,
		LastErr:  err,
	}
}
