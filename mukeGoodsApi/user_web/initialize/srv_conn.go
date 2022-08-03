package initialize

import (
	"fmt"
	"webApi/user_web/global"
	userpb "webApi/user_web/proto/v1"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// 负载调用
func InitSrvConn() {

	// 调用负载应用服务
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.UserSrv.Name),
		// "consul://172.16.3.208:8500/user_srv?wait=5s",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[GetUserList] cannot dial user server", "mes", err.Error())
	}
	// defer userConn.Close()
	// 调用服务
	global.UserSrvClient = userpb.NewUserClient(userConn)
}

// 单体应用调用
func InitSrvConn1() {

	// 连接Consul  从注册信息获取 用户srv服务 的信息
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	serverData, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrv.Name))
	if err != nil {
		panic("InitSrvConn:" + err.Error())
	}
	userSrvHost := ""
	userSrvPort := 0
	for _, v := range serverData {
		userSrvHost = v.Address
		userSrvPort = v.Port
		break
	}

	if userSrvHost == "" {
		return
	}
	zap.S().Infof("{userSrvAddr} %s:%d", userSrvHost, userSrvPort)

	// 正常调用单体应用服务
	// Dial grpc
	// userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrv.Host, global.ServerConfig.UserSrv.Port), grpc.WithInsecure())
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] cannot dial user server", "mes", err.Error())
	}
	// 调用服务
	global.UserSrvClient = userpb.NewUserClient(userConn)
}
