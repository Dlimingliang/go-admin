package services

import (
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
)

type OperationRecordService struct{}

func (operationRecordService *OperationRecordService) GetOperationRecordPage(req request.OperationRecordSearch) ([]model.OperationRecord, int64, error) {

	db := global.GaDb.Model(model.OperationRecord{})
	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}
	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	var operationRecordList []model.OperationRecord
	err = db.Order("id desc").Scopes(Paginate(req.Page, req.PageSize)).Preload("User").Find(&operationRecordList).Error
	return operationRecordList, total, err
}

func (operationRecordService *OperationRecordService) GetOperationRecord(id int) (model.OperationRecord, error) {
	var operationRecord model.OperationRecord
	err := global.GaDb.Where("id = ?").First(&operationRecord).Error
	return operationRecord, err
}

func (operationRecordService *OperationRecordService) CreateOperationRecord(req model.OperationRecord) (model.OperationRecord, error) {
	err := global.GaDb.Create(&req).Error
	return req, err
}
