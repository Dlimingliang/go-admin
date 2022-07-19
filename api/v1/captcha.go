package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model/response"
	"github.com/Dlimingliang/go-admin/utils/captcha"
)

var store = captcha.NewDefaultRedisStore()

// Captcha
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.CaptchaResponse,msg=string} "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router /base/captcha [post]
func (baseApi BaseApi) Captcha(ctx *gin.Context) {
	captchaConfig := global.ServerConfig.CaptchaConfig
	driver := base64Captcha.NewDriverDigit(captchaConfig.ImgHeight, captchaConfig.ImgWidth, captchaConfig.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(ctx))

	if id, b64s, err := cp.Generate(); err != nil {
		HandlerErr(err, "获取验证码失败", ctx)
	} else {
		response.SuccessWithDetailed(response.CaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: captchaConfig.KeyLong,
		}, "获取验证码成功", ctx)
	}
}
