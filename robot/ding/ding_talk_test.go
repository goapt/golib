package ding

import (
	"os"
	"testing"
)

func TestDingTalkRobot_Message(t *testing.T) {
	robot := NewRobot()
	robot.SetToken(os.Getenv("DING_ROBOT_TOKEN"))
	err := robot.Message("ccccc")

	if err != nil {
		t.Error(err)
	}
}

func TestDingTalkRobot_MarkdownMessage(t *testing.T) {
	robot := NewRobot()
	robot.SetToken(os.Getenv("DING_ROBOT_TOKEN"))

	err := robot.MarkdownMessage("## 呵呵\n\n > Hello \n\n")

	if err != nil {
		t.Error(err)
	}
}
