package router

import (
	"github.com/gin-gonic/gin"
	"sportshot/internal/webserver/api"
)

func InitAuthRouter(router *gin.RouterGroup) {
	Router := router.Group("/auth")

	{
		Router.Static("/static", "./internal/webserver/static/auth")
		Router.GET("/login", api.GetLoginPage)
		Router.POST("/login", api.ValidateLogin)
		Router.POST("/register", api.Register)
	}
}
