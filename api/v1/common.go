package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strings"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model/response"
)

func HandlerValidatorErr(err error, ctx *gin.Context) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		global.GaLog.Error("数据绑定错误", zap.Error(errs))
		response.FailWithMessage("数据绑定错误", ctx)
		return
	}
	response.FailWithMessage(removeTopStruct(errs.Translate(global.ValidatorTrans)), ctx)
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
