package api

import (
	"fmt"
	"go-base-task-4/common/r"
	"go-base-task-4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 测试获取所有用户
func GetAllUsers(c *gin.Context) {
	user := models.User{}
	userList, err := user.GetUserList()
	fmt.Println(userList, err)
	if err != nil {
		userList = []models.User{}
	}
	r.OK(c, userList, "")
}

// 新增用户
func AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		//todo 参数错误处理
		r.Error(c, http.StatusBadRequest, err, "参数错误")
		return
	}

}
