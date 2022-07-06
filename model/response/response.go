package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Ok = 0
)

const (
	ValidationErrorCode         = 300
	BusinessValidationErrorCode = 301
	InternalErrorCode           = 399
)

type Response struct {
	Code          int         `json:"code"`          //编码 本来是string，但是由于已经这么设计了，没办法
	Msg           string      `json:"msg"`           //错误信息
	ValidationMsg interface{} `json:"validationMsg"` //信息
	Data          interface{} `json:"data"`          //数据
}

func ReturnHttpCodeAndMessage(httpCode int, code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(httpCode, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func BusinessValidationError(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusConflict, Response{
		Code: BusinessValidationErrorCode,
		Msg:  msg,
		Data: map[string]interface{}{},
	})
}

func ValidationError(validationMsg interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code:          ValidationErrorCode,
		Msg:           "参数验证失败",
		ValidationMsg: validationMsg,
		Data:          map[string]interface{}{},
	})
}

func Success(ctx *gin.Context) {
	SuccessWithDetailed(map[string]interface{}{}, "操作成功", ctx)
}

func SuccessWithMessage(msg string, ctx *gin.Context) {
	SuccessWithDetailed(map[string]interface{}{}, msg, ctx)
}

func SuccessWithDetailed(data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: Ok,
		Msg:  msg,
		Data: data,
	})
}
