package utils

import (
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) string {
	return c.Request.Header.Get("x-token")
}

func GetClaims(c *gin.Context) (*JwtClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		return claims, err
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, err := GetClaims(c); err != nil {
		return 0
	} else {
		return claims.ID
	}
}
