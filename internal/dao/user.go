package dao

import (
	"gorm.io/gorm"
)

const StudyUserTable = "user"

type (
	StudyUserModel interface {
		Find(id int) (*User, error)
		Finds(ids []int) ([]*User, error)
		FindByUserName(name string) (*User, error)
		Insert(userName string, password string, nikeName string) (*User, error)
	}

	defaultStudyUserModel struct {
		db *gorm.DB
	}

	// 用户表
	User struct {
		gorm.Model
		UserName string `gorm:"column:user_name"` // 用户名密码
		Password string `gorm:"column:password"`  // 用户名密码
		NikeName string `gorm:"column:nike_name"` // 用户昵称
	}
)

func NewStudyUserModel(db *gorm.DB) StudyUserModel {
	return &defaultStudyUserModel{
		db: db,
	}
}

func (*User) TableName() string {
	return StudyUserTable
}

func (m *defaultStudyUserModel) Find(id int) (v *User, err error) {
	err = m.db.Where("id = ?", id).Find(&v).Error
	return
}

func (m *defaultStudyUserModel) Finds(ids []int) (v []*User, err error) {
	err = m.db.Where("id in (?)", ids).Find(&v).Error
	return
}

func (m *defaultStudyUserModel) FindByUserName(name string) (v *User, err error) {
	err = m.db.Where("user_name = ?", name).Find(&v).Error
	return
}

func (m *defaultStudyUserModel) Insert(userName string, password string, nikeName string) (v *User, err error) {
	v = &User{
		Model:    gorm.Model{},
		UserName: userName,
		Password: password,
		NikeName: nikeName,
	}

	err = m.db.Table(StudyUserTable).Create(v).Error
	return
}
