package service

import (
	"github.com/Dlimingliang/go-admin/core/business"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
)

type RoleService struct {
}

var RoleServiceInstance = new(RoleService)

func (roleService *RoleService) GetRoleList() ([]model.Role, error) {
	var roleList []model.Role
	err := global.GaDb.Find(&roleList).Error
	if err != nil {
		return roleList, err
	}
	roleMap := make(map[string][]model.Role)
	for _, role := range roleList {
		roleMap[role.ParentId] = append(roleMap[role.ParentId], role)
	}
	rootRoleList := roleMap["0"]
	for i := 0; i < len(rootRoleList); i++ {
		setChildrenRole(&rootRoleList[i], roleMap)
	}
	return rootRoleList, err
}

func setChildrenRole(role *model.Role, roleMap map[string][]model.Role) {
	role.Children = roleMap[role.AuthorityId]
	for i := 0; i < len(role.Children); i++ {
		setChildrenRole(&role.Children[i], roleMap)
	}
}

func (roleService *RoleService) CreateRole(req model.Role) (model.Role, error) {
	var role model.Role
	if result := global.GaDb.Where("authority_id = ?", req.AuthorityId).First(&role); result.RowsAffected > 0 {
		return role, business.New("该角色code已存在")
	}
	if result := global.GaDb.Where("authority_name = ?", req.AuthorityName).First(&role); result.RowsAffected > 0 {
		return role, business.New("该角色名称已存在")
	}

	err := global.GaDb.Create(&req).Error
	return req, err
}

func (roleService *RoleService) UpdateRole(req model.Role) error {
	return global.GaDb.Where("authority_id = ?", req.AuthorityId).Updates(&req).Error
}

func (roleService *RoleService) SetMenuAuthority(req model.Role) error {
	var role model.Role
	global.GaDb.Preload("Menus").First(&role, "authority_id = ?", req.AuthorityId)
	err := global.GaDb.Model(&role).Association("Menus").Replace(&req.Menus)
	return err
}

func (roleService *RoleService) DeleteRole(id string) error {
	if result := global.GaDb.Where("parent_id = ?", id).First(&model.Role{}); result.RowsAffected > 0 {
		return business.New("此角色存在子角色不可删除")
	}
	return global.GaDb.Where("authority_id = ?", id).Delete(&model.Role{}).Error
}
