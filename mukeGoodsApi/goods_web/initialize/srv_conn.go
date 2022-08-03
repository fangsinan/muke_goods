package initialize

import (
	"fmt"
	"webApi/goods_web/global"
	proto "webApi/goods_web/proto/v1"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// 负载调用
func InitSrvConn() {

	fmt.Println(global.ServerConfig)
	// 调用负载应用服务
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.GoodsSrv.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[GoodsSrv] cannot dial goods server", "mes", err.Error())
	}
	// 调用服务
	global.GoodsSrvClient = proto.NewGoodsClient(userConn)
}
