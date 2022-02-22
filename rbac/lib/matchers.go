package lib

import (
	"fmt"
	"strings"
)


// 添加匹配规则函数
var Admin=[]string{"admin","root"}
func InitRBACMatch() {
	// 判断普通用户权限
	E.AddFunction("methodMatch", func(arguments ...interface{}) (i interface{}, e error) {
		if len(arguments)==2{
			k1,k2:=arguments[0].(string),arguments[1].(string)
			return MethodMatch(k1,k2),nil
		}
		return nil,fmt.Errorf("methodMatch error")
	})

	// 判断是否是超级管理员
	E.AddFunction("IsSuper", func(arguments ...interface{}) (i interface{}, e error) {
		if len(arguments)==1{
			user := arguments[0].(string)
			return IsSpuerAdmin(user),nil
		}
		return nil,fmt.Errorf("SuperMatch error")
	})
}


// 超级管理员匹配
func IsSpuerAdmin(userName string) bool {
	for _,user := range Admin {
		if user == userName {
			return true
		}
	}
	return false
}

// 普通用户匹配
func MethodMatch(key1 string, key2 string) bool{
	ks:=strings.Split(key2," ")
	for _,s:=range ks{
		if s==key1{
			return true
		}
	}
	return false
}
