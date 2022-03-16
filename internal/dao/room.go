package dao

import (
	"gorm.io/gorm"
)

const StudyRoomTable = "room"

type (
	StudyRoomModel interface {
		Find(id int) (*Room, error)
		Finds(ids []int) ([]*Room, error)
		FindAll() ([]*Room, error)
	}

	defaultStudyRoomModel struct {
		db *gorm.DB
	}

	Room struct {
		gorm.Model
		Name       string `gorm:"column:name"`
		SeatNumber int    `gorm:"column:seat_number"` // 最大座位数量
	}
)

func NewStudyRoomModel(db *gorm.DB) StudyRoomModel {
	return &defaultStudyRoomModel{
		db: db,
	}
}

func (*Room) TableName() string {
	return StudyRoomTable
}

func (m *defaultStudyRoomModel) Find(id int) (v *Room, err error) {
	err = m.db.Where("id = ?", id).Find(&v).Error
	return
}

func (m *defaultStudyRoomModel) Finds(ids []int) (v []*Room, err error) {
	err = m.db.Where("id in (?)", ids).Find(&v).Error
	return
}
func (m *defaultStudyRoomModel) FindAll() (v []*Room, err error) {
	err = m.db.Find(&v).Error
	return
}
