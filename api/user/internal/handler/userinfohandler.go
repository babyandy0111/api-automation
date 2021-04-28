package handler

import (
	"go-zero-demo/utils"
	"net/http"

	"go-zero-demo/api/user/internal/logic"
	"go-zero-demo/api/user/internal/svc"
)

func UserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("x-user-id")
		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo(userId)
		if err != nil {
			utils.ParamErrorResult(r, w, err)
		} else {
			utils.HttpResult(r, w, resp, err)
		}
	}
}
