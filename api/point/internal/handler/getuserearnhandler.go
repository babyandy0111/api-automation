package handler

import (
	"net/http"

	"go-zero-demo/api/point/internal/logic"
	"go-zero-demo/api/point/internal/svc"
	"go-zero-demo/api/point/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserEarnHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserEarnRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserEarnLogic(r.Context(), ctx)
		resp, err := l.GetUserEarn(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
