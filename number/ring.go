package number

import (
	"time"
)

func NewRing(min, max int) *Ring {
	r := &Ring{
		ch:  make(chan int),
		Min: min,
		Max: max,
	}

	r.init()
	return r
}

func NewShuffleRing(min, max int) *Ring {
	r := &Ring{
		ch:  make(chan int),
		Min: min,
		Max: max,
	}

	r.initShuffle()
	return r
}

type Ring struct {
	ch  chan int
	Min int
	Max int
}

func (r *Ring) initShuffle() {
	m := make(map[int]struct{})
	for i := r.Min; i <= r.Max; i++ {
		m[i] = struct{}{}
	}

	go func() {
		for {
			for k, _ := range m {
				r.Push(k)
			}
			//保证同一个周期内不会重复
			<-time.After(1 * time.Second)
		}
	}()
}

func (r *Ring) init() {
	go func() {
		for {
			for i := r.Min; i <= r.Max; i++ {
				r.Push(i)
			}
			// 取完一个周期，则停顿一秒钟，保证不会重复
			<-time.After(1 * time.Second)
		}
	}()
}

func (r *Ring) Next() int {
	return r.Pull()
}

func (r *Ring) Pull() int {
	return <-r.ch
}

func (r *Ring) Push(n int) {
	r.ch <- n
}
