package config

import (
	"errors"
	"go-starter/internal/pkg/log"
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
	Address string     `json:"address"`
	Prom    *PlugsItem `json:"prom"`
}

// PlugsItem 插件配置
type PlugsItem struct {
	Name   string `json:"name"`
	Enable bool   `json:"enable" yaml:"enable"`
	Path   string `json:"path"`
}

// Config 系统配置
type Config struct {
	Address    string       `json:"address" yaml:"address"`
	Env        ENV          `json:"env" yaml:"env"`
	Version    string       `json:"version" yaml:"version"`
	CronEnable bool         `json:"cron_enable" yaml:"cron_enable"`
	Plugs      *PlugsConfig `json:"plugs" yaml:"plugs"`
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

func Parse() *Config {
	var conf Config
	if err := Unmarshal(*confPath, &conf); err != nil {
		panic(err)
	}
	SysConfig = conf
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
