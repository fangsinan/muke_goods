package router

import (
	"webApi/user_web/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user") //.Use(middlewares.JWTAuth())
	{
		UserRouter.GET("/list", api.GetUserList)
		UserRouter.GET("/login", api.PasswordLogin)
	}

}
