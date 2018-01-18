package nsqclient

import (
	"github.com/nsqio/go-nsq"
	"time"
)

type INsqHandler interface {
	GetTopic() string
	GetChannel() string
	SetHandle(HandleFunc)
	GetHandle() HandleFunc
	SetShouldRequeue(fn func(message *nsq.Message) (bool, time.Duration))
	GetShouldRequeue(message *nsq.Message) (bool, time.Duration)
	GetMaxAttepts() uint16
	Group(group string)
	IsOpenChannelTopic() bool
	GetChannelTopic() string
	GetSize() int
}
