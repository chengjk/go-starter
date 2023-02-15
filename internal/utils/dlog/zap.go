package dlog

import (
	"go-starter/internal/utils/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// NewZapPlain 输出纯文本
func NewZapPlain(conf *log.Config) *zap.Logger {
	var c *log.Config
	c = conf
	var (
		encoderConfig = zapcore.EncoderConfig{
			MessageKey: "msg",
			LineEnding: zapcore.DefaultLineEnding,
		}
		lowLevel = zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= log.AtomicLevel(c.Level).Level() && l < zapcore.ErrorLevel
		})
		infoWrite  = []zapcore.WriteSyncer{log.TimeBaseRotate(c, log.InfoLevel)}
		lowEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	)

	if c.Stdout {
		infoWrite = append(infoWrite, zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewTee(
		zapcore.NewCore(lowEncoder, zapcore.NewMultiWriteSyncer(infoWrite...), lowLevel),
	)
	var opts = []zap.Option{
		zap.WithCaller(false),
	}
	return zap.New(core, opts...)
}
