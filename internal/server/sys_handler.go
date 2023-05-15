package server

import (
	"github.com/gin-gonic/gin"
	"go-starter/internal/pkg/log"
	"syscall"
)

func PreStop(ctx *gin.Context) {
	log.Info("pre stop.")
	Instance().QuitChan <- syscall.SIGQUIT
}
