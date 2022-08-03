package main

import (
	"fmt"
	"webApi/user_web/global"
	"webApi/user_web/initialize"

	"go.uber.org/zap"
)

func main() {
	// 初始化logger 日志
	initialize.InitLogger()

	// 初始化 config 配置
	initialize.InitConfig()

	// 初始化router 路由
	Router := initialize.Routers()

	// 初始化validator翻译
	err := initialize.InitValidator("zh")
	if err != nil {
		panic(err)
	}

	// 初始化srv 服务连接
	initialize.InitSrvConn()

	// 初始化 手机号 正则验证
	// initialize.InitMobileValidator("mobileValidator")

	// 将web服务注册至consul中心
	initialize.InitRegister()

	// zap.S().Debug("serve start port: %d", global.ServerConfig.WebPort)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.WebPort)); err != nil {
		zap.S().Panic("cannot start server:%v", err.Error())
	}

}
