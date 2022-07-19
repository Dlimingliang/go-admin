package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/Dlimingliang/go-admin/global"
)

func InitRedis() {
	redisConfig := global.ServerConfig.RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GaLog.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GaLog.Info("redis connect ping response:", zap.String("pong", pong))
		global.Ga_Redis = client
	}
}
