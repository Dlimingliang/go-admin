package v1

import "github.com/Dlimingliang/go-admin/service"

type ApiGroup struct {
	UserApi
	RoleApi
	MenuApi
	DictionaryApi
	DictionaryDetailApi
}

var ApiGroupAPPs = new(ApiGroup)

var (
	userService             = service.GroupApps.UserService
	menuService             = service.GroupApps.MenuService
	roleService             = service.GroupApps.RoleService
	dictionaryService       = service.GroupApps.DictionaryService
	dictionaryDetailService = service.GroupApps.DictionaryDetailService
)
