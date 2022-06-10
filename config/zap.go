package config

type Zap struct {
	Level         string `mapstructure:"level"`
	Prefix        string `mapstructure:"prefix"`
	Format        string `mapstructure:"format"`
	Director      string `mapstructure:"director"`
	EncodeLevel   string `mapstructure:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key"`
	MaxAge        int    `mapstructure:"max-age"`
	MaxSize       int    `mapstructure:"max-size"`
	MaxBackups    int    `mapstructure:"max-backups"`
	ShowLine      bool   `mapstructure:"show-line"`
	LogInFile     bool   `mapstructure:"log-in-file"`
}
