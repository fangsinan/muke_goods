package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"goods/user_srv/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() {
	// 设置Mysql全局的info级别的logger  打印sql
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	var err error

	// var (
	// 	Host         = "localhost:3306"
	// 	DBName       = "muxue_goods"
	// 	Username_pwd = "root:123123"
	// )
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := Username_pwd + "@tcp(" + Host + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.ServerConfig.Mysql.User,
		global.ServerConfig.Mysql.Password,
		global.ServerConfig.Mysql.Host,
		global.ServerConfig.Mysql.Port,
		global.ServerConfig.Mysql.Name,
	)

	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "muxue_",
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
