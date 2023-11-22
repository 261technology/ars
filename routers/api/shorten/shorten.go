package shorten

import (
	"github.com/gin-gonic/gin"
	controller "github.com/harisaginting/gwyn/controllers"
	service "github.com/harisaginting/gwyn/services"
)

func Add(group *gin.RouterGroup) {
	c := controller.ProviderShortenController(&service.Shorten{})
	rgroup := group.Group("shorten")
	rgroup.POST("/", c.Create)
	rgroup.POST("/shorten", c.Create)
	rgroup.GET("/:code", c.Execute)
	rgroup.GET("/:code/stats", c.Status)
}
