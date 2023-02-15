package log

import (
	"go.uber.org/zap"
)

var (
	Logger  *zap.SugaredLogger
	_logger *zap.Logger
	_sugar  *zap.SugaredLogger
)

//default setting. should be override
func init() {
	_c := &Config{
		Path:    "./logs",
		MaxSize: 2000,
		Stdout:  true,
		Level:   "debug",
	}
	DoInit(_c)
}

func DoInit(conf *Config) {
	_logger = NewZap(conf)
	Logger = _logger.Sugar()
	_sugar = _logger.Sugar()
}

func Debug(args ...interface{}) {
	_sugar.Debug("", args)
}

func Info(args ...interface{}) {
	_sugar.Info(args)
}

func Warn(args ...interface{}) {
	_sugar.Warn(args)
}

func Error(args ...interface{}) {
	_sugar.Error(args)
}

func Debugf(template string, args ...interface{}) {
	_sugar.Debugf(template, args)
}

func Infof(template string, args ...interface{}) {
	_sugar.Infof(template, args)
}

func Warnf(template string, args ...interface{}) {
	_sugar.Warnf(template, args)
}

func Errorf(template string, args ...interface{}) {
	_sugar.Errorf(template, args)
}

func Debugln(args ...interface{}) {
	_sugar.Debugln(args)
}

func Infoln(args ...interface{}) {
	_sugar.Infoln(args)
}

func Warnln(args ...interface{}) {
	_sugar.Warnln(args)
}
func Errorln(args ...interface{}) {
	_sugar.Errorln(args)
}

func Fatal(args ...interface{}) {
	_sugar.Fatal(args)
}
