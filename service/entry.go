package service

type Group struct {
	UserService
	RoleService
	MenuService
}

var GroupApps = new(Group)
