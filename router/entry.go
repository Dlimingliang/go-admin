package router

type Group struct {
	UserRouter
	MenuRouter
	RoleRouter
	DictionaryRouter
	DictionaryDetailRouter
}

var GroupApps = new(Group)
