package config

type Zap struct {
	Level         string `mapstructure:"level"`
	Prefix        string `mapstructure:"prefix"`
	Format        string `mapstructure:"format"`
	Director      string `mapstructure:"director"`
	EncodeLevel   string `mapstructure:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key"`
	MaxAge        int    `mapstructure:"max-age"` //日志最大保留天数
	ShowLine      bool   `mapstructure:"show-line"`
	LogInFile     bool   `mapstructure:"log-in-file"` //日志是否写入文件中
}
