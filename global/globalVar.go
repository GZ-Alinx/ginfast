package global

import (
	//"go.uber.org/zap"
	"ginjujin/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局参数
var (
	RunMode string = "dev" // 运行模式 dev/sit/pre/prd
	Settings config.ServerConfig
	Lg *zap.Logger
	Trans ut.Translator
	DB *gorm.DB
	Redis *redis.Client
)