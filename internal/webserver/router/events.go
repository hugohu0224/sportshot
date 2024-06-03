package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sportshot/internal/webserver/api"
)

func InitEventRouter(router *gin.RouterGroup) {
	Router := router.Group("/events")

	{
		Router.Static("/static", "./internal/webserver/static")
		Router.GET("/search", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "search-events.tmpl", gin.H{})
		})
		Router.GET("/query", api.GetEvents)
	}
}
