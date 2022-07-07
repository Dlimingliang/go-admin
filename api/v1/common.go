package v1

import (
	"encoding/json"
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
		response.BusinessValidationError(err.Error(), ctx)
	case validator.ValidationErrors:
		errs, _ := err.(validator.ValidationErrors)
		response.ValidationError(removeTopStruct(errs.Translate(global.ValidatorTrans)), ctx)
	case *json.UnmarshalTypeError:
		errs, _ := err.(*json.UnmarshalTypeError)
		response.ValidationError(errs.Error(), ctx)
	default:
		global.GaLog.Error(msg, zap.Error(err))
		response.ReturnHttpCodeAndMessage(http.StatusInternalServerError, response.InternalErrorCode, map[string]interface{}{}, msg, ctx)
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
