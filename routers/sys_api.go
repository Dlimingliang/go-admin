package routers

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
	"github.com/Dlimingliang/go-admin/middleware"
)

type SysApiRouter struct{}

func (s *SysApiRouter) InitSysApiRouter(group *gin.RouterGroup) {
	sysApi := v1.ApiGroupAPP.SysApiApi
	sysApiRouter := group.Group("api").Use(middleware.OperationRecord())
	{
		sysApiRouter.POST("createApi", sysApi.CreateSysApi)   // 创建Api
		sysApiRouter.POST("deleteApi", sysApi.DeleteSysApi)   // 删除Api
		sysApiRouter.POST("getApiById", sysApi.GetSysApiById) // 获取单条Api消息
		sysApiRouter.POST("updateApi", sysApi.UpdateSysApi)   // 更新Api
	}

	sysApiRouterWithoutRecord := group.Group("api")
	{
		sysApiRouterWithoutRecord.POST("getAllApis", sysApi.GetAllSysApi)  //获取所有Api
		sysApiRouterWithoutRecord.POST("getApiList", sysApi.GetSysApiPage) //分页获取Api
	}
}
