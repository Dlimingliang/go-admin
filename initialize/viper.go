package initialize

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/Dlimingliang/go-admin/config"
	"github.com/Dlimingliang/go-admin/global"
)

/**
viper
优先级: 命令行 > 环境变量 > 默认配置
*/

func InitViper() {

	var configFile string
	//依次从命令行>环境变量>默认配置读取配置文件地址
	flag.StringVar(&configFile, config.FileEnv, "", "配置文件地址.")
	flag.Parse()
	if configFile == "" {
		if configEnv := os.Getenv(config.FileEnv); configEnv == "" {
			switch gin.Mode() {
			case gin.DebugMode:
				configFile = config.DefaultFile
			case gin.ReleaseMode:
				configFile = config.ProdFile
			case gin.TestMode:
				configFile = config.TestFile
			}
			fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, configFile)
		} else {
			configFile = configEnv
			fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", config.FileEnv, configFile)
		}
	} else {
		fmt.Printf("您正在使用命令行的-%s参数传递的值,config的路径为%s\n", config.FileEnv, configFile)
	}

	//读取配置文件
	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	//检测配置文件变化，自动发生变更
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变:", in.Name)
		if err = v.Unmarshal(global.ServerConfig); err != nil {
			fmt.Println(err)
		}
	})

	//转换配置到struct
	if err = v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
}
