package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-starter/internal/config"
	"go-starter/internal/pkg/dlog"
	"go-starter/internal/pkg/log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var server *Server

type Server struct {
	Http     *http.Server
	QuitChan chan os.Signal
}

func (s Server) Close() error {
	//do some last words
	log.Info("server exiting... ")
	return nil
}

func (s Server) Start() error {
	c := &config.SysConfig
	//start job
	if c.CronEnable {
		StartCron()
	}
	//start http server
	go func() {
		if err := server.Http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server listen err:%s", err)
		}
	}()

	//add close event listener
	signal.Notify(server.QuitChan, syscall.SIGINT, syscall.SIGTERM)

	//blocking
	<-server.QuitChan
	ctx, channel := context.WithTimeout(context.Background(), 10*time.Second)
	defer channel()
	//shutdown http server
	if err := server.Http.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error")
		return err
	}
	return server.Close()
}

//-----------

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
