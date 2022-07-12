package model

type DictionaryDetail struct {
	BaseModel
	Label        string `json:"label" gorm:"type:varchar(255);not null;comment:展示值"`
	Value        int    `json:"value" gorm:"not null;comment:使用值"`
	Status       *bool  `json:"status" gorm:"not null;comment:状态"`
	Sort         int    `json:"sort" gorm:"not null;comment:排序"`
	DictionaryID int    `json:"sysDictionaryID" gorm:"column:dictionary_id;not null;comment:字典ID"`
}
