package router

import (
	"webApi/live_ws/api"

	"github.com/gin-gonic/gin"
)

func InitWsRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("live")
	{
		UserRouter.GET("/ws", api.WsHandler)
	}
}
