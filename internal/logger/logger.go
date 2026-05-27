package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var levelNames = map[Level]string{
	LevelDebug: "DEBUG",
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
	LevelFatal: "FATAL",
}

var levelColors = map[Level]string{
	LevelDebug: "\033[36m",
	LevelInfo:  "\033[32m",
	LevelWarn:  "\033[33m",
	LevelError: "\033[31m",
	LevelFatal: "\033[31m",
}

const (
	colorReset   = "\033[0m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
)

type Logger struct {
	out      io.Writer
	pid      int
	minLevel Level
}

var globalLogger *Logger

func Init(levelStr string) {
	var minLvl Level
	cleaned := strings.ToUpper(strings.TrimSpace(levelStr))

	switch cleaned {
	case "DEBUG":
		minLvl = LevelDebug
	case "INFO":
		minLvl = LevelInfo
	case "WARN", "WARNING":
		minLvl = LevelWarn
	case "ERROR":
		minLvl = LevelError
	case "FATAL":
		minLvl = LevelFatal
	default:
		minLvl = LevelInfo
	}

	globalLogger = &Logger{
		out:      os.Stdout,
		pid:      os.Getpid(),
		minLevel: minLvl,
	}
}

func (l *Logger) log(level Level, format string, v ...interface{}) {
	if globalLogger == nil {
		return
	}

	if level < l.minLevel {
		return
	}

	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	lvlName := fmt.Sprintf("%5s", levelNames[level])
	colorCode := levelColors[level]
	coloredLevel := fmt.Sprintf("%s%s%s", colorCode, lvlName, colorReset)
	coloredPID := fmt.Sprintf("%s%5d%s", colorMagenta, l.pid, colorReset)
	threadName := "main"

	loggerName := "unknown"
	for i := 1; i < 10; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok {
			if strings.Contains(file, "internal/logger") || strings.Contains(file, "logger.go") {
				continue
			}
			loggerName = fmt.Sprintf("%s:%d", filepath.Base(file), line)
			break
		}
	}

	if len(loggerName) > 40 {
		loggerName = loggerName[len(loggerName)-40:]
	}
	coloredLogger := fmt.Sprintf("%s%-40s%s", colorCyan, loggerName, colorReset)

	var msg string
	if format == "" {
		msg = fmt.Sprintln(v...)
	} else {
		msg = fmt.Sprintf(format, v...)
	}
	msg = strings.TrimSuffix(msg, "\n")

	result := fmt.Sprintf("%s %s %s --- [%15.15s] %s : %s\n",
		timeStr,
		coloredLevel,
		coloredPID,
		threadName,
		coloredLogger,
		msg,
	)

	_, _ = l.out.Write([]byte(result))
}

func Printf(format string, v ...interface{}) { globalLogger.log(LevelInfo, format, v...) }
func Println(v ...interface{})               { globalLogger.log(LevelInfo, "", v...) }

func Info(format string, v ...interface{})  { globalLogger.log(LevelInfo, format, v...) }
func Warn(format string, v ...interface{})  { globalLogger.log(LevelWarn, format, v...) }
func Error(format string, v ...interface{}) { globalLogger.log(LevelError, format, v...) }
func Debug(format string, v ...interface{}) { globalLogger.log(LevelDebug, format, v...) }

func Fatalf(format string, v ...interface{}) {
	globalLogger.log(LevelFatal, format, v...)
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	globalLogger.log(LevelFatal, "", v...)
	os.Exit(1)
}

func Panicf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	globalLogger.log(LevelFatal, format, v...)
	panic(msg)
}
