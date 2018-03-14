package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var _ ILogger = (*StdLogger)(nil)

type StdLogger struct {
	lg *log.Logger
}

func NewStdLogger() ILogger {
	return &StdLogger{
		lg: log.New(os.Stderr, "", log.Lshortfile | log.LstdFlags),
	}
}

func (l *StdLogger) Fatal(format string, args ...interface{}) {
	l.lg.Output(3,fmt.Sprintf(format, args...))
}

func (l *StdLogger) Debug(format string, args ...interface{}) {
	l.lg.Output(3,fmt.Sprintf(format, args...))
}

func (l *StdLogger) Info(format string, args ...interface{}) {
	l.lg.Output(3,fmt.Sprintf(format, args...))
}

func (l *StdLogger) Error(format string, args ...interface{}) {
	l.lg.Output(3,fmt.Sprintf(format, args...))
}

func (l *StdLogger) Log(level string, format string, args ...interface{}) {
	l.lg.Output(3,fmt.Sprintf(format, args...))
}

func (l *StdLogger) SetFormatter(format logrus.Formatter) {

}
