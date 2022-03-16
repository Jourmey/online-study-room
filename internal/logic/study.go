package study

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"online-study-room/internal/dao"
	"online-study-room/internal/svc"
	"online-study-room/internal/types"
	"time"
)

type Logic struct {
	logx.Logger
	ctx context.Context

	room dao.StudyRoomModel
	seat dao.StudySeatModel
	user dao.StudyUserModel
}

func NewLogic(ctx context.Context, svcCtx *svc.ServiceContext) Logic {
	return Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		room:   svcCtx.Room,
		seat:   svcCtx.Seat,
		user:   svcCtx.User,
	}
}

// 获取所有房间的拓扑结构
func (c *Logic) RoomTopology() (types.RoomTopology, error) {
	rooms, err := c.room.FindAll()
	if err != nil {
		return types.RoomTopology{}, err
	}

	roomIds := make([]int, 0, len(rooms))
	for _, room := range rooms {
		roomIds = append(roomIds, int(room.ID))
	}

	seats, err := c.seat.FindByRoomIds(roomIds)
	if err != nil {
		return types.RoomTopology{}, err
	}

	r := types.RoomTopology{
		Rooms: make([]types.Room, 0, len(rooms)),
	}

	for _, room := range rooms {
		oneRoom := types.Room{
			ID:         int(room.ID),
			Name:       room.Name,
			SeatNumber: room.SeatNumber,
			Seats:      make([]types.Seat, 0),
		}
		for _, seat := range seats {
			if seat.RoomId == int(room.ID) {
				oneRoom.Seats = append(oneRoom.Seats, types.Seat{
					ID:   int(seat.ID),
					X:    seat.X,
					Y:    seat.Y,
					Wide: seat.Wide,
					High: seat.High,
				})
			}
		}
		r.Rooms = append(r.Rooms, oneRoom)
	}

	return r, nil
}

// 获取所有有人的座位的数据
func (c *Logic) Seats() (types.SeatAllResponse, error) {
	o, err := c.seat.FindAllOnline()
	if err != nil {
		return types.SeatAllResponse{}, err
	}
	r := types.SeatAllResponse{
		Seats: make([]types.OnlineSeat, 0, len(o)),
	}
	for _, oo := range o {
		r.Seats = append(r.Seats, types.OnlineSeat{
			UserId:    oo.UserId,
			EnteredAt: oo.EnteredAt.Format(types.GoTimeFormat),
			Until:     oo.Until.Format(types.GoTimeFormat),
			WorkName:  oo.WorkName,
			ColorCode: oo.ColorCode,
		})
	}
	return r, nil
}

// 用户进入座位
func (c *Logic) InfoSeat(seatId int,
	userId int,
	workName string,
	seatColorCode string,
) error {
	exitDate := time.Now()
	enterDate := exitDate.Add(2 * time.Hour)

	return c.seat.IntoSeat(seatId,
		userId,
		workName,
		enterDate,
		exitDate,
		seatColorCode)
}

// 用户离开座位
func (c *Logic) LeaveSeat(seatId int,
	userId int,
) error {

	return c.seat.LeaveSeat(seatId,
		userId)
}

//
//func (c *StudyController) SetLastEnteredDate(_ context.Context, tx *firestore.Transaction, userId string, enteredDate time.Time) error {
//
//	return nil
//}
//
//func (c *StudyController) SetLastExitedDate(tx *firestore.Transaction, userId string, exitedDate time.Time) error {
//
//	return nil
//}
//
//func (c *StudyController) UnSetSeatInRoom(tx *firestore.Transaction, seat Seat) error {
//
//	return nil
//}
//
//func (c *StudyController) SetMyRankVisible(_ context.Context, tx *firestore.Transaction, userId string, rankVisible bool) error {
//
//	return nil
//}
//
//func (c *StudyController) SetMyDefaultStudyMin(tx *firestore.Transaction, userId string, defaultStudyMin int) error {
//
//	return nil
//}
//
//func (c *StudyController) AddUserHistory(tx *firestore.Transaction, userId string, action string, details interface{}) error {
//
//	return nil
//}
//
//func (c *StudyController) RetrieveUser(ctx context.Context, tx *firestore.Transaction, userId string) (UserDoc, error) {
//
//	return userData, nil
//}
//
//func (c *StudyController) UpdateTotalTime(
//	tx *firestore.Transaction,
//	userId string,
//	newTotalTimeSec int,
//	newDailyTotalTimeSec int,
//) error {
//
//}
//
//func (c *StudyController) SaveLiveChatId(ctx context.Context, tx *firestore.Transaction, liveChatId string) error {
//
//	return nil
//}
//
//func (c *StudyController) InitializeUser(tx *firestore.Transaction, userId string, userData UserDoc) error {
//
//}
//
//func (c *StudyController) RetrieveAllUserDocRefs(ctx context.Context) ([]*firestore.DocumentRef, error) {
//
//}
//
//func (c *StudyController) RetrieveAllNonDailyZeroUserDocs(ctx context.Context) *firestore.DocumentIterator {
//
//}
//
//func (c *StudyController) ResetDailyTotalStudyTime(tx *firestore.Transaction, userRef *firestore.DocumentRef) error {
//
//	return nil
//}
//
//func (c *StudyController) SetLastResetDailyTotalStudyTime(tx *firestore.Transaction, date time.Time) error {
//
//	return nil
//}
//
//func (c *StudyController) SetDesiredMaxSeats(tx *firestore.Transaction, desiredMaxSeats int) error {
//
//	return nil
//}
//
//func (c *StudyController) SetMaxSeats(tx *firestore.Transaction, maxSeats int) error {
//
//	return nil
//}
//
//func (c *StudyController) SetAccessTokenOfChannelCredential(tx *firestore.Transaction, accessToken string, expireDate time.Time) error {
//
//	return nil
//}
//
//func (c *StudyController) SetAccessTokenOfBotCredential(ctx context.Context, tx *firestore.Transaction, accessToken string, expireDate time.Time) error {
//
//	return nil
//}
//
//func (c *StudyController) UpdateSeats(tx *firestore.Transaction, seats []Seat) error {
//
//}
