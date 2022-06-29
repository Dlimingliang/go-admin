package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
)

// GetRoleList
// @tags 角色相关接口
// @summary 获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "角色列表信息"
// @Router /authority/getAuthorityList [post]
func GetRoleList(ctx *gin.Context) {
	if roleList, err := roleService.GetRoleList(); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(response.PageResult{
			List:     roleList,
			Total:    0,
			Page:     1,
			PageSize: 999,
		}, "获取成功", ctx)
	}
}

// CreateRole
// @tags 角色相关接口
// @summary 新增角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RoleInfo true "角色信息"
// @Success 200 {object} response.Response{data=response.RoleResult} "角色信息"
// @Router /authority/createAuthority [post]
func CreateRole(ctx *gin.Context) {
	req := request.RoleInfo{}
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	createRole := model.Role{
		AuthorityId:   req.AuthorityId,
		AuthorityName: req.AuthorityName,
		ParentId:      req.ParentId,
	}
	if role, err := roleService.CreateRole(createRole); err != nil {
		HandlerErr(err, "创建失败", ctx)
	} else {
		response.SuccessWithDetailed(response.RoleResult{Authority: role}, "创建成功", ctx)
	}
}

// UpdateRole
// @tags 角色相关接口
// @summary 修改角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RoleInfo true "角色信息"
// @Success 200 {object} response.Response{msg=string} "更新角色"
// @Router /authority/updateAuthority [post]
func UpdateRole(ctx *gin.Context) {
	req := request.RoleInfo{}
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	updateRole := model.Role{
		AuthorityId:   req.AuthorityId,
		AuthorityName: req.AuthorityName,
		ParentId:      req.ParentId,
	}

	if err := roleService.UpdateRole(updateRole); err != nil {
		HandlerErr(err, "更新失败", ctx)
	} else {
		response.SuccessWithMessage("更新成功", ctx)
	}
}

// DeleteRole
// @tags 角色相关接口
// @summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ById true "ID"
// @Success 200 {object} response.Response{msg=string} "删除角色"
// @Router /authority/deleteAuthority [post]
func DeleteRole(ctx *gin.Context) {
	reqId := request.ById{}
	if err := ctx.ShouldBind(&reqId); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := roleService.DeleteRole(reqId.ID); err != nil {
		HandlerErr(err, "删除失败", ctx)
	} else {
		response.SuccessWithMessage("删除成功", ctx)
	}
}
