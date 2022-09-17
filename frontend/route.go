package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harisaginting/guin/common/utils/helper"
	"github.com/harisaginting/guin/model"
)

var page model.Page

func init() {
	page = model.Page{
		Now:    helper.Now().Format("2006-01-02 15:04:05"),
		Domain: helper.MustGetEnv("HOST"),
	}
}

func Page(r *gin.RouterGroup) {
	// routing page
	r.GET("/", homepage)
	return
}

func homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"d": page,
	})
	return
}
