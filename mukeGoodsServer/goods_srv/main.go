package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"goods/goods_srv/global"
	"goods/goods_srv/handler"
	"goods/goods_srv/initialize"
	pb "goods/goods_srv/proto/v1"
	"goods/goods_srv/utils"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// 配置命令行化的 ip 和port  eg: main -id -port
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "port")
	// flag.Parse()
	// addRess := fmt.Sprintf("%s:%s", *IP, *Port)
	// fmt.Println("> listen to address:", addRess)

	//初始化工作
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}

	if *IP == "0.0.0.0" {
		*IP = global.ServerConfig.Host
	}

	flag.Parse()
	addRess := fmt.Sprintf("%s:%d", *IP, *Port)
	zap.S().Infof("listen: %s", addRess)

	g := grpc.NewServer()

	pb.RegisterGoodsServer(g, &handler.GoodsServer{})
	lis, err := net.Listen("tcp", addRess)
	if err != nil {
		panic("listen:" + err.Error())
	}

	// 注册健康检查
	grpc_health_v1.RegisterHealthServer(g, health.NewServer())

	// 服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",
		global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port,
	)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// grpc 服务检查项
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", *IP, *Port),
		Timeout:                        "5s", // 超时机制
		Interval:                       "5s", // 检测间隔时长
		DeregisterCriticalServiceAfter: "15s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name

	serviceID := fmt.Sprintf("%s - %s", global.ServerConfig.Name, uuid.NewV4())
	registration.ID = serviceID
	registration.Port = *Port
	registration.Tags = []string{"goods", "goods_srv"}
	registration.Address = *IP
	registration.Check = check
	client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	// 启用携程  部署负载服务
	go func() {
		err = g.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接收终止信号  删除注册在consul的服务
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 删除服务
	if err = client.Agent().ServiceDeregister(serviceID); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}
