package request

import "go-base-task-4/models"

type PostPage struct {
	models.PageInfo
	Title   string `json:"email" "`
	Content string `json:"email" `
	UserID  uint
}
