package domain

import (
	"encoding/json"
	"io"
)

type CarRequest struct {
	ID           string  `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	FuelType     string  `json:"fuelType"`
	Year         float32 `json:"year"`
	Price        float32 `json:"price"`
	IdDealerShip string  `json:"idDealership"`
}

func FromJSONCarRequest(body io.Reader) (*CarRequest, error) {
	carRequest := CarRequest{}
	if err := json.NewDecoder(body).Decode(&carRequest); err != nil {
		return nil, err
	}

	return &carRequest, nil
}
