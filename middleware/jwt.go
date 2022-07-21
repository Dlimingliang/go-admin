package middleware

import (
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model/response"
	"github.com/Dlimingliang/go-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"

	"github.com/Dlimingliang/go-admin/services"
)

var jwtService = services.ServiceGroupApp.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			response.BusinessValidationError("未登录,请去登录", ctx)
			ctx.Abort()
			return
		}

		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			response.BusinessValidationError(err.Error(), ctx)
			ctx.Abort()
			return
		}

		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = jwt.NewNumericDate(time.Unix(time.Now().Unix()+global.ServerConfig.JWTConfig.ExpiresTime, 0))
			newToken, _ := j.RefreshToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			ctx.Header("new-token", newToken)
			ctx.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			jwtService.SetRedisJwt(newToken, newClaims.UserName)
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
