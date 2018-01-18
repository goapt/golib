package crontab

import (
	"fmt"
	"testing"
	"time"
)

func TestCrontab_IsDue(t *testing.T) {

	exps := map[string]string{
		"@every 1m":  "",
		"@every 2m":  "",
		"@every 3m":  "",
		"@every 65m": "",
		"@every 5m":  "",
		"@every 3h":  "",
		"18:53":      "",
	}

	fmt.Println(exps)
}

func TestCrontab_NewCrontab(t *testing.T) {
	cn := NewCrontab(func(cron *Crontab) {
		cron.Name = "test"
		cron.Date = "@every 1m"

		cron.SetHandle(func(ctx interface{}) {
			fmt.Println(time.Now().Format("2006-01-02 13:04:05"))
		})
	})

	Register(cn)
	Run(time.Now(), nil)
}
