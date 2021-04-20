package logic

import (
	"context"

	"go-micro-project/rpc/user/internal/svc"
	"go-micro-project/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserinfoLogic) Userinfo(in *user.UserinfoRequest) (*user.Response, error) {
	// todo: add your logic here and delete this line

	return &user.Response{}, nil
}
