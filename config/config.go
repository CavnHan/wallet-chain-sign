package config

import (
	"github.com/CavnHan/wallet-chain-sign/flags"
	"github.com/urfave/cli/v2"
)

type Config struct {
	LevelDbPath string
	RpcServer   ServerConfig
}

type ServerConfig struct {
	Host string
	Port int
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		LevelDbPath: ctx.String(flags.LevelDbPathFlag.Name),
		RpcServer: ServerConfig{
			Host: ctx.String(flags.RpcHostFlag.Name),
			Port: ctx.Int(flags.RpcPortFlag.Name),
		},
	}
}
