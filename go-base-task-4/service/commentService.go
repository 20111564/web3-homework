package service

import (
	db2 "go-base-task-4/db"
	"go-base-task-4/models"
	"go-base-task-4/models/request"
)

type CommentService struct{}

func (c *CommentService) GetCommentList(commentPageInfo request.CommentPageInfo) (list interface{}, total int64, err error) {
	limit := commentPageInfo.PageSize
	offset := commentPageInfo.PageSize * (commentPageInfo.Page - 1)

	db := db2.SqlDB.Model(&models.Comment{})

	var commentList []models.Comment

	if error := db.Count(&total).Error; error != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&commentList).Error
	return commentList, total, err
}
