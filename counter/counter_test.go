package counter

import (
	"math/rand"
	"testing"
	"time"

	"github.com/verystar/golib/debug"
	"fmt"
)

func TestNewCounter(t *testing.T) {

	l := NewTimeline(86400)
	var c = NewCounter(5, 1, 3)
	if err := l.AddCounter(c); err != nil {
		fmt.Println(err)
	}
	l.Start()


	go func() {
		for {
			n := rand.Int31n(10)
			debug.Tag("add", n)
			c.Add(n)
			fmt.Println( "CounterSum" , l.CounterSum(c , time.Now()) )
			time.Sleep(time.Second / 2)
		}
	}()

	c.AddHandle(func(c *Counter, t time.Time, sum int32) {
		debug.Tag("handle", fmt.Sprintf("%+v", c), t, sum)
	})
	<-time.After(10 * time.Second)
}
