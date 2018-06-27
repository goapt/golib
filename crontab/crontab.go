package crontab

import (
	"fmt"
	"log"
	"math"
	"strings"
	"sync"
	"time"
)

type ICrontab interface {
	SetHandle(fn func(ctx interface{}))
	GetHandle() func(ctx interface{})
	GetName() string
	IsDue(currentDate time.Time) bool
}

var _ ICrontab = (*Crontab)(nil)

// Date support every day:"00:05" or every month first day:"01 00:05" or every minute:"@every 1m"
// @every 1[m|h]
type Crontab struct {
	Name    string
	Date    string
	handler func(ctx interface{})
}

func NewCrontab(options ... func(*Crontab)) *Crontab {
	cron := new(Crontab)
	for _, option := range options {
		option(cron)
	}

	return cron
}

func (c *Crontab) GetName() string {
	return c.Name
}

func (c *Crontab) SetHandle(fn func(ctx interface{})) {
	c.handler = fn
}

func (c *Crontab) GetHandle() func(ctx interface{}) {
	return c.handler
}

func (c *Crontab) IsDue(currentDate time.Time) bool {
	descriptor := c.Date

	if descriptor == "" {
		return false
	}

	currentDate = currentDate.Add(1*time.Second - time.Duration(currentDate.Nanosecond())*time.Nanosecond)

	const every = "@every "
	if strings.HasPrefix(descriptor, every) {
		startDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.Location())
		diff := currentDate.Sub(startDate)
		diffMinutes := math.Floor(diff.Minutes())

		duration, err := time.ParseDuration(descriptor[len(every):])
		if err != nil {
			log.Printf(fmt.Sprintf("Failed to parse duration %s: %s", descriptor, err.Error()))
			return false
		}

		if duration.Minutes() == 0 {
			return false
		}

		//fmt.Println(diffMinutes,duration.Minutes())
		return int(diffMinutes)%int(duration.Minutes()) == 0
	}

	d := currentDate.Format("2006-01-02 15:04")
	return strings.Index(d, descriptor) != -1
}

var Crontabs [] ICrontab

func Register(c ICrontab) {
	Crontabs = append(Crontabs, c)
}

func Run(t time.Time, ctx interface{}) {
	//批量执行当前crontab
	var wg sync.WaitGroup
	for _, c := range Crontabs {
		if c.IsDue(t) {
			wg.Add(1)
			go func(c ICrontab, ctx interface{}) {
				defer wg.Done()
				fmt.Println("Run:" + c.GetName())
				log.Printf("[Go Cron]Run:%s", c.GetName())
				handle := c.GetHandle()
				handle(ctx)
			}(c, ctx)
		}
	}
	wg.Wait()
}
