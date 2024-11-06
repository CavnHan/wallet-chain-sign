package opio

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// DefaultInterruptSignals 定义了一组默认的中断信号，包括 SIGINT、SIGKILL、SIGTERM 和 SIGQUIT。
var DefaultInterruptSignals = []os.Signal{
	os.Interrupt,
	os.Kill,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

// BlockOnInterrupts 函数阻塞执行，直到接收到指定的中断信号。
// 如果没有指定信号，使用默认的中断信号。
func BlockOnInterrupts(signals ...os.Signal) {
	if len(signals) == 0 {
		signals = DefaultInterruptSignals
	}
	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, signals...)
	<-interruptChannel
}

// BlockOnInterruptsContext 函数阻塞执行，直到接收到指定的中断信号或上下文被取消。
// 如果没有指定信号，使用默认的中断信号。
func BlockOnInterruptsContext(ctx context.Context, signals ...os.Signal) {
	if len(signals) == 0 {
		signals = DefaultInterruptSignals
	}
	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, signals...)
	select {
	case <-interruptChannel:
	case <-ctx.Done():
		signal.Stop(interruptChannel)
	}
}

// interruptContextKeyType 是用于在上下文中存储中断处理函数的键类型。
type interruptContextKeyType struct{}

// blockerContextKey 是用于在上下文中存储中断处理函数的键值。
var blockerContextKey = interruptContextKeyType{}

// interruptCatcher 结构体包含一个用于接收中断信号的通道。
type interruptCatcher struct {
	incoming chan os.Signal
}

// Block 方法阻塞执行，直到接收到中断信号或上下文被取消。
func (c *interruptCatcher) Block(ctx context.Context) {
	select {
	case <-c.incoming:
	case <-ctx.Done():
	}
}

// WithInterruptBlocker 函数返回一个新的上下文，其中包含一个中断处理函数。
// 如果上下文已经包含中断处理函数，则直接返回原上下文。
func WithInterruptBlocker(ctx context.Context) context.Context {
	if ctx.Value(blockerContextKey) != nil { // already has an interrupt handler
		return ctx
	}
	catcher := &interruptCatcher{
		incoming: make(chan os.Signal, 10),
	}
	signal.Notify(catcher.incoming, DefaultInterruptSignals...)

	return context.WithValue(ctx, blockerContextKey, BlockFn(catcher.Block))
}

// WithBlocker 函数返回一个新的上下文，其中包含指定的中断处理函数。
func WithBlocker(ctx context.Context, fn BlockFn) context.Context {
	return context.WithValue(ctx, blockerContextKey, fn)
}

// BlockFn 类型表示一个接受上下文的函数，用于处理中断信号。
type BlockFn func(ctx context.Context)

// BlockerFromContext 函数从上下文中提取中断处理函数。
func BlockerFromContext(ctx context.Context) BlockFn {
	v := ctx.Value(blockerContextKey)
	if v == nil {
		return nil
	}
	return v.(BlockFn)
}

// CancelOnInterrupt 函数返回一个新的上下文，当接收到中断信号时，取消该上下文。
// 使用上下文中的中断处理函数（如果存在），否则使用默认的中断处理函数。
func CancelOnInterrupt(ctx context.Context) context.Context {
	inner, cancel := context.WithCancel(ctx)

	blockOnInterrupt := BlockerFromContext(ctx)
	if blockOnInterrupt == nil {
		blockOnInterrupt = func(ctx context.Context) {
			BlockOnInterruptsContext(ctx) // default signals
		}
	}

	go func() {
		blockOnInterrupt(ctx)
		cancel()
	}()

	return inner
}
