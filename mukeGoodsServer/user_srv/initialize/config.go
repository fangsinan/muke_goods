package initialize

import (
	"goods/user_srv/global"

	"github.com/spf13/viper"
)

// 初始化config
func InitConfig() {
	configFile := "user_srv/config/user_srv_dev.yaml"
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
}
