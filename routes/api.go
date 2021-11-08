package routes

import (
	"github.com/gin-gonic/gin"
	"jassue-gin/app/common/request"
	"jassue-gin/app/controllers/app"
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

	router.POST("/user/register", func(c *gin.Context) {
		var form request.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSONP(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		c.JSONP(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	router.POST("/auth/register", app.Register)
}
