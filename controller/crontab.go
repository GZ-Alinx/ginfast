package controller

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"time"
)



// 计划任务实现 秒、分、时、日、月、周
func CronTest() {
	// date := time.Now().Format("2006-01-02 15:04:05")
	cron.New(cron.WithSeconds())
	c := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
		)))
	c.AddFunc("1 05 01 01 * *", func() {
		zap.L().Info(time.Now().Format("2006-01-02 15:04:05"))
		//fmt.Println("计划任务执行",time.Now().Format("2006-01-02 15:04:05"))
	})
	c.Start()
}
