package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/Dlimingliang/go-admin/core/business"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model/response"
)

func HandlerErr(err error, msg string, ctx *gin.Context) {
	switch err.(type) {
	case business.GAValidateError:
		response.FailWithMessage(err.Error(), ctx)
	case validator.ValidationErrors:
		errs, _ := err.(validator.ValidationErrors)
		response.ReturnResultWithHttpCode(http.StatusBadRequest, response.ValidateError, map[string]interface{}{}, removeTopStruct(errs.Translate(global.ValidatorTrans)), ctx)
	default:
		global.GaLog.Error(msg, zap.Error(err))
		response.ReturnResultWithHttpCode(http.StatusInternalServerError, response.Error, map[string]interface{}{}, msg, ctx)
	}
	return
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
