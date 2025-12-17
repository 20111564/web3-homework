package service

import (
	db2 "go-base-task-4/db"
	"go-base-task-4/models"
	"go-base-task-4/models/request"
)

type PostService struct{}

func (p *PostService) GetUserInfoList(postPage request.PostPage) (list interface{}, total int64, err error) {
	limit := postPage.PageSize
	offset := postPage.PageSize * (postPage.Page - 1)

	db := db2.SqlDB.Model(&models.Post{})

	var postList []models.Post

	if postPage.Title != "" {
		db = db.Where("title LIKE ?", "%"+postPage.Title+"%")
	}
	if error := db.Count(&total).Error; error != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&postList).Error
	return postList, total, err
}
