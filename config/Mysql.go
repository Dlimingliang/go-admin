package config

import "fmt"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"db-name"`
	Config   string `yaml:"config"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogMode  string `yaml:"log-mode"`
}

func (mysqlConfig *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host,
		mysqlConfig.Port, mysqlConfig.DbName, mysqlConfig.Config)
}
