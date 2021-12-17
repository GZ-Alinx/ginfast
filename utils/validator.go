package utils

import (
	"ginjujin/global"
	"ginjujin/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
	"strings"
)


//  HandleValidatorError 处理字段校验异常
func HandleValidatorError(c *gin.Context, err error) {
	// 如何返回错误信息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		response.Err(c, http.StatusBadRequest, 400, "字段校验错误", err.Error())
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": err.Error(),
		//})
	}


	msg := removeTopStruct(errs.Translate(global.Trans))
	response.Err(c,http.StatusBadRequest,400, "字段校验错误",msg)
	//c.JSON(http.StatusBadRequest, gin.H{
	//	"error": removeTopStruct(errs.Translate(global.Trans)),
	//})
	return
}

// removeTopStruct 定义一个去掉结构体名称前缀的自定义方案
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	// ps:removeTopStruct主要用作字符串的切分,因为用翻译成中文,返回的key里面前半段还是有英文,做切分处理
	for field,err := range fileds {
		// 从文本的逗号开始切分， 处理后"mobile":"mobile"为必填字段， 处理签： "PasswordLoginForm.mobile": "mobile为必填字段"
		rsp[field[strings.Index(field,".")+1:]] = err
	}
	return rsp
}


// 验证手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	// 利用反射拿到结构体tag含有mobile的key字段
	mobile := fl.Field().String()
	// 使用正则表达式判断是否合法
	ok,_ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
