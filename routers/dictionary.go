package routers

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

type DictionaryRouter struct{}

func (dictionaryRouter *DictionaryRouter) InitDictionaryRouter(group *gin.RouterGroup) {
	dictionaryApi := v1.ApiGroupAPP.DictionaryApi
	dictionaryGroup := group.Group("sysDictionary")
	{
		dictionaryGroup.POST("createSysDictionary", dictionaryApi.CreateDictionary)
		dictionaryGroup.DELETE("deleteSysDictionary", dictionaryApi.DeleteDictionary)
		dictionaryGroup.PUT("updateSysDictionary", dictionaryApi.UpdateDictionary)
	}
	dictionaryGroupWithoutRecord := group.Group("sysDictionary")
	{
		dictionaryGroupWithoutRecord.GET("findSysDictionary", dictionaryApi.GetDictionaryById)
		dictionaryGroupWithoutRecord.GET("getSysDictionaryList", dictionaryApi.GetDictionaryPage)
	}
}
