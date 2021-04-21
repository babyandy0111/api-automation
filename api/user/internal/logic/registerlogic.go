package logic

import (
	"context"
	"go-micro-project/rpc/user/user"

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
	resp, err := l.svcCtx.User.Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	token := types.JwtToken{
		AccessToken:  resp.AccessToken,
		AccessExpire: resp.AccessExpire,
		RefreshAfter: resp.RefreshAfter,
	}

	response := types.UserReply{
		Id:       resp.Id,
		Email:    resp.Email,
		JwtToken: token,
	}

	return &types.RegisterResponse{
		UserReply: response,
	}, nil
}
