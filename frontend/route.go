package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harisaginting/gwyn/common/utils/helper"
	"github.com/harisaginting/gwyn/model"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var page model.Page

func init() {
	page = model.Page{
		Now:    helper.Now().Format("2006-01-02 15:04:05"),
		Domain: helper.MustGetEnv("DOMAIN"),
	}
}

func Page(r *gin.RouterGroup) {
	// routing page
	r.GET("/", homepage)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"p": page,
	})
}
