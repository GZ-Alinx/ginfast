package dao

import (
	"fmt"
	"ginjujin/forms"
	"ginjujin/global"
	"ginjujin/models"
	"github.com/fatih/color"
)

var users []models.User

// GetUserList 获取用户列表(page第几页，page——size每页几条数据)
func GetUserListDao(page int, page_size int)(int,[]interface{}) {
	// 分页用户列表数据
	userList := make([]interface{},0, len(users))
	// 计算偏移量
	offset := (page -1) * page_size
	// 查询所有的user
	result := global.DB.Offset(offset).Limit(page_size).Find(&users)

	if result.RowsAffected == 0 {
		return 0, userList
	}
	// 获取user总数
	total := len(users)

	// 查询数据
	result.Offset(offset).Limit(page_size).Find(&users)

	for _,useSingle := range users {
		birthday := ""
		if useSingle.Birthday == nil {
			birthday = ""
		} else {
			birthday = useSingle.Birthday.Format("2006-01-02")
		}
		userItemMap := map[string]interface{} {
			"id": useSingle.ID,
			"password":  useSingle.Password,
			"username":  useSingle.Username,
			"head_url":  useSingle.HeadUrl,
			"birthday":  birthday,
			"address":   useSingle.Address,
			"desc":      useSingle.Desc,
			"gender":    useSingle.Gender,
			"role":      useSingle.Role,
			"mobile":    useSingle.Mobile,
		}
		userList = append(userList, userItemMap)
	}
	// ps: 注释写的很清楚,最后返回的是 user表中总数和查询到数据
	return total, userList
}


// UsernameFindUserInfo 通过username找到用户信息
func FindUserInfo(username string, password string) (*models.User,bool) {
	var user models.User
	// 查询用户
	rows := global.DB.Where(&models.User{Username: username, Password:password}).Find(&user)
	fmt.Println(&user)
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}

func RegisterInsert(Userinfo *forms.UserRegistry) bool {
	var user models.User
	user.Username = Userinfo.Username
	user.Role = Userinfo.Role
	user.Password = Userinfo.Password
	user.Address = Userinfo.Address
	user.Desc = Userinfo.Desc
	user.Mobile = Userinfo.Mobile
	user.ID = 123

	result := global.DB.Create(&user)
	if result.Error != nil {
		color.Red("写入失败", result.Error)
		return false
	}
	return  true
}