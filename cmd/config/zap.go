package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Zap struct {
	Level         string `json:"level" yaml:"level" mapstructure:"level"`                            // 级别
	Prefix        string `json:"prefix" yaml:"prefix" mapstructure:"prefix"`                         // 日志前缀
	Format        string `json:"format" yaml:"format" mapstructure:"format"`                         // 输出
	Director      string `json:"director"  yaml:"director" mapstructure:"director"`                  // 日志文件夹
	EncodeLevel   string `json:"encode-level" yaml:"encode-level" mapstructure:"encode-level"`       // 编码级
	StacktraceKey string `json:"stacktrace-key" yaml:"stacktrace-key" mapstructure:"stacktrace-key"` // 栈名

	MaxAge       int  `json:"max-age" yaml:"max-age" mapstructure:"max-age"`                      // 日志留存时间
	ShowLine     bool `json:"show-line" yaml:"show-line" mapstructure:"show-line"`                // 显示行
	LogInConsole bool `json:"log-in-console" yaml:"log-in-console" mapstructure:"log-in-console"` // 输出控制台
}

// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 根据字符串转化为 zapcore.Level
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
