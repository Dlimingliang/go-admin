package config

type Zap struct {
	Level         string `mapstructure:"level"`
	Prefix        string `mapstructure:"prefix"`
	Format        string `mapstructure:"format"`
	Director      string `mapstructure:"director"`
	EncodeLevel   string `mapstructure:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key"`
}
