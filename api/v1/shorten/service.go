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
		return
	}
	res.Items = shortens
	res.Total = len(shortens)
	return
}

func (service *Service) Create(ctx context.Context, req RequestCreate) (res ResponseCreate, status int, err error) {
	status = 500
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
		status = 400
		err = errors.New("invalid url host")
		return
	}
	if !(resCheckUrl.StatusCode >= 200 && resCheckUrl.StatusCode <= 300) {
		status = 400
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
			status = 422
			return
		} else {
			check := Shorten{Shortcode: req.Shortcode}
			service.repo.Get(ctx, &check)
			if check.ID != 0 {
				err = errors.New("The desired shortcode is already in use. ")
				status = 409
				return
			}
		}
	}
	shorten, err := service.repo.Create(ctx, req)
	if err != nil {
		status = 500
		return
	}
	res.Shortcode = shorten.Shortcode
	status = 201
	return
}

func (service *Service) Status(ctx context.Context, code string) (res Shorten, status int, err error) {
	status = 500
	res.Shortcode = code
	err = service.repo.Get(ctx, &res)
	if err != nil {
		log.Error(ctx, err)
		status = 500
		return
	}
	if res.ID == 0 {
		status = 404
		err = errors.New("The shortcode cannot be found in the system")
		log.Error(ctx, err)
	}
	status = 200
	return
}

func (service *Service) Execute(ctx context.Context, code string) (res Shorten, status int, err error) {
	status = 500
	res.Shortcode = code
	err = service.repo.Get(ctx, &res)
	if err != nil {
		status = 500
		return
	}

	if res.ID == 0 {
		status = 404
		err = errors.New("The shortcode cannot be found in the system")
	}
	service.repo.Execute(ctx, res)
	status = 302
	return
}
