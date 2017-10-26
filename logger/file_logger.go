package logger

import (
	"fmt"
	"os"
	"path/filepath"
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

func (l *FileLogger) SetFormatter(format logrus.Formatter) {
	l.Formatter = format
}
