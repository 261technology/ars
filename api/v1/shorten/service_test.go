package shorten

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/harisaginting/guin/common/utils/helper"
)

// func TestExecute(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()

// 	w := httptest.NewRecorder()

//		req, _ := http.NewRequest(http.MethodGet, "/", nil)
//		router.ServeHTTP(w, req)
//		//Assertion
//		assert.Equal(t, http.StatusOK, w.Code)
//	}
var ctrlTest Controller
var ctx context.Context
var code string

const projectDirName = "guin"
const urltest = "www.google.com"

func init() {
	ctx = context.Background()
	helper.LoadEnv(projectDirName)
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	req := RequestCreate{
		URL: urltest,
	}
	res, status, _ := ctrlTest.service.Create(ctx, req)
	t.Logf("test create shorten with data %v", req)

	//Assertion
	assert.Equal(t, http.StatusCreated, status)
	code = res.Shortcode
}

func TestStatus(t *testing.T) {
	_, status, _ := ctrlTest.service.Status(ctx, code)
	t.Logf("test get code : %s with status %d", code, status)

	//Assertion
	assert.Equal(t, http.StatusOK, status)
}

func TestExecute(t *testing.T) {
	_, status, _ := ctrlTest.service.Execute(ctx, code)
	t.Logf("test get url with code : %s", code)
	//Assertion
	assert.Equal(t, http.StatusFound, status)
}
