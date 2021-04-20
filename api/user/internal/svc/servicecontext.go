package svc

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"go-micro-project/api/user/internal/config"
	"go-micro-project/rpc/user/userclient"
)

type ServiceContext struct {
	Config   config.Config
	User     userclient.User
	Log      logx.LogConf
	RestConf rest.RestConf
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		User:     userclient.NewUser(zrpc.MustNewClient(c.User)),
		Log:      logx.LogConf{Path: "./logs", Mode: "file"}, // 紀錄檔路徑
		RestConf: rest.RestConf{Host: "test"},
	}
}
