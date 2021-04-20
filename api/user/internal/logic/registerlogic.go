package logic

import (
	"context"

	"go-micro-project/api/user/internal/svc"
	"go-micro-project/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterRequest) (*types.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	return &types.RegisterResponse{}, nil
}
