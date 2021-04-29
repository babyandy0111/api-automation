package handler

import (
	"go-zero-demo/utils"
	"net/http"

	"go-zero-demo/api/user/internal/logic"
	"go-zero-demo/api/user/internal/svc"
	"go-zero-demo/api/user/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			utils.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		utils.HttpResult(r, w, resp, err)
	}
}
