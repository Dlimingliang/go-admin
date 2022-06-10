package initialize

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/Dlimingliang/go-admin/config"
	"github.com/Dlimingliang/go-admin/global"
)

/**
gorm配置项
1. 数据库连接及配置
2. gorm日志配置
3. gorm连接池配置
*/

func InitGorm() {

	//数据库配置
	mysqlConfigInfo := global.ServerConfig.MysqlConfig
	mysqlConfig := mysql.Config{
		DSN:                       mysqlConfigInfo.Dsn(),
		SkipInitializeWithVersion: false,
	}
	//gorm配置
	gormConfig := &gorm.Config{
		AllowGlobalUpdate:                        false, //不允许全局更新和删除
		DisableForeignKeyConstraintWhenMigrating: true,  //禁用外键
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //配置单数表名,默认如果struct为User,生成的表为users.配置单数表名生成的表为user
		},
	}
	//gorm日志配置
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 禁用彩色打印
		},
	)
	switch mysqlConfigInfo.LogMode {
	case config.MysqlSlientLog:
		gormConfig.Logger = gormLogger.LogMode(logger.Silent)
	case config.MysqlErrorLog:
		gormConfig.Logger = gormLogger.LogMode(logger.Error)
	case config.MysqlWarnLog:
		gormConfig.Logger = gormLogger.LogMode(logger.Warn)
	case config.MysqlInfoLog:
		gormConfig.Logger = gormLogger.LogMode(logger.Info)
	default:
		gormConfig.Logger = gormLogger.LogMode(logger.Info)
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		return
	} else {
		//设置连接池
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(mysqlConfigInfo.MaxIdleConn) //最大空闲连接数
		sqlDB.SetMaxOpenConns(mysqlConfigInfo.MaxOpenConn) //最大连接数
		sqlDB.SetConnMaxLifetime(time.Hour)                //连接空闲超时
		global.GaDb = db
	}
}
