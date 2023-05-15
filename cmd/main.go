package main

import (
	"flag"
	"go-starter/internal/config"
	"go-starter/internal/http"
	"go-starter/internal/server"
)

func main() {
	flag.Parse()
	config.Parse()
	serv := server.Start()
	http.New(serv)
}
