package services

import (
	"context"
	"time"

	"github.com/Dlimingliang/go-admin/global"
)

type JwtService struct{}

func (receiver *JwtService) GetRedisJwt(userName string) (string, error) {
	return global.Ga_Redis.Get(context.Background(), userName).Result()
}

func (receiver *JwtService) SetRedisJwt(jwt string, userName string) error {
	timer := time.Duration(global.ServerConfig.JWTConfig.ExpiresTime) * time.Second
	err := global.Ga_Redis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
