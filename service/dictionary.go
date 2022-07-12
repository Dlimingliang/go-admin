package service

import (
	"errors"
	"github.com/Dlimingliang/go-admin/core/business"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"

	"gorm.io/gorm"
)

type DictionaryService struct{}

func (dictionaryService *DictionaryService) GetDictionaryPage(req request.DictionarySearch) ([]model.Dictionary, int64, error) {
	db := global.GaDb.Model(model.Dictionary{})
	if req.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+req.Name+"%")
	}
	if req.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+req.Type+"%")
	}
	if req.Status != nil {
		db = db.Where("`status` = ?", req.Status)
	}
	if req.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+req.Desc+"%")
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	var dictionaryList []model.Dictionary
	err = db.Scopes(Paginate(req.Page, req.PageSize)).Find(&dictionaryList).Error
	return dictionaryList, total, err
}

func (dictionaryService *DictionaryService) GetDictionaryById(id int) (model.Dictionary, error) {
	var dictionary model.Dictionary
	err := global.GaDb.Where("id = ? and status = ?", id, true).Preload("DictionaryDetails", "status = ?", true).First(&dictionary).Error
	return dictionary, err
}

func (dictionaryService *DictionaryService) CreateDictionary(req model.Dictionary) (model.Dictionary, error) {
	var dictionary model.Dictionary
	res := global.GaDb.Where("type = ?", req.Type).First(&dictionary)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return dictionary, res.Error
	}
	if res.RowsAffected > 0 {
		return dictionary, business.New("type已存在")
	}

	err := global.GaDb.Create(&req).Error
	return req, err
}

func (dictionaryService *DictionaryService) UpdateDictionary(req model.Dictionary) error {
	var dictionary model.Dictionary
	res := global.GaDb.Where("d <> ? AND type = ?", req.ID, req.Type).First(&dictionary)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	if res.RowsAffected > 0 {
		return business.New("type已存在")
	}
	err := global.GaDb.Model(req).Updates(req).Error
	return err
}

func (dictionaryService *DictionaryService) DeleteDictionary(id int) error {
	err := global.GaDb.Transaction(func(tx *gorm.DB) error {
		//删除字典
		if err := tx.Delete(&model.Dictionary{BaseModel: model.BaseModel{ID: id}}).Error; err != nil {
			return err
		}
		//删除详细
		if err := tx.Delete(&model.DictionaryDetail{}, "dictionary_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
