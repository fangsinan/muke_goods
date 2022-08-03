package handler

import (
	"context"
	"encoding/json"
	"goods/goods_srv/global"
	"goods/goods_srv/model"
	pb "goods/goods_srv/proto/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

// //商品分类
func (s *GoodsServer) GetAllCategorysList(c context.Context, req *emptypb.Empty) (*pb.CategoryListResponse, error) {
	categorys := []model.Category{}

	res := global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	// res := global.DB.Preload("SubCategory.SubCategory").Find(&categorys)
	if res.Error != nil {
		return nil, res.Error
	}
	b, _ := json.Marshal(categorys)

	CategoryInfos := []*pb.CategoryInfoResponse{}
	for _, v := range categorys {
		info := &pb.CategoryInfoResponse{
			Id:    v.ID,
			Name:  v.Name,
			Level: v.Level,
		}
		CategoryInfos = append(CategoryInfos, info)
	}
	// fmt.Printf("\n\n\n\n%v\n\n\n\n", string(b))
	return &pb.CategoryListResponse{
		JsonData: string(b),
		Data:     CategoryInfos,
		Total:    int32(res.RowsAffected),
	}, nil
}

// //获取子分类
// func (s *GoodsServer) GetSubCategory(context.Context, *pb.CategoryListRequest) (*pb.SubCategoryListResponse, error)

// func (s *GoodsServer) CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error)
// func (s *GoodsServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
// func (s *GoodsServer) UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error)
