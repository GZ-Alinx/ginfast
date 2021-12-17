package forms


/*
注意几个点:
	1. Captcha的最大值和最小值都为5,应为后续设置获取验证码的位数是5
	2. PasswordLoginForm修改后,login的接口就必须传入验证码和验证id,否则报错
*/

type LoginForm struct {
	//用户名
	Username string `form:"username" json:"username" binding:"required"`
	// 密码  binding:"required"为必填字段,长度大于3小于20
	Password  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	// 手机号验证
	//Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Captcha string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"` // 验证码
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"` // 验证码ID
}


type UserRegistry struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Mobile string 	`form:"mobile" json:"mobile"`
	Address string  `form:"address" json:"address"`
	Desc string 	`form:"desc" json:"desc"`
	Role int 		`form:"role" json:"role"`
}


type UserListForm struct {
	Page  int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"pagesize" json:"pagesize" binding:"required"`
}