package initinal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sportshot/internal/webserver/router"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.LoadHTMLGlob("./internal/webserver/templates/*.tmpl")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	ApiGroup := r.Group("/v1")

	router.InitEventRouter(ApiGroup)
	router.InitAuthRouter(ApiGroup)

	return r
}
