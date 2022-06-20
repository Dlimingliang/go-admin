package service

import (
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
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
	err = db.Scopes(Paginate(page.Page, page.PageSize)).Find(&userList).Error
	return userList, total, err
}
