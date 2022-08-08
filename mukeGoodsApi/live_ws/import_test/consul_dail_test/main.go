package main

import (
	"context"
	"fmt"
	"log"

	userpb "webApi/user_web/proto/v1"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
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
		userC := userpb.NewUserClient(conn)
		data, _ := userC.GetUserList(context.Background(), &userpb.PageInfo{
			Pn:    uint32(1),
			PSize: uint32(10),
		})
		fmt.Println(data)
		i++
	}

}
