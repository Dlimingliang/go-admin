package services

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Dlimingliang/go-admin/core/business"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
)

type SysApiService struct{}

func (sysApiService *SysApiService) GetAllSysApis() ([]model.SysApi, error) {
	var sysApiList []model.SysApi
	err := global.GaDb.Find(&sysApiList).Error
	return sysApiList, err
}

func (sysApiService *SysApiService) GetSysApiPage(req model.SysApi, pageInfo request.PageInfo) ([]model.SysApi, int64, error) {
	db := global.GaDb.Model(model.SysApi{})
	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}
	if req.Description != "" {
		db = db.Where("description LIKE ?", "%"+req.Description+"%")
	}
	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}
	if req.ApiGroup != "" {
		db = db.Where("api_group = ?", req.ApiGroup)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	var sysApiList []model.SysApi
	err = db.Scopes(Paginate(pageInfo.Page, pageInfo.PageSize)).Order("api_group").Find(&sysApiList).Error
	return sysApiList, total, err
}

func (sysApiService *SysApiService) GetSysApiById(id int) (model.SysApi, error) {
	var sysApi model.SysApi
	err := global.GaDb.Where("id = ?", id).First(&sysApi).Error
	return sysApi, err
}

func (sysApiService *SysApiService) CreateSysApi(req model.SysApi) (model.SysApi, error) {
	var sysApi model.SysApi
	if !errors.Is(global.GaDb.Where("path = ? AND method = ?", req.Path, req.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return sysApi, business.New("存在相同api")
	}

	err := global.GaDb.Create(req).Error
	return req, err
}

func (sysApiService *SysApiService) UpdateSysApi(req model.SysApi) error {
	var oldSysApi model.SysApi
	err := global.GaDb.Where("id = ?", req.ID).First(&oldSysApi).Error
	if err != nil {
		return err
	}

	err = CasbinServiceApp.UpdateCasbinApi(oldSysApi.Path, req.Path, oldSysApi.Method, req.Method)
	if err != nil {
		return err
	}
	err = global.GaDb.Save(req).Error
	return err
}

func (sysApiService *SysApiService) DeleteSysApi(req model.SysApi) error {
	err := global.GaDb.Delete(&req).Error
	CasbinServiceApp.ClearCasbin(0, req.Path, req.Method)
	return err
}
