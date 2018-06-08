package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var _ ILogger = (*ChanLogger)(nil)

type logdata struct {
	log    *ChanLogger
	level  string
	format string
	args   []interface{}
}

var logchan chan *logdata

type ChanLogger struct {
	conf *Config
}

func NewChanLogger(conf *Config) ILogger {
	return &ChanLogger{
		conf: conf,
	}
}

func (l *ChanLogger) Fatal(format string, args ...interface{}) {
	logchan <- &logdata{
		log:    l,
		level:  LevelFatal,
		format: format,
		args:   args,
	}
	<-time.After(2 * time.Second)
	os.Exit(1)
}

func (l *ChanLogger) Debug(format string, args ...interface{}) {
	logchan <- &logdata{
		log:    l,
		level:  LevelDebug,
		format: format,
		args:   args,
	}
}

func (l *ChanLogger) Info(format string, args ...interface{}) {
	logchan <- &logdata{
		log:    l,
		level:  LevelInfo,
		format: format,
		args:   args,
	}
}

func (l *ChanLogger) Error(format string, args ...interface{}) {
	logchan <- &logdata{
		log:    l,
		level:  LevelError,
		format: format,
		args:   args,
	}
}

func (l *ChanLogger) Log(level string, format string, args ...interface{}) {
	switch level {
	case "debug":
		l.Debug(format, args...)
	case "info":
		l.Info(format, args...)
	case "error":
		l.Error(format, args...)
	case "fatal":
		l.Fatal(format, args...)
	default:
		l.Error(format, args...)
	}
}

func (l *ChanLogger) SetFormatter(format logrus.Formatter) {
}
