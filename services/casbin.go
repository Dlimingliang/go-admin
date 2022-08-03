package services

import (
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/Dlimingliang/go-admin/core/business"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model/request"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

//获取当前角色的权限列表
func (casbinService CasbinService) GetPolicyPathByAuthorityId(roleId string) []request.CasbinInfo {
	var casbinInfos []request.CasbinInfo
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, roleId)
	for _, value := range list {
		casbinInfos = append(casbinInfos, request.CasbinInfo{Path: value[1], Method: value[2]})
	}
	return casbinInfos
}

//根据角色更新相关api
func (casbinService *CasbinService) UpdateCasbin(roleId string, casbinInfos []request.CasbinInfo) error {
	casbinService.ClearCasbin(0, roleId)

	rules := [][]string{}
	for _, info := range casbinInfos {
		rules = append(rules, []string{roleId, info.Path, info.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return business.New("存在相同的api，添加失败，请联系管理员")
	}
	return nil
}

//更新api路径
func (casbinService *CasbinService) UpdateCasbinApi(oldPath, newPath, oldMethod, newMethod string) error {
	err := global.GaDb.Model(&gormadapter.CasbinRule{}).
		Where("v1 = ? AND v2 = ?", oldPath, oldMethod).
		Updates(map[string]interface{}{"v1": newPath, "v2": newMethod}).Error
	return err
}

//清除casbin
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

//持久化casbin到数据库 引入自定义规则
var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		adapter, _ := gormadapter.NewAdapterByDB(global.GaDb)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.ServerConfig.CasbinConfig.ModelPath, adapter)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
