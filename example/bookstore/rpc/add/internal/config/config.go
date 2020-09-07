package config

import (
	"github.com/gofaith/go-zero/core/stores/cache"
	"github.com/gofaith/go-zero/rpcx"
)

type Config struct {
	rpcx.RpcServerConf
	DataSource string
	Table      string
	Cache      cache.CacheConf
}