package log

import (
	"fmt"
	"os"
	"time"

	"fake-btc-markets/pkg/config"
)

var (
	levelMap = map[int]string {
		LevelDebug: "DEBUG",
		LevelInfo:  "INFO",
		LevelWarn:  "WARN",
		LevelError: "ERROR",
		LevelFatal: "FATAL",
	}
)

const (
	LevelDebug = 10
	LevelInfo  = 20
	LevelWarn  = 30
	LevelError = 40
	LevelFatal = 50
)

func getLevelString(level int) string {
	levelString, ok := levelMap[level]
	if ok {
		return levelString
	}

	return "INFO"
}

func Debug(format string, a ...interface{}) {
	log(LevelDebug, format, a...)
}

func Info(format string, a ...interface{}) {
	log(LevelInfo, format, a...)
}

func Warn(err error, a ...interface{}) {
	log(LevelWarn, err.Error(), a...)
}

func Error(err error, a ...interface{}) {
	log(LevelError, err.Error(), a...)
}

func Fatal(err error, a ...interface{}) {
	log(LevelFatal, err.Error(), a...)
	os.Exit(1)
}

func log(level int, format string, a ...interface{}) {
	if level >= config.Get().LogLevel {
		now := time.Now().Format(time.RFC3339)
		fmt.Println(
			fmt.Sprintf("%s [%s]", now, getLevelString(level)),
			fmt.Sprintf(format, a...),
		)
	}
}
