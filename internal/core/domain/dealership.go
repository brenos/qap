package domain

import (
	"encoding/json"
	"fmt"
	"io"
)

type Dealership struct {
	ID      string     `json:"id"`
	Name    string     `json:"name"`
	Address string     `json:"address"`
	State   string     `json:"state"`
	Country string     `json:"country"`
	Cars    []CleanCar `json:"cars,omitempty"`
}

type CleanDealership struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type CreateDealershipRequest struct {
	Name    string `json:"name" binding:"required,min=3"`
	Address string `json:"address" binding:"required,min=3"`
	State   string `json:"state" binding:"required,min=3"`
	Country string `json:"country" binding:"required,min=3"`
}

func NewDealership(id, name, address, state, country string) *Dealership {
	return &Dealership{
		ID:      id,
		Name:    name,
		Address: address,
		State:   state,
		Country: country,
	}
}

func FromJSONDealershipRequest(body io.Reader) (*Dealership, error) {
	dealership := Dealership{}
	if err := json.NewDecoder(body).Decode(&dealership); err != nil {
		return nil, err
	}

	return &dealership, nil
}

func (d *Dealership) String() string {
	return fmt.Sprintf("%s - %s - %s - %s - %s - Cars: %d", d.ID, d.Name, d.Address, d.State, d.Country, len(d.Cars))
}
