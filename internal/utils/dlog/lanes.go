package dlog

import (
	"go-starter/internal/utils/log"
	"time"
)

var cache map[string]DLog

func Lane(name string) DLog {
	return cache[name]
}

func Init(config *log.Config) {
	cache = make(map[string]DLog, 0)
	lanes := []string{"aaa", "bbb"}
	for _, lane := range lanes {
		dLog := DLog{
			Lane:       lane,
			Format:     "%Y%m%d-%H%M.json",
			RotateTime: time.Minute,
		}
		initDLog(&dLog, config)
		cache[lane] = dLog
	}
	InitSingle(config.Path)
}
