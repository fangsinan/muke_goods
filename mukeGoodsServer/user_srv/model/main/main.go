package main

// 处理mysql同步数据
import (
	"goods/user_srv/model"
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
			TablePrefix:   "muxue_",
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

func main() {

	db.AutoMigrate(&model.User{})

	get()
	add()
	save()
}

func get() {
	// fmt.Println("sql info:")
	// pro := NlsgProduct{}
	// res := db.First(&pro)

	// 条件语句  注意零值
	// res := db.Where(&NlsgProduct{Code: "111"}).First(&pro, []int{1}) // 推荐
	// res := db.Where("id = ?", 1).First(&pro) // 如果上述无法实现 则使用此方法

	// res := db.First(&pro, []int{1, 2}) // SELECT * FROM `nlsg_products` WHERE `nlsg_products`.`id` IN (1,2) ORDER BY `nlsg_products`.`id` LIMIT 1
	// res := db.Take(&pro)
	// res := db.Last(&pro)
	// fmt.Println(pro.ID)
	// fmt.Println(res.RowsAffected)

	// 多个查询需要slice
	// var pros []NlsgProduct
	// pres := db.Find(&pros) // 检索所有数据
	// fmt.Println()
	// if errors.Is(pres.Error, gorm.ErrRecordNotFound) {
	// 	fmt.Println("Sql Find() > 数据为空")
	// }
	// for i, pro := range pros {
	// 	fmt.Printf(" > 第%d条数据id为：%d \n", i, pro.ID)
	// }
	// fmt.Printf(" > 总数据为：%d \n", pres.RowsAffected)

}

func add() {
	// 迁移 schema
	// db.AutoMigrate(&NlsgProduct{})

	// product := NlsgProduct{
	// 	Code: "D44",
	// }
	// // Create
	// res := db.Create(&product)
	// fmt.Println(product.ID)
	// fmt.Println(res.RowsAffected)

	// 批量插入
	// products := []NlsgProduct{
	// 	{Code: "code1"},
	// 	{Code: "code2"},
	// }
	// // Create
	// res := db.Create(&products)
	// fmt.Println(res.RowsAffected)

}

func save() {
	// product := NlsgProduct{}
	// // 没有则新增 有则修改
	// db.First(&product)
	// // product.Code = "ddd"
	// // product.Price = 7
	// // res := db.Save(&product)
	// // fmt.Println(product.ID)
	// // fmt.Println(res.RowsAffected)

	// // 根据model 的值和where 修改 first 查询的语句
	// db.Model(&product).Where("code = ?", "xxx").Update("code", "yyy")
}
