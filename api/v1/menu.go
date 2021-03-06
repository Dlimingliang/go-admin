package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
)

type MenuApi struct{}

// GetMenuList
// @tags 菜单相关接口
// @summary 获取菜单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "菜单列表信息"
// @Router /menu/getMenuList [post]
func (menuApi *MenuApi) GetMenuList(ctx *gin.Context) {
	if menuList, err := menuService.GetMenuTree(); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(response.PageResult{
			List:     menuList,
			Total:    0,
			Page:     1,
			PageSize: 999,
		}, "获取成功", ctx)
	}
}

// GetMenuTree
// @tags 菜单相关接口
// @summary 获取菜单树
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.MenuTree,msg=string} "菜单树信息"
// @Router /menu/getBaseMenuTree [post]
func (menuApi *MenuApi) GetMenuTree(ctx *gin.Context) {
	if menuList, err := menuService.GetMenuTree(); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		//为了前端不报错
		for i := 0; i < len(menuList); i++ {
			menuTest(&menuList[i])
		}
		response.SuccessWithDetailed(response.MenuTree{MenuList: menuList}, "获取成功", ctx)
	}
}
func menuTest(menu *model.Menu) {
	menu.MenuBtn = []interface{}{}
	menu.Parameters = []interface{}{}
	children := menu.Children
	if len(children) > 0 {
		for i := 0; i < len(children); i++ {
			menuTest(&children[i])
		}
	}
}

// GetMenuByRole
// @tags 菜单相关接口
// @summary 根据角色获取菜单树
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色id"
// @Success 200 {object} response.Response{data=response.MenuTree,msg=string} "菜单树信息"
// @Router /menu/getMenuAuthority [post]
func (menuApi *MenuApi) GetMenuByRole(ctx *gin.Context) {
	var req request.GetAuthorityId
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if menuList, err := menuService.GetMenuByRole(req.AuthorityId); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		for i := 0; i < len(menuList); i++ {
			menuByRoleTest(&menuList[i])
		}
		response.SuccessWithDetailed(response.MenuTree{MenuList: menuList}, "获取成功", ctx)
	}
}

func menuByRoleTest(menu *model.Menu) {
	menu.MenuId = strconv.Itoa(menu.ID)
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
func (menuApi *MenuApi) GetMenuById(ctx *gin.Context) {
	var reqId request.ById
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
func (menuApi *MenuApi) CreateMenu(ctx *gin.Context) {
	var req request.MenuInfo
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	parentId, _ := strconv.Atoi(req.ParentId)
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
		ParentId:  parentId,
	}

	if menu, err := menuService.CreateMenu(createMenu); err != nil {
		HandlerErr(err, "创建失败", ctx)
	} else {
		response.SuccessWithDetailed(response.MenuResult{Menu: menu}, "创建成功", ctx)
	}
}

// AddMenuAuthority
// @tags 菜单相关接口
// @summary 更新角色菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthorityInfo true "角色菜单信息"
// @Success 200 {object} response.Response{msg=string} "更新角色菜单"
// @Router /menu/addMenuAuthority [post]
func (menuApi *MenuApi) AddMenuAuthority(ctx *gin.Context) {
	var req request.AddMenuAuthorityInfo
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}
	if err := menuService.AddMenuAuthority(req.Menus, req.AuthorityId); err != nil {
		HandlerErr(err, "添加失败", ctx)
	} else {
		response.SuccessWithMessage("添加成功", ctx)
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
func (menuApi *MenuApi) UpdateMenu(ctx *gin.Context) {
	var req request.MenuInfo
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	parentId, _ := strconv.Atoi(req.ParentId)
	updateMenu := model.Menu{
		BaseModelNoDelete: model.BaseModelNoDelete{ID: req.ID},
		Meta: model.Meta{
			Name: req.MetaInfo.Name,
			Icon: req.MetaInfo.Icon,
		},
		RoutePath: req.RoutePath,
		RouteName: req.RouteName,
		Hidden:    *req.Hidden,
		Component: req.Component,
		Sort:      req.Sort,
		ParentId:  parentId,
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
func (menuApi *MenuApi) DeleteMenu(ctx *gin.Context) {
	var reqId request.ById
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
