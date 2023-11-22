package services

import (
	"context"
	"strings"

	model "github.com/harisaginting/gwyn/models"
	httpModel "github.com/harisaginting/gwyn/models/http"
)

type CarService interface {
	List(ctx context.Context, res *httpModel.ResponseListCars, fileter model.Car) (err error)
	List2(ctx context.Context, res *httpModel.ResponseListCars2) (err error)
}

type Car struct {
}

func (service *Car) List(ctx context.Context, res *httpModel.ResponseListCars, filter model.Car) (err error) {
	cars := []model.Car{
		{Brand: "Ford", Model: "Fiesta", Transmission: "Manual", Price: 165000000},
		{Brand: "Ford", Model: "Fiesta", Transmission: "Manual", Price: 175000000},
		{Brand: "Ford", Model: "Fiesta", Transmission: "Automatic", Price: 18000000},
		{Brand: "Ford", Model: "Fiesta", Transmission: "Manual", Price: 155000000},
		{Brand: "VW", Model: "Polo", Transmission: "Manual", Price: 170000000},
		{Brand: "VW", Model: "Beetle", Transmission: "Manual", Price: 265000000},
		{Brand: "VW", Model: "Polo", Transmission: "Automatic", Price: 165000000},
	}

	for i := range cars {
		if i == 0 {
			continue
		}
		prev := i - 1
		if cars[i].Brand == cars[prev].Brand {
			cars[i].Brand = ""
		}
		if cars[i].Model == cars[prev].Model {
			cars[i].Model = ""
		}
		if cars[i].Transmission == cars[prev].Transmission {
			cars[i].Transmission = ""
		}
		if cars[i].Price == cars[prev].Price {
			cars[i].Price = 0
		}
	}
	result := []model.Car{}
	for _, c := range cars {
		if filter.Brand != "" || filter.Model != "" || filter.Transmission != "" {
			apply := true
			if filter.Brand != "" && strings.EqualFold(filter.Brand, c.Brand) {
				apply = false
			}
			if filter.Model != "" && strings.EqualFold(filter.Model, c.Model) {
				apply = false
			}
			if filter.Transmission != "" && strings.EqualFold(filter.Transmission, c.Transmission) {
				apply = false
			}
			if apply {
				result = append(result, c)
			}
		} else {
			result = append(result, c)
		}
	}
	res.Items = cars
	res.Total = len(cars)
	return
}

func (service *Car) List2(ctx context.Context, res *httpModel.ResponseListCars2) (err error) {
	cars := [][]string{
		{"Ford", "Fiesta", "Manual", "165000000"},
		{"Ford", "Fiesta", "Manual", "175000000"},
		{"Ford", "Fiesta", "Automatic", "18000000"},
		{"Ford", "Fiesta", "Manual", "155000000"},
		{"VW", "Polo", "Manual", "170000000"},
		{"VW", "Beetle", "Manual", "265000000"},
		{"VW", "Polo", "Automatic", "165000000"},
	}

	for i := range cars {
		if i == 0 {
			continue
		}
		prev := i - 1
		if cars[i][0] == cars[prev][0] {
			cars[i][0] = ""
		}
		if cars[i][1] == cars[prev][1] {
			cars[i][1] = ""
		}
		if cars[i][2] == cars[prev][2] {
			cars[i][2] = ""
		}
		if cars[i][3] == cars[prev][3] {
			cars[i][3] = "0"
		}
	}

	res.Items = cars
	res.Total = len(cars)
	return
}
