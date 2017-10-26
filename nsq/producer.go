package nsq

import (
	"log"
	"os"
	"strings"
	gonsq "github.com/nsqio/go-nsq"
	"github.com/verystar/nsqpool"
	"github.com/verystar/golib/logger"
)

var (
	nsqList map[string]pool.Pool
	errs    []string
)

type Config struct {
	Host     string
	Port     string
	InitSize int `toml:"init_size"`
	MaxSize  int `toml:"max_size"`
}

func Connect(configs map[string]Config) {
	defer func() {
		if len(errs) > 0 {
			panic("[nsq] " + strings.Join(errs, "\n"))
		}
	}()

	nsqList = make(map[string]pool.Pool)

	for name, conf := range configs {
		n, err := NewProducerPool(conf.Host+":"+conf.Port, conf.InitSize, conf.MaxSize)
		logger.Debug("[nsq] connect:" + conf.Host + ":" + conf.Port)
		if err == nil {
			nsqList[name] = n
		} else {
			errs = append(errs, err.Error())
		}
	}
}

func Client(name ... string) (pool.Pool, bool) {
	key := "default"
	if len(name) > 0 {
		key = name[0]
	}
	n, ok := nsqList[key]
	return n, ok
}

// CreateNSQProducerPool create a nwq producer pool
func NewProducerPool(addr string, initSize, maxSize int) (pool.Pool, error) {
	factory := func() (*gonsq.Producer, error) {
		return NewProducer(addr)
	}
	nsqPool, err := pool.NewChannelPool(initSize, maxSize, factory)
	if err != nil {
		return nil, err
	}
	return nsqPool, nil
}

// CreateNSQProducer create nsq producer
func NewProducer(addr string) (*gonsq.Producer, error) {
	cfg := gonsq.NewConfig()
	producer, err := gonsq.NewProducer(addr, cfg)
	if err != nil {
		return nil, err
	}
	producer.SetLogger(log.New(os.Stderr, "", log.Flags()), gonsq.LogLevelError)
	return producer, nil
}