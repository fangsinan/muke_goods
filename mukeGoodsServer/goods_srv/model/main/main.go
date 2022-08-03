package main

// 处理mysql同步数据
import (
	"goods/goods_srv/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	host         = "localhost:3306"
	dbName       = "muxue_goods"
	username_pwd = "root:123123"
)

// 定义db
var db *gorm.DB

func init() {
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
	dsn := username_pwd + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
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

func main() {

	db.AutoMigrate(
		&model.Category{},
		&model.Brands{},
		&model.GoodsCategoryBrand{},
		&model.Banner{},
		&model.Goods{},
	)
}
