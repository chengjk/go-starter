package log

import (
	"go-starter/internal/pkg/log/loglevel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type Config struct {
	Level      string `json:"level" yaml:"level"`           // 日志级别
	Host       string `json:"host" yaml:"host"`             // os.Hostname()
	Path       string `json:"path" yaml:"path"`             // 输出路径
	Format     string `json:"format" yaml:"format"`         // 文件名格式
	Stdout     bool   `json:"stdout" yaml:"stdout"`         // 标准输出
	MaxSize    int    `json:"maxsize" yaml:"maxsize"`       // 单个文件日志大小M
	MaxAge     int    `json:"maxAge" yaml:"maxAge"`         // 最大有效期 Hour
	MaxBackups int    `json:"maxBackups" yaml:"maxBackups"` // 最大文件数量
	Compress   bool   `json:"compress" yaml:"compress"`     // 是否压缩
	//time based rotate
	RotateTime time.Duration
}

var c *Config

// NewZap 输出标准日志
func NewZap(conf *Config) *zap.Logger {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	conf.Host = hostname
	c = conf
	var (
		encoderConfig = zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "ts",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		}
		lowLevel = zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= loglevel.AtomicLevel(c.Level).Level() && l < zapcore.ErrorLevel
		})
		highLevel = zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= zapcore.ErrorLevel
		})
		infoWrite   = []zapcore.WriteSyncer{SizeBasedRotate(c, loglevel.Info)}
		errorWrite  = []zapcore.WriteSyncer{SizeBasedRotate(c, loglevel.Error)}
		lowEncoder  zapcore.Encoder
		highEncoder zapcore.Encoder
	)

	if c.Stdout {
		infoWrite = append(infoWrite, zapcore.AddSync(os.Stdout))
		errorWrite = append(errorWrite, zapcore.AddSync(os.Stderr))
	}
	lowEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	highEncoder = zapcore.NewConsoleEncoder(encoderConfig)

	//switch envs.Mode {
	//case envs.PROD:
	//	lowEncoder = zapcore.NewJSONEncoder(encoderConfig)
	//	highEncoder = zapcore.NewJSONEncoder(encoderConfig)
	//case envs.TEST, envs.DEV:
	//	fallthrough
	//default:
	//	lowEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	//	highEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	//}

	core := zapcore.NewTee(
		zapcore.NewCore(lowEncoder, zapcore.NewMultiWriteSyncer(infoWrite...), lowLevel),
		zapcore.NewCore(highEncoder, zapcore.NewMultiWriteSyncer(errorWrite...), highLevel),
	)
	var opts = []zap.Option{
		zap.WithCaller(true),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(highLevel),
		zap.Fields(zap.String("host", c.Host)),
	}
	return zap.New(zapcore.NewSamplerWithOptions(core, time.Second, 100, 100), opts...)
}
