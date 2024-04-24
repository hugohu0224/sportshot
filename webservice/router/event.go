package router

import (
	"github.com/gin-gonic/gin"
	"sportshot/webservice/api"
)

func InitEventRouter(router *gin.RouterGroup) {
	Router := router.Group("event")
	{
		Router.GET("/", api.GetEvents)
	}

}
