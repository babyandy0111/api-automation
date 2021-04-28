// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-demo/api/point/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/points/earn",
				Handler: GetUserEarnHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/points/use",
				Handler: GetUserUseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/users/:user_id/points",
				Handler: GetUserPointHandler(serverCtx),
			},
		},
	)
}
