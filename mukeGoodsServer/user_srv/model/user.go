package model

import (
	"errors"
	"goods/user_srv/global"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

func NewMysql() *UserModel {
	// 提至init
	return &UserModel{
		DB: global.DB,
	}
	// 原来写法
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // 慢 SQL 阈值
	// 		LogLevel:                  logger.Info, // 日志级别
	// 		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
	// 		Colorful:                  true,        // 禁用彩色打印
	// 	},
	// )
	// var err error
	// // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := username_pwd + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   "muxue_",
	// 		SingularTable: true,
	// 	},
	// 	Logger: newLogger,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// return &UserModel{
	// 	DB: mysqlDB,
	// }

}

// 公共字段  //muxue_goods
type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20);not null"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女 male表示男'"`
	Role     int        `gorm:"column:role;default:1;type:tinyint(1) comment '1表示普通用户 2表示管理员'"`
}

func (m UserModel) GetList(page int, size int) (int32, []User, error) {
	var users []User
	var total int32 = 0
	result := m.DB.Find(&users)

	userRes := m.DB.Scopes(Paginate(page, size)).Find(&users)
	if userRes.Error != nil {
		//无数据
		if errors.Is(userRes.Error, gorm.ErrRecordNotFound) {
			return 0, users, nil
		}
		return 0, nil, userRes.Error
	}
	total = int32(result.RowsAffected)
	return total, users, nil
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 2
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
