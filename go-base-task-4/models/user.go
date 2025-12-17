package models

import (
	"go-base-task-4/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username";gorm:"unique;not null"`
	Password string `json:"password";gorm:"not null"`
	Email    string `json:"email";gorm:"unique;not null"`
}

// 查询所有用户
func (u *User) GetUserList() (userList []User, err error) {
	err = db.SqlDB.Debug().Model(&User{}).Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}
