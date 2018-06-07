package monitor

import (
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/verystar/golib/logger"
	"github.com/verystar/golib/redis"
)

type RedisMonitor struct {
	data     chan map[string]interface{}
	dbPrefix string
}

var redisMonitor *RedisMonitor
var once sync.Once

func NewRedisMonitor(dbPrefix string) *RedisMonitor {
	//Single
	once.Do(func() {
		redisMonitor = new(RedisMonitor)
		redisMonitor.data = make(chan map[string]interface{}, 100)
		redisMonitor.dbPrefix = dbPrefix
	})
	return redisMonitor
}

func (m *RedisMonitor) Run() {
	var err error
	var marshaledBytes []byte

	client := redis.Client("stat")

	var count uint32
	pipe := client.Pipeline()
	for data := range m.data {
		n := atomic.AddUint32(&count, 1)
		if marshaledBytes, err = json.Marshal(data); err != nil {
			continue
		}
		pipe.RPush("__stat__", string(marshaledBytes))
		if n == 100 {
			atomic.StoreUint32(&count, 0)
			_, err = pipe.Exec()
			if err != nil {
				logger.Error("Stat Pipeline error", err)
			}
		}
	}
}

func FormatTime(diff_time float64) string {
	var diff_time_str string
	if diff_time < 0.05 {
		diff_time_str = "0.00s到0.05s"
	} else if diff_time < 0.1 {
		diff_time_str = "0.05s到0.1s"
	} else if diff_time < 0.5 {
		diff_time_str = "0.1s到0.5s"
	} else if diff_time < 1 {
		diff_time_str = "0.5s到1s"
	} else if diff_time < 2 {
		diff_time_str = "1s到2s"
	} else if diff_time < 3 {
		diff_time_str = "2s到3s"
	} else if diff_time < 4 {
		diff_time_str = "3s到4s"
	} else if diff_time < 5 {
		diff_time_str = "4s到5s"
	} else if diff_time < 10 {
		diff_time_str = "5s到10s"
	} else {
		diff_time_str = "10s到∞秒"
	}

	return diff_time_str
}

func Stat(num int64, v1, v2, v3 string) {
	if redisMonitor == nil || redisMonitor.data == nil {
		return
	}

	if num < 0 || v1 == "" || v2 == "" || v3 == "" {
		return
	}

	data := map[string]interface{}{
		"dbf":     redisMonitor.dbPrefix,
		"num":     num,
		"v1":      v1,
		"v2":      v2,
		"v3":      v3,
		"v4":      nil,
		"replace": false,
		"time":    time.Now().Unix(),
	}

	redisMonitor.data <- data
}
