package logic

import (
	"context"
	"errors"
	"go-micro-project/utils"
	"time"

	"go-micro-project/rpc/user/internal/svc"
	"go-micro-project/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.Response, error) {
	res, err := l.svcCtx.Model.FindOneByEmail(in.Email)
	if err == nil {
		passwords := utils.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
		if passwords == res.Password {
			now := time.Now().Unix()
			accessExpire := l.svcCtx.Config.AccessExpire
			jwtToken, err := utils.GetJwtToken(l.svcCtx.Config.AccessSecret, now, accessExpire, res.Id)
			if err != nil {
				return nil, err
			}
			response := user.Response{
				Email:        res.Email,
				Id:           res.Id,
				AccessToken:  jwtToken,
				AccessExpire: now + accessExpire,
				RefreshAfter: now + accessExpire/2,
			}
			return &response, nil
		} else {
			return nil, errors.New("密碼錯誤")
		}

	}
	return nil, err
}
