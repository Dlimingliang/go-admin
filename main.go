package main

import (
	"fmt"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/initialize"
	"github.com/Dlimingliang/go-admin/model"
)

func main() {
	initialize.InitViper()
	initialize.InitZap()
	initialize.InitGorm()
	ginRouter := initialize.InitRouters()
	//migrate()

	if err := ginRouter.Run(fmt.Sprintf(":%d", global.ServerConfig.SystemConfig.ServerPort)); err != nil {
		global.GaLog.Panic(err.Error())
	}
}

func migrate() {
	global.GaDb.AutoMigrate(&model.User{})
}
