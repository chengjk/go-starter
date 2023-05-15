package http

import "C"
import (
	"github.com/gin-gonic/gin"
	"go-starter/internal/config"
	"go-starter/internal/pkg/http/middleware"
	"go-starter/internal/pkg/log"
	"go-starter/internal/server"
	"go-starter/internal/server/handler"
	v1 "go-starter/internal/server/handler/v1"
	"net/http"
)

func New(server *server.Server) {
	root := gin.Default()
	server.Http = &http.Server{
		Addr:    config.SysConfig.Address,
		Handler: root,
	}
	go func() {
		if err := server.Http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server listen err:%s", err)
		}
	}()

	globalMiddleware(root)
	registerRouter(root)
}

//使用全局中间件
func globalMiddleware(engine *gin.Engine) {
	engine.Use(middleware.CORS())
}

//注册路由
func registerRouter(root *gin.Engine) *gin.Engine {

	root.GET("/ping", handler.Ping)
	root.GET("/test", handler.Test)
	root.GET("/pre-stop", server.PreStop)
	//使用局部中间件
	v1Group := root.Group("v1").Use(middleware.Limiter(10))
	v1Group.GET("/ping", v1.Ping)
	return root
}
