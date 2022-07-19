package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"github.com/Dlimingliang/go-admin/core/business"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
	"github.com/Dlimingliang/go-admin/utils"
)

type BaseApi struct{}

// Login
// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data=response.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /base/login [post]
func (baseApi *BaseApi) Login(ctx *gin.Context) {
	var req request.Login
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if store.Verify(req.CaptchaId, req.Captcha, true) {
		//查询用户及密码是否正确
		if user, err := userService.Login(model.User{Username: req.Username, Password: req.Password}); err != nil {
			HandlerErr(err, "登录失败", ctx)
		} else {
			baseApi.GenerateToken(ctx, user)
		}
	} else {
		HandlerErr(business.New("验证码错误"), "", ctx)
	}

}

// 登录以后签发jwt
func (baseApi *BaseApi) GenerateToken(ctx *gin.Context, user model.User) {
	j := utils.NewJWT()
	claims := j.CreateClaims(request.BaseClaims{
		ID:       user.ID,
		UserName: user.Username,
		NickName: user.NickName,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		HandlerErr(err, "获取token失败", ctx)
		return
	}

	if _, err := jwtService.GetRedisJwt(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJwt(token, user.Username); err != nil {
			HandlerErr(err, "设置登录状态失败", ctx)
			return
		}
		response.SuccessWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", ctx)
	} else if err != nil {
		HandlerErr(err, "设置登录状态失败", ctx)
		return
	} else {
		global.GaLog.Panic("不知道现在有啥用，有了token，还要进行登录!")
	}
}
