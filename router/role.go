package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

type RoleRouter struct{}

func (roleRouter *RoleRouter) InitRoleRouter(group *gin.RouterGroup) {
	roleApi := v1.ApiGroupAPPs.RoleApi
	roleGroup := group.Group("authority")
	{
		roleGroup.POST("createAuthority", roleApi.CreateRole)
		roleGroup.PUT("updateAuthority", roleApi.UpdateRole)
		roleGroup.POST("deleteAuthority", roleApi.DeleteRole)
	}
	roleGroupWithoutRecord := group.Group("authority")
	{
		roleGroupWithoutRecord.POST("getAuthorityList", roleApi.GetRoleList)
	}
}
