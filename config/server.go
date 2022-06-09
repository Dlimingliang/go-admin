package config

type Server struct {
	MysqlConfig Mysql `mapstructure:"mysql"`
	ZapConfig   Zap   `mapstructure:"zap"`
}
