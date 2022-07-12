package service

type Group struct {
	UserService
	RoleService
	MenuService
	DictionaryService
	DictionaryDetailService
}

var GroupApps = new(Group)
