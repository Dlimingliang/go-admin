package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BusinessError = 7
	Ok            = 0
)

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnResultWithHttpCode(httpCode int, code int, data interface{}, msg interface{}, ctx *gin.Context) {
	ctx.JSON(httpCode, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func ReturnResult(code int, data interface{}, msg interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(c *gin.Context) {
	ReturnResult(Ok, map[string]interface{}{}, "操作成功", c)
}

func SuccessWithMessage(message interface{}, c *gin.Context) {
	ReturnResult(Ok, map[string]interface{}{}, message, c)
}

func SuccessWithDetailed(data interface{}, message interface{}, c *gin.Context) {
	ReturnResult(Ok, data, message, c)
}

func Fail(c *gin.Context) {
	ReturnResult(BusinessError, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message interface{}, c *gin.Context) {
	ReturnResult(BusinessError, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message interface{}, c *gin.Context) {
	ReturnResult(BusinessError, data, message, c)
}
