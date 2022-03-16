package svc

import (
	"online-study-room/internal/config"
	"online-study-room/internal/dao"
)

type ServiceContext struct {
	Config config.Config

	Room dao.StudyRoomModel
	Seat dao.StudySeatModel
	User dao.StudyUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := dao.MustOpenSqlite(c.Sqlite)

	return &ServiceContext{
		Config: c,
		Room:   dao.NewStudyRoomModel(db),
		Seat:   dao.NewStudySeatModel(db),
		User:   dao.NewStudyUserModel(db),
	}
}
