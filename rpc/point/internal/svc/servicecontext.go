package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	point "go-zero-demo/model/point"
	"go-zero-demo/rpc/point/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	PointEarnModel point.PointEarnModel
	PointUseModel  point.PointUseModel
	UserPointModel point.UserPointModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:         c,
		PointEarnModel: point.NewPointEarnModel(conn, c.CacheRedis),
		PointUseModel:  point.NewPointUseModel(conn, c.CacheRedis),
		UserPointModel: point.NewUserPointModel(conn, c.CacheRedis),
	}
}
