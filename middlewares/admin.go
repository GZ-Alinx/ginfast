package middlewares




import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// isAdminAuth 判断权限  ps; 原理判断token的AuthorityId
func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取token信息
		claims,_ := ctx.Get("claims")
		// 获取现在用户信息
		currenUser := claims.(*CustomClaims)

		// 判断role权限
		if currenUser.AuthorityId != 1 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "用户没有权限",
			})
			// 中断下面中间件
			ctx.Abort()
			return
		}
		// 继续执行其他中间件
		ctx.Next()
	}
}