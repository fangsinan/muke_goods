package handler

import (
	pb "goods/goods_srv/proto/v1"

	"gorm.io/gorm"
)

type GoodsServer struct {
	pb.UnimplementedGoodsServer
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
