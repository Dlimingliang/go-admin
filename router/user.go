package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

func InitUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("user")
	{
		userGroup.GET("getUserList", v1.GetUserList)
	}
}
