package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var logHandler = sync.Map{}

// FileHook to send logs via syslog.
type FileHook struct {
	conf    *Config
	logFile string
	mu      sync.RWMutex
}

func NewFileHook(conf *Config) (*FileHook, error) {

	if _, err := os.Stat(conf.LogPath); err != nil {
		err = os.MkdirAll(conf.LogPath, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("can't mkdirall directory: path = %v, err = %v", conf.LogPath, err)
		}
	}

	hook := &FileHook{
		conf: conf,
	}
	return hook, nil
}

func (hook *FileHook) Fire(entry *logrus.Entry) error {
	hook.mu.Lock()
	defer hook.mu.Unlock()
	if hook.conf.LogMaxFiles > 0 {
		delDate := time.Now().AddDate(0, 0, -hook.conf.LogMaxFiles).Format("2006-01-02")
		os.Remove(hook.conf.LogPath + hook.conf.LogName + "-" + delDate + ".log")
	}

	d := time.Now().Format("2006-01-02")
	logFile := filepath.Join(hook.conf.LogPath, hook.conf.LogName+"-"+d+".log")

	var logWriter *os.File
	f, ok := logHandler.Load(logFile)
	if !ok {
		var err error
		logWriter, err = os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("can't open file: path = %v, err = %v", logFile, err)
		}
		logHandler.Store(logFile, logWriter)
	}else {
		logWriter = f.(*os.File)
	}

	entry.Logger.Out = logWriter
	return nil
}

func (hook *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
