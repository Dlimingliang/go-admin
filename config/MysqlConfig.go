package config

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"db-name"`
	Config   string `yaml:"config"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
