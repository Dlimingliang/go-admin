package services

import (
	"errors"
	"gorm.io/gorm"

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
	res := global.GaDb.Where("authority_id = ?", req.AuthorityId).First(&role)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return role, res.Error
	}
	if res.RowsAffected > 0 {
		return role, business.New("该角色code已存在")
	}

	res = global.GaDb.Where("authority_name = ?", req.AuthorityName).First(&role)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return role, res.Error
	}
	if res.RowsAffected > 0 {
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
	//如果有用户绑定，不可以进行删除
	var roleUser model.Role
	res := global.GaDb.Preload("Users").Where("authority_id = ?", id).First(&roleUser)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	if len(roleUser.Users) != 0 {
		return business.New("此角色有用户正在使用禁止删除")
	}

	//如果有子角色不可以删除
	var existSubRoles []model.Role
	if err := global.GaDb.Where("parent_id = ?", id).Find(&existSubRoles).Error; err != nil {
		return err
	}
	if len(existSubRoles) > 0 {
		return business.New("此角色存在子角色不允许删除")
	}

	//删除角色
	err := global.GaDb.Transaction(func(tx *gorm.DB) error {
		var role model.Role
		if err := tx.Preload("Menus").Where("authority_id = ?", id).First(&role).Delete(&role).Error; err != nil {
			return err
		}

		if len(role.Menus) > 0 {
			if err := tx.Model(&role).Association("Menus").Delete(&role.Menus); err != nil {
				return err
			}
		}
		return nil
	})
	CasbinServiceApp.ClearCasbin(0, id)
	return err
}
