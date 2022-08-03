package routers

import (
	v1 "github.com/Dlimingliang/go-admin/api/v1"
	"github.com/Dlimingliang/go-admin/middleware"
	"github.com/gin-gonic/gin"
)

type CasbinRouter struct{}

func (c *CasbinRouter) InitCasbinRouter(group *gin.RouterGroup) {
	casbinApi := v1.ApiGroupAPP.CasbinApi

	casbinRouter := group.Group("casbin").Use(middleware.OperationRecord())
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
	casbinRouterWithoutRecord := group.Group("casbin").Use(middleware.OperationRecord())
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
