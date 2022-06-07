package global

import (
	"gorm.io/gorm"

	"github.com/Dlimingliang/go-admin/config"
)

var (
	ServerConfig = &config.Server{}
	DB           *gorm.DB
)
