package routers

import (
	"github.com/Dlimingliang/go-admin/middleware"
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

type DictionaryDetailRouter struct{}

func (dictionaryDetailRouter *DictionaryDetailRouter) InitDictionaryDetailRouter(group *gin.RouterGroup) {
	dictionaryDetailApi := v1.ApiGroupAPP.DictionaryDetailApi
	dictionaryGroup := group.Group("sysDictionaryDetail").Use(middleware.OperationRecord())
	{
		dictionaryGroup.POST("createSysDictionaryDetail", dictionaryDetailApi.CreateDictionaryDetail)
		dictionaryGroup.DELETE("deleteSysDictionaryDetail", dictionaryDetailApi.DeleteDictionaryDetail)
		dictionaryGroup.PUT("updateSysDictionaryDetail", dictionaryDetailApi.UpdateDictionaryDetail)
	}
	dictionaryGroupWithoutRecord := group.Group("sysDictionaryDetail")
	{
		dictionaryGroupWithoutRecord.GET("findSysDictionaryDetail", dictionaryDetailApi.GetDictionaryDetailById)
		dictionaryGroupWithoutRecord.GET("getSysDictionaryDetailList", dictionaryDetailApi.GetDictionaryDetailPage)
	}
}
