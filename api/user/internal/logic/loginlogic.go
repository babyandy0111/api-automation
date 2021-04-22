package logic

import (
	"context"
	"go-zero-demo/rpc/user/user"

	"go-zero-demo/api/user/internal/svc"
	"go-zero-demo/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginRequest) (*types.LoginResponse, error) {

	resp, err := l.svcCtx.User.Login(l.ctx, &user.LoginRequest{
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

	return &types.LoginResponse{UserReply: response}, nil
}
