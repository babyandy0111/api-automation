package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-demo/api/user/internal/config"
	"go-zero-demo/rpc/user/userclient"
)

type ServiceContext struct {
	Config config.Config
	User   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.User)),
	}
}
