package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-starter/internal/config"
	"go-starter/internal/utils/dlog"
	"go-starter/internal/utils/http/middleware"
	"go-starter/internal/utils/log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var server *Server

func Start(c *config.Config) {
	log.DoInit(c.Logs)
	dlog.Init(c.Logs)
	loadPlugs(c)
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

func loadPlugs(c *config.Config) {
	if c.Plugs != nil && c.Plugs.Prom != nil && c.Plugs.Prom.Enable {
		//prom metrics
		http.Handle(c.Plugs.Prom.Path, promhttp.Handler())
	}
	// pprof
	// /debug/pprof
	go http.ListenAndServe(c.Plugs.Address, nil)

}

func useMiddleware(engine *gin.Engine) {
	engine.Use(middleware.Limiter(1))
	engine.Use(middleware.CORS())
}

func addRouterHandler() *gin.Engine {
	r := gin.Default()
	useMiddleware(r)

	r.GET("/ping", ping)
	r.GET("/test", test)
	r.GET("/pre-stop", preStop)
	v1 := r.Group("v1")
	v1.GET("/ping", ping)
	return r
}
