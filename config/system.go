package config

type System struct {
	ServerName string `mapstructure:"server-name"`
	ServerPort int    `mapstructure:"server-port"`
}
