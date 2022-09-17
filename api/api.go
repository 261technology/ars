package api

import (
	"github.com/gin-gonic/gin"
	shorten "github.com/harisaginting/guin/api/v1/shorten"
)

// Swagger Config
// @title GUIN
// @version 1.0
// @description GUIN
// @host localhost:4000
// @BasePath /
// @schemes http
// @query.collection.format multi
// @contact.name Harisa Ginting
// @contact.url ”
func V1(r *gin.RouterGroup) {
	// Dependency injection
	var shortenController shorten.Controller

	// group v1
	v1 := r.Group("v1")
	{
		// config
		apiShortenGroup := v1.Group("shorten")
		{
			apiShortenGroup.POST("/", shortenController.Create)
		}
	}
	r.POST("/shorten", shortenController.Create)
	r.GET("/:code", shortenController.Execute)
	r.GET("/:code/stats", shortenController.Status)
}
