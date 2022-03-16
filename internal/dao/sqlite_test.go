package dao

import (
	"online-study-room/internal/config"
	"testing"
)

func Test_sqlite(t *testing.T) {
	db := MustOpenSqlite(config.Sqlite{
		Dsn: "/Users/zhengjm/product/jourmey/online-study-room/study.db",
	})

	//创建表,自动迁移(把结构体和数据表相对应)
	err := db.AutoMigrate(&Room{}, &User{}, &Seat{})
	if err != nil {
		t.Fatal(err)
	}
}
