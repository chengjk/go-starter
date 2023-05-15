package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-starter/internal/pkg/log"
)

func PreStop(ctx *gin.Context) {
	log.Info("pre stop.")
	Instance().Close(context.Background())
}
