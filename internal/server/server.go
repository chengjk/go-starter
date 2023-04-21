package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-starter/internal/config"
	"go-starter/internal/utils/dlog"
	"go-starter/internal/utils/log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var server *Server

func Start() *Server {
	c := &config.SysConfig
	log.DoInit(c.Logs)
	dlog.Init(c.Logs)
	loadPlugs(c)
	server = newServer(c)
	panic(server.Start())
}

func Instance() *Server {
	return server
}

func newServer(c *config.Config) *Server {
	if c.Env == config.PROD {
		gin.SetMode(gin.ReleaseMode)
	}
	config.SysConfig = *c
	server = &Server{
		Http: &http.Server{
			Addr:    c.Address,
			Handler: registerRouter(),
		},
		QuitChan: make(chan os.Signal, 1),
	}
	return server
}

func loadPlugs(c *config.Config) {
	if c.Plugs != nil && c.Plugs.Prom != nil && c.Plugs.Prom.Enable {
		//prom metrics
		http.Handle(c.Plugs.Prom.Path, promhttp.Handler())
	}
	// pprof
	// /debug/pprof
	go http.ListenAndServe(c.Plugs.Address, nil)

}
