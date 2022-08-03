package initialize

import (
	"fmt"
	"webApi/user_web/global"

	"github.com/hashicorp/consul/api"
)

func InitRegister() {

	Server := global.ServerConfig
	cfg := api.DefaultConfig()
	// consol 访问地址
	cfg.Address = fmt.Sprintf("%s:%d", Server.ConsulInfo.Host, Server.ConsulInfo.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 注册服务需要 健康检查的项
	check := &api.AgentServiceCheck{
		// HTTP:                           "http://172.16.1.163:8021/health", // 服务响应入口
		HTTP:                           fmt.Sprintf("http://%s:%d/health", Server.WebHost, Server.WebPort), // 服务响应入口
		Timeout:                        "5s",                                                               // 超时机制
		Interval:                       "5s",                                                               // 检测间隔时长
		DeregisterCriticalServiceAfter: "15s",
	}
	// 生成对应的检查服务对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = Server.Name
	registration.ID = Server.Name
	registration.Port = Server.WebPort
	registration.Tags = []string{"users", "web_user_api"}
	registration.Address = Server.WebHost
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}
