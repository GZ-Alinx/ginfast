package utils

import (
	"ginjujin/middlewares"
	"ginjujin/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateToken(c *gin.Context,id uint, UserName string, Role int) string {
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID: uint(id),
		UserName: UserName,
		AuthorityId: uint(Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// TODO 设置token过期时间
			ExpiresAt: time.Now().Unix() + 60*50*24*30, // token过期时间
			Issuer: "test",
		},
	}

	// 生成token
	token,err := j.CreateToken(claims)
	if err != nil {
		response.Success(c,401, "token生成失败,重新再试", "test")
		return ""
	}
	return token
}