package robot

import (
	"testing"
	"time"

	"github.com/goapt/golib/robot/ding"
)

func TestLimitedAlarm(t *testing.T) {
	Init(ding.NewRobot())

	for i := 0; i < 10; i++ {
		LimitedAlarm("test", time.Second*10, "123123")
	}
}
