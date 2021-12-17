package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 成功返回
func Success(c *gin.Context, code int, msg interface{},data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{} {
		"code":code, // 自定义代码
		"msg":msg,	 // 消息
		"data":data, // 数据
	})
	return
}


// 错误返回
func Err(c *gin.Context, httpCode int, code int, msg string, jsonStr interface{}) {
	c.JSON(httpCode, map[string]interface{} {
		"code":code,
		"msg":msg,
		"data":jsonStr,
	})
	return
}