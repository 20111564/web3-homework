package e

import (
	"github.com/gin-gonic/gin"
)

/*
*
捕获全局异常返回
*/
func GlobalError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	c.Next()
}
