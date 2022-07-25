package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/initialize"
	"github.com/Dlimingliang/go-admin/model"
)

// @title GA-API接口文档
// @description GA管理平台
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	initialize.InitViper()
	initialize.InitZap()
	initialize.InitGorm()
	initialize.InitRedis()
	ginRouter := initialize.InitRouters()
	initialize.InitValidator()
	//migrate()

	//启动服务
	address := fmt.Sprintf(":%d", global.ServerConfig.SystemConfig.ServerPort)
	srv := &http.Server{
		Addr:    address,
		Handler: ginRouter,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			global.GaLog.Error(err.Error())
		}
	}()
	time.Sleep(10 * time.Microsecond)
	global.GaLog.Info("go-admin启动成功", zap.String("address", address))

	//优雅地终止服务
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.GaLog.Error("go-admin退出失败!", zap.Error(err))
	}
	global.GaLog.Info("go-admin退出成功")
}

func migrate() {
	global.GaDb.AutoMigrate(&model.User{})
	global.GaDb.AutoMigrate(&model.Menu{})
	global.GaDb.AutoMigrate(&model.Role{})
	global.GaDb.AutoMigrate(&model.Dictionary{})
	global.GaDb.AutoMigrate(&model.DictionaryDetail{})
	global.GaDb.AutoMigrate(&model.OperationRecord{})
}
