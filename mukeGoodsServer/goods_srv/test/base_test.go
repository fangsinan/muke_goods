package main

import (
	"fmt"
	"goods/goods_srv/global"
	"goods/goods_srv/initialize"
	"goods/goods_srv/proto/v1"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// 单体应用调用
func init() {
	// 读config

	initialize.InitLogger()
	initialize.InitConfig(fmt.Sprintf("/Users/nlsg/sinan/workCode/golang/www/muke/mukeGoodsServer/goods_srv/config/goods_srv_dev.yaml"))
	initialize.InitDB()

	// 连接Consul  从注册信息获取 用户srv服务 的信息
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	serverData, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.Name))
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
	global.SrvClient = proto.NewGoodsClient(userConn)
}
