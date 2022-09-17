package shorten

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/harisaginting/guin/common/http/response"
)

type Controller struct {
	service Service
}

func ProviderController(s Service) Controller {
	return Controller{
		service: s,
	}
}

func (ctrl *Controller) Get(c *gin.Context) {
	ctx := c.Request.Context()

	var responseBody ResponseList
	ctrl.service.List(ctx, &responseBody)

	response.Json(c, responseBody)
}

func (ctrl *Controller) Status(c *gin.Context) {
	ctx := c.Request.Context()

	code := c.Param("code")
	d, status, err := ctrl.service.Status(ctx, code)
	switch status {
	case 200:
		res := ResponseStatus{
			StartDate:     d.StartDate,
			LastSeenDate:  d.LastSeenDate,
			RedirectCount: d.RedirectCount,
		}
		response.StatusOK(c, res)
	case 404:
		response.StatusNotFound(c, err)
	default:
		response.StatusError(c, err)
	}
}

func (ctrl *Controller) Execute(c *gin.Context) {
	ctx := c.Request.Context()

	code := c.Param("code")
	d, status, err := ctrl.service.Execute(ctx, code)
	switch status {
	case 302:
		response.StatusRedirect(c, d.Url)
	case 404:
		response.StatusNotFound(c, err)
	default:
		response.StatusError(c, err)
	}
}

func (ctrl *Controller) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var requestBody RequestCreate
	request, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.StatusError(c, err)
		return
	}
	err = json.Unmarshal([]byte(request), &requestBody)
	if err != nil {
		response.BadRequest(c)
		return
	}

	d, status, err := ctrl.service.Create(ctx, requestBody)
	switch status {
	case 201:
		response.StatusCreated(c, ResponseCreate{Shortcode: d.Shortcode})
	case 400:
		response.BadRequest(c, err.Error())
	case 409:
		response.StatusConflict(c, err)
	case 422:
		response.StatusUnprocessableEntity(c, err)
	default:
		response.StatusError(c, err)
	}
}
