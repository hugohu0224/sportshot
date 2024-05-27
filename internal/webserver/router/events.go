package router

import (
	"github.com/gin-gonic/gin"
	"sportshot/internal/webserver/api"
)

func InitEventRouter(router *gin.RouterGroup) {
	Router := router.Group("events")

	{
		Router.GET("/", api.GetEvents)
		Router.Static("/search-events", "internal/webserver/static/index.html")
	}
}
