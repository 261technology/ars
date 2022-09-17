package shorten

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"

	"github.com/harisaginting/guin/common/goflake/generator"
	"github.com/harisaginting/guin/common/log"
	"github.com/harisaginting/guin/common/utils/helper"
)

type Service struct {
	repo Repository
}

func (service *Service) List(ctx context.Context, res *ResponseList) (err error) {
	shortens, err := service.repo.FindAll(ctx)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	res.Items = shortens
	res.Total = len(shortens)
	return
}

func (service *Service) Create(ctx context.Context, req RequestCreate) (res ResponseCreate, status int, err error) {
	status = http.StatusInternalServerError
	req.URL = helper.AdjustUrl(req.URL)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	checkUrl, err := http.NewRequest("GET", req.URL, nil)
	if err != nil {
		log.Error(ctx, err, "Failed initiate request to storageService Service")
		return
	}

	resCheckUrl, err := client.Do(checkUrl)
	if err != nil {
		status = http.StatusBadRequest
		err = errors.New("invalid url host")
		log.Error(ctx, err)
		return
	}
	if !(resCheckUrl.StatusCode >= 200 && resCheckUrl.StatusCode <= 300) {
		status = http.StatusBadRequest
		err = errors.New("url host not found")
		return
	}

	if req.Shortcode == "" {
		for {
			req.Shortcode = generator.GenerateIdentifier()
			check := Shorten{Shortcode: req.Shortcode}
			service.repo.Get(ctx, &check)
			if check.ID == 0 {
				break
			}
		}
	} else {
		if !helper.IsMatchRegex(req.Shortcode) {
			err = errors.New("The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{6}$.")
			log.Error(ctx, err)
			status = http.StatusUnprocessableEntity
			return
		} else {
			check := Shorten{Shortcode: req.Shortcode}
			service.repo.Get(ctx, &check)
			if check.ID != 0 {
				err = errors.New("The desired shortcode is already in use. ")
				status = http.StatusConflict
				return
			}
		}
	}
	shorten, err := service.repo.Create(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		status = http.StatusInternalServerError
		return
	}
	res.Shortcode = shorten.Shortcode
	status = http.StatusCreated
	return
}

func (service *Service) Status(ctx context.Context, code string) (res Shorten, status int, err error) {
	status = http.StatusInternalServerError
	res.Shortcode = code
	err = service.repo.Get(ctx, &res)
	if err != nil {
		log.Error(ctx, err)
		status = http.StatusInternalServerError
		return
	}
	if res.ID == 0 {
		status = http.StatusNotFound
		err = errors.New("The shortcode cannot be found in the system")
		log.Error(ctx, err)
		return
	}
	status = http.StatusOK
	return
}

func (service *Service) Execute(ctx context.Context, code string) (res Shorten, status int, err error) {
	status = http.StatusInternalServerError
	res.Shortcode = code
	err = service.repo.Get(ctx, &res)
	if err != nil {
		status = http.StatusInternalServerError
		log.Error(ctx, err)
		return
	}

	if res.ID == 0 {
		status = http.StatusNotFound
		err = errors.New("The shortcode cannot be found in the system")
		log.Error(ctx, err)
		return
	}
	service.repo.Execute(ctx, res)
	status = http.StatusFound
	return
}
