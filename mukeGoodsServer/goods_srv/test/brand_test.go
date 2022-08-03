package main

import (
	"context"
	"goods/goods_srv/global"
	"goods/goods_srv/model"
	"goods/goods_srv/proto/v1"
	"testing"
)

func TestBrandList(t *testing.T) {
	t.Run("list", func(t *testing.T) {
		rsp, err := global.SrvClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
		if err != nil {
			t.Fatalf("[BrandList]  %v", err)
		}
		t.Logf("total: %v", rsp.Total)
		t.Logf("data :%v", rsp.Data)
	})

}
func TestCreateBrand(t *testing.T) {

	t.Run("CreateBrand", func(t *testing.T) {
		t.Log("创建brand开始")
		rsp, err := global.SrvClient.CreateBrand(context.Background(), &proto.BrandRequest{
			Name: "pin1",
		})
		if err != nil {
			t.Fatalf("[CreateBrand]  %v", err)
		}

		t.Logf("创建brand id是 :%v", rsp.Id)
	})

}

func TestDeleteBrand(t *testing.T) {

	// var id int32 = 1113
	t.Run("DeleteBrand", func(t *testing.T) {
		var id int32 = 1113
		t.Log("删除brand开始")
		_, err := global.SrvClient.DeleteBrand(context.Background(), &proto.BrandRequest{
			Id: id,
		})
		if err != nil {
			t.Fatalf("[DeleteBrand]  %v", err)
		}

		t.Log("删除brand完成")

		t.Log("验证brand开始")
		brand := &model.Brands{}
		res := global.DB.First(brand, id)
		t.Logf("total: %v  data:%v", res.RowsAffected, brand)
	})

}
