package handler

import (
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/verystar/golib/debug"
	"github.com/verystar/golib/logger"
	"github.com/verystar/golib/nsqclient/delay"
)

var _ INsqHandler = (*NsqHandler)(nil)

type HandleFunc func(debug *debug.DebugTag, log logger.ILogger, message *nsq.Message) error

var NsqGroups = make(map[string][]INsqHandler)

type NsqHandler struct {
	Topic            string
	Channel          string
	Size             int
	MaxAttepts       uint16
	OpenChannelTopic bool // 是否开启独立的topic [Topic.Channel]
	Handler          HandleFunc
	shouldRequeue    func(message *nsq.Message) (bool, time.Duration)
}

func NewNsqHandler(options ... func(*NsqHandler)) *NsqHandler {
	handler := new(NsqHandler)
	for _, option := range options {
		option(handler)
	}
	return handler
}

func (this *NsqHandler) GetTopic() string {
	return this.Topic
}

func (this *NsqHandler) IsOpenChannelTopic() bool {
	return this.OpenChannelTopic
}

func (this *NsqHandler) GetChannelTopic() string {
	return this.Topic + "." + this.Channel
}

func (this *NsqHandler) GetChannel() string {
	return this.Channel
}

func (this *NsqHandler) SetHandle(fn HandleFunc) {
	this.Handler = fn
}

func (this *NsqHandler) GetHandle() HandleFunc {
	return this.Handler
}

func (this *NsqHandler) GetMaxAttepts() uint16 {
	if this.MaxAttepts == 0 {
		this.MaxAttepts = 100
	}
	return this.MaxAttepts
}

func (this *NsqHandler) SetShouldRequeue(fn func(message *nsq.Message) (bool, time.Duration)) {
	this.shouldRequeue = fn
}

func (this *NsqHandler) GetShouldRequeue(message *nsq.Message) (bool, time.Duration) {
	if this.shouldRequeue == nil {
		return delay.DefaultDelay(message)
	}

	return this.shouldRequeue(message)
}

func (this *NsqHandler) Group(group string) {
	NsqGroups[group] = append(NsqGroups[group], this)
}

func (h *NsqHandler) GetSize() int {
	if h.Size == 0 {
		return 1
	}
	return h.Size
}
