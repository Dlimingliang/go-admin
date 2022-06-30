package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

func InitMenuRouter(group *gin.RouterGroup) {
	menuGroup := group.Group("menu")
	{
		menuGroup.POST("addBaseMenu", v1.CreateMenu)            //新建菜单
		menuGroup.POST("updateBaseMenu", v1.UpdateMenu)         //修改菜单
		menuGroup.POST("deleteBaseMenu", v1.DeleteMenu)         //删除菜单
		menuGroup.POST("addMenuAuthority", v1.AddMenuAuthority) //增加角色和菜单关系
	}
	menuGroupWithoutRecord := group.Group("menu")
	{
		menuGroupWithoutRecord.POST("getMenuList", v1.GetMenuList)        //分页获取菜单列表
		menuGroupWithoutRecord.POST("getBaseMenuTree", v1.GetMenuTree)    //获取菜单树
		menuGroupWithoutRecord.POST("getBaseMenuById", v1.GetMenuById)    //根据id获取菜单信息
		menuGroupWithoutRecord.POST("getMenuAuthority", v1.GetMenuByRole) //根据角色获取菜单
	}
}
