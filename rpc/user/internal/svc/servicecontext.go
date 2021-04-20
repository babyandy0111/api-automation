package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	model "go-micro-project/model/user"
	"go-micro-project/rpc/user/internal/config"
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
