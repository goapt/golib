package nsqclient

import (
	"context"
	"fmt"
	"time"

	gonsq "github.com/nsqio/go-nsq"
	"github.com/verystar/golib/color"
	"github.com/verystar/golib/debug"
	"github.com/verystar/golib/logger"
)

func Run(group string, ctx context.Context, conf Config) {
	stop := make(chan struct{})
	defer close(stop)
	if !runMulti(ctx, group, conf) {
		panic("未匹配到NSQ的运行时参数")
	}

	for {
		select {
		case <-ctx.Done():
			debug.Tag("ctx2.Done")
			// 给3秒的处理时间
			fmt.Println(color.Yellow("%s", "----------------------------------"))
			fmt.Println(color.Yellow("%s , %s", "| give nsq consumer 3 second", "... |"))
			fmt.Println(color.Yellow("%s", "----------------------------------"))
			time.AfterFunc(time.Second*3, func() {
				stop <- struct{}{}
			})
			goto END
		}
	}
END:
	<-stop
}

func Register(h INsqHandler, group string) {
	h.Group(group)
}

func runMulti(ctx context.Context, group string, conf Config) bool {

	if _, check := NsqGroups[group]; !check {
		return false
	}

	for _, h := range NsqGroups[group] {
		for i := 0; i < h.GetSize(); i++ {
			go runNsqConsumer(ctx, h, conf, false)

			if h.IsOpenChannelTopic() {
				go runNsqConsumer(ctx, h, conf, true)
			}
		}
	}
	return true
}

func runNsqConsumer(ctx context.Context, h INsqHandler, conf Config, isChannelTopic bool) {
	var topic string
	if isChannelTopic {
		topic = h.GetChannelTopic()
	} else {
		topic = h.GetTopic()
	}

	manager, err := NewNsqConsumer(ctx, topic, h.GetChannel(), func(nc *gonsq.Config) {
		nc.MaxAttempts = h.GetMaxAttepts()
	})

	if err != nil {
		fmt.Println("NewNsqConsumer err:", err)
	}

	log := logger.NewLogger(func(config *logger.Config) {
		config.LogName = h.GetChannel()
	})

	var fn gonsq.HandlerFunc = func(m *gonsq.Message) error {
		m.DisableAutoResponse()
		d := debug.NewDebugTag()
		defer func() {
			endHandler(h.GetTopic()+":"+h.GetChannel(), d)
			if err := recover(); err != nil {
				log.Error("[Nsq Consumer Handler Recover]%s%s", err, debug.Stack(1))
				should, t := h.GetShouldRequeue(m)
				if should {
					m.Requeue(t)
				}
			}
		}()

		hd := h.GetHandle()
		err := hd(d, log, m)
		if err != nil {
			d.Tag(h.GetTopic()+":"+h.GetChannel(), err.Error())

			log.Log(logger.Compile(err.Error()), "[NSQ Consumer Error:"+h.GetTopic()+":"+h.GetChannel()+"]%v", map[string]string{
				"error":   err.Error(),
				"channel": h.GetTopic() + ":" + h.GetChannel(),
				"data":    string(m.Body),
			})

			should, t := h.GetShouldRequeue(m)
			if should {
				m.Requeue(t)
			}
		}
		m.Finish()
		return nil
	}

	manager.AddHandler(fn)
	go manager.Run(conf)
}

func endHandler(dir string, d *debug.DebugTag) {
	d.SaveToHour(dir)
}
