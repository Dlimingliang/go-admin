package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	CreateTime    time.Time      `json:"CreatedAt" gorm:"column:create_time;type:datetime(3);not null;comment:创建时间;autoCreateTime:milli"`
	CreateUser    int            `json:"-" gorm:"column:create_user;comment:创建人"`
	UpdateTime    time.Time      `json:"UpdatedAt" gorm:"column:update_time;not null;comment:更新时间;autoUpdateTime:milli"`
	UpdateUser    int            `json:"-" gorm:"column:update_user;comment:更新人"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index;uniqueIndex:name_delete_unique;comment:软删除"`
	AuthorityId   string         `json:"authorityId" gorm:"primary_key;column:authority_id;type:varchar(10);comment:角色编码"`                                 //角色编码
	AuthorityName string         `json:"authorityName" gorm:"column:authority_name;type:varchar(10);not null;uniqueIndex:name_delete_unique;comment:角色名称"` //角色名称
	ParentId      string         `json:"parentId" gorm:"column:parent_id;type:varchar(10);not null;comment:父角色ID"`                                         //父角色ID
	DefaultRouter string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`                                                              // 默认菜单(默认dashboard)
	Children      []Role         `json:"children" gorm:"-"`                                                                                                //子角色
	Users         []User         `json:"-" gorm:"many2many:user_roles;"`
	Menus         []Menu         `json:"menus" gorm:"many2many:role_menus;"`
}
