package handler

import (
	"net/http"

	"go-zero-demo/api/point/internal/logic"
	"go-zero-demo/api/point/internal/svc"
	"go-zero-demo/api/point/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserUseHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserUseRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserUseLogic(r.Context(), ctx)
		resp, err := l.GetUserUse(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
