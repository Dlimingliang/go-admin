package initialize

import (
	"github.com/Dlimingliang/go-admin/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/Dlimingliang/go-admin/docs"
	"github.com/Dlimingliang/go-admin/routers"
)

func InitRouters() *gin.Engine {
	ginRouter := gin.Default()
	//设置swagger
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//通过环境变量禁用swagger
	//ginRouter.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "prod"))

	//跨域中间件设置
	ginRouter.Use(cors.New(cors.Config{
		//直接放行全部跨域请求
		AllowAllOrigins: true,
		//AllowOrigins:           nil,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		MaxAge:           12 * time.Hour,
	}))

	systemRouterGroup := routers.RouterGroupApp
	PublicApiGroup := ginRouter.Group("")
	{
		systemRouterGroup.BaseRouter.InitBaseRouter(PublicApiGroup)
	}

	PrivateApiGroup := ginRouter.Group("")
	PrivateApiGroup.Use(middleware.JWTAuth())
	systemRouterGroup.InitMenuRouter(PrivateApiGroup)
	systemRouterGroup.InitRoleRouter(PrivateApiGroup)
	systemRouterGroup.InitUserRouter(PrivateApiGroup)
	systemRouterGroup.InitDictionaryRouter(PrivateApiGroup)
	systemRouterGroup.InitDictionaryDetailRouter(PrivateApiGroup)
	return ginRouter
}
