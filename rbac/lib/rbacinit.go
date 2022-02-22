package lib

import (
	"ginjujin/global"
	"ginjujin/rbac/models"
)

type RoleRel struct {
	PRole string
	Role string
}

func (this *RoleRel) String() string {
	return this.PRole+":"+this.Role
}

// 获取角色
func GetRoles(pid int, m *[]*RoleRel, pname string) {
	proles := make([]*models.Role,0)
	global.DB.Where("role_pid=?",pid).Find(&proles)
	if len(proles)==0{
		return
	}
	for _,item := range proles {
		if pname!=""{
			*m=append(*m, &RoleRel{pname, item.RoleName})
		}
		GetRoles(item.RoleId, m,item.RoleName)
	}
}

// 获取用户和角色对应关系
func GetUserRoles() (users []*models.Users) {
	global.DB.Select("a.user_name, c.role_name").Table("users a,user_roles b, roles c").
		Where("a.user_id=b.user_id and b.role_id=c.role_id").
		Order("a.user_id desc").Find(&users)
	return
}


// 获取路由和角色对应关系
func GetRouterRoles() (routers []*models.Routers) {
	global.DB.Select("a.r_uri,a.r_method,c.role_name").
		Table("routers a,router_roles b,roles c ").
		Where(" a.r_id=b.router_id and b.role_id=c.role_id").
		Order(" role_name").Find(&routers)
	return
}

