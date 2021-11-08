package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	router.GET("/test", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "success")
	})
}
