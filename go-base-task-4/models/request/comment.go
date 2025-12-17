package request

import "go-base-task-4/models"

type CommentPageInfo struct {
	models.PageInfo
	Content string `json:"content" `
	UserID  uint   `json:"userId" "`
	PostID  uint   `json:"postId" "`
}
