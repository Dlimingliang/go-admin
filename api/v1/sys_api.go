package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
)

type SysApiApi struct{}

func (s *SysApiApi) GetSysApiPage(ctx *gin.Context) {
	var req request.SysApiSearch
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if sysApiList, total, err := sysApiService.GetSysApiPage(req.SysApi, req.PageInfo); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(response.PageResult{
			List:     sysApiList,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", ctx)
	}
}

func (s *SysApiApi) GetAllSysApi(ctx *gin.Context) {
	if sysApiList, err := sysApiService.GetAllSysApis(); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(gin.H{"apis": sysApiList}, "获取成功", ctx)
	}
}

func (s *SysApiApi) GetSysApiById(ctx *gin.Context) {
	var req request.ById
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if sysApi, err := sysApiService.GetSysApiById(req.ID); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(gin.H{"api": sysApi}, "获取成功", ctx)
	}
}

func (s *SysApiApi) CreateSysApi(ctx *gin.Context) {
	var req model.SysApi
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if sysApi, err := sysApiService.CreateSysApi(req); err != nil {
		HandlerErr(err, "创建失败", ctx)
	} else {
		response.SuccessWithDetailed(gin.H{"sysApi": sysApi}, "创建成功", ctx)
	}
}

func (s *SysApiApi) UpdateSysApi(ctx *gin.Context) {
	var req model.SysApi
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := sysApiService.UpdateSysApi(req); err != nil {
		HandlerErr(err, "更新失败", ctx)
	} else {
		response.SuccessWithMessage("更新成功", ctx)
	}
}

func (s *SysApiApi) DeleteSysApi(ctx *gin.Context) {
	var req model.SysApi
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := sysApiService.DeleteSysApi(req); err != nil {
		HandlerErr(err, "删除失败", ctx)
	} else {
		response.SuccessWithMessage("删除成功", ctx)
	}
}
