package api

import (
	"github.com/gin-gonic/gin"
	"github.com/harisaginting/gwyn/routers/api/shorten"
)

type Api struct {
	group *gin.RouterGroup
}

func New(group *gin.RouterGroup) (v Api) {
	v.group = group.Group("api")

	// add route shorten
	shorten.Add(group)

	return
}
