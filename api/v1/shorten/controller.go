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

/**
 * @Description list all shorten url
 */
func (ctrl *Controller) Get(c *gin.Context) {
	ctx := c.Request.Context()

	var responseBody ResponseList
	ctrl.service.List(ctx, &responseBody)

	response.Json(c, responseBody)
}

// @Summary check status shortcode
// @Tags shorten
// @Description get status shortcode and redirect count
// @Param code path string true "shortcode url"
// @Success 200 {object} ResponseList "success get shortcode status"
// @Failure 404 {object} response.Message "shotcode not found"
// @Failure 500 {object} response.Message "internal server error"
// @Produce json
// @Router /{code}/status [get]
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

// @Summary execute shortcode
// @Tags shorten
// @Description redirect to url by shortcode
// @Param code path string true "shortcode url"
// @Success 302 "redirect to shorten url"
// @Failure 404 {object} response.Message "shotcode not found"
// @Failure 500 {object} response.Message "internal server error"
// @Produce json
// @Router /{code} [get]
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

// @Summary create shortcode
// @Tags shorten
// @Description create shorten url and get shortcode
// @Param bodyRequest body RequestCreate true  "payload create shorten url"
// @Success 201	{object} ResponseCreate "success"
// @Failure 400 {object} response.Message "bad request"
// @Failure 409 {object} response.Message "shortcode already used or not available"
// @Failure 422 {object} response.Message "shortcode format is invalid"
// @Failure 500 {object} response.Message "internal server error"
// @Produce json
// @Router / [post]
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
