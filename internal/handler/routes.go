package handler

import (
	"net/http"

	hello "online-study-room/internal/handler/hello"
	"online-study-room/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/hello",
				Handler: hello.HelloHandler(serverCtx),
			},
		},
	)
}
