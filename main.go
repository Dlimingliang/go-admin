package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/initialize"
)

func main() {
	initialize.InitViper()
	initialize.InitZap()
	initialize.InitGorm()
	ginRouter := initialize.InitGin()
	ginRouter.GET("/test", test)
	if err := ginRouter.Run(); err != nil {
		global.GaLog.Panic(err.Error())
	}
}

func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "123")
}
