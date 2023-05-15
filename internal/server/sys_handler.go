package server

import (
	"github.com/gin-gonic/gin"
	"go-starter/internal/pkg/log"
)

func PrettyStop(ctx *gin.Context) {
	log.Info("pre stop.")
	server.Close()
}
