package robot

import (
	"os"
	"testing"

	"github.com/goapt/golib/robot/ding"
	"github.com/goapt/golib/robot/wechat"
)

func TestRobot_Ding(t *testing.T) {
	if testing.Short() {
		t.Skip("skip")
	}
	r := ding.NewRobot()
	r.SetToken(os.Getenv("DING_ROBOT_TOKEN"))
	Init(r)
	err := Message("test")
	if err != nil {
		t.Error(err)
	}
}
func TestRobot_Wechat(t *testing.T) {
	if testing.Short() {
		t.Skip("skip")
	}
	r := wechat.NewRobot()
	r.SetToken(os.Getenv("WECHAT_ROBOT_TOKEN"))
	Init(r)
	err := Message("test")
	if err != nil {
		t.Error(err)
	}
}
