package robot

import (
	"errors"
	"fmt"
)

var (
	InfoColor    = "info"
	CommentColor = "comment"
	WarningColor = "warning"

	instance RobotInterface
)

type RobotInterface interface {
	SetToken(token string)
	Message(text string, at ...string) error
	MarkdownMessage(text string, at ...string) error
	CardMessage(title, text string, btns []map[string]string) error
}

func Init(robot RobotInterface) {
	instance = robot
}

func SetToken(token string) {
	if instance != nil {
		instance.SetToken(token)
	}
}

func Message(text string, at ...string) error {
	if instance == nil {
		return errors.New("robot instance is nil")
	}

	return instance.Message(text, at...)
}

func MarkdownMessage(text string, at ...string) error {
	if instance == nil {
		return errors.New("robot instance is nil")
	}

	return instance.MarkdownMessage(text, at...)
}

func CardMessage(title, text string, btns []map[string]string) error {
	if instance == nil {
		return errors.New("robot instance is nil")
	}

	return instance.CardMessage(title, text, btns)
}

func Color(text, color string) string {
	return fmt.Sprintf(`<font color="%s">%s</font>`, color, text)
}
