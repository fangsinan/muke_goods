package handler

import (
	"context"
	pb "goods/goods_srv/proto/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

// //轮播图
func (s *GoodsServer) BannerList(context.Context, *emptypb.Empty) (*pb.BannerListResponse, error) {
	return &pb.BannerListResponse{}, nil
}

// func (s *GoodsServer) CreateBanner(context.Context, *BannerRequest) (*BannerResponse, error)
// func (s *GoodsServer) DeleteBanner(context.Context, *BannerRequest) (*emptypb.Empty, error)
// func (s *GoodsServer) UpdateBanner(context.Context, *BannerRequest) (*emptypb.Empty, error)
