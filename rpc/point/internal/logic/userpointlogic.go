package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-demo/rpc/point/internal/svc"
	"go-zero-demo/rpc/point/point"
)

type UserPointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPointLogic {
	return &UserPointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPointLogic) UserPoint(in *point.GetUserPointRequest) (*point.GetUserPointResponse, error) {
	total, err := l.svcCtx.PointEarnModel.QueryEarnSUM(in.UserID)
	if err == nil {
		response := point.GetUserPointResponse{
			Point: total,
		}
		return &response, nil
	} else {
		return nil, err
	}
}
