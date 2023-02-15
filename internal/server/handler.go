package server

import (
	"github.com/gin-gonic/gin"
	"go-starter/internal/utils/log"
	"go-starter/internal/utils/monitor"
	"syscall"
)

func ping(c *gin.Context) {
	monitor.System.IncComplete()
	c.JSON(200, gin.H{
		"name":    "go-starter",
		"version": server.Conf.Version,
		"message": "pong",
	})
}

func preStop(ctx *gin.Context) {
	log.Info("pre stop.")
	server.QuitChan <- syscall.SIGQUIT
}

func test(ctx *gin.Context) {
	log.Infoln("test info msg")
	log.Warnln("test warn msg")
	log.Errorln("test error msg")
}
