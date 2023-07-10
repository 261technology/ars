package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/harisaginting/gwyn/controllers"
)

// Swagger Config
// @title gwyn
// @version 1.0
// @description gwyn
// @host localhost:4000
// @BasePath /
// @schemes http
// @query.collection.format multi
// @contact.name Harisa Ginting
// @contact.url ”
func Api(r *gin.RouterGroup) {
	// Dependency injection
	shortenController := controller.ShortenController{}

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
