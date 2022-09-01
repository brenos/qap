package domain

import (
	"encoding/json"
	"io"
)

type CreateCarRequest struct {
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	FuelType     string  `json:"fuelType"`
	Year         float32 `json:"year"`
	Price        float32 `json:"price"`
	IdDealerShip string  `json:"idDealership"`
}

func FromJSONCreateCarRequest(body io.Reader) (*CreateCarRequest, error) {
	createCarRequest := CreateCarRequest{}
	if err := json.NewDecoder(body).Decode(&createCarRequest); err != nil {
		return nil, err
	}

	return &createCarRequest, nil
}
