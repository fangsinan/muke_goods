package main

import (
	"fmt"
	"webApi/live_ws/global"
	"webApi/live_ws/initialize"

	"go.uber.org/zap"
)

func main() {
	// 初始化logger 日志
	initialize.InitLogger()

	// 初始化 config 配置
	initialize.InitConfig()

	// 初始化router 路由
	Router := initialize.Routers()

	// zap.S().Debug("serve start port: %d", global.ServerConfig.WebPort)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.WebPort)); err != nil {
		zap.S().Panic("cannot start server:%v", err.Error())
	}

}
