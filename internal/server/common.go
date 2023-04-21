package server

import (
	"github.com/gin-gonic/gin"
	"go-starter/internal/config"
	"go-starter/internal/utils/log"
	"go-starter/internal/utils/monitor"
	"go-starter/internal/utils/resp"
	"syscall"
)

func Ping(c *gin.Context) {
	monitor.System.IncComplete()
	resp.Success(c, gin.H{
		"name":    "go-starter",
		"version": config.SysConfig.Version,
		"message": "pong",
	})
}

func PreStop(ctx *gin.Context) {
	log.Info("pre stop.")
	Instance().QuitChan <- syscall.SIGQUIT
}

func Test(ctx *gin.Context) {
	log.Infoln("test info msg")
	log.Warnln("test warn msg")
	log.Errorln("test error msg")
	resp.Success(ctx, nil)
}
