package logic

import (
	"context"

	"go-zero-demo/api/point/internal/svc"
	"go-zero-demo/api/point/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserUseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserUseLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserUseLogic {
	return GetUserUseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserUseLogic) GetUserUse(req types.GetUserUseRequest) (*types.GetUserUseResponse, error) {
	// todo: add your logic here and delete this line

	return &types.GetUserUseResponse{}, nil
}
