package routers

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

type OperationRecordRouter struct{}

func (o *OperationRecordRouter) InitOperationRecordRouter(group *gin.RouterGroup) {
	operationApi := v1.ApiGroupAPP.OperationRecordApi
	operationRecordRouter := group.Group("sysOperationRecord")
	{
		operationRecordRouter.GET("findSysOperationRecord", operationApi.GetOperationRecord)
		operationRecordRouter.GET("getSysOperationRecordList", operationApi.GetOperationRecordPage)
	}

}
