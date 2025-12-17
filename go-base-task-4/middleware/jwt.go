package middleware

import (
	"github.com/gin-gonic/gin"
	"go-base-task-4/common/r"
	"go-base-task-4/utils"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := utils.GetToken(c)
		if token == "" {
			r.Error(c, http.StatusUnauthorized, nil, "未登录或非法访问，请登录")
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			r.Error(c, http.StatusUnauthorized, nil, "无效token")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
