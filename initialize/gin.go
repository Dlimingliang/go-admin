package initialize

import "github.com/gin-gonic/gin"

func InitGin() *gin.Engine {
	ginRouter := gin.Default()
	return ginRouter
}
