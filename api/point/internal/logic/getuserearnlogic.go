package logic

import (
	"context"

	"go-zero-demo/api/point/internal/svc"
	"go-zero-demo/api/point/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserEarnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserEarnLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserEarnLogic {
	return GetUserEarnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserEarnLogic) GetUserEarn(req types.GetUserEarnRequest) (*types.GetUserEarnResponse, error) {
	// todo: add your logic here and delete this line

	return &types.GetUserEarnResponse{}, nil
}
