package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-demo/api/point/internal/config"
	"go-zero-demo/rpc/point/pointclient"
)

type ServiceContext struct {
	Config config.Config
	Point  pointclient.Point
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Point:  pointclient.NewPoint(zrpc.MustNewClient(c.Point)),
	}
}
