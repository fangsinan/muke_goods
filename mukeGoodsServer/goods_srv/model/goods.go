package model

//类型， 这个字段是否能为null， 这个字段应该设置为可以为null还是设置为空， 0
//实际开发过程中 尽量设置为不为null
//https://zhuanlan.zhihu.com/p/73997266
//这些类型我们使用int32还是int
type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32       `json:"parent_category_id"`
	ParentCategory   *Category   `json:"parent_category"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null"`
	ShipFree bool `gorm:"default:false;not null"`
	IsNew    bool `gorm:"default:false;not null"`
	IsHot    bool `gorm:"default:false;not null"`

	Name            string   `gorm:"type:varchar(50);not null"`
	GoodsSn         string   `gorm:"type:varchar(50);not null"`
	ClickNum        int32    `gorm:"type:int;default:0;not null"`
	SoldNum         int32    `gorm:"type:int;default:0;not null"`
	FavNum          int32    `gorm:"type:int;default:0;not null"`
	MarketPrice     float32  `gorm:"not null"`
	ShopPrice       float32  `gorm:"not null"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null"`
	Images          GormList `gorm:"type:varchar(1000);not null"`
	DescImages      GormList `gorm:"type:varchar(1000);not null"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null"`
}

// package model

// // 分类
// type Category struct {
// 	BaseModel
// 	Name             string `gorm:"type:varchar(50) comment '分类名';not null"`
// 	Level            int32  `gorm:"type:int;not null;default:1"`
// 	IsTab            bool   `gorm:"not null;default:false"`
// 	ParentCategoryID int32
// 	ParentCategory   *Category
// }

// // 品牌
// type Brands struct {
// 	BaseModel
// 	Name string `gorm:"type:varchar(50) comment '品牌名';not null"`
// 	Logo string `gorm:"type:varchar(200);not null;"`
// }

// // 品牌 分类关系表
// type GoodsCategoryBrands struct {
// 	BaseModel
// 	CategoryID int32 `gorm:"index:idx_category_brand,unique;type:int;not null"`
// 	//外键连接表
// 	Category Category
// 	BrandsID int32 `gorm:"index:idx_category_brand,unique;type:int;not null"`
// 	Brands   Brands
// }

// // 重载表名 默认 大写处下划线分割
// // func (GoodsCategoryBrands) TableName() string {
// // 	return "category_brands"
// // }

// // 轮播
// type Banner struct {
// 	BaseModel
// 	Image string `gorm:"type:varchar(50) comment '品牌名';not null"`
// 	Url   string `gorm:"type:varchar(200);not null;"`
// 	Index int32  `gorm:"type:int comment '排序';not null;default:1"`
// }

// // 商品表
// type Goods struct {
// 	BaseModel
// 	CategoryID int32 `gorm:"type:int;not null"`
// 	Category   Category
// 	BrandsID   int32 `gorm:"type:int;not null"`
// 	Brands     Brands

// 	Name             string   `gorm:"type:varchar(200);not null"`
// 	GoodsSn          string   `gorm:"type:varchar(200);not null"`
// 	ClickNum         int32    `gorm:"type:int comment '点击量';not null "`
// 	FavkNum          int32    `gorm:"type:int  comment '售卖量';not null"`
// 	MarketPrice      int32    `gorm:"type:int comment '价格';not null "`
// 	ShopPrice        int32    `gorm:"type:int comment '售卖价格';not null "`
// 	GoodsBrief       string   `gorm:"type:varchar(1000) comment '简介';not null"`
// 	Images           GormList `gorm:"type:varchar(1000) comment '图片展示';not null"`
// 	DescImages       GormList `gorm:"type:varchar(1000) comment '详情多图';not null "`
// 	GoodsFrontImages string   `gorm:"type:varchar(200) comment '商品封面图片';not null"`

// 	OnSale   bool `gorm:"type:tinyint(1) comment '是否上架';not null;default:false"`
// 	ShipFree bool `gorm:"not null;default:false;type:tinyint(1) comment '是否免运费'"`
// 	IsNew    bool `gorm:"type:tinyint(1) comment '是否新品'; default:false;not null"`
// 	IsHot    bool `gorm:"default:false;type:tinyint(1) comment '是否热卖'"`
// }
