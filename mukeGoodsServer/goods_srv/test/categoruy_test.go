package main

import (
	"context"
	"goods/goods_srv/global"
	"testing"

	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGetAllCategorysList(t *testing.T) {
	t.Run("GetAllCategorysLis", func(t *testing.T) {
		rsp, err := global.SrvClient.GetAllCategorysList(context.Background(), &emptypb.Empty{})
		if err != nil {
			t.Fatalf("[GetAllCategorysList]  %v", err)
		}
		t.Logf("total: %v", rsp.Total)
		t.Logf("data :%v", rsp.JsonData)
	})

}
