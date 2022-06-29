package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

func InitRoleRouter(group *gin.RouterGroup) {
	roleGroup := group.Group("authority")
	{
		roleGroup.POST("createAuthority", v1.CreateRole)
		roleGroup.PUT("updateAuthority", v1.UpdateRole)
		roleGroup.POST("deleteAuthority", v1.DeleteRole)
	}
	roleGroupWithoutRecord := group.Group("authority")
	{
		roleGroupWithoutRecord.POST("getAuthorityList", v1.GetRoleList)
	}
}
