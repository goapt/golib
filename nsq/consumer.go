package nsq

import (
	gonsq "github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
	"context"
	"github.com/verystar/golib/logger"
	"fmt"
	"github.com/verystar/golib/color"
)

type NsqConsumer struct {
	Consumer *gonsq.Consumer
	Handlers []gonsq.Handler
	ctx      context.Context
	topic    string
	channel  string
}

func NewNsqConsumer(ctx context.Context, topic, channel string, options ...func(*gonsq.Config)) (*NsqConsumer, error) {
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
		ctx:      ctx,
		topic:    topic,
		channel:  channel,
	}, nil
}

func (this *NsqConsumer) AddHandler(handler gonsq.Handler) {
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
			fmt.Println(color.Yellow("[%s] %s,%s","stop consumer", this.topic, this.channel))
			this.Consumer.Stop()
			fmt.Println(color.Yellow("[%s] %s,%s" , "stop consumer success", this.topic, this.channel))
			return
		}
	}
}
