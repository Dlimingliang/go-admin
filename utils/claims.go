package utils

import (
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/gin-gonic/gin"
)

func GetClaims(ctx *gin.Context) (*request.CustomClaims, error) {
	token := ctx.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GaLog.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

func GetUserId(ctx *gin.Context) int {
	if claims, exists := ctx.Get("claims"); !exists {
		if cl, err := GetClaims(ctx); err != nil {
			return 0
		} else {
			return cl.Id
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.Id

	}
}
