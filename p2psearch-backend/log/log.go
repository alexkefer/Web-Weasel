package log

import (
	"fmt"
	"time"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
)

func Debug(format string, a ...any) { logAtLevel(LevelDebug, format, a...) }
func Info(format string, a ...any)  { logAtLevel(LevelInfo, format, a...) }
func Warn(format string, a ...any)  { logAtLevel(LevelWarn, format, a...) }
func Error(format string, a ...any) { logAtLevel(LevelError, format, a...) }

func logAtLevel(level int, format string, a ...any) {
	now := time.Now()
	levelStr := ""

	switch level {
	case LevelDebug:
		levelStr = "[DEBUG]"
	case LevelInfo:
		levelStr = "[INFO] "
	case LevelWarn:
		levelStr = "\u001B[33m[WARN]\u001B[0m "
	case LevelError:
		levelStr = "\u001B[31m[ERROR]\u001B[0m"
	default:
	}

	fmt.Printf("%s[%s] ", levelStr, now.Format(time.RFC3339))
	fmt.Printf(format, a...)
	fmt.Printf("\n")
}
