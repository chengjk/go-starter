package dlog

import (
	"go-starter/internal/pkg/log"
	"go.uber.org/zap"
	"time"
)

var (
	dataLogger *zap.Logger
)

func InitSingle(path string) {
	dataLogger = NewZapPlain(
		&log.Config{
			Path:       path + "/data",
			Format:     "data_%Y%m%d-%H.json",
			MaxSize:    256,
			MaxAge:     24,
			Stdout:     false,
			Level:      "info",
			RotateTime: time.Hour,
		})
}

func Info(line string) {
	dataLogger.Info(line)
}
