package crontab

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/verystar/golib/logger"
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

func (this *Crontab) GetName() string {
	return this.Name
}

func (this *Crontab) SetHandle(fn func(ctx interface{})) {
	this.handler = fn
}

func (this *Crontab) GetHandle() func(ctx interface{}) {
	return this.handler
}

func (this *Crontab) IsDue(currentDate time.Time) bool {
	descriptor := this.Date

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
			logger.Error(fmt.Sprintf("Failed to parse duration %s: %s", descriptor, err.Error()))
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
	log := logger.NewLogger(func(c *logger.Config) {
		c.LogName = "cron"
	})

	//批量执行当前crontab
	var wg sync.WaitGroup
	for _, c := range Crontabs {
		if c.IsDue(t) {
			wg.Add(1)
			go func(c ICrontab, ctx interface{}) {
				defer wg.Done()
				fmt.Println("Run:" + c.GetName())
				log.Info("[Go Cron]Run:%s", c.GetName())
				handle := c.GetHandle()
				handle(ctx)
			}(c, ctx)
		}
	}
	wg.Wait()
}
