package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
)

type DictionaryDetailApi struct{}

// GetDictionaryDetailPage
// @tags 字典详细相关接口
// @summary 分页获取字典详细列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DictionaryDetailSearch true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取字典列表,返回包括列表,总数,页码,每页数量"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
func (dictionaryDetailApi *DictionaryDetailApi) GetDictionaryDetailPage(ctx *gin.Context) {
	var req request.DictionaryDetailSearch
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if dictionaryDetailList, total, err := dictionaryDetailService.GetDictionaryDetailPage(req); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(response.PageResult{
			List:     dictionaryDetailList,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", ctx)
	}
}

// GetDictionaryDetailById
// @tags 字典详细相关接口
// @summary 根据ID获取字典详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ById true "ID"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "字典详细信息"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
func (dictionaryDetailApi *DictionaryDetailApi) GetDictionaryDetailById(ctx *gin.Context) {
	var reqId request.ById
	if err := ctx.ShouldBind(&reqId); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}
	if dictionaryDetail, err := dictionaryDetailService.GetDictionaryDetailById(reqId.ID); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(gin.H{"reSysDictionaryDetail": dictionaryDetail}, "获取成功", ctx)
	}
}

// CreateDictionaryDetail
// @tags 字典详细相关接口
// @summary 新增字典详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "字典详细信息"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "字典详细信息"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
func (dictionaryDetailApi *DictionaryDetailApi) CreateDictionaryDetail(ctx *gin.Context) {
	var req model.DictionaryDetail
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if dictionaryDetail, err := dictionaryDetailService.CreateDictionaryDetail(req); err != nil {
		HandlerErr(err, "创建失败", ctx)
	} else {
		response.SuccessWithDetailed(gin.H{"resysDictionary": dictionaryDetail}, "创建成功", ctx)
	}
}

// UpdateDictionaryDetail
// @tags 字典详细相关接口
// @summary 修改字典详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "字典详细信息"
// @Success 200 {object} response.Response{msg=string} "更新字典详细"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
func (dictionaryDetailApi *DictionaryDetailApi) UpdateDictionaryDetail(ctx *gin.Context) {
	var req model.DictionaryDetail
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := dictionaryDetailService.UpdateDictionaryDetail(req); err != nil {
		HandlerErr(err, "更新失败", ctx)
	} else {
		response.SuccessWithMessage("更新成功", ctx)
	}
}

// DeleteDictionaryDetail
// @tags 字典详细相关接口
// @summary 删除字典详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ById true "ID"
// @Success 200 {object} response.Response{msg=string} "删除字典详细"
// @Router /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func (dictionaryDetailApi *DictionaryDetailApi) DeleteDictionaryDetail(ctx *gin.Context) {
	var req request.ById
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := dictionaryDetailService.DeleteDictionaryDetail(req.ID); err != nil {
		HandlerErr(err, "删除失败", ctx)
	} else {
		response.SuccessWithMessage("删除成功", ctx)
	}
}
