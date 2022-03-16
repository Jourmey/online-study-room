package dao

import (
	"online-study-room/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSqlite(config config.Sqlite) (*gorm.DB, error) {
	return open(config)
}

func MustOpenSqlite(config config.Sqlite) *gorm.DB {
	db, err := open(config)
	if err != nil {
		logx.Error("开启Sqlite数据库失败", config)
		panic(err)
	}
	return db
}

func open(config config.Sqlite) (db *gorm.DB, err error) {
	return gorm.Open(sqlite.Open(config.Dsn), &gorm.Config{})
}
