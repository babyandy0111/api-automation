// Code generated by goctl. DO NOT EDIT!
// Source: point.proto

package server

import (
	"context"

	"go-zero-demo/rpc/point/internal/logic"
	"go-zero-demo/rpc/point/internal/svc"
	"go-zero-demo/rpc/point/point"
)

type PointServer struct {
	svcCtx *svc.ServiceContext
}

func NewPointServer(svcCtx *svc.ServiceContext) *PointServer {
	return &PointServer{
		svcCtx: svcCtx,
	}
}

func (s *PointServer) UserEarn(ctx context.Context, in *point.GetUserEarnRequest) (*point.GetUserEarnResponse, error) {
	l := logic.NewUserEarnLogic(ctx, s.svcCtx)
	return l.UserEarn(in)
}

func (s *PointServer) UserUse(ctx context.Context, in *point.GetUserUseRequest) (*point.GetUserUseResponse, error) {
	l := logic.NewUserUseLogic(ctx, s.svcCtx)
	return l.UserUse(in)
}

func (s *PointServer) UserPoint(ctx context.Context, in *point.GetUserPointRequest) (*point.GetUserPointResponse, error) {
	l := logic.NewUserPointLogic(ctx, s.svcCtx)
	return l.UserPoint(in)
}
