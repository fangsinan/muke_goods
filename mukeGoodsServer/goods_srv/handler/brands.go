package handler

import (
	"context"
	"goods/goods_srv/global"
	"goods/goods_srv/model"
	pb "goods/goods_srv/proto/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// //品牌和轮播图
func (s *GoodsServer) BrandList(c context.Context, req *pb.BrandFilterRequest) (*pb.BrandListResponse, error) {
	// brandListResponse := pb.BrandListResponse{}

	var brands []model.Brands
	// res := global.DB.Find(&brands)
	// if res.Error != nil {
	// 	return nil, res.Error
	// }
	// 分页
	res := global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if res.Error != nil {
		return nil, res.Error
	}
	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)

	var brandResponses []*pb.BrandInfoResponse
	for _, brand := range brands {
		brandResponses = append(brandResponses, &pb.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}

	return &pb.BrandListResponse{
		// Total: int32(res.RowsAffected),
		Total: int32(total),
		Data:  brandResponses,
	}, nil
}

// 新建
func (s *GoodsServer) CreateBrand(c context.Context, req *pb.BrandRequest) (*pb.BrandInfoResponse, error) {
	if res := global.DB.First(&model.Brands{}, "name = ?", req.Name); res.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "brand is exists")
	}
	// 创建
	brand := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	global.DB.Save(brand)
	return &pb.BrandInfoResponse{
		Id: brand.ID,
	}, nil
}

func (s *GoodsServer) DeleteBrand(c context.Context, req *pb.BrandRequest) (*emptypb.Empty, error) {
	if res := global.DB.Delete(&model.Brands{}, req.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "brand is not exist")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateBrand(c context.Context, req *pb.BrandRequest) (*emptypb.Empty, error) {
	brand := &model.Brands{}
	if res := global.DB.First(brand, req.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "brand is not exist")
	}

	if req.Name != "" {
		brand.Name = req.Name
	}
	if req.Logo != "" {
		brand.Logo = req.Logo
	}
	global.DB.Save(brand)
	return &emptypb.Empty{}, nil
}
