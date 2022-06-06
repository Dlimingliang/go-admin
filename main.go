package main

import (
	"fmt"

	"github.com/Dlimingliang/go-admin/config"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/initialize"
)

func main() {
	initConfig()
	//初始化数据库
	initialize.InitGorm()
	testGorm()
}

func initConfig() {
	global.ServerConfig = &config.ServerConfig{MysqlConfigInfo: config.MysqlConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		DbName:   "go-admin",
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
		User:     "root",
		Password: "123456!",
	}}
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
