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

			{http.MethodGet, "/api/room/topology", study.RoomTopologyHandler(serverCtx)}, // 获取房间拓扑
			{http.MethodGet, "/api/seat/all", study.SeatsHandler(serverCtx)},             // 获取所有在线的课桌信息

			{http.MethodPost, "/api/seat/into", study.IntoSeatHandler(serverCtx)},   // 学生入座
			{http.MethodPost, "/api/seat/leave", study.LeaveSeatHandler(serverCtx)}, // 学生离座
		},
	)
}
