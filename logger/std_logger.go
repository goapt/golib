package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var _ ILogger = (*StdLogger)(nil)

type StdLogger struct {
}

func NewStdLogger() ILogger {
	return &StdLogger{}
}

func (l *StdLogger) Fatal(format string, args ...interface{}) {
	log.Printf(format, args...)
	os.Exit(1)
}

func (l *StdLogger) Debug(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *StdLogger) Info(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *StdLogger) Error(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *StdLogger) Log(level string, format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *StdLogger) SetFormatter(format logrus.Formatter) {

}
