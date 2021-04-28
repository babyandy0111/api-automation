package handler

import (
	"go-zero-demo/utils"
	"net/http"

	"go-zero-demo/api/point/internal/logic"
	"go-zero-demo/api/point/internal/svc"
	"go-zero-demo/api/point/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserPointHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserPointRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserPointLogic(r.Context(), ctx)
		resp, err := l.GetUserPoint(req)
		if err != nil {
			utils.ParamErrorResult(r, w, err)
		} else {
			utils.HttpResult(r, w, resp, err)
		}
	}
}
