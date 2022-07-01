package v1

import "github.com/Dlimingliang/go-admin/service"

type ApiGroup struct {
	UserApi
	RoleApi
	MenuApi
}

var ApiGroupAPPs = new(ApiGroup)

var (
	userService = service.GroupApps.UserService
	menuService = service.GroupApps.MenuService
	roleService = service.GroupApps.RoleService
)
