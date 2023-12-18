package config

import (
	"flag"
	"go-starter/internal/envs"
)

var SysConfig Config

func init() {
	confPath = flag.String("f", envs.ProjectDir()+"/profiles/conf_dev.yaml", "select a config file. e.g. cmd -f ./profiles/conf_dev.yaml")
	if confPath == nil || *confPath == "" {
		flag.Usage()
	}
}
