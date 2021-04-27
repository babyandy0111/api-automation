package logic

import (
	"context"

	"go-zero-demo/rpc/point/internal/svc"
	"go-zero-demo/rpc/point/point"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserUseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUseLogic {
	return &UserUseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUseLogic) UserUse(in *point.GetUserUseRequest) (*point.GetUserUseResponse, error) {
	// todo: add your logic here and delete this line

	return &point.GetUserUseResponse{}, nil
}
