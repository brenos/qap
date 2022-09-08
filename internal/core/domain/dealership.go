package domain

import "fmt"

type Dealership struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	State   string `json:"state"`
	Country string `json:"country"`
	Cars    []Car  `json:"cars"`
}

func NewDealershipWithoutCars(id, name, address, state, country string) *Dealership {
	return &Dealership{
		ID:      id,
		Name:    name,
		Address: address,
		State:   state,
		Country: country,
	}
}

func NewDealership(id, name, address, state, country string, cars []Car) *Dealership {
	return &Dealership{
		ID:      id,
		Name:    name,
		Address: address,
		State:   state,
		Country: country,
		Cars:    cars,
	}
}

func (d *Dealership) String() string {
	return fmt.Sprintf("%s - %s - %s - %s - %s - Cars: %d", d.ID, d.Name, d.Address, d.State, d.Country, len(d.Cars))
}
