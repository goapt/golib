package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/sirupsen/logrus"
)

var _ ILogger = (*FileLogger)(nil)
// FileLogger file logger
type FileLogger struct {
	*logrus.Logger
}

// NewFileLogger providers a file logger based on logrus
func NewFileLogger(filename string, options ...func(*logrus.Logger)) (ILogger, error) {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("can't get file abs path: filename = %v, err = %v", filename, err)
	}

	dirPath := filepath.Dir(absPath)
	if _, err := os.Stat(dirPath); err != nil {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("can't mkdirall directory: path = %v, err = %v", absPath, err)
		}
	}

	f, err := os.OpenFile(absPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("can't open file: path = %v, err = %v", absPath, err)
	}

	l := &logrus.Logger{
		Out:       f,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
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

func (l *FileLogger) Compile(format string, args ...interface{}) {
	r, _ := regexp.Compile(`^<(debug|info|error|fatal)>(.*)`)
	match := r.FindStringSubmatch(format)

	if len(match) > 2 {
		l.Log(match[1], format, args...)
	} else {
		l.Error(format, args...)
	}
}

func (l *FileLogger) SetFormatter(format logrus.Formatter) {
	l.Formatter = format
}
