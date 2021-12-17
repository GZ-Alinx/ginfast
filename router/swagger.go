package router

import (
	_ "ginjujin/docs"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)




// @title 基础开发框架
// @version 0.0.1
// @description 通用gin框架开发基础框架/前后端分离
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /v1
func SwaggerRouter(Router *gin.RouterGroup) {
	// 访问 http://127.0.0.1:8080/v1/user/list
	UserRouter := Router.Group("swagger")
	{
		UserRouter.GET("/*any", gs.WrapHandler(swaggerFiles.Handler))
	}
}
