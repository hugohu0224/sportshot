package initinal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sportshot/internal/webserver/router"
)

func InitRouters() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	ApiGroup := r.Group("v1")

	router.InitEventRouter(ApiGroup)

	return r
}
