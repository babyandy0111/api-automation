package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"go-micro-project/api/user/internal/config"
	"go-micro-project/rpc/user/userclient"
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
