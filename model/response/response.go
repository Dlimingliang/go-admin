package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnResult(code int, data interface{}, msg interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(c *gin.Context) {
	ReturnResult(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func SuccessWithMessage(message interface{}, c *gin.Context) {
	ReturnResult(SUCCESS, map[string]interface{}{}, message, c)
}

func SuccessWithDetailed(data interface{}, message interface{}, c *gin.Context) {
	ReturnResult(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	ReturnResult(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message interface{}, c *gin.Context) {
	ReturnResult(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message interface{}, c *gin.Context) {
	ReturnResult(ERROR, data, message, c)
}
