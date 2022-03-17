package study

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"online-study-room/internal/handler/base"
	"online-study-room/internal/logic/room"
	"online-study-room/internal/svc"
)

func RoomTopologyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := room.NewLogic(r.Context(), ctx)
		resp, err := l.RoomTopology()
		base.Result(w, resp, err)
	}
}

func SeatByIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Roomid int `path:"roomid"`
		}

		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := room.NewLogic(r.Context(), ctx)
		resp, err := l.SeatsByRoomId(req.Roomid)
		base.Result(w, resp, err)
	}
}

func IntoSeatHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			SeatId        int    `json:"seat_id"`
			UserId        int    `json:"user_id"`
			WorkName      string `json:"work_name"`
			SeatColorCode string `json:"seat_color_code"`
		}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := room.NewLogic(r.Context(), ctx)
		err := l.InfoSeat(req.SeatId,
			req.UserId,
			req.WorkName,
			req.SeatColorCode)
		base.Result(w, nil, err)
	}
}

func LeaveSeatHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			SeatId int `json:"seat_id"`
			UserId int `json:"user_id"`
		}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := room.NewLogic(r.Context(), ctx)
		err := l.LeaveSeat(req.SeatId,
			req.UserId)
		base.Result(w, nil, err)
	}
}
