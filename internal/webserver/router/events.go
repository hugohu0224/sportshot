package router

import (
	"github.com/gin-gonic/gin"
	"sportshot/internal/webserver/api"
)

func InitEventRouter(router *gin.RouterGroup) {
	Router := router.Group("/events")

	{
		Router.Static("/search", "./internal/webserver/static")
		Router.GET("/query", api.GetEvents)
	}
}
