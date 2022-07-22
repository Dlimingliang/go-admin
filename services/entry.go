package services

type Group struct {
	UserService
	RoleService
	MenuService
	DictionaryService
	DictionaryDetailService
	JwtService
	OperationRecordService
}

var ServiceGroupApp = new(Group)
