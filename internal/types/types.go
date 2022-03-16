package types

import "online-study-room/internal/dao"

type CommonResult struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

type GetRequest struct {
	Id int `path:"id"`
}

type RoomTopology struct {
	Rooms []Room
}

type Room struct {
	dao.Room
	Seats []Seat
}

type Seat struct {
	dao.Seat
}
