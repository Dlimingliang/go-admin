package v1

import "github.com/Dlimingliang/go-admin/services"

type ApiGroup struct {
	BaseApi
	UserApi
	RoleApi
	MenuApi
	DictionaryApi
	DictionaryDetailApi
	OperationRecordApi
}

var ApiGroupAPP = new(ApiGroup)

var (
	userService             = services.ServiceGroupApp.UserService
	menuService             = services.ServiceGroupApp.MenuService
	roleService             = services.ServiceGroupApp.RoleService
	dictionaryService       = services.ServiceGroupApp.DictionaryService
	dictionaryDetailService = services.ServiceGroupApp.DictionaryDetailService
	jwtService              = services.ServiceGroupApp.JwtService
	operationRecordService  = services.ServiceGroupApp.OperationRecordService
)
