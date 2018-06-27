package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var _ ILogger = (*FileLogger)(nil)
// FileLogger file logger
type FileLogger struct {
	*logrus.Logger
}

// NewFileLogger providers a file logger based on logrus
func NewFileLogger(options ...func(*logrus.Logger)) (ILogger, error) {
	l := &logrus.Logger{
		Out: os.Stderr,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		Hooks: make(logrus.LevelHooks),
	}

	for _, option := range options {
		option(l)
	}

	return &FileLogger{
		l,
	}, nil
}

func (l *FileLogger) Debug(format string, args ...interface{}) {
	l.WithFields(logrus.Fields{
		"fingerprint": []string{format},
	}).Debugf(format, args...)
}

func (l *FileLogger) Info(format string, args ...interface{}) {
	l.WithFields(logrus.Fields{
		"fingerprint": []string{format},
	}).Infof(format, args...)
}

func (l *FileLogger) Error(format string, args ...interface{}) {
	l.WithFields(logrus.Fields{
		"fingerprint": []string{format},
	}).Errorf(format, args...)
}

func (l *FileLogger) Fatal(format string, args ...interface{}) {
	l.WithFields(logrus.Fields{
		"fingerprint": []string{format},
	}).Fatalf(format, args...)
}

func (l *FileLogger) Log(level string, format string, args ...interface{}) {
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

func (l *FileLogger) SetFormatter(format logrus.Formatter) {
	l.Formatter = format
}
