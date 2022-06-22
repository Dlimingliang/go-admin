package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

func InitUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("user")
	{
		userGroup.POST("admin_register", v1.RegisterAdmin)
		userGroup.PUT("setUserInfo")
		userGroup.DELETE("deleteUser")
		userGroup.PUT("resetPassword")
	}

	userGroupWithoutRecord := group.Group("user")
	{
		userGroupWithoutRecord.POST("getUserList", v1.GetUserList)
	}
}
