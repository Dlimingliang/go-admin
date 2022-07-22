package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
)

type OperationRecordApi struct{}

func (operationRecordApi *OperationRecordApi) GetOperationRecordPage(ctx *gin.Context) {
	var req request.OperationRecordSearch
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if operationRecordList, total, err := operationRecordService.GetOperationRecordPage(req); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(response.PageResult{
			List:     operationRecordList,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", ctx)
	}
}

func (operationRecordApi *OperationRecordApi) GetOperationRecord(ctx *gin.Context) {
	var reqId request.ById
	if err := ctx.ShouldBind(&reqId); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}
	if operationRecord, err := operationRecordService.GetOperationRecord(reqId.ID); err != nil {
		HandlerErr(err, "获取失败", ctx)
	} else {
		response.SuccessWithDetailed(gin.H{"reSysOperationRecord": operationRecord}, "获取成功", ctx)
	}
}
