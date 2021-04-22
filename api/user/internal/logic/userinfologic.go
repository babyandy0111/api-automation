package logic

import (
	"context"
	"go-zero-demo/rpc/user/user"

	"go-zero-demo/api/user/internal/svc"
	"go-zero-demo/api/user/internal/types"

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

func (l *UserInfoLogic) UserInfo(userid string) (*types.UserinfoResponse, error) {

	resp, err := l.svcCtx.User.Userinfo(l.ctx, &user.UserinfoRequest{
		Userid: userid,
	})

	if err != nil {
		return nil, err
	}
	response := types.UserReply{
		Id:       resp.Id,
		Email:    resp.Email,
		Username: resp.Name,
	}

	return &types.UserinfoResponse{
		UserReply: response,
	}, nil
}
