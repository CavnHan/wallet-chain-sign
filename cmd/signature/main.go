package main

import (
	"context"
	"os"

	"github.com/CavnHan/wallet-chain-sign/common/opio"
	"github.com/ethereum/go-ethereum/log"
)

// GitCommit 和 GitData 用于存储 Git 提交信息和数据
var (
	GitCommit = ""
	GitData   = ""
)

func main() {
	// 设置默认的日志记录器
	// log.NewLogger 创建一个新的日志记录器
	// log.NewTerminalHandlerWithLevel 创建一个终端处理器，输出到标准错误输出，日志级别为 Info
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))

	// 创建一个新的 CLI 应用程序实例
	// NewCli 函数接受 GitCommit 和 GitData 作为参数
	app := NewCli(GitCommit, GitData)

	// 创建一个带有中断阻塞功能的上下文
	// opio.WithInterruptBlocker 创建一个上下文，当接收到中断信号时阻塞
	ctx := opio.WithInterruptBlocker(context.Background())

	// 运行 CLI 应用程序
	// app.RunContext 运行应用程序，传入上下文和命令行参数
	if err := app.RunContext(ctx, os.Args); err != nil {
		// 如果运行过程中出现错误，记录错误日志并退出程序
		log.Error("Application failed")
		os.Exit(1)
	}
}
