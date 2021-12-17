package router

import (
	"ginjujin/controller"
	"ginjujin/middlewares"
	"github.com/gin-gonic/gin"
)


func UserRouter(Router *gin.RouterGroup) {
	// 访问 http://127.0.0.1:8080/v1/user/list
	UserRouter := Router.Group("user")
	{
		//UserRouter.GET("list", func(c *gin.Context) {
		//	c.JSON(200, gin.H{
		//		"message": "pong",
		//	})
		//})
		UserRouter.POST("login", controller.Login)
		UserRouter.POST("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(),controller.GetUserList)
		UserRouter.POST("registry", controller.UserRegistry)
	}
}
