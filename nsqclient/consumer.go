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

func (n *NsqConsumer) AddHandler(handler nsq.Handler) {
	n.Handlers = append(n.Handlers, handler)
}

func (n *NsqConsumer) Run(conf Config) {
	if len(n.Handlers) == 0 {
		errors.New("Handler Is Empty")
	}
	for _, handler := range n.Handlers {
		n.Consumer.AddHandler(handler)
	}
	if err := n.Consumer.ConnectToNSQD(conf.Host + ":" + conf.Port); err != nil {
		logger.Error("nsq:ConnectToNSQD", err)
		return
	}
	for {
		select {
		case <-n.ctx.Done():
			fmt.Println(color.Yellow("[%s] %s,%s", "stop consumer", n.topic, n.channel))
			n.Consumer.Stop()
			fmt.Println(color.Yellow("[%s] %s,%s", "stop consumer success", n.topic, n.channel))
			return
		}
	}
}
