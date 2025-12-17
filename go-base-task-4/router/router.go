package router

import (
	"github.com/gin-gonic/gin"
	"go-base-task-4/api"
	"go-base-task-4/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//公开路由
	publicGroup := router.Group("api/")
	{
		//用户
		publicUserGroup := publicGroup.Group("post/")
		{
			publicUserGroup.POST("api/user/register", api.RegisterUser)
			publicUserGroup.POST("api/user/login", api.Login)
		}
		//文章
		publicPostGroup := publicGroup.Group("post/")
		{
			publicPostGroup.POST("page", api.PostPage)
			publicPostGroup.GET("detail/:id", api.PostDetail)
		}
		//评论
		publicCommentGroup := publicGroup.Group("comment/")
		{
			publicCommentGroup.POST("page", api.CommentPage)
		}

	}

	//鉴权路由
	privateGroup := router.Group("api/")
	{
		//文章
		postGroup := privateGroup.Group("post/")
		{
			postGroup.POST("create", api.CreatePost)
			postGroup.POST("edit", api.PostEdit)
			postGroup.POST("del", api.PostDel)
		}
		//评论
		commentGroup := privateGroup.Group("comment/")
		{
			commentGroup.POST("add", api.CommentAdd)
		}

	}
	//jwt鉴权
	privateGroup.Use(middleware.JWTAuth())

	return router
}
