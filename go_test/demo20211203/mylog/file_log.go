package mylog

import (
	"fmt"
	"os"
	"time"
)

type FileLogger struct {
	level       int
	logFilePath string
	logFileName string
	logFile     *os.File
}

// NewFileLogger 是一个生成文件日志结构体示例的构造函数
func NewFileLogger(level int, logFilePath, logFileName string) *FileLogger {
	flObj := &FileLogger{
		level:       level,
		logFilePath: logFilePath,
		logFileName: logFileName,
	}
	flObj.initFileLogger()
	return flObj
}

// initFileLogger 用来初始化文件日志句柄
func (f *FileLogger) initFileLogger() {
	filePath := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("open file:%s failed", filePath))
	}
	f.logFile = file
}

// Debug 写日志
func (f *FileLogger) Debug(format string, a ...interface{}) {
	file, funcName, line := getCallerInfo()
	//往文件里写
	//f.logFile.WriteString()
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%d] %s", nowStr, getLevelStr(f.level),
		file, funcName, line,
		format)
	fmt.Fprintf(f.logFile, format, a...)
	fmt.Fprintln(f.logFile)
}
