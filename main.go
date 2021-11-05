package main

import (
	"github.com/gin-gonic/gin"
	"jassue-gin/bootstrap"
	"jassue-gin/global"
	"net/http"
)

func main() {
	// 初始化配置文件
	bootstrap.InitializeConfig()

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务
	err := r.Run(":" + global.App.Config.App.Port)
	if err != nil {
		return
	}
}
