package router

import (
	"github.com/gin-gonic/gin"
	"sportshot/internal/webserver/api"
	"sportshot/internal/webserver/auth"
)

func InitEventRouter(router *gin.RouterGroup) {
	{
		Router := router.Group("/events")
		Router.Static("/static", "./internal/webserver/static/events")
		Router.GET("/search", api.GetSearchPage)
	}
	{
		protectedRouter := router.Group("/events")
		protectedRouter.Use(auth.JwtAuthMiddleware())
		protectedRouter.GET("/query", api.GetEvents)
	}
}
