package main

import (
	"context"
	"fmt"
	"log"

	proto "webApi/goods_web/proto/v1"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.Dial(
		"consul://172.16.3.208:8500/user_srv?wait=5s",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	i := 0
	for i < 3 {
		userC := proto.NewGoodsClient(conn)
		data, _ := userC.GetAllCategorysList(context.Background(), &emptypb.Empty{})
		fmt.Println(data)
		i++
	}

}
