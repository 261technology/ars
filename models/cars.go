package models

type Car struct {
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Transmission string  `json:"transmission"`
	Price        float64 `json:"price"`
}
