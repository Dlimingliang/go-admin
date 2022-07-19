package routers

type Group struct {
	BaseRouter
	UserRouter
	MenuRouter
	RoleRouter
	DictionaryRouter
	DictionaryDetailRouter
}

var RouterGroupApp = new(Group)
