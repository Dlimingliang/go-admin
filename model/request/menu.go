package request

type MenuInfo struct {
	ID        int    `json:"ID"`
	Name      string `json:"title" binding:"required,min=1,max=20"`     //菜单名称
	Icon      string `json:"icon" binding:"required,min=1,max=20"`      //icon
	RoutePath string `json:"path" binding:"required,min=1,max=20"`      //路由path
	RouteName string `json:"name" binding:"required,min=1,max=20"`      // 路由名称
	Hidden    *bool  `json:"hidden" binding:"required"`                 // 是否隐藏
	Component string `json:"component" binding:"required,min=1,max=40"` //前端文件路径
	Sort      int    `json:"sort" binding:"required,min=1,max=999"`     //排序
	ParentId  *int   `json:"parentId" binding:"required"`               //父ID
}
