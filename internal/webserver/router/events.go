package router

import (
	"github.com/gin-gonic/gin"
	"sportshot/internal/webserver/api"
)

func InitEventRouter(router *gin.RouterGroup) {
	Router := router.Group("/events")

	{
		Router.Static("/static", "./internal/webserver/static/events")
		Router.GET("/search", api.GetSearchPage)
		Router.GET("/query", api.GetEvents)
	}
}
