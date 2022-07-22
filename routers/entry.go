package routers

type Group struct {
	BaseRouter
	UserRouter
	MenuRouter
	RoleRouter
	DictionaryRouter
	DictionaryDetailRouter
	OperationRecordRouter
}

var RouterGroupApp = new(Group)
