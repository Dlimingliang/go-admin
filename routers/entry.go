package routers

type Group struct {
	BaseRouter
	UserRouter
	MenuRouter
	RoleRouter
	DictionaryRouter
	DictionaryDetailRouter
	OperationRecordRouter
	SysApiRouter
	CasbinRouter
}

var RouterGroupApp = new(Group)
