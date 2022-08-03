package global

import (
	"goods/goods_srv/config"
	"goods/goods_srv/proto/v1"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	Host         = "localhost:3306"
	DBName       = "muxue_goods"
	Username_pwd = "root:123123"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	SrvClient    proto.GoodsClient
)

// 无用
func Dbinit() {
	// 设置全局的info级别的logger  打印sql
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
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := Username_pwd + "@tcp(" + Host + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "muxue_",
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
