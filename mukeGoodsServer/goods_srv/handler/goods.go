package handler

import (
	"context"
	pb "goods/goods_srv/proto/v1"
)

// //商品接口
func (s *GoodsServer) GoodsList(context.Context, *pb.GoodsFilterRequest) (*pb.GoodsListResponse, error) {
	return &pb.GoodsListResponse{}, nil
}

// //现在用户提交订单有多个商品，你得批量查询商品的信息吧
// func (s *GoodsServer) BatchGetGoods(context.Context, *pb.BatchGoodsIdInfo) (*pb.GoodsListResponse, error)
// func (s *GoodsServer) CreateGoods(context.Context, *pb.CreateGoodsInfo) (*pb.GoodsInfoResponse, error)
// func (s *GoodsServer) DeleteGoods(context.Context, *pb.DeleteGoodsInfo) (*emptypb.Empty, error)
// func (s *GoodsServer) UpdateGoods(context.Context, *pb.CreateGoodsInfo) (*emptypb.Empty, error)
// func (s *GoodsServer) GetGoodsDetail(context.Context, *pb.GoodInfoRequest) (*pb.GoodsInfoResponse, error)
