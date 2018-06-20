package crontab

import (
	"fmt"
	"testing"
	"time"
)

func TestCrontab_IsDue(t *testing.T) {

	exps := map[string]string{
		"@every 1m":  "2012-10-24 07:47:01",
		"@every 2m":  "2012-10-24 07:02:01",
		"@every 3m":  "2012-10-24 07:03:01",
		"@every 1h":  "2012-10-24 07:00:01",
		"@every 2h":  "2012-10-24 08:00:01",
		"18:53":      "2012-10-24 18:53:01",
	}

	cn := NewCrontab()

	for k, v := range exps {
		cn.Date = k
		tt, _ := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
		ok := cn.IsDue(tt)

		if !ok {
			t.Errorf("Crontab not match:%s->%s", k, v)
		}
	}
}

func TestCrontab_NewCrontab(t *testing.T) {
	cn := NewCrontab(func(cron *Crontab) {
		cron.Name = "test"
		cron.Date = "@every 1m"

		cron.SetHandle(func(ctx interface{}) {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		})
	})

	Register(cn)
	Run(time.Now(), nil)
}
