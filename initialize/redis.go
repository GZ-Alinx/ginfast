package initialize

import (
	"fmt"
	"ginjujin/global"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

func InitRedis() error {
	addr := fmt.Sprintf("%s:%d", global.Settings.RedisInfo.Host, global.Settings.RedisInfo.Port)
	// 生成redis客户端
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.Settings.RedisInfo.Password,
		DB:       global.Settings.RedisInfo.DB,
	})

	// 链接redis
	_, err := global.Redis.Ping().Result()
	if err != nil {
		color.Red("[InitRedis] 链接redis异常:")
		color.Yellow(err.Error())
		return err
	}
	return nil
}
