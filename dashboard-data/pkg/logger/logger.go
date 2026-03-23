package logger

import (
	"log"
	"os"
)

var (
	// InfoLogger 信息日志
	InfoLogger *log.Logger
	// ErrorLogger 错误日志
	ErrorLogger *log.Logger
)

// InitLogger 初始化日志
func InitLogger() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info 记录信息日志
func Info(format string, v ...interface{}) {
	if InfoLogger != nil {
		InfoLogger.Printf(format, v...)
	}
}

// Error 记录错误日志
func Error(format string, v ...interface{}) {
	if ErrorLogger != nil {
		ErrorLogger.Printf(format, v...)
	}
}
