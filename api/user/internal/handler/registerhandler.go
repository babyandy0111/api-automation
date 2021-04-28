package handler

import (
	"go-zero-demo/utils"
	"net/http"

	"go-zero-demo/api/user/internal/logic"
	"go-zero-demo/api/user/internal/svc"
	"go-zero-demo/api/user/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			httpx.OkJson(w, utils.FailureResponse(nil, err.Error(), 1000))
			return
		}

		l := logic.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {
			utils.ParamErrorResult(r, w, err)
		} else {
			utils.HttpResult(r, w, resp, err)
		}
	}
}
