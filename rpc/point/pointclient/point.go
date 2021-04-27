// Code generated by goctl. DO NOT EDIT!
// Source: point.proto

//go:generate mockgen -destination ./point_mock.go -package pointclient -source $GOFILE

package pointclient

import (
	"context"

	"go-zero-demo/rpc/point/point"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	GetUserEarnRequest   = point.GetUserEarnRequest
	GetUserEarnResponse  = point.GetUserEarnResponse
	GetUserUseRequest    = point.GetUserUseRequest
	GetUserUseResponse   = point.GetUserUseResponse
	GetUserPointRequest  = point.GetUserPointRequest
	GetUserPointResponse = point.GetUserPointResponse

	Point interface {
		UserEarn(ctx context.Context, in *GetUserEarnRequest) (*GetUserEarnResponse, error)
		UserUse(ctx context.Context, in *GetUserUseRequest) (*GetUserUseResponse, error)
		UserPoint(ctx context.Context, in *GetUserPointRequest) (*GetUserPointResponse, error)
	}

	defaultPoint struct {
		cli zrpc.Client
	}
)

func NewPoint(cli zrpc.Client) Point {
	return &defaultPoint{
		cli: cli,
	}
}

func (m *defaultPoint) UserEarn(ctx context.Context, in *GetUserEarnRequest) (*GetUserEarnResponse, error) {
	client := point.NewPointClient(m.cli.Conn())
	return client.UserEarn(ctx, in)
}

func (m *defaultPoint) UserUse(ctx context.Context, in *GetUserUseRequest) (*GetUserUseResponse, error) {
	client := point.NewPointClient(m.cli.Conn())
	return client.UserUse(ctx, in)
}

func (m *defaultPoint) UserPoint(ctx context.Context, in *GetUserPointRequest) (*GetUserPointResponse, error) {
	client := point.NewPointClient(m.cli.Conn())
	return client.UserPoint(ctx, in)
}
