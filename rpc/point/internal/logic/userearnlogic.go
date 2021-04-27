package logic

import (
	"context"

	"go-zero-demo/rpc/point/internal/svc"
	"go-zero-demo/rpc/point/point"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserEarnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserEarnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserEarnLogic {
	return &UserEarnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserEarnLogic) UserEarn(in *point.GetUserEarnRequest) (*point.GetUserEarnResponse, error) {
	// todo: add your logic here and delete this line

	return &point.GetUserEarnResponse{}, nil
}
