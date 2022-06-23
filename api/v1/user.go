package v1

import (
	"github.com/Dlimingliang/go-admin/core/business"
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

	createUser := model.User{
		Username:  register.Username,
		Password:  register.Password,
		NickName:  register.NickName,
		Mobile:    register.Mobile,
		Email:     register.Email,
		HeaderImg: register.HeaderImg,
	}
	if user, err := userService.RegisterAdmin(createUser); err != nil {
		errs, ok := err.(business.GAValidateError)
		if ok {
			response.FailWithMessage(errs.Error(), ctx)
		} else {
			global.GaLog.Error("注册失败", zap.Error(err))
			response.FailWithMessage("注册失败", ctx)
		}
	} else {
		response.SuccessWithDetailed(response.UserResult{User: user}, "注册成功", ctx)
	}
}

func SetUserInfo(ctx *gin.Context) {
	changeUserInfo := request.ChangeUserInfo{}
	if err := ctx.ShouldBind(&changeUserInfo); err != nil {
		HandlerValidatorErr(err, ctx)
		return
	}

	user := model.User{
		BaseModel: model.BaseModel{
			ID: changeUserInfo.ID,
		},
		Mobile:    changeUserInfo.Mobile,
		NickName:  changeUserInfo.NickName,
		Email:     changeUserInfo.Email,
		HeaderImg: changeUserInfo.HeaderImg,
	}
	if err := userService.SetUserInfo(user); err != nil {
		global.GaLog.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.SuccessWithMessage("更新成功", ctx)
	}
}

func ResetPassword(ctx *gin.Context) {
	reqId := request.ById{}
	if err := ctx.ShouldBind(&reqId); err != nil {
		HandlerValidatorErr(err, ctx)
		return
	}

	if err := userService.ResetPassword(reqId.ID); err != nil {
		global.GaLog.Error("重置失败", zap.Error(err))
		response.FailWithMessage("重置失败", ctx)
	} else {
		response.SuccessWithMessage("重置成功", ctx)
	}
}

func DeleteUser(ctx *gin.Context) {
	reqId := request.ById{}
	if err := ctx.ShouldBind(&reqId); err != nil {
		HandlerValidatorErr(err, ctx)
		return
	}

	if err := userService.DeleteUser(reqId.ID); err != nil {
		global.GaLog.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.SuccessWithMessage("删除成功", ctx)
	}
}
