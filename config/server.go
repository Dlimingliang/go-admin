package config

type Server struct {
	SystemConfig System `mapstructure:"system"`
	MysqlConfig  Mysql  `mapstructure:"mysql"`
	ZapConfig    Zap    `mapstructure:"zap"`
}
