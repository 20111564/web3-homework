package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content";gorm:"not null"`
	UserID  uint   `json:"userId";gorm:"not null"`
	User    User
	PostID  uint `json:"postId";gorm:"not null"`
	Post    Post
}
