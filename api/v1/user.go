package v1

import (
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {
	pageInfo := request.PageInfo{}
	if err := ctx.ShouldBind(&pageInfo); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if userList, total, err := userService.GetUserList(pageInfo); err != nil {
		HandlerErr(err, "获取失败", ctx)
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
		HandlerErr(err, "数据绑定错误", ctx)
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
		HandlerErr(err, "注册失败", ctx)
	} else {
		response.SuccessWithDetailed(response.UserResult{User: user}, "注册成功", ctx)
	}
}

func SetUserInfo(ctx *gin.Context) {
	changeUserInfo := request.ChangeUserInfo{}
	if err := ctx.ShouldBind(&changeUserInfo); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
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
		HandlerErr(err, "更新失败", ctx)
	} else {
		response.SuccessWithMessage("更新成功", ctx)
	}
}

func ResetPassword(ctx *gin.Context) {
	reqId := request.ById{}
	if err := ctx.ShouldBind(&reqId); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := userService.ResetPassword(reqId.ID); err != nil {
		HandlerErr(err, "重置失败", ctx)
	} else {
		response.SuccessWithMessage("重置成功", ctx)
	}
}

func DeleteUser(ctx *gin.Context) {
	reqId := request.ById{}
	if err := ctx.ShouldBind(&reqId); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := userService.DeleteUser(reqId.ID); err != nil {
		HandlerErr(err, "删除失败", ctx)
	} else {
		response.SuccessWithMessage("删除成功", ctx)
	}
}
