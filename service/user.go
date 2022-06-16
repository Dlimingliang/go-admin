package service

import (
	"gorm.io/gorm"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/common"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetUserList(page common.PageInfo) ([]model.User, int, error) {
	var userList []model.User
	result := global.GaDb.Scopes(Paginate(page.Page, page.PageSize)).Find(&userList)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64
	result = global.GaDb.Model(model.User{}).Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return userList, int(total), nil
}
