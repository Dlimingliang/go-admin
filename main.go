package main

import (
	"fmt"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/initialize"
)

func main() {
	initialize.InitViper()
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
	global.DB.Find(&test, 1)
	fmt.Println(test.Id, test.Name)
}
