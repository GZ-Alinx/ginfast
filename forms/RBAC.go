package forms





type PermissionArgs struct {
	//用户名
	Username string `form:"username" json:"username" binding:"required"`
	// 请求路由  binding:"required"为必填字段,长度大于3小于20
	URL  string `form:"url" json:"url"  binding:"required"`
	// 请求方法
	Method string `form:"mothod" json:"mothod" binding:"required"` // 验证码
}
