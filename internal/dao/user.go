package dao

import (
	"gorm.io/gorm"
)

const StudyUserTable = "user"

type (
	StudyUserModel interface {
		Find(id int) (*User, error)
		Finds(ids []int) ([]*User, error)
	}

	defaultStudyUserModel struct {
		db *gorm.DB
	}

	// 用户表
	User struct {
		gorm.Model
		UserName string `json:"user_name"` // 用户名密码
		//Password string `json:"password"`  // 用户名密码
		NikeName string `json:"nike_name"` // 用户昵称
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
