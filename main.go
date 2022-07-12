package main

import (
	"fmt"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/initialize"
	"github.com/Dlimingliang/go-admin/model"
)

// @title GA-API接口文档
// @description GA管理平台
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	initialize.InitViper()
	initialize.InitZap()
	initialize.InitGorm()
	ginRouter := initialize.InitRouters()
	initialize.InitValidator()
	migrate()

	if err := ginRouter.Run(fmt.Sprintf(":%d", global.ServerConfig.SystemConfig.ServerPort)); err != nil {
		global.GaLog.Panic(err.Error())
	}
}

func migrate() {
	global.GaDb.AutoMigrate(&model.User{})
	global.GaDb.AutoMigrate(&model.Menu{})
	global.GaDb.AutoMigrate(&model.Role{})
	global.GaDb.AutoMigrate(&model.Dictionary{})
	global.GaDb.AutoMigrate(&model.DictionaryDetail{})
}
