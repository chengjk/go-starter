package config

import (
	"errors"
	"flag"
	"go-starter/internal/utils/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ENV string

const (
	DEV  ENV = "dev"
	TEST ENV = "test"
	PROD ENV = "prod"
)

// PlugsConfig 插件配置
type PlugsConfig struct {
	Enable  bool   `json:"enable" yaml:"enable"`
	Address string `json:"address" yaml:"address"`
}

// Config 系统配置
type Config struct {
	Address    string       `json:"address" yaml:"address"`
	Env        ENV          `json:"env" yaml:"env"`
	Version    string       `json:"version" yaml:"version"`
	CronEnable bool         `json:"cron_enable" yaml:"cron_enable"`
	Prom       *PlugsConfig `json:"prom" yaml:"prom"`
	Pprof      *PlugsConfig `json:"pprof" yaml:"pprof"`
	Logs       *log.Config  `json:"logs" yaml:"logs"`
}

func (c *Config) Valid() error {
	if c == nil {
		panic(c)
	}
	//other validation
	return nil
}

var confPath *string

func init() {
	confPath = flag.String("f", "./configs/conf_dev.yaml", "select a config file. e.g. cmd -f ./configs/conf_dev.yaml")
	if confPath == nil || *confPath == "" {
		flag.Usage()
	}
}

func Parse() *Config {
	var conf Config
	if err := Unmarshal(*confPath, &conf); err != nil {
		panic(err)
	}
	return &conf
}

func Unmarshal(path string, out interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, out); err != nil {
		return err
	}
	if o, ok := out.(interface{ Valid() error }); ok {
		return o.Valid()
	}
	return errors.New("invalid type")
}
