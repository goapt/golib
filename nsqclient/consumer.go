package nsqclient

import (
	"context"
	"fmt"

	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
	"github.com/verystar/golib/color"
	"github.com/verystar/golib/logger"
)

type NsqConsumer struct {
	Consumer *nsq.Consumer
	Handlers []nsq.Handler
	ctx      context.Context
	topic    string
	channel  string
}

func NewNsqConsumer(ctx context.Context, topic, channel string, options ...func(*nsq.Config)) (*NsqConsumer, error) {
	conf := nsq.NewConfig()
	conf.MaxAttempts = 0

	for _, option := range options {
		option(conf)
	}

	consumer, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		return nil, err
	}
	return &NsqConsumer{
		Consumer: consumer,
		ctx:      ctx,
		topic:    topic,
		channel:  channel,
	}, nil
}

func (this *NsqConsumer) AddHandler(handler nsq.Handler) {
	this.Handlers = append(this.Handlers, handler)
}

func (this *NsqConsumer) Run(conf Config) {
	if len(this.Handlers) == 0 {
		errors.New("Handler Is Empty")
	}
	for _, handler := range this.Handlers {
		this.Consumer.AddHandler(handler)
	}
	if err := this.Consumer.ConnectToNSQD(conf.Host + ":" + conf.Port); err != nil {
		logger.Error("nsq:ConnectToNSQD", err)
		return
	}
	for {
		select {
		case <-this.ctx.Done():
			fmt.Println(color.Yellow("[%s] %s,%s", "stop consumer", this.topic, this.channel))
			this.Consumer.Stop()
			fmt.Println(color.Yellow("[%s] %s,%s", "stop consumer success", this.topic, this.channel))
			return
		}
	}
}
