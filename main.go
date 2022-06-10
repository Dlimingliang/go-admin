package main

import (
	"github.com/Dlimingliang/go-admin/initialize"
)

func main() {
	initialize.InitViper()
	initialize.InitZap()
	//初始化数据库
	initialize.InitGorm()
}
