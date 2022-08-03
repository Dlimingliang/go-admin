package services

type Group struct {
	UserService
	RoleService
	MenuService
	DictionaryService
	DictionaryDetailService
	JwtService
	OperationRecordService
	SysApiService
	CasbinService
}

var ServiceGroupApp = new(Group)
