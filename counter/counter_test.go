package counter

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewCounter(t *testing.T) {

	ctx := context.Background()
	l := NewTimeline(86400, ctx)
	var c = NewCounter(5, 1, 3)
	if err := l.AddCounter(c); err != nil {
		t.Error(err)
	}
	l.Start()

	go func() {
		for {
			n := rand.Int31n(10)
			c.Add(n)
			fmt.Println("CounterSum", l.CounterSum(c, time.Now()))
			time.Sleep(time.Second / 2)
		}
	}()

	c.AddHandle(func(c *Counter, t time.Time, sum int32) {
		fmt.Println("handle", fmt.Sprintf("%+v", c), t, sum)
	})
	<-time.After(10 * time.Second)
}
