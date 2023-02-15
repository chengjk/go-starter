package dlog

import (
	"go-starter/internal/utils/log"
	"go.uber.org/zap"
	"time"
)

type DataLog interface {
	Info(line string)
	Error(line string)
}

type DLog struct {
	logger     *zap.Logger
	errLogger  *zap.Logger
	Lane       string
	Format     string
	RotateTime time.Duration
}

func (d DLog) Info(line string) {
	d.logger.Info(line)
}

func (d DLog) Error(line string) {
	d.errLogger.Info(line)
}

func initDLog(d *DLog, conf *log.Config) {
	format := d.Lane + "_" + d.Format
	d.logger = NewZapPlain(&log.Config{
		Path:       conf.Path + "/data/" + d.Lane,
		Format:     format,
		MaxSize:    256,
		MaxAge:     24,
		Stdout:     false,
		Level:      "info",
		RotateTime: d.RotateTime,
	})
	d.errLogger = NewZapPlain(&log.Config{
		Path:       conf.Path + "/error/" + d.Lane,
		Format:     format,
		MaxSize:    256,
		MaxAge:     24,
		Stdout:     false,
		Level:      "info",
		RotateTime: d.RotateTime,
	})
}
