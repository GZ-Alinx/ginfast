package controller

import (
	"ginjujin/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)


// base64Captcha 缓存对象
var store = base64Captcha.DefaultMemStore

// GetCaptcha 获取校验码
func GetCaptcha(ctx *gin.Context) {
	// ps: base64Captcha.NewDriverDigit可以根据参数调节验证码参数
	driver := base64Captcha.NewDriverDigit(80,240,5,0.7,80)
	cp := base64Captcha.NewCaptcha(driver, store)

	// b64s 是图片的base64编码
	id,b64s,err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误：%s", err.Error())
		response.Err(ctx, http.StatusInternalServerError, 500, "生成验证码错误", "")
		return
	}

	// 返回图片验证码数据
	response.Success(ctx, 200, "生成验证码成功", gin.H{
		"captchaId": id,
		"picPath": b64s,
	})

}