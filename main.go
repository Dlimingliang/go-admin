package main

import (
	"fmt"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/initialize"
)

func main() {
	initialize.InitViper()
	initialize.InitZap()
	global.GaLog.Info("zap 打印Info日志测试")
	global.GaLog.Warn("zap 打印Warn日志测试")
	global.GaLog.Error("zap 打印Error日志测试")
	//初始化数据库
	initialize.InitGorm()
	testGorm()
}

type Test struct {
	Id   int
	Name string
}

func testGorm() {
	var test Test
	global.GaDb.Find(&test, 1)
	fmt.Println(test.Id, test.Name)
}
