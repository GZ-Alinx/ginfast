package main

import (
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"time"

	_ "embed"
	"ginjujin/global"
	"ginjujin/initialize"
)

func main() {
	// 1.初始化yaml配置
	initialize.InitConfig()

	// 2. 初始化routers
	Router := initialize.Routers()
	color.Cyan("路由配置初始化成功~")

	// 3. 初始化日志信息
	initialize.InitLogger()
	color.Cyan("日志配置初始化成功~")

	// 4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	color.Cyan("翻译配置初始化成功~")

	// 5. 初始化mysql
	if err := initialize.InitMysqlDB();err != nil {
		color.Red("数据库初始化异常")
		panic(err)
	}
	color.Cyan("数据库初始化成功~")

	// 6. 初始化redis
	if err := initialize.InitRedis(); err != nil {
		color.Red("redis初始化失败")
		panic(err)
	}
	color.Cyan("Redis初始化成功~")

	// 7. Swagger初始化


	color.Yellow(">>>>>>>>>>>>>>>>>>>>>Gin后台服务启动成功~")
	global.Redis.Set("test","testValue", 500 * time.Second)
	value := global.Redis.Get("tets")
	color.Blue(value.Val())


	// 启动gin 并配置端口,global.Settings.Port是yaml配置的
	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("this is hello func", zap.String("error", "启动错误"))
	}
}
