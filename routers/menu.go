package routers

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

type MenuRouter struct{}

func (menuRouter *MenuRouter) InitMenuRouter(group *gin.RouterGroup) {
	menuApi := v1.ApiGroupAPP.MenuApi
	menuGroup := group.Group("menu")
	{
		menuGroup.POST("addBaseMenu", menuApi.CreateMenu)            //新建菜单
		menuGroup.POST("updateBaseMenu", menuApi.UpdateMenu)         //修改菜单
		menuGroup.POST("deleteBaseMenu", menuApi.DeleteMenu)         //删除菜单
		menuGroup.POST("addMenuAuthority", menuApi.AddMenuAuthority) //增加角色和菜单关系
	}
	menuGroupWithoutRecord := group.Group("menu")
	{
		menuGroupWithoutRecord.POST("getMenuList", menuApi.GetMenuList)        //分页获取菜单列表
		menuGroupWithoutRecord.POST("getBaseMenuTree", menuApi.GetMenuTree)    //获取菜单树
		menuGroupWithoutRecord.POST("getBaseMenuById", menuApi.GetMenuById)    //根据id获取菜单信息
		menuGroupWithoutRecord.POST("getMenuAuthority", menuApi.GetMenuByRole) //根据角色获取菜单
	}
}
