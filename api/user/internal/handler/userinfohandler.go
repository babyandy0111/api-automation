package handler

import (
	"go-micro-project/utils"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-micro-project/api/user/internal/logic"
	"go-micro-project/api/user/internal/svc"
)

func UserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("x-user-id")
		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo(userId)
		if err != nil {
			// httpx.Error(w, err)
			httpx.OkJson(w, utils.FailureResponse(nil, err.Error(), 1000))
		} else {
			// httpx.OkJson(w, resp)
			httpx.OkJson(w, utils.SuccessResponse(resp, "获取成功"))
		}
	}
}
