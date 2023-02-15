package main

import (
	"flag"
	"go-starter/internal/config"
	"go-starter/internal/server"
)

func main() {
	flag.Parse()
	conf := config.Parse()
	server.Start(conf)
}
