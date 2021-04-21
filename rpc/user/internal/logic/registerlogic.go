package logic

import (
	"context"
	"errors"
	model "go-micro-project/model/user"
	"go-micro-project/utils"
	"time"

	"go-micro-project/rpc/user/internal/svc"
	"go-micro-project/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.Response, error) {
	_, err := l.svcCtx.Model.FindOneByEmail(in.Email)

	if err == model.ErrNotFound {
		newuser := model.User{
			Name:     in.Username,
			Email:    in.Email,
			Password: utils.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		if result, err := l.svcCtx.Model.Insert(newuser); err != nil {
			return nil, err
		} else {
			newuser.Id, err = result.LastInsertId()
			if err != nil {
				return nil, err
			} else {
				now := time.Now().Unix()
				accessExpire := l.svcCtx.Config.AccessExpire
				jwtToken, err := utils.GetJwtToken(l.svcCtx.Config.AccessSecret, now, accessExpire, newuser.Id)
				if err != nil {
					return nil, err
				}
				response := user.Response{
					Email:        newuser.Email,
					Id:           newuser.Id,
					AccessToken:  jwtToken,
					AccessExpire: now + accessExpire,
					RefreshAfter: now + accessExpire/2,
				}
				return &response, nil
			}
		}
	} else if err == nil {
		return nil, errors.New("該用戶已存在")
	} else {
		return nil, err
	}
}
