package request

import "github.com/Dlimingliang/go-admin/model"

type MenuInfo struct {
	ID        int    `json:"ID"`
	RoutePath string `json:"path" binding:"required,min=1,max=20"`      //路由path
	RouteName string `json:"name" binding:"required,min=1,max=20"`      // 路由名称
	Hidden    *bool  `json:"hidden" binding:"required"`                 // 是否隐藏
	Component string `json:"component" binding:"required,min=1,max=40"` //前端文件路径
	Sort      int    `json:"sort" binding:"required,min=1,max=999"`     //排序
	ParentId  string `json:"parentId" binding:"required"`               //父ID
	MetaInfo  `json:"meta"`
}

type MetaInfo struct {
	Name string `json:"title" binding:"required,min=1,max=20"` //菜单名称
	Icon string `json:"icon" binding:"required,min=1,max=20"`  //icon
}

type AddMenuAuthorityInfo struct {
	Menus       []model.Menu `json:"menus"`       //菜单列表
	AuthorityId string       `json:"authorityId"` // 角色ID
}
