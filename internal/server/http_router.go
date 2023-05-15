package server

import (
	"github.com/gin-gonic/gin"
	"go-starter/internal/pkg/http/middleware"
	"go-starter/internal/server/handler"
	v1 "go-starter/internal/server/handler/v1"
)

//使用全局中间件
func globalMiddleware(engine *gin.Engine) {
	engine.Use(middleware.CORS())
}

//注册路由
func registerRouter() *gin.Engine {
	root := gin.Default()

	globalMiddleware(root)
	root.GET("/ping", handler.Ping)
	root.GET("/test", handler.Test)
	root.GET("/pre-stop", PreStop)
	//使用局部中间件
	v1Group := root.Group("v1").Use(middleware.Limiter(10))
	v1Group.GET("/ping", v1.Ping)
	return root
}
