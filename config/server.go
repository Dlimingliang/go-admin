package config

type Server struct {
	SystemConfig  System  `mapstructure:"system"`
	MysqlConfig   Mysql   `mapstructure:"mysql"`
	ZapConfig     Zap     `mapstructure:"zap"`
	JWTConfig     JWT     `mapstructure:"jwt"`
	RedisConfig   Redis   `mapstructure:"redis"`
	CaptchaConfig Captcha `mapstructure:"captcha"`
	CasbinConfig  Casbin  `mapstructure:"casbin"`
}
