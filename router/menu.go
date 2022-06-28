package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

func InitMenuRouter(group *gin.RouterGroup) {
	menuGroup := group.Group("menu")
	{
		menuGroup.POST("addBaseMenu", v1.CreateMenu)
		menuGroup.POST("updateBaseMenu", v1.UpdateMenu)
		menuGroup.POST("deleteBaseMenu", v1.DeleteMenu)
	}
	menuGroupWithoutRecord := group.Group("menu")
	{
		menuGroupWithoutRecord.POST("getMenu", v1.GetMenuTree)
		menuGroupWithoutRecord.POST("getBaseMenuById", v1.GetMenuById)
	}
}
