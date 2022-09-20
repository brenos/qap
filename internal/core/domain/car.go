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
	IdDealerShip string     `json:"idDealership"`
	Dealership   Dealership `json:"dealership"`
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
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	FuelType     string  `json:"fuelType"`
	Year         float32 `json:"year"`
	Price        float32 `json:"price"`
	IdDealerShip string  `json:"idDealership"`
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

func FromJSONCreateCarRequest(body io.Reader) (*CreateCarRequest, error) {
	createCarRequest := CreateCarRequest{}
	if err := json.NewDecoder(body).Decode(&createCarRequest); err != nil {
		return nil, err
	}

	return &createCarRequest, nil
}
