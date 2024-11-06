package main

import (
	"context"
	"fmt"

	"github.com/CavnHan/wallet-chain-sign/common/cliapp"
	"github.com/CavnHan/wallet-chain-sign/config"
	flags2 "github.com/CavnHan/wallet-chain-sign/flags"
	"github.com/CavnHan/wallet-chain-sign/leveldb"
	"github.com/CavnHan/wallet-chain-sign/services/rpc"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli/v2"
)

// runRpc 函数用于启动 RPC 服务
// 它接受 CLI 上下文和一个取消函数作为参数，返回一个 Lifecycle 实例和一个错误
func runRpc(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	fmt.Println("running grpc services...")

	// 从 CLI 上下文中读取配置
	cfg := config.NewConfig(ctx)

	// 创建 RPC 服务器配置
	grpcServerCfg := &rpc.RpcServerConfig{
		GrpcHostname: cfg.RpcServer.Host,
		GrpcPort:     cfg.RpcServer.Port,
	}

	// 创建 LevelDB 实例
	db, err := leveldb.NewKeyStore(cfg.LevelDbPath)
	if err != nil {
		log.Error("new key store level db", "err", err)
	}

	// 返回一个新的 RPC 服务器实例
	return rpc.NewRpcServer(db, grpcServerCfg)
}

// NewCli 函数创建并返回一个新的 CLI 应用程序实例
// 它接受 Git 提交和 Git 数据作为参数
func NewCli(GitCommit string, GitData string) *cli.App {
	// 获取命令行标志
	flags := flags2.Flags

	// 创建并返回 CLI 应用程序实例
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitData),                         // 设置版本信息
		Description:          "An exchange wallet scanner services with rpc and rest api services", // 设置描述信息
		EnableBashCompletion: true,                                                                 // 启用 Bash 自动补全
		Commands: []*cli.Command{
			{
				Name:        "rpc",                       // 命令名称
				Flags:       flags,                       // 命令行标志
				Description: "Run rpc services",          // 命令描述
				Action:      cliapp.LifecycleCmd(runRpc), // 命令操作
			},
			{
				Name:        "version",              // 命令名称
				Description: "Show project version", // 命令描述
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx) // 显示版本信息
					return nil
				},
			},
		},
	}
}
