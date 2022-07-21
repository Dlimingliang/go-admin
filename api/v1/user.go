package v1

import (
	"github.com/Dlimingliang/go-admin/utils"
	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
)

type UserApi struct{}

// GetUserList
// @tags 用户相关接口
// @summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router /user/getUserList [post]
func (userApi *UserApi) GetUserList(ctx *gin.Context) {

	var pageInfo request.PageInfo
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

// GetUserInfo
// @Tags SysUser
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取用户信息"
// @Router /user/getUserInfo [get]
func (userApi *UserApi) GetUserInfo(ctx *gin.Context) {
	userId := utils.GetUserId(ctx)
	if user, err := userService.GetUserInfo(userId); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(gin.H{"userInfo": user}, "获取成功", ctx)
	}
}

// RegisterAdmin
// @tags 用户相关接口
// @summary 注册管理员
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {object} response.Response{data=response.UserResult} "用户注册账号,返回包括用户信息"
// @Router /user/admin_register [post]
func (userApi *UserApi) RegisterAdmin(ctx *gin.Context) {
	var register request.Register
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

	for _, id := range register.AuthorityIds {
		createUser.Roles = append(createUser.Roles, model.Role{AuthorityId: id})
	}

	if user, err := userService.RegisterAdmin(createUser); err != nil {
		HandlerErr(err, "注册失败", ctx)
	} else {
		response.SuccessWithDetailed(response.UserResult{User: user}, "注册成功", ctx)
	}
}

// SetUserInfo
// @tags 用户相关接口
// @summary 修改管理员信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ChangeUserInfo true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {object} response.Response{msg=string} "修改管理员信息"
// @Router /user/setUserInfo [put]
func (userApi *UserApi) SetUserInfo(ctx *gin.Context) {
	var changeUserInfo request.ChangeUserInfo
	if err := ctx.ShouldBind(&changeUserInfo); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if len(changeUserInfo.AuthorityIds) > 0 {
		err := userService.SetUserAuthorities(changeUserInfo.ID, changeUserInfo.AuthorityIds)
		if err != nil {
			HandlerErr(err, "更新失败", ctx)
		}
	}

	user := model.User{
		BaseModelNoDelete: model.BaseModelNoDelete{
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

// SetUserInfo
// @tags 用户相关接口
// @summary 修改管理员角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuthorities true "修改管理员角色"
// @Success 200 {object} response.Response{msg=string} "修改管理员角色"
// @Router /user/setUserAuthorities [post]
func (userApi *UserApi) SetUserAuthorities(ctx *gin.Context) {
	var req request.SetUserAuthorities
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := userService.SetUserAuthorities(req.ID, req.AuthorityIds); err != nil {
		HandlerErr(err, "更新失败", ctx)
	} else {
		response.SuccessWithMessage("更新成功", ctx)
	}

}

// ResetPassword
// @tags 用户相关接口
// @summary 重置密码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ById true "ID"
// @Success 200 {object} response.Response{msg=string} "重置密码"
// @Router /user/resetPassword [post]
func (userApi *UserApi) ResetPassword(ctx *gin.Context) {
	var reqId request.ById
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

// DeleteUser
// @tags 用户相关接口
// @summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ById true "ID"
// @Success 200 {object} response.Response{msg=string} "删除用户"
// @Router /user/deleteUser [delete]
func (userApi *UserApi) DeleteUser(ctx *gin.Context) {
	var reqId request.ById
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
