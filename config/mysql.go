package config

import "fmt"

const (
	MysqlSlientLog = "silent"
	MysqlErrorLog  = "businesserror"
	MysqlWarnLog   = "warn"
	MysqlInfoLog   = "info"
)

type Mysql struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	DbName      string `mapstructure:"db-name"`
	Config      string `mapstructure:"config"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	LogMode     string `mapstructure:"log-mode"`
	MaxIdleConn int    `mapstructure:"max-idle-conn"`
	MaxOpenConn int    `mapstructure:"max-open-conn"`
}

func (mysqlConfig *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host,
		mysqlConfig.Port, mysqlConfig.DbName, mysqlConfig.Config)
}
