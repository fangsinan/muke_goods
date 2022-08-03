package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

// 测试consul 工具的  服务注册、服务发现

// consul  服务注册
func Registe1r(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	// consol 访问地址
	cfg.Address = "192.168.3.4:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 注册服务需要 健康检查的项
	check := &api.AgentServiceCheck{
		HTTP:                           "http://192.168.3.4:8021/health", // 服务响应入口
		Timeout:                        "5s",                             // 超时机制
		Interval:                       "5s",                             // 检测间隔时长
		DeregisterCriticalServiceAfter: "5s",
	}
	// 生成对应的检查服务对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil

}

// consul  服务发现
func AllServices() {
	cfg := api.DefaultConfig()
	// consol 访问地址
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for k, v := range data {
		fmt.Printf("\nAllServices: key: %s  val: %v \n", k, v)
	}

}

// consul  服务过滤发现
func FilterService() {
	cfg := api.DefaultConfig()
	// consol 访问地址
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(`Service == "web_user_api"`)
	if err != nil {
		panic(err)
	}
	for k, v := range data {
		fmt.Printf("\nFilterService: key: %s  val: %v \n", k, v)
	}

}

func main() {

	// Register("192.168.3.4", 8021, "web_user_api", []string{"users"}, "web_user_api")
	Register("172.16.3.208", 8021, "user-web", []string{"mxshop", "bobby"}, "user-web")
	// AllServices()
	// FilterService()
}

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "172.16.3.208:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://172.16.3.208:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	// client.Agent().ServiceDeregister(id) // 注销服务发现  并不关闭进程
	if err != nil {
		panic(err)
	}
	return nil
}
