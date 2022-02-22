package router

import (
	"ginjujin/rbac"
	"github.com/gin-gonic/gin"
)

func RbacAddToUser(context *gin.RouterGroup) {
	RBACRouter := context.Group("rbac")
	{
		RBACRouter.POST("add", rbac.AddPermissionForUsers)
	}
}
