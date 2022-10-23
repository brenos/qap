package domain

import (
	"encoding/json"
	"io"
)

type Car struct {
	ID           string     `json:"id"`
	Brand        string     `json:"brand"`
	Model        string     `json:"model"`
	FuelType     string     `json:"fuelType"`
	Year         float32    `json:"year"`
	Price        float32    `json:"price"`
	IdDealerShip string     `json:"idDealership,omitempty"`
	Dealership   Dealership `json:"dealership,omitempty"`
}

type CleanCar struct {
	ID       string  `json:"id"`
	Brand    string  `json:"brand"`
	Model    string  `json:"model"`
	FuelType string  `json:"fuelType"`
	Year     float32 `json:"year"`
	Price    float32 `json:"price"`
}

type CreateCarRequest struct {
	Brand        string  `json:"brand" binding:"required,min=3"`
	Model        string  `json:"model" binding:"required,min=3"`
	FuelType     string  `json:"fuelType" binding:"required,min=3"`
	Year         float32 `json:"year" binding:"required,gte=1900"`
	Price        float32 `json:"price" binding:"required,gte=0"`
	IdDealerShip string  `json:"idDealership" binding:"required"`
}

func NewCar(id, brand, model, fuelType, idDealership string, year, price float32) *Car {
	return &Car{
		ID:           id,
		Brand:        brand,
		Model:        model,
		FuelType:     fuelType,
		Year:         year,
		Price:        price,
		IdDealerShip: idDealership,
	}
}

func FromJSONCarRequest(body io.Reader) (*Car, error) {
	car := Car{}
	if err := json.NewDecoder(body).Decode(&car); err != nil {
		return nil, err
	}

	return &car, nil
}
