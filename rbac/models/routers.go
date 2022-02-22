package models

import "fmt"

type Routers struct {
	RouterId int `gorm:"column:user_id;primaryKey"`
	RouterName string `gorm:"column:r_name"`
	RouterUri string `gorm:"column:r_uri;"`
	RouterMethod string `gorm:"column:r_method"`
	RoleName string
}

func(this *Routers) TableName() string{
	return "roles"
}
func(this *Routers) String() string{
	return fmt.Sprintf("%s-%s",this.RouterName,this.RoleName)
}
