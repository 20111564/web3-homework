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

// 新建文章
func CreatePost(c *gin.Context) {
	post := models.Post{}
	if err := c.ShouldBind(&post); err != nil {
		r.ErrMsg(c, "参数错误")
		return
	}
	userId := utils.GetUserID(c)
	post.UserID = uint(userId)
	if error := db.SqlDB.Create(&post).Error; error != nil {
		r.ErrMsg(c, "保存失败")
		return
	}
	r.OK(c, "", "")
}

// 文章列表
func PostPage(c *gin.Context) {
	var postPage request.PostPage
	if err := c.ShouldBind(&postPage); err != nil {
		r.ErrMsg(c, "参数错误")
		return
	}
	postService := &service.PostService{}
	list, total, err := postService.GetUserInfoList(postPage)
	if err != nil {
		r.ErrMsg(c, "查询失败")
		return
	}
	r.OK(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     postPage.Page,
		PageSize: postPage.PageSize,
	}, "")
}

// 文章详情
func PostDetail(c *gin.Context) {
	postId := c.Param("id")
	var post models.Post
	userId := utils.GetUserID(c)
	if err := db.SqlDB.Where("id = ? and user_id = ?", postId, userId).First(&post).Error; err != nil {
		r.ErrMsg(c, "文章正在创作中")
		return
	}
	r.OK(c, post, "")
}

// 文章编辑
func PostEdit(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBind(&post); err != nil {
		r.ErrMsg(c, err.Error())
		return
	}
	if post.ID <= 0 {
		r.ErrMsg(c, "文章ID必传")
	}
	userId := utils.GetUserID(c)
	if err := db.SqlDB.Model(&models.Post{}).Where("id = ? and user_id = ?", post.ID, userId).Updates(&post).Error; err != nil {
		r.ErrMsg(c, err.Error())
		return
	}
	r.OK(c, "", "")
}

// 文章删除
func PostDel(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBind(&post); err != nil {
		r.ErrMsg(c, err.Error())
	}
	if post.ID == 0 {
		r.ErrMsg(c, "参数错误")
		return
	}
	userId := utils.GetUserID(c)
	if err := db.SqlDB.Where("id = ? and user_id = ?", post.ID, userId).Delete(&models.Post{}).Error; err != nil {
		r.ErrMsg(c, err.Error())
		return
	}
	r.OK(c, "", "")
}
