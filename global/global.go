package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/Dlimingliang/go-admin/config"
)

var (
	ServerConfig = &config.Server{}
	GaDb         *gorm.DB
	GaLog        *zap.Logger
)
