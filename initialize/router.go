package initialize

import (
	"ginjujin/middlewares"
	"ginjujin/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	Router.Use(middlewares.Core())

	// 路由分组
	ApiGroup := Router.Group("/v1/")
	router.UserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	router.SwaggerRouter(ApiGroup)

	return Router
}
