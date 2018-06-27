package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/verystar/golib/contract"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelError = "error"
	LevelFatal = "fatal"
)

// ILogger is the logger interface
type ILogger interface {
	contract.ILogger
	Log(string, string, ...interface{})
	SetFormatter(format logrus.Formatter)
}