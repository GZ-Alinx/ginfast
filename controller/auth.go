package controller

import (
	"ginjujin/dao"
	"ginjujin/forms"
	"ginjujin/global"
	"ginjujin/models"
	"ginjujin/response"
	"ginjujin/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)


// 用户登录
// @Summary 用户登录
// @Tags 用户认证/登录/注册
// @Router /user/login [post]
// @Param {object} body forms.LoginForm true "登录用户名"
// @Success 200
func Login(c *gin.Context) {
	LoginForm := forms.LoginForm{}
	// ps:ShouldBind()函数里面传入PasswordLoginForm实例,会根据结构体的tag校验字段
	if err := c.ShouldBind(&LoginForm); err != nil {
		color.Blue(err.Error(),err)
		utils.HandleValidatorError(c,err)
		return
	}

	// 数字验证码验证
	if !store.Verify(LoginForm.CaptchaId, LoginForm.Captcha, true) {
		response.Err(c, 400, 400,"验证码错误","")
		return
	}

	// 查询数据库是否有该用户
	user,ok := dao.FindUserInfo(LoginForm.Username,LoginForm.Password)
	if !ok {
		response.Err(c,401, 401,"未注册该用户","")
		return
	}
	token := utils.CreateToken(c, user.ID, user.Username, user.Role)
	userinfoMap := HandleUserModelToMap(user)
	userinfoMap["token"] = token
	global.Redis.Set(user.Username,token, 172800 * time.Second)

	// 返回数据
	response.Success(c,200,"success", userinfoMap)
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "success",
	//})
}


// 用户信息格式化
func HandleUserModelToMap(user *models.User) map[string]interface{} {
	birthday := ""
	if user.Birthday == nil {
		birthday = ""
	} else {
		birthday = user.Birthday.Format("2006-01-02")
	}
	userItemMap := map[string]interface{}{
		"id":        user.ID,
		"username":  user.Username,
		"head_url":  user.HeadUrl,
		"birthday":  birthday,
		"address":   user.Address,
		"desc":      user.Desc,
		"gender":    user.Gender,
		"role":      user.Role,
		"mobile":    user.Mobile,
	}
	return userItemMap
}


// 用户注册
// @Summary 用户注册
// @Tags 用户认证/登录/注册
// @Router /user/registry [post]
// @Param {object} body forms.UserRegistry true "注册参数"
// @Success 200
func UserRegistry(c *gin.Context) {
	var User = forms.UserRegistry{}
	if err := c.ShouldBind(&User); err != nil {
		color.Red("请求信息绑定错误")
		color.Blue(err.Error(),err)
		utils.HandleValidatorError(c,err)
		return
	}
	if ok := dao.RegisterInsert(&User); !ok {
		color.Red("注册失败")
		zap.L().Error("用户注册失败")
		return
	}
	response.Success(c, http.StatusOK,"注册成功", User)
}

