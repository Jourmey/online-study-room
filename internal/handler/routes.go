package handler

import (
	"net/http"
	"online-study-room/internal/handler/study"

	hello "online-study-room/internal/handler/hello"
	"online-study-room/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{http.MethodGet, "/api/hello", hello.HelloHandler(serverCtx)},

			{http.MethodGet, "/api/room/topology", study.RoomTopologyHandler(serverCtx)},

			{http.MethodPost, "/api/seat/set", study.SetSeatHandler(serverCtx)},
			{http.MethodPost, "/api/seat/leave", study.LeaveSeatHandler(serverCtx)},
		},
	)
}
