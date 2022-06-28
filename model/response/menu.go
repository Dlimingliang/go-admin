package response

import "github.com/Dlimingliang/go-admin/model"

type MenuResult struct {
	Menu model.Menu `json:"menu"` //菜单信息
}

type MenuTree struct {
	MenuList []model.Menu `json:"menuList"` //菜单树
}
