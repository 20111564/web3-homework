package router

import (
	"github.com/gin-gonic/gin"
	"go-base-task-4/api"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//路由群组
	users := router.Group("api/user")
	{
		users.GET("/allUserList", api.GetAllUsers)
	}
	return router
}
