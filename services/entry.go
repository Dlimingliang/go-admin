package services

type Group struct {
	UserService
	RoleService
	MenuService
	DictionaryService
	DictionaryDetailService
	JwtService
}

var ServiceGroupApp = new(Group)
