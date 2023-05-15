package v1

import (
	"github.com/gin-gonic/gin"
	"go-starter/internal/pkg/resp"
)

func Ping(c *gin.Context) {
	resp.Success(c, "/v1/ping")
}
