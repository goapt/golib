package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var _ ILogger = (*StdLogger)(nil)

type StdLogger struct {
	log *log.Logger
}

func NewStdLogger() ILogger {
	return &StdLogger{
		log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *StdLogger) Fatal(format string, args ...interface{}) {
	l.log.Fatalf(format, args...)
}

func (l *StdLogger) Debug(format string, args ...interface{}) {
	l.log.Printf(format, args...)
}

func (l *StdLogger) Info(format string, args ...interface{}) {
	l.log.Printf(format, args...)
}

func (l *StdLogger) Error(format string, args ...interface{}) {
	l.log.Printf(format, args...)
}

func (l *StdLogger) Log(level string, format string, args ...interface{}) {
	l.log.Printf(format, args...)
}

func (l *StdLogger) SetFormatter(format logrus.Formatter) {

}
