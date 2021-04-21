package handler

import (
	"go-micro-project/utils"
	"net/http"

	"go-micro-project/api/user/internal/logic"
	"go-micro-project/api/user/internal/svc"
	"go-micro-project/api/user/internal/types"

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
			// httpx.Error(w, err)
			httpx.OkJson(w, utils.FailureResponse(nil, err.Error(), 1000))
		} else {
			// httpx.OkJson(w, resp)
			httpx.OkJson(w, utils.SuccessResponse(resp, ""))
		}
	}
}
