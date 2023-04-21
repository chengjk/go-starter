package server

import (
	"context"
	"go-starter/internal/config"
	"go-starter/internal/utils/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
