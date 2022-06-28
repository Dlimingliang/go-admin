package service

import (
	"github.com/Dlimingliang/go-admin/core/business"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
)

type MenuService struct{}

func (menuService *MenuService) GetMenuTree() ([]model.Menu, error) {
	var menuList []model.Menu
	err := global.GaDb.Where(&model.Menu{ParentId: 0}).Preload("Children").Find(&menuList).Error
	return menuList, err
}

func (menuService *MenuService) GetMenuById(id int) (model.Menu, error) {
	var menu model.Menu
	err := global.GaDb.First(&menu, id).Error
	return menu, err
}

func (menuService *MenuService) CreateMenu(req model.Menu) (model.Menu, error) {

	var menu model.Menu
	//验证菜单名称重复
	if result := global.GaDb.Where(&model.Menu{Name: req.Name}).First(&menu); result.RowsAffected > 0 {
		return menu, business.New("该菜单名称已存在")
	}
	//创建菜单
	err := global.GaDb.Create(&req).Error
	return req, err
}

func (menuService *MenuService) UpdateMenu(req model.Menu) error {
	return global.GaDb.Updates(req).Error
}

func (menuService *MenuService) DeleteMenu(id int) error {
	return global.GaDb.Delete(&model.Menu{}, id).Error
}
