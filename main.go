package main

import (
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/initialize"
)

func main() {
	initialize.InitViper()
	initialize.InitZap()
	initialize.InitGorm()
	ginRouter := initialize.InitGin()
	if err := ginRouter.Run(); err != nil {
		global.GaLog.Panic(err.Error())
	}
}
