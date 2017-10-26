package nsq

import (
	gonsq "github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
)

type NsqConsumer struct {
	Consumer *gonsq.Consumer
	Handlers []gonsq.Handler
}

func NewNsqConsumer(topic, channel string, options ...func(*gonsq.Config)) (*NsqConsumer, error) {
	conf := gonsq.NewConfig()
	conf.MaxAttempts = 0

	for _, option := range options {
		option(conf)
	}

	consumer, err := gonsq.NewConsumer(topic, channel, conf)
	if err != nil {
		return nil, err
	}
	return &NsqConsumer{
		Consumer: consumer,
	}, nil
}

func (this *NsqConsumer) AddHandler(handler gonsq.Handler) {
	this.Handlers = append(this.Handlers, handler)
}

func (this *NsqConsumer) Run(conf Config) error {
	if len(this.Handlers) == 0 {
		errors.New("Handler Is Empty")
	}
	for _, handler := range this.Handlers {
		this.Consumer.AddHandler(handler)
	}

	return this.Consumer.ConnectToNSQD(conf.Host + ":" + conf.Port)
}