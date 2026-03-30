package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Logger 全局日志实例
	Logger *zap.Logger
	// SugarLogger 全局 Sugar 日志实例
	SugarLogger *zap.SugaredLogger
)

// InitLogger 初始化日志组件
func InitLogger(serviceName string) {
	// 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 创建核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	// 创建日志实例
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Logger = Logger.With(zap.String("service", serviceName))

	// 创建 Sugar 日志实例
	SugarLogger = Logger.Sugar()
}

// Debug 记录调试级别日志
func Debug(msg string, fields ...zapcore.Field) {
	Logger.Debug(msg, fields...)
}

// Debugf 使用 Sugar 记录调试级别日志（printf 风格）
func Debugf(format string, args ...interface{}) {
	SugarLogger.Debugf(format, args...)
}

// Info 记录信息级别日志
func Info(msg string, fields ...zapcore.Field) {
	Logger.Info(msg, fields...)
}

// Infof 使用 Sugar 记录信息级别日志（printf 风格）
func Infof(format string, args ...interface{}) {
	SugarLogger.Infof(format, args...)
}

// Warn 记录警告级别日志
func Warn(msg string, fields ...zapcore.Field) {
	Logger.Warn(msg, fields...)
}

// Warnf 使用 Sugar 记录警告级别日志（printf 风格）
func Warnf(format string, args ...interface{}) {
	SugarLogger.Warnf(format, args...)
}

// Error 记录错误级别日志
func Error(msg string, fields ...zapcore.Field) {
	Logger.Error(msg, fields...)
}

// Errorf 使用 Sugar 记录错误级别日志（printf 风格）
func Errorf(format string, args ...interface{}) {
	SugarLogger.Errorf(format, args...)
}

// Fatal 记录致命级别日志
func Fatal(msg string, fields ...zapcore.Field) {
	Logger.Fatal(msg, fields...)
}

// Fatalf 使用 Sugar 记录致命级别日志（printf 风格）
func Fatalf(format string, args ...interface{}) {
	SugarLogger.Fatalf(format, args...)
}

// WithRequestID 添加请求 ID 到日志字段
func WithRequestID(requestID string) *zap.Logger {
	return Logger.With(zap.String("request_id", requestID))
}

// WithFields 添加自定义字段到日志
func WithFields(fields ...zapcore.Field) *zap.Logger {
	return Logger.With(fields...)
}

// Sync 刷新日志缓冲区
func Sync() error {
	return Logger.Sync()
}