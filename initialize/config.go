package initialize

import (
	"ginjujin/config"
	"ginjujin/global"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()

	// 文件的路径如何设置
	switch global.RunMode {
		case "dev":
			v.SetConfigFile("./conf/settings-dev.yaml")
		case "sit":
			v.SetConfigFile("./conf/settings-sit.yaml")
		case "pro":
			v.SetConfigFile("./conf/settings-pro.yaml")
	}
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	// 给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}

	// 传递给全局变量
	global.Settings = serverConfig
	color.Blue("开始配置初始化中")
}