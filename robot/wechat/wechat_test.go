package wechat

import (
	"os"
	"testing"
)

func TestWechatRobot_Message(t *testing.T) {
	if testing.Short() {
		t.Skip("skip")
	}
	robot := NewRobot()
	robot.SetToken(os.Getenv("WECHAT_ROBOT_TOKEN"))
	err := robot.Message("test")
	if err != nil {
		t.Error(err)
	}
}

func TestWechatRobot_MarkdownMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("skip")
	}
	robot := NewRobot()
	robot.SetToken(os.Getenv("WECHAT_ROBOT_TOKEN"))
	err := robot.MarkdownMessage("## 呵呵\n\n > Hello \n\n")
	if err != nil {
		t.Error(err)
	}
}
