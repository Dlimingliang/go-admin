package initialize

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Dlimingliang/go-admin/config"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/utils"
)

/**
zap 配置项
1. 日志格式
2. 日志输出地 控制台、文件控制台都有-俩种方式
3. 写入文件的时候，要有大小分割
*/

func InitZap() {

	if ok, _ := utils.PathExists(global.ServerConfig.ZapConfig.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("创建 %v 目录\n", global.ServerConfig.ZapConfig.Director)
		_ = os.Mkdir(global.ServerConfig.ZapConfig.Director, os.ModePerm)
	}

	zapConfigInfo := global.ServerConfig.ZapConfig
	var core zapcore.Core

	switch zapConfigInfo.Level {
	case config.ZapDebug:
		core = getEncoderCore(zapcore.DebugLevel)
	case config.ZapInfo:
		core = getEncoderCore(zapcore.InfoLevel)
	case config.ZapWarn:
		core = getEncoderCore(zapcore.WarnLevel)
	case config.ZapError:
		core = getEncoderCore(zapcore.ErrorLevel)
	default:
		core = getEncoderCore(zapcore.DebugLevel)
	}
	logger := zap.New(core)
	if zapConfigInfo.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	global.GaLog = logger
}

func getEncoderCore(level zapcore.Level) (core zapcore.Core) {
	var writeSyncer zapcore.WriteSyncer
	if global.ServerConfig.ZapConfig.LogInFile {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   fmt.Sprintf("./%s/%s.log", global.ServerConfig.ZapConfig.Director, global.ServerConfig.SystemConfig.ServerName),
			MaxSize:    global.ServerConfig.ZapConfig.MaxSize,
			MaxBackups: global.ServerConfig.ZapConfig.MaxBackups,
			LocalTime:  true,
			MaxAge:     global.ServerConfig.ZapConfig.MaxAge,
			Compress:   true,
		}
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		writeSyncer = zapcore.AddSync(os.Stdout)
	}

	return zapcore.NewCore(getEncoder(), writeSyncer, level)
}

func getEncoder() zapcore.Encoder {
	if global.ServerConfig.ZapConfig.Format == config.ZapJsonFormat {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderConfig() zapcore.EncoderConfig {
	zapConfigInfo := global.ServerConfig.ZapConfig
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  zapConfigInfo.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case zapConfigInfo.EncodeLevel == config.ZapLE: // 小写编码器(默认)
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	case zapConfigInfo.EncodeLevel == config.ZapLCE: // 小写编码器带颜色
		encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case zapConfigInfo.EncodeLevel == config.ZapCE: // 大写编码器
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	case zapConfigInfo.EncodeLevel == config.ZapCCE: // 大写编码器带颜色
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return encoderConfig
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.ServerConfig.ZapConfig.Prefix + " 2006-01-02 15:04:05.000"))
}
