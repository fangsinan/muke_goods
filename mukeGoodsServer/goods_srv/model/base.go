package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// 公共字段  //muxue_goods
type BaseModel struct {
	ID        int32          `gorm:"primarykey;type:int"  json:"id"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"-"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	IsDeleted bool           `json:"-"`
}

// 重定义数据类型
type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GormList) Scan(value interface{}) error {
	err := json.Unmarshal(value.([]byte), &g)
	return err
}
