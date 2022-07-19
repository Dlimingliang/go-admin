package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model/request"
)

var (
	TokenExpired     = errors.New("过期的Token")
	TokenNotValidYet = errors.New("未生效的Token")
	TokenMalformed   = errors.New("畸形Token")
	TokenInvalid     = errors.New("无效Token")
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ServerConfig.JWTConfig.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.ServerConfig.JWTConfig.BufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.ServerConfig.JWTConfig.Issuer,
			NotBefore: jwt.NewNumericDate(time.Unix(time.Now().Unix()-1000, 0)),                                      // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+global.ServerConfig.JWTConfig.ExpiresTime, 0)), // 过期时间 7天  配置文件
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) RefreshToken(token string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.GaConcurrencyControl.Do("JWT:"+token, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
