package types

const (
	GoTimeFormat = "2006-01-02 15:04:05"
)

type CommonResult struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

type GetRequest struct {
	Id int `path:"id"`
}

type RoomTopology struct {
	Rooms []Room `json:"rooms"`
}

type Room struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SeatNumber int    `json:"seat_number"`
	Seats      []Seat `json:"seats"`
}

type (
	SeatAllResponse struct {
		Seats []OnlineSeat `json:"seats"`
	}
	OnlineSeat struct { // 用户携带的信息
		UserId    string `json:"user_id"`    // 用户ID
		EnteredAt string `json:"entered_at"` // 入住时间
		Until     string `json:"until"`      // 自动退房预定时间
		WorkName  string `json:"work_name"`  // 工作名称
		ColorCode string `json:"color_code"` // 座位背景色的彩色编码
	}
)

type Seat struct {
	ID   int `json:"id"`
	X    int `json:"x"`
	Y    int `json:"y"`
	Wide int `json:"wide"`
	High int `json:"high"`
}
