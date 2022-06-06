package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/Dlimingliang/go-admin/global"
)

func InitGorm() {
	//主要配置点
	//1. 连接及数据库配置
	//2. 连接池配置
	//3. 日志配置

	//日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	//生成数据库连接
	mysqlConfig := global.ServerConfig.MysqlConfigInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host,
		mysqlConfig.Port, mysqlConfig.DbName, mysqlConfig.Config)

	//连接数据库设置
	global.DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: false, //禁用全局更新和删除
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //配置单数表名,默认如果struct为User,生成的表为users.配置单数表名生成的表为user
		},
		Logger: newLogger,
	})

	//设置连接池
	sqlDB, _ := global.DB.DB()
	sqlDB.SetMaxIdleConns(10)           //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          //最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //连接空闲超时
}
