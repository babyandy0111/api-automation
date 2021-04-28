package logic

import (
	"context"
	"go-zero-demo/api/point/internal/svc"
	"go-zero-demo/api/point/internal/types"
	"go-zero-demo/rpc/point/point"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserPointLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPointLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserPointLogic {
	return GetUserPointLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPointLogic) GetUserPoint(req types.GetUserPointRequest) (*types.GetUserPointResponse, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.Point.UserPoint(l.ctx, &point.GetUserPointRequest{
		UserID: req.UserId,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetUserPointResponse{Point: resp.Point}, nil
}
