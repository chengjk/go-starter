package handler

import (
	"github.com/gin-gonic/gin"
	"go-starter/internal/config"
	"go-starter/internal/pkg/log"
	"go-starter/internal/pkg/monitor"
	"go-starter/internal/pkg/resp"
)

func Ping(c *gin.Context) {
	monitor.System.IncComplete()
	resp.Success(c, gin.H{
		"name":    "go-starter",
		"version": config.SysConfig.Version,
		"message": "pong",
	})
}

func Test(ctx *gin.Context) {
	log.Infoln("test info msg")
	log.Warnln("test warn msg")
	log.Errorln("test error msg")
	resp.Success(ctx, nil)
}
