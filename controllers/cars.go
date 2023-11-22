package controllers

import (
	"github.com/gin-gonic/gin"
	model "github.com/harisaginting/gwyn/models"
	httpModel "github.com/harisaginting/gwyn/models/http"
	service "github.com/harisaginting/gwyn/services"
	"github.com/harisaginting/gwyn/utils/http/response"
)

type CarController struct {
	service service.CarService
}

func ProviderCarController(s service.CarService) CarController {
	return CarController{
		service: s,
	}
}

// @Summary
// @Tags cars
// @Description redirect to url by shortcode
// @Param code path string true "cars url"
// @Success 200 "show list cars"
// @Failure 404 {object} response.Message "route not found"
// @Failure 500 {object} response.Message "internal server error"
// @Produce json
// @Router /cars [get]
func (ctrl *CarController) List(c *gin.Context) {
	ctx := c.Request.Context()

	car := model.Car{
		Brand:        c.Query("brand"),
		Model:        c.Query("model"),
		Transmission: c.Query("transmission"),
	}
	var responseBody httpModel.ResponseListCars
	ctrl.service.List(ctx, &responseBody, car)
	response.Json(c, responseBody)
}

func (ctrl *CarController) List2(c *gin.Context) {
	ctx := c.Request.Context()

	var responseBody httpModel.ResponseListCars2
	ctrl.service.List2(ctx, &responseBody)
	response.Json(c, responseBody)
}
