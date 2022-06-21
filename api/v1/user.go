package v1

import (
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/global"
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
