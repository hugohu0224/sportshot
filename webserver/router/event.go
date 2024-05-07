package router

import (
	"github.com/gin-gonic/gin"
	"sportshot/webserver/api"
)

func InitEventRouter(router *gin.RouterGroup) {
	Router := router.Group("event")
	{
		Router.GET("/", api.GetEvents)
	}
}
