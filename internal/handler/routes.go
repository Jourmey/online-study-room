package handler

import (
	"net/http"
	handler "online-study-room/internal/handler/hello"
	"online-study-room/internal/handler/study"
	"online-study-room/internal/handler/user"
	"online-study-room/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{http.MethodGet, "/api/hello", handler.HelloHandler(serverCtx)},

			{http.MethodPost, "/api/user/registration", user.UserRegisterHandler(serverCtx)}, // 用户注册
			{http.MethodPost, "/api/user/signin", user.UserSigninHandler(serverCtx)},         // 用户登录验证
			{http.MethodGet, "/api/user/:id", user.UserByIdHandler(serverCtx)},               // 查询用户信息

			{http.MethodGet, "/api/room/topology", study.RoomTopologyHandler(serverCtx)}, // 获取房间拓扑
			{http.MethodGet, "/api/seat/:roomid", study.SeatByIdHandler(serverCtx)},      // 获取房间的在线课桌信息

			{http.MethodPost, "/api/seat/into", study.IntoSeatHandler(serverCtx)},   // 学生入座
			{http.MethodPost, "/api/seat/leave", study.LeaveSeatHandler(serverCtx)}, // 学生离座
		},
	)
}
