package initialize

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/Dlimingliang/go-admin/global"
)

func InitViper() {
	v := viper.New()
	v.SetConfigFile("config.yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = v.Unmarshal(global.ServerConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(global.ServerConfig)
}
