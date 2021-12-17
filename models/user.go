package models


import (
	"time"
)


// 用户表字段
type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Password string `json:"password" gorm:"password"`
	Username string `json:"username" gorm:"username"`
	HeadUrl string `json:"head_url" gorm:"head_url"`
	Birthday *time.Time `json:"birthday" gorm:"type:date"`
	Address string `json:"address" gorm:"address"`
	Desc string `json:"desc" gorm:"desc"`
	Gender string `json:"gender" gorm:"gender"`
	Role int `json:"role" gorm:"role"`
	Mobile string `json:"mobile" gorm:"mobile"`
}

//用户表名称
func (User) TableName() string {
	return "user"
}
