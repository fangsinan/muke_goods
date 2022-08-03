package initialize

import "go.uber.org/zap"

func InitLogger() {
	// dev 环境
	logger, _ := zap.NewDevelopment()

	// logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}
