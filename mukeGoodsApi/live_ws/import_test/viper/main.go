package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type UserSrv struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name    string  `mapstructure:"name"`
	Port    int     `mapstructure:"webPort"`
	UserSrv UserSrv `mapstructure:"user_srv"`
}

func main() {
	v := viper.New()
	v.SetConfigFile("live_ws/import_test/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	ser := &ServerConfig{}
	if err := v.Unmarshal(ser); err != nil {
		panic(err)
	}
	fmt.Printf("%v", ser)
	// port := v.Get("webPort")
	// fmt.Printf("%t", port)
}
