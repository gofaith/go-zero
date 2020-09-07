package config

import (
	"github.com/gofaith/go-zero/rest"
	"github.com/gofaith/go-zero/rpcx"
)

type Config struct {
	rest.RestConf
	Add   rpcx.RpcClientConf
	Check rpcx.RpcClientConf
}