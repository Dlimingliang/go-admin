package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/**
创建字段 要指定的gorm字段
主键: primarykey、autoIncrement
非主键: column、type及size、not null、default、 unique、comment
*/

type BaseModel struct {
	ID         int                   `json:"ID" gorm:"primarykey"`
	CreateTime time.Time             `json:"CreatedAt" gorm:"column:create_time;type:datetime(3);not null;comment:创建时间;autoCreateTime:milli"`
	CreateUser int                   `json:"-" gorm:"column:create_user;not null;comment:创建人"`
	UpdateTime time.Time             `json:"UpdatedAt" gorm:"column:update_time;not null;comment:更新时间;autoUpdateTime:milli"`
	UpdateUser int                   `json:"-" gorm:"column:update_user;not null;comment:更新人"`
	Enable     soft_delete.DeletedAt `json:"-" gorm:"type:tinyint(1);not null;default:0;comment:是否删除 0:否 1:是;softDelete:flag"`
}
