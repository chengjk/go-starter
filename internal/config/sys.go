package config

import "flag"

var SysConfig Config

func init() {
	confPath = flag.String("f", "./profiles/conf_dev.yaml", "select a config file. e.g. cmd -f ./profiles/conf_dev.yaml")
	if confPath == nil || *confPath == "" {
		flag.Usage()
	}
}
