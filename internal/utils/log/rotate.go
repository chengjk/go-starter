package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go-starter/internal/utils/log/loglevel"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
	"time"
)

func TimeBaseRotate(c *Config, l loglevel.level) zapcore.WriteSyncer {
	logFile := filepath.Join(c.Path, c.Format)
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithRotationSize(int64(c.MaxSize*1024*1024)),
		rotatelogs.WithMaxAge(time.Duration(c.MaxAge)*time.Hour),
		rotatelogs.WithRotationTime(c.RotateTime))
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(rotator)
}

func SizeBasedRotate(c *Config, l loglevel.level) zapcore.WriteSyncer {
	var (
		format = fmt.Sprintf("%s.log", l)
		name   = filepath.Join(c.Path, format)
	)
	rotator := &lumberjack.Logger{
		Filename:   name,          // 日志文件名字
		MaxSize:    c.MaxSize,     // 日志文件的大小 单位是M 默认是100M
		MaxAge:     c.MaxAge / 24, // 日志文件的生命周期(天).
		MaxBackups: c.MaxBackups,  // 保留旧的日志文件最大数量
		LocalTime:  true,          // 是否使用本地时间进行日志格式化等,默认false代表使用UTC时间
		Compress:   c.Compress,    // 开启gzip压缩 默认不开启
	}
	return zapcore.AddSync(rotator)
}
