package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
)

type DictionaryApi struct{}

// GetDictionaryPage
// @tags 字典相关接口
// @summary 分页获取字典列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DictionarySearch true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取字典列表,返回包括列表,总数,页码,每页数量"
// @Router /sysDictionary/getSysDictionaryList [post]
func (dictionaryApi *DictionaryApi) GetDictionaryPage(ctx *gin.Context) {
	var req request.DictionarySearch
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if dictionaryList, total, err := dictionaryService.GetDictionaryPage(req); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(response.PageResult{
			List:     dictionaryList,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", ctx)
	}
}

// GetDictionaryById
// @tags 字典相关接口
// @summary 根据ID获取字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ById true "ID"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "字典信息"
// @Router /sysDictionary/findSysDictionary [get]
func (dictionaryApi *DictionaryApi) GetDictionaryById(ctx *gin.Context) {
	var reqId request.ById
	if err := ctx.ShouldBind(&reqId); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}
	if dictionary, err := dictionaryService.GetDictionaryById(reqId.ID); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(gin.H{"resysDictionary": dictionary}, "获取成功", ctx)
	}
}

// CreateDictionary
// @tags 字典相关接口
// @summary 新增字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dictionary true "字典信息"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "字典信息"
// @Router /sysDictionary/createSysDictionary [post]
func (dictionaryApi *DictionaryApi) CreateDictionary(ctx *gin.Context) {
	var req model.Dictionary
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if dictionary, err := dictionaryService.CreateDictionary(req); err != nil {
		HandlerErr(err, "创建失败", ctx)
	} else {
		response.SuccessWithDetailed(gin.H{"resysDictionary": dictionary}, "创建成功", ctx)
	}
}

// UpdateDictionary
// @tags 字典相关接口
// @summary 修改字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dictionary true "字典信息"
// @Success 200 {object} response.Response{msg=string} "更新字典"
// @Router /sysDictionary/updateSysDictionary [put]
func (dictionaryApi *DictionaryApi) UpdateDictionary(ctx *gin.Context) {
	var req model.Dictionary
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := dictionaryService.UpdateDictionary(req); err != nil {
		HandlerErr(err, "更新失败", ctx)
	} else {
		response.SuccessWithMessage("更新成功", ctx)
	}
}

// DeleteDictionary
// @tags 字典相关接口
// @summary 删除字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ById true "ID"
// @Success 200 {object} response.Response{msg=string} "删除字典"
// @Router /sysDictionary/deleteSysDictionary [delete]
func (dictionaryApi *DictionaryApi) DeleteDictionary(ctx *gin.Context) {
	var req request.ById
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := dictionaryService.DeleteDictionary(req.ID); err != nil {
		HandlerErr(err, "删除失败", ctx)
	} else {
		response.SuccessWithMessage("删除成功", ctx)
	}
}
