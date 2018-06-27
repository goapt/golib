package logger

import "testing"

func TestNewStdLogger(t *testing.T) {
	log := NewStdLogger()
	log.Debug("test debug log")
}