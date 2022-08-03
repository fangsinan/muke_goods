package initialize

import (
	"goods/goods_srv/global"

	"github.com/spf13/viper"
)

// 初始化config
func InitConfig(configFiles ...string) {

	var configFile string
	if len(configFiles) > 0 {
		configFile = configFiles[0]
	} else {
		configFile = "goods_srv/config/goods_srv_dev.yaml"
	}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
}
