package initialize

import (
	"net/http"
	"webApi/user_web/middlewares"
	"webApi/user_web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 配置 consul的服务注册入口
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	// 配置跨域问题
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/v1")
	router.InitUserRouter(ApiGroup)
	return Router
}
