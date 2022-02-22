package rbac

import (
	"fmt"
	"ginjujin/forms"
	"ginjujin/rbac/lib"
	"ginjujin/response"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"regexp"
)


//  服务级别的API根据官方文档进行调整编写  https://casbin.org/docs/zh-CN/rbac-api


// 检查是否登录  这里需要拿到用户识别标识符

func CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.RequestURI != "/v1/user/registry" && context.Request.Header.Get("token")==""{
			context.AbortWithStatusJSON(400, gin.H{"message":"token required"})
		}else {
			context.Set("user_name", context.Request.Header.Get("token"))
			context.Next()
		}
	}
}


// 验证所有访问用户和路由权限，并排除注册接口
func RBAC() gin.HandlerFunc {
	return func(context *gin.Context) {
		user,_:= context.Get("user_name")
		fmt.Println(user,context.Request.RequestURI, context.Request.Method)
		accessRegistry,err := regexp.MatchString(context.Request.RequestURI,"/v1/user/registry")
		accessLogin,err := regexp.MatchString(context.Request.RequestURI,"/v1/user/login")
		access,err := lib.E.Enforce(user,context.Request.RequestURI, context.Request.Method)
		if accessRegistry || accessLogin {
			context.Next()
		} else if err != nil || !access  {
			context.AbortWithStatusJSON(403, gin.H{"message":"forbidden"})
		} else {
			context.Next()
		}
	}
}

// 添加用户权限
func AddPermissionForUsers(c *gin.Context) {
	PermissionArgs := forms.PermissionArgs{}
	if err:=c.ShouldBind(&PermissionArgs);err != nil {
		color.Red(err.Error(),err)
		return
	}


	ok,err := lib.E.AddPolicy(PermissionArgs.Username,PermissionArgs.URL, PermissionArgs.Method)
	if err != nil || !ok {
		zap.L().Error("权限增加失败", zap.Error(err))
		response.Err(c,403,403,"权限添加失败","")
		return
	}
	// 成功返回
	response.Success(c, 200, "success","")
	return
}

// 添加用户权限并对接角色
func RBACAddForRole() {

}









func RBACMiddlewares() (fs []gin.HandlerFunc) {
	lib.InitRBACDBConfig()
	lib.InitRBACMatch()
	fs = append(fs, CheckLogin(), RBAC())
	color.Cyan("RBAC模型初始化成功~")
	return
}

