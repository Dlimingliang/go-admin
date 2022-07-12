package model

type Dictionary struct {
	BaseModel
	Name              string             `json:"name" form:"name"  gorm:"type:varchar(255);not null;comment:字典中文名"`
	Type              string             `json:"type" form:"type" gorm:"type:varchar(255);not null;comment:字典英文名"`
	Status            *bool              `json:"status" form:"status" gorm:"not null;comment:状态"`
	Desc              string             `json:"desc" form:"desc" gorm:"type:varchar(255);not null;comment:描述"`
	DictionaryDetails []DictionaryDetail `json:"dictionaryDetails"`
}
