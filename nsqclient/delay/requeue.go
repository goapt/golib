package delay

import (
	"time"

	"github.com/nsqio/go-nsq"
)

func DefaultDelay(m *nsq.Message) (bool, time.Duration) {
	return true, 60 * time.Second
}

var Deferred = map[int]int{
	1: 15,
	2: 15,
	3: 30,
	4: 180,
	5: 1800,
	6: 1800,
	7: 1800,
	8: 1800,
	9: 3600,
}

func DeferredDelay(m *nsq.Message) (bool, time.Duration) {
	a := int(m.Attempts)
	l := len(Deferred)
	delay, ok := Deferred[a]
	if !ok {
		if a < len(Deferred) {
			return false, -1
		} else {
			return true, time.Duration(Deferred[l-1]) * time.Second
		}
	}

	return true, time.Duration(delay) * time.Second
}
