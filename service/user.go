package service

import (
	"github.com/Dlimingliang/go-admin/core/business"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/utils"
)

type UserService struct{}

func (userService *UserService) GetUserList(page request.PageInfo) ([]model.User, int64, error) {
	db := global.GaDb.Model(model.User{})

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	var userList []model.User
	err = db.Scopes(Paginate(page.Page, page.PageSize)).Preload("Roles").Find(&userList).Error
	return userList, total, err
}

func (userService *UserService) RegisterAdmin(req model.User) (model.User, error) {

	var user model.User
	//验证用户和电话是否存在
	if result := global.GaDb.Where(&model.User{Username: req.Username}).First(&user); result.RowsAffected > 0 {
		return user, business.New("该用户名已存在")
	}
	if result := global.GaDb.Where(&model.User{Mobile: req.Mobile}).First(&user); result.RowsAffected > 0 {
		return user, business.New("该电话已存在")
	}

	//生成用户密码
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return user, err
	}
	req.Password = hashPassword
	//创建用户并返回用户信息
	err = global.GaDb.Create(&req).Error
	return user, err
}

func (userService *UserService) SetUserInfo(user model.User) error {
	return global.GaDb.Updates(&user).Error
}

func (userService *UserService) SetUserAuthorities(id int, authorityIds []string) error {
	var user model.User
	if err := global.GaDb.Preload("Roles").First(&user, id).Error; err != nil {
		return err
	}
	//更换角色
	roles := make([]model.Role, 0)
	for _, authorityId := range authorityIds {
		roles = append(roles, model.Role{AuthorityId: authorityId})
	}
	err := global.GaDb.Model(&user).Association("Roles").Replace(&roles)
	return err
}

func (userService *UserService) DeleteUser(id int) error {
	return global.GaDb.Delete(&model.User{}, id).Error
}

func (userService *UserService) ResetPassword(id int) error {
	hashPassword, err := utils.HashPassword("123456")
	if err != nil {
		return err
	}
	return global.GaDb.Updates(&model.User{BaseModelNoDelete: model.BaseModelNoDelete{ID: id}, Password: hashPassword}).Error
}
