package logic

import (
	"context"
	"go-zero-demo/rpc/user/user"
	"strconv"

	"go-zero-demo/rpc/user/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserinfoLogic) Userinfo(in *user.UserinfoRequest) (*user.Response, error) {

	newid, _ := strconv.ParseInt(in.Userid, 10, 64)
	result, err := l.svcCtx.Model.FindOne(newid)
	if err == nil {
		return &user.Response{
			Id:    result.Id,
			Email: result.Email,
			Name:  result.Name,
		}, nil
	} else {
		return nil, err
	}
}
