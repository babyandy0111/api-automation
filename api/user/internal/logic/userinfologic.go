package logic

import (
	"context"

	"go-micro-project/api/user/internal/svc"
	"go-micro-project/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req types.UserinfoRequest) (*types.UserinfoResponse, error) {
	// todo: add your logic here and delete this line

	return &types.UserinfoResponse{}, nil
}
