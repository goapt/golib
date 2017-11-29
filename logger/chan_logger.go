package logger

import (
	"github.com/sirupsen/logrus"
)

var _ ILogger = (*ChanLogger)(nil)

type logdata struct {
	log    *ChanLogger
	level  string
	format string
	args []interface {}
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

func (l *ChanLogger) SetFormatter(format logrus.Formatter) {
}
