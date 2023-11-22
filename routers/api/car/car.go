package car

import (
	"github.com/gin-gonic/gin"
	controller "github.com/harisaginting/gwyn/controllers"
	service "github.com/harisaginting/gwyn/services"
)

func Add(group *gin.RouterGroup) {
	c := controller.ProviderCarController(&service.Car{})
	rgroup := group.Group("cars")
	rgroup.GET("/", c.List)
	rgroup.GET("/v2", c.List2)
}
