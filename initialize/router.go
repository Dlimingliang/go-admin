package initialize

import (
	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/router"
)

func InitRouters() *gin.Engine {
	ginRouter := gin.Default()

	apiGroup := ginRouter.Group("")
	router.InitUserRouter(apiGroup)
	return ginRouter
}
