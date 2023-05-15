package middleware

import (
	"github.com/gin-gonic/gin"
	"go-starter/internal/pkg/resp"
	"go-starter/internal/pkg/resp/codes"
	"golang.org/x/time/rate"
)

func Limiter(bust int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(bust), bust)
	return func(context *gin.Context) {
		if !limiter.Allow() {
			resp.SuccessWith(context, codes.TooManyReq, nil)
			context.Abort()
			return
		}
		context.Next()
	}
}
