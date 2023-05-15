package server

import (
	"github.com/robfig/cron/v3"
	"go-starter/internal/pkg/dlog"
	"go-starter/internal/pkg/log"
	"time"
)

func StartCron() {
	c := cron.New()
	_, err := c.AddFunc("0/1 * * * *", cronJob)
	if err != nil {
		panic(err)
	}
	c.Start()
}

func cronJob() {
	line := "corn job:" + time.Now().String()
	log.Info(line)
	dlog.Info(line)
	dlog.Lane("aaa").Info(line)
	dlog.Lane("aaa").Error(line)
}
