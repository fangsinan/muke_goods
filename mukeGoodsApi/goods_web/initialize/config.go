package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"webApi/goods_web/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	// 设置 debug 的环境变量  如果为true
	debug := GetEnvInfo("debug")
	configFix := "goods_web/config/goods_web"
	configFileName := fmt.Sprintf("%s-dev.yaml", configFix)
	if debug {
		configFileName = fmt.Sprintf("%s-pro.yaml", configFix)
	}

	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 使用config时  需要先初始化数据  所以提取到 global 初始化参数
	serConfig := global.ServerConfig
	if err := v.Unmarshal(serConfig); err != nil {
		panic(err)
	}
	// zap.S().Infof("config init: %v", global.ServerConfig)

	// 动态监控config变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		zap.S().Infof("configFile changes Reload: %v", in.Name)
		// 重新加载 写入配置
		_ = v.ReadInConfig()
		v.Unmarshal(serConfig)
	})
}
