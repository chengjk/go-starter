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

func Start(c *config.Config) {
	log.DoInit(c.Logs)
	dlog.Init(c.Logs)

	//pprof 和 prometheus的端口可以合并

	if c.Pprof != nil && c.Pprof.Enable {
		// /debug/pprof
		go http.ListenAndServe(c.Pprof.Address, nil)
	}

	if c.Prom != nil && c.Prom.Enable {
		//prom metrics
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			go http.ListenAndServe(c.Prom.Address, nil)
		}()
	}
	server = NewServer(c)
	panic(server.Start(c))
}

func NewServer(c *config.Config) *Server {
	if c.Env == config.PROD {
		gin.SetMode(gin.ReleaseMode)
	}
	server = &Server{
		Conf: c,
		Http: &http.Server{
			Addr:    c.Address,
			Handler: addRouterHandler(),
		},
		QuitChan: make(chan os.Signal, 1),
	}
	return server
}

func addRouterHandler() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/test", test)
	r.GET("/pre-stop", preStop)
	v1 := r.Group("v1")
	v1.GET("/ping", ping)
	return r
}
