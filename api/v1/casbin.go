package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/model/response"
)

type CasbinApi struct{}

func (c *CasbinApi) GetPolicyPathByAuthorityId(ctx *gin.Context) {
	var req request.CasbinInReceive
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	paths := casbinService.GetPolicyPathByAuthorityId(req.AuthorityId)
	response.SuccessWithDetailed(gin.H{"paths": paths}, "获取成功", ctx)
}

func (c *CasbinApi) UpdateCasbin(ctx *gin.Context) {
	var req request.CasbinInReceive
	if err := ctx.ShouldBind(&req); err != nil {
		HandlerErr(err, "数据绑定错误", ctx)
		return
	}

	if err := casbinService.UpdateCasbin(req.AuthorityId, req.CasbinInfos); err != nil {
		HandlerErr(err, "更新失败", ctx)
	} else {
		response.SuccessWithMessage("更新成功", ctx)
	}
}
