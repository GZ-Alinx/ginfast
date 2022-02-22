package models

import "fmt"

type Users struct {
	UserID int `gorm:"column:user_id;primaryKey"`
	UserName string `gorm:"column:user_name"`
	RoleName string
}

func(this *Users) TableName() string {
	return "users"
}

func (this *Users) string() string  {
	return fmt.Sprintf("%s-%s", this.UserName, this.RoleName)
}