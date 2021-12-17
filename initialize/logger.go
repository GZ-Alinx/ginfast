package initialize


import (
	"fmt"
	"ginjujin/utils"

	//"github.com/fatih/color"
	"go.uber.org/zap"
	"ginjujin/global"
)


// InitLogger 初始化Logger
func InitLogger() {
	// 实例化zap配置
	cfg := zap.NewDevelopmentConfig()

	// 配置日志输出的地址
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog%s.log", global.Settings.LogsAddress, utils.GetNowFormatTodayTime()),
		"stdout",
	}

	// 创建logger实例
	logg,_ := cfg.Build()
	zap.ReplaceGlobals(logg)
	global.Lg = logg
}





