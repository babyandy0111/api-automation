package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	model "go-zero-demo/model/user"
	"go-zero-demo/rpc/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserModel(conn, c.CacheRedis),
	}
}
