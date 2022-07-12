package model

import (
	"time"

	"gorm.io/gorm"
)

/**
创建字段 要指定的gorm字段
主键: primarykey、autoIncrement
非主键: column、type及size、not null、default、 unique、index、comment
*/

type BaseModel struct {
	ID         int            `json:"ID" gorm:"primary_key"`
	CreateTime time.Time      `json:"CreatedAt" gorm:"column:create_time;type:datetime(3);not null;comment:创建时间;autoCreateTime:milli"`
	CreateUser int            `json:"-" gorm:"column:create_user;comment:创建人"`
	UpdateTime time.Time      `json:"UpdatedAt" gorm:"column:update_time;not null;comment:更新时间;autoUpdateTime:milli"`
	UpdateUser int            `json:"-" gorm:"column:update_user;comment:更新人"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index;comment:软删除"`
}

type BaseModelNoDelete struct {
	ID         int       `json:"ID" gorm:"primary_key"`
	CreateTime time.Time `json:"CreatedAt" gorm:"column:create_time;type:datetime(3);not null;comment:创建时间;autoCreateTime:milli"`
	CreateUser int       `json:"-" gorm:"column:create_user;comment:创建人"`
	UpdateTime time.Time `json:"UpdatedAt" gorm:"column:update_time;not null;comment:更新时间;autoUpdateTime:milli"`
	UpdateUser int       `json:"-" gorm:"column:update_user;comment:更新人"`
}
