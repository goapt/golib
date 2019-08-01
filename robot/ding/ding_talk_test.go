package ding

import (
	"os"
	"testing"
)

func TestDingTalkRobot_Message(t *testing.T) {
	robot := NewRobot()
	robot.SetToken(os.Getenv("DING_ROBOT_TOKEN"))
	err := robot.Message("test")

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

func TestDingTalkRobot_CardMessage(t *testing.T) {
	robot := NewRobot()
	robot.SetToken(os.Getenv("DING_ROBOT_TOKEN"))

	err := robot.CardMessage("要事提醒", "该起床了", []map[string]string{
		{
			"title":     "收到提醒",
			"actionURL": "https://www.fifsky.com/",
		},
		{
			"title":     "小睡一下",
			"actionURL": "https://www.fifsky.com/",
		},
	})

	if err != nil {
		t.Error(err)
	}
}
