package models


import (
	"fmt"
)

type Role struct {
	RoleId int `gorm:"column:role_id;primaryKey"`
	RoleName string `gorm:"column:role_name"`
	RolePid int `gorm:"column:role_pid"`
	RoleComment string `gorm:"column:role_comment"`
}
func (this *Role) TableName() string {
	return "roles"
}

func (this *Role) String() string {
	return fmt.Sprintf("ID:%d 角色名:%s", this.RoleId, this.RoleName)
}
