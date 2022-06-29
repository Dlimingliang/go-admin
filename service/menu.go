package service

import (
	"gorm.io/gorm"

	"github.com/Dlimingliang/go-admin/core/business"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
)

type MenuService struct{}

func (menuService *MenuService) GetMenuTree() ([]model.Menu, error) {
	var menuList []model.Menu
	err := global.GaDb.Where("parent_id = ?", 0).Preload("Children", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort, create_time desc")
	}).Order("sort, create_time desc").Find(&menuList).Error
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
	if result := global.GaDb.Where("name = ?", req.Meta.Name).First(&menu); result.RowsAffected > 0 {
		return menu, business.New("该菜单名称已存在")
	}
	//创建菜单
	err := global.GaDb.Create(&req).Error
	return req, err
}

func (menuService *MenuService) UpdateMenu(req model.Menu) error {
	updateMap := make(map[string]interface{})
	updateMap["parent_id"] = req.ParentId
	updateMap["route_path"] = req.RoutePath
	updateMap["route_name"] = req.RouteName
	updateMap["hidden"] = req.Hidden
	updateMap["component"] = req.Component
	updateMap["sort"] = req.Sort
	updateMap["name"] = req.Name
	updateMap["icon"] = req.Icon
	return global.GaDb.Updates(updateMap).Error
}

func (menuService *MenuService) DeleteMenu(id int) error {
	if result := global.GaDb.Where("parent_id = ?", id).First(&model.Menu{}); result.RowsAffected > 0 {
		return business.New("此菜单存在子菜单不可删除")
	}
	//删除菜单及关联的角色
	return global.GaDb.Delete(&model.Menu{}, id).Error
}
