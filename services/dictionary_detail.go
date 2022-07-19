package services

import (
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
)

type DictionaryDetailService struct{}

func (dictionaryDetailService *DictionaryDetailService) GetDictionaryDetailPage(req request.DictionaryDetailSearch) ([]model.DictionaryDetail, int64, error) {
	db := global.GaDb.Model(model.DictionaryDetail{})
	if req.Label != "" {
		db = db.Where("`label` LIKE ?", "%"+req.Label+"%")
	}
	if req.Value != 0 {
		db = db.Where("value = ?", req.Value)
	}
	if req.Status != nil {
		db = db.Where("`status` = ?", req.Status)
	}
	if req.DictionaryID != 0 {
		db = db.Where("dictionary_id = ?", req.DictionaryID)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	var dictionaryDetailList []model.DictionaryDetail
	err = db.Scopes(Paginate(req.Page, req.PageSize)).Find(&dictionaryDetailList).Error
	return dictionaryDetailList, total, err
}

func (dictionaryDetailService *DictionaryDetailService) GetDictionaryDetailById(id int) (model.DictionaryDetail, error) {
	var dictionaryDetail model.DictionaryDetail
	err := global.GaDb.Where("id = ?", id).First(&dictionaryDetail).Error
	return dictionaryDetail, err
}

func (dictionaryDetailService *DictionaryDetailService) CreateDictionaryDetail(req model.DictionaryDetail) (model.DictionaryDetail, error) {
	err := global.GaDb.Create(&req).Error
	return req, err
}

func (dictionaryDetailService *DictionaryDetailService) UpdateDictionaryDetail(req model.DictionaryDetail) error {
	err := global.GaDb.Model(req).Updates(req).Error
	return err
}

func (dictionaryDetailService *DictionaryDetailService) DeleteDictionaryDetail(id int) error {
	err := global.GaDb.Delete(&model.DictionaryDetail{}, id).Error
	return err
}
