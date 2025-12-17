package api

import (
	"github.com/gin-gonic/gin"
	"go-base-task-4/common/r"
	"go-base-task-4/db"
	"go-base-task-4/models"
	"go-base-task-4/models/request"
	"go-base-task-4/models/response"
	"go-base-task-4/service"
	"go-base-task-4/utils"
)

// 新增评论
func CommentAdd(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBind(&comment); err != nil {
		r.ErrMsg(c, "参数错误")
		return
	}
	if comment.PostID <= 0 {
		r.ErrMsg(c, "文章ID不能为空")
		return
	}
	if err := db.SqlDB.Find(&models.Post{}, "id=?", comment.PostID).Error; err != nil {
		r.ErrMsg(c, "文章不存在，不能评论")
		return
	}
	userId := utils.GetUserID(c)
	comment.UserID = userId
	if error := db.SqlDB.Create(&comment).Error; error != nil {
		r.ErrMsg(c, "评论失败")
		return
	}
	r.OK(c, "", "")
}

// 评论列表
func CommentPage(c *gin.Context) {
	var commentPageInfo request.CommentPageInfo
	if err := c.ShouldBind(&commentPageInfo); err != nil {
		r.ErrMsg(c, "参数错误")
		return
	}
	var commentService = &service.CommentService{}
	list, total, error := commentService.GetCommentList(commentPageInfo)
	if error != nil {
		r.ErrMsg(c, "查询失败")
		return
	}
	r.OK(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     commentPageInfo.Page,
		PageSize: commentPageInfo.PageSize,
	}, "")
}
