package robot

import (
	"sync"
	"time"

	"github.com/goapt/golib/ding/ding"
)

// 有限的发送

//var send_map = make(map[string]time.Time)
var sendMap sync.Map

func LimitedAlarm(key string, duration time.Duration, content string, at ...string) error {
	current := time.Now()
	if data, ok := sendMap.Load(key); ok {
		// 如果没有超过时间限制 则不发送
		t := data.(time.Time)
		if t.Add(duration).Unix() > current.Unix() {
			return nil
		}
	}

	err := Message(content, at...)
	if err == nil {
		// 发送成功
		sendMap.Store(key, current)
	}
	return err
}
