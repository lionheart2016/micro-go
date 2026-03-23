package logger

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// Level 日志级别
type Level string

const (
	Debug Level = "DEBUG"
	Info  Level = "INFO"
	Warn  Level = "WARN"
	Error Level = "ERROR"
	Fatal Level = "FATAL"
)

// Logger 日志结构
type Logger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

// LogEntry 日志条目
type LogEntry struct {
	Level   Level   `json:"level"`
	Time    string  `json:"time"`
	Message string  `json:"message"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
}

// New 创建新的日志记录器
func New() *Logger {
	return &Logger{
		debugLogger: log.New(os.Stdout, "", 0),
		infoLogger:  log.New(os.Stdout, "", 0),
		warnLogger:  log.New(os.Stdout, "", 0),
		errorLogger: log.New(os.Stderr, "", 0),
		fatalLogger: log.New(os.Stderr, "", 0),
	}
}

// Debug 记录调试级别日志
func (l *Logger) Debug(message string, fields map[string]interface{}) {
	l.log(Debug, message, fields, l.debugLogger)
}

// Info 记录信息级别日志
func (l *Logger) Info(message string, fields map[string]interface{}) {
	l.log(Info, message, fields, l.infoLogger)
}

// Warn 记录警告级别日志
func (l *Logger) Warn(message string, fields map[string]interface{}) {
	l.log(Warn, message, fields, l.warnLogger)
}

// Error 记录错误级别日志
func (l *Logger) Error(message string, fields map[string]interface{}) {
	l.log(Error, message, fields, l.errorLogger)
}

// Fatal 记录致命级别日志并退出程序
func (l *Logger) Fatal(message string, fields map[string]interface{}) {
	l.log(Fatal, message, fields, l.fatalLogger)
	os.Exit(1)
}

// log 记录日志
func (l *Logger) log(level Level, message string, fields map[string]interface{}, logger *log.Logger) {
	entry := LogEntry{
		Level:   level,
		Time:    time.Now().Format(time.RFC3339),
		Message: message,
		Fields:  fields,
	}

	entryJSON, err := json.Marshal(entry)
	if err != nil {
		logger.Println("Failed to marshal log entry:", err)
		return
	}

	logger.Println(string(entryJSON))
}
