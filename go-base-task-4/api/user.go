package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-base-task-4/common/r"
	"go-base-task-4/db"
	"go-base-task-4/models"
	"go-base-task-4/models/response"
	"go-base-task-4/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

// 注册用户
func RegisterUser(c *gin.Context) {
	var user = models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		r.ErrMsg(c, "用户参数错误")
		return
	}
	if err := db.SqlDB.Where("username = ?", user.Username).First(&user).Error; err == nil { // 判断用户名是否注册
		r.ErrMsg(c, "用户已注册")
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		r.ErrMsg(c, "密码解析错误")
		return
	}
	user.Password = string(hashedPassword)
	if err := db.SqlDB.Create(&user).Error; err != nil {
		r.ErrMsg(c, "用户创建失败")
		return
	}
	r.OK(c, "", "")

}

// 用户登录
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		r.ErrMsg(c, "用户参数错误")
		return
	}

	var storedUser models.User
	if err := db.SqlDB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		r.ErrMsg(c, "用户名错误")
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		r.ErrMsg(c, "密码错误")
		return
	}

	token, err := utils.LoginToken(storedUser)
	if err != nil {
		r.ErrMsg(c, err.Error())
		return
	}
	loginToken := response.LoginToken{
		Token: token,
	}
	r.OK(c, loginToken, "")
}
