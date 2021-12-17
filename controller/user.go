package controller

import (
	"ginjujin/dao"
	"ginjujin/forms"
	"ginjujin/response"
	"ginjujin/utils"
	"github.com/gin-gonic/gin"
)

// 对用户进行增删改查




// GetUserList 获取用户列表

// @Summary 获取用户列表
// @Tags 用户管理
// @Router /user/list [post]
// @Param {object} body forms.UserListForm true "查询用户参数"
// @Success 200
// @Security ApiKeyAuth
func GetUserList(c *gin.Context) {
	UserListForm := forms.UserListForm{}
	if err := c.ShouldBind(&UserListForm); err != nil {
		utils.HandleValidatorError(c,err)
		return
	}
	// 获取数据
	total, userlist := dao.GetUserListDao(UserListForm.Page, UserListForm.PageSize)
	if (total + len(userlist)) == 0 {
		response.Err(c,400,400, "未获取到数据", map[string]interface{} {
			"total": total,
			"userlist": userlist,
		})
		return
	}

	response.Success(c,200, "获取用户列表成功", map[string]interface{}{
		"total": total,
		"userlist": userlist,
	})
}