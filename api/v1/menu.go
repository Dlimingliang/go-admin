package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
)

// GetMenuTree
// @tags 菜单相关接口
// @summary 获取菜单树
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.MenuTree,msg=string} "菜单树信息"
// @Router /menu/getMenu [post]
func GetMenuTree(ctx *gin.Context) {
	if menuList, err := menuService.GetMenuTree(); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(response.MenuTree{MenuList: menuList}, "获取成功", ctx)
	}
}

// GetMenuById
// @tags 菜单相关接口
// @summary 根据ID获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ById true "ID"
// @Success 200 {object} response.Response{data=response.MenuResult,msg=string} "菜单信息"
// @Router /menu/getBaseMenuById [post]
func GetMenuById(ctx *gin.Context) {
	reqId := request.ById{}
	if err := ctx.ShouldBind(&reqId); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if menu, err := menuService.GetMenuById(reqId.ID); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(response.MenuResult{Menu: menu}, "获取成功", ctx)
	}
}

// CreateMenu
// @tags 菜单相关接口
// @summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MenuInfo true "菜单信息"
// @Success 200 {object} response.Response{data=response.MenuResult} "菜单信息"
// @Router /menu/addBaseMenu [post]
func CreateMenu(ctx *gin.Context) {
	req := request.MenuInfo{}
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	createMenu := model.Menu{
		Meta: model.Meta{
			Name: req.MetaInfo.Name,
			Icon: req.MetaInfo.Icon,
		},
		RoutePath: req.RoutePath,
		RouteName: req.RouteName,
		Hidden:    *req.Hidden,
		Component: req.Component,
		Sort:      req.Sort,
		ParentId:  *req.ParentId,
	}

	if menu, err := menuService.CreateMenu(createMenu); err != nil {
		HandlerErr(err, "创建失败", ctx)
	} else {
		response.SuccessWithDetailed(response.MenuResult{Menu: menu}, "创建成功", ctx)
	}
}

// UpdateMenu
// @tags 菜单相关接口
// @summary 修改菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MenuInfo true "菜单信息"
// @Success 200 {object} response.Response{msg=string} "更新菜单"
// @Router /menu/updateBaseMenu [post]
func UpdateMenu(ctx *gin.Context) {
	req := request.MenuInfo{}
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	updateMenu := model.Menu{
		BaseModel: model.BaseModel{ID: req.ID},
		Meta: model.Meta{
			Name: req.MetaInfo.Name,
			Icon: req.MetaInfo.Icon,
		},
		RoutePath: req.RoutePath,
		RouteName: req.RouteName,
		Hidden:    *req.Hidden,
		Component: req.Component,
		Sort:      req.Sort,
		ParentId:  *req.ParentId,
	}

	if err := menuService.UpdateMenu(updateMenu); err != nil {
		HandlerErr(err, "更新失败", ctx)
	} else {
		response.SuccessWithMessage("更新成功", ctx)
	}
}

// DeleteMenu
// @tags 菜单相关接口
// @summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ById true "ID"
// @Success 200 {object} response.Response{msg=string} "删除菜单"
// @Router /menu/deleteBaseMenu [post]
func DeleteMenu(ctx *gin.Context) {
	reqId := request.ById{}
	if err := ctx.ShouldBind(&reqId); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := menuService.DeleteMenu(reqId.ID); err != nil {
		HandlerErr(err, "删除失败", ctx)
	} else {
		response.SuccessWithMessage("删除成功", ctx)
	}
}
