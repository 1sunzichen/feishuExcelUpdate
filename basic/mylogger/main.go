package mylogger

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	Unknown LogLevel = iota
	DEBUG
	TRACE
	Info
	Warning
	Error
	Fatal
)

type Logger struct {
	Level LogLevel
}

func parseLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return Info, nil
	case "trace":
		return TRACE, nil
	case "warning":
		return Warning, nil
	case "error":
		return Error, nil
	case "fatal":
		return Fatal, nil
	default:
		return Unknown, errors.New("无效")
	}

}

func GetInfo(n int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Print("runtimefaile\n")
		return "", "", 0
	}

	// fmt.Print("\n", file)
	funcName := runtime.FuncForPC(pc).Name()
	// fmt.Printf("\n funcName:%s", funcName)
	// fmt.Printf("\n filename:%v", path.Base(file))
	// fmt.Printf("\n line:%d", line)
	return funcName, file, line
}

// NewLog
func NewLog(l string) Logger {
	level, err := parseLevel(l)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}
func (l Logger) returnLog(s LogLevel) bool {
	return l.Level >= s
}

func (I Logger) Debug(format string, arg ...interface{}) {
	if I.returnLog(DEBUG) {
		msg := fmt.Sprintf(format, arg...)
		now := time.Now()
		funcName, fileName, line := GetInfo(2)
		fmt.Printf("[%s][Debug] %s \nfuncName:%s\nfilePath: %s\nline: %d", now.Format("2006-01-02 15:04:05"), msg, funcName, fileName, line)
	}
}
func (I Logger) Info(msg string) {
	if I.Level == Info {
		now := time.Now()
		fmt.Printf("[%s][Info] %s", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (I Logger) Warning(msg string) {
	fmt.Println(msg)
}
func (I Logger) Error(msg string) {
	fmt.Println(msg)
}
