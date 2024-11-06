package cliapp

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/CavnHan/wallet-chain-sign/common/opio"
)

// Lifecycle 接口定义了应用程序生命周期的三个方法：Start、Stop 和 Stopped。
type Lifecycle interface {
	Start(ctx context.Context) error // 启动应用程序
	Stop(ctx context.Context) error  // 停止应用程序
	Stopped() bool                   // 检查应用程序是否已停止
}

// LifecycleAction 类型定义了一个函数类型，用于设置和返回 Lifecycle 实例。
type LifecycleAction func(ctx *cli.Context, close context.CancelCauseFunc) (Lifecycle, error)

// LifecycleCmd 函数接受一个 LifecycleAction 函数，返回一个 cli.ActionFunc，用于处理命令行操作。
func LifecycleCmd(fn LifecycleAction) cli.ActionFunc {
	return lifecycleCmd(fn, opio.BlockOnInterruptsContext)
}

// waitSignalFn 类型定义了一个函数类型，用于等待中断信号。
type waitSignalFn func(ctx context.Context, signals ...os.Signal)

// interruptErr 变量表示中断信号错误。
var interruptErr = errors.New("interrupt signal")

// lifecycleCmd 函数接受一个 LifecycleAction 函数和一个 waitSignalFn 函数，返回一个 cli.ActionFunc。
func lifecycleCmd(fn LifecycleAction, blockOnInterrupt waitSignalFn) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		// 创建一个带有取消功能的上下文 appCtx
		hostCtx := ctx.Context
		appCtx, appCancel := context.WithCancelCause(hostCtx)
		ctx.Context = appCtx

		// 启动一个 goroutine，等待中断信号并取消上下文
		go func() {
			blockOnInterrupt(appCtx)
			appCancel(interruptErr)
		}()

		// 调用 LifecycleAction 函数，设置并返回 Lifecycle 实例
		appLifecycle, err := fn(ctx, appCancel)
		if err != nil {
			return errors.Join(
				fmt.Errorf("failed to setup: %w", err),
				context.Cause(appCtx),
			)
		}

		// 启动 Lifecycle 实例
		if err := appLifecycle.Start(appCtx); err != nil {
			return errors.Join(
				fmt.Errorf("failed to start: %w", err),
				context.Cause(appCtx),
			)
		}

		// 等待上下文完成
		<-appCtx.Done()

		// 创建一个新的上下文 stopCtx，用于停止 Lifecycle 实例
		stopCtx, stopCancel := context.WithCancelCause(hostCtx)
		go func() {
			blockOnInterrupt(stopCtx)
			stopCancel(interruptErr)
		}()

		// 停止 Lifecycle 实例
		stopErr := appLifecycle.Stop(stopCtx)
		stopCancel(nil)
		if stopErr != nil {
			return errors.Join(
				fmt.Errorf("failed to stop: %w", stopErr),
				context.Cause(stopCtx),
			)
		}
		return nil
	}
}
