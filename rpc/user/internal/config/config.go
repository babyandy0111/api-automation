package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource   string
	CacheRedis   cache.CacheConf
	AccessSecret string
	AccessExpire int64
	Salt         string
}
