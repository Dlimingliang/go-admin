package router

type Group struct {
	UserRouter
	MenuRouter
	RoleRouter
}

var GroupApps = new(Group)
