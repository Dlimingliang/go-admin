package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"github.com/Dlimingliang/go-admin/config"
)

var (
	ServerConfig         = &config.Server{}
	GaDb                 *gorm.DB
	GaLog                *zap.Logger
	ValidatorTrans       ut.Translator
	Ga_Redis             *redis.Client
	GaConcurrencyControl = &singleflight.Group{}
)
