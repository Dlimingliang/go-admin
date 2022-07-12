package router

type Group struct {
	UserRouter
	MenuRouter
	RoleRouter
	DictionaryRouter
}

var GroupApps = new(Group)
