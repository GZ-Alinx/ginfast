package initialize


import (
	"fmt"
	"ginjujin/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() error {
	// 参考资料： https://github.com/go-sql-driver/mysql#dsn-data-source-name
	mysqlInfo := global.Settings.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.User,
		mysqlInfo.PassWord,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.DBName)
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("数据库链接错误，err:",zap.String("err",err.Error()))
		return err
	}
	global.DB = db
	return nil
}