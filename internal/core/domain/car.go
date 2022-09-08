package domain

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
