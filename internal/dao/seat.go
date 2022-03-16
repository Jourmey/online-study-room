package dao

import (
	"gorm.io/gorm"
	"time"
)

const StudySeatTable = "seat"

type (
	StudySeatModel interface {
		Find(id int) (*Seat, error)
		Finds(ids []int) ([]*Seat, error)
		FindByRoomIds(ids []int) ([]*Seat, error)
		SetSeat(seatId int,
			userId int,
			workName string,
			enterDate time.Time,
			exitDate time.Time,
			seatColorCode string) (err error)
		LeaveSeat(seatId int,
			userId int) (err error)
	}

	defaultStudySeatModel struct {
		db *gorm.DB
	}

	// 座位表
	Seat struct {
		gorm.Model
		RoomId int `gorm:"column:room_id"` // 所属房间

		// 用户携带的信息
		UserId    string    `gorm:"column:user_id"`    // 用户ID
		EnteredAt time.Time `gorm:"column:entered_at"` // 入住时间
		Until     time.Time `gorm:"column:until"`      // 自动退房预定时间
		WorkName  string    `gorm:"column:work_name"`  // 工作名称
		ColorCode string    `gorm:"column:color_code"` // 座位背景色的彩色编码

		// 座位本身的信息
		X    int `gorm:"column:x"`
		Y    int `gorm:"column:y"`
		Wide int `gorm:"column:wide"`
		High int `gorm:"column:high"`
	}
)

func NewStudySeatModel(db *gorm.DB) StudySeatModel {
	return &defaultStudySeatModel{
		db: db,
	}
}

func (*Seat) TableName() string {
	return StudySeatTable
}

func (m *defaultStudySeatModel) Find(id int) (v *Seat, err error) {
	err = m.db.Where("id = ?", id).Find(&v).Error
	return
}

func (m *defaultStudySeatModel) Finds(ids []int) (v []*Seat, err error) {
	err = m.db.Where("id in (?)", ids).Find(&v).Error
	return
}

func (m *defaultStudySeatModel) FindByRoomIds(ids []int) (v []*Seat, err error) {
	err = m.db.Where("room_id in (?)", ids).Find(&v).Error
	return
}

func (m *defaultStudySeatModel) SetSeat(seatId int,
	userId int,
	workName string,
	enterDate time.Time,
	exitDate time.Time,
	seatColorCode string) (err error) {

	update := map[string]interface{}{
		"user_id":    userId,
		"entered_at": enterDate,
		"until":      exitDate,
		"work_name":  workName,
		"color_code": seatColorCode,
	}
	err = m.db.Table(StudySeatTable).Where("id = ?", seatId).Updates(update).Error
	return
}

func (m *defaultStudySeatModel) LeaveSeat(seatId int, userId int) (err error) {
	update := map[string]interface{}{
		"entered_at": nil,
		"until":      nil,
		"work_name":  nil,
		"color_code": nil,
	}
	err = m.db.Table(StudySeatTable).Where("id = ? and user_id = ?", seatId, userId).Updates(update).Error
	return
}
