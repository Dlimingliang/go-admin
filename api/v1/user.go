package v1

import (
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
)

func GetUserList(ctx *gin.Context) {
	pageInfo := request.PageInfo{}
	if err := ctx.ShouldBind(&pageInfo); err != nil {
		HandlerValidatorErr(err, ctx)
		return
	}

	if userList, total, err := userService.GetUserList(pageInfo); err != nil {
		global.GaLog.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.SuccessWithDetailed(response.PageResult{
			List:     userList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

func RegisterAdmin(ctx *gin.Context) {
	register := request.Register{}
	if err := ctx.ShouldBind(&register); err != nil {
		HandlerValidatorErr(err, ctx)
		return
	}

	user := model.User{
		Username:  register.Username,
		Password:  register.Password,
		NickName:  register.NickName,
		Mobile:    register.Mobile,
		Email:     register.Email,
		HeaderImg: register.HeaderImg,
	}
	if user, err := userService.RegisterAdmin(user); err != nil {
		global.GaLog.Error("注册失败", zap.Error(err))
		response.FailWithMessage("注册失败", ctx)
	} else {
		response.SuccessWithDetailed(response.UserResult{User: user}, "注册成功", ctx)
	}
}
