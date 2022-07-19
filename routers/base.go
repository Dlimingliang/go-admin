package routers

import (
	v1 "github.com/Dlimingliang/go-admin/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (receiver BaseRouter) InitBaseRouter(group *gin.RouterGroup) {
	baseApi := v1.ApiGroupAPP.BaseApi
	baseRouter := group.Group("base")
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
}
