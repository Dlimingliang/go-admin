package service

type Group struct {
	UserService
	RoleService
	MenuService
	DictionaryService
}

var GroupApps = new(Group)
