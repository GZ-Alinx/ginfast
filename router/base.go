package router

import (
	"ginjujin/controller"
	"github.com/gin-gonic/gin"

)


func InitBaseRouter(Router *gin.RouterGroup) {
	// 图片编码转换 http://tool.chinaz.com/tools/imgtobase/
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", controller.GetCaptcha)
	}
}
