package model

type DictionaryDetail struct {
	BaseModel
	Label        string `json:"label" form:"label" gorm:"type:varchar(255);not null;comment:展示值"`
	Value        int    `json:"value" form:"value" gorm:"not null;comment:使用值"`
	Status       *bool  `json:"status" form:"status" gorm:"not null;comment:状态"`
	Sort         int    `json:"sort" form:"sort" gorm:"not null;comment:排序"`
	DictionaryID int    `json:"sysDictionaryID" form:"sysDictionaryID" gorm:"column:dictionary_id;not null;comment:字典ID"`
}
