package logger

import "github.com/sirupsen/logrus"

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelError = "error"
	LevelFatal = "fatal"
)

const (
	SENTRY_DSN = "https://25b0e63c81264cf0b708c9d5b34749ac:a0024c692dbc4bab884cece2b5afc952@sentry.verystar.cn/3"
)

// ILogger is the logger interface
type ILogger interface {
	Fatal(string, ...interface{})
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Error(string, ...interface{})
	SetFormatter(format logrus.Formatter)
}
