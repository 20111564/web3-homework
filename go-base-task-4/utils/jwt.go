package utils

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v5"
	"go-base-task-4/models"
	"time"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("aaaaa"),
	}
}

type JwtClaims struct {
	ID       uint
	Username string
	jwt.RegisteredClaims
}

func (j *JWT) CreateClaims(jwtClaims JwtClaims) JwtClaims {
	claims := JwtClaims{
		ID: jwtClaims.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                                // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),              // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // 过期时间 7天
			Issuer:    "lu",                                                   // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		return token.Claims.(*JwtClaims), nil
	}
	return nil, errors.New("invalid token")
}

func LoginToken(user models.User) (token string, err error) {
	j := NewJWT()
	claims := j.CreateClaims(JwtClaims{
		ID:       user.ID,
		Username: user.Username,
	})
	token, err = j.CreateToken(claims)
	return
}
