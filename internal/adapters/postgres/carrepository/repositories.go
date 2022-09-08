package carrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/brenos/qap/internal/core/domain"
	ports "github.com/brenos/qap/internal/core/ports/car"
)

type carPostgre struct {
	ID           string
	Brand        string
	Model        string
	FuelType     string
	Year         float32
	Price        float32
	IdDealerShip string
	Dealership   domain.Dealership
}

type carListPostgre []carPostgre

func (p *carPostgre) ToDomain() *domain.Car {
	return &domain.Car{
		ID:           p.ID,
		Brand:        p.Brand,
		Model:        p.Model,
		FuelType:     p.FuelType,
		Year:         p.Year,
		Price:        p.Price,
		IdDealerShip: p.IdDealerShip,
		Dealership:   p.Dealership,
	}
}

func (p *carPostgre) FromDomain(car *domain.Car) {
	if p == nil {
		p = &carPostgre{}
	}

	p.ID = car.ID
	p.Brand = car.Brand
	p.Model = car.Model
	p.FuelType = car.FuelType
	p.Year = car.Year
	p.Price = car.Price
	p.IdDealerShip = car.IdDealerShip
	p.Dealership = car.Dealership
}

func (p carListPostgre) ToDomain() []domain.Car {
	cars := make([]domain.Car, len(p))
	for k, carIt := range p {
		car := carIt.ToDomain()
		cars[k] = *car
	}

	return cars
}

type carPostgreRepo struct {
	db *sql.DB
}

func NewCarPostgreRepo(db *sql.DB) ports.CarRepository {
	return &carPostgreRepo{
		db: db,
	}
}

func (p *carPostgreRepo) Get(id string) (*domain.Car, error) {
	var car carPostgre = carPostgre{}
	stmt := fmt.Sprintf("SELECT c.id, c.brand, c.model, c.fueltype, c.\"year\", c.price, c.iddealership, d.id, d.\"name\", d.address, d.state, d.country "+
		"FROM cars c join dealerships d on c.iddealership = d.id WHERE c.id = '%s'", id)

	result := p.db.QueryRow(stmt)
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Scan(&car.ID, &car.Brand, &car.Model, &car.FuelType, &car.Year, &car.Price,
		&car.IdDealerShip, &car.Dealership.ID, &car.Dealership.Name, &car.Dealership.Address,
		&car.Dealership.State, &car.Dealership.Country)
	if err != nil {
		return nil, err
	}

	return car.ToDomain(), nil
}

func (p *carPostgreRepo) list(stmt string) ([]domain.Car, error) {
	var cars carListPostgre

	result, err := p.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		car := carPostgre{}

		err := result.Scan(&car.ID, &car.Brand, &car.Model, &car.FuelType, &car.Year, &car.Price,
			&car.IdDealerShip, &car.Dealership.ID, &car.Dealership.Name, &car.Dealership.Address,
			&car.Dealership.State, &car.Dealership.Country)
		if err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	return cars.ToDomain(), nil
}

func (p *carPostgreRepo) ListByDealership(idDealership string) ([]domain.Car, error) {
	if strings.TrimSpace(idDealership) == "" {
		return nil, errors.New("Dealership ID is empty")
	}

	stmt := fmt.Sprintf("SELECT c.id, c.brand, c.model, c.fueltype, c.\"year\", c.price, c.iddealership, d.id, d.\"name\", d.address, d.state, d.country "+
		"FROM cars c join dealerships d on c.iddealership = d.id WHERE c.iddealership like '%s'", idDealership)
	return p.list(stmt)
}

func (p *carPostgreRepo) listByBrand(brand string) ([]domain.Car, error) {
	if strings.TrimSpace(brand) == "" {
		return nil, errors.New("Brand is empty")
	}

	stmt := fmt.Sprintf("SELECT c.id, c.brand, c.model, c.fueltype, c.\"year\", c.price, c.iddealership, d.id, d.\"name\", d.address, d.state, d.country "+
		"FROM cars c join dealerships d on c.iddealership = d.id WHERE c.brand like '%s'", brand)
	return p.list(stmt)
}

func (p *carPostgreRepo) listByModel(model string) ([]domain.Car, error) {
	if strings.TrimSpace(model) == "" {
		return nil, errors.New("Model is empty")
	}

	stmt := fmt.Sprintf("SELECT c.id, c.brand, c.model, c.fueltype, c.\"year\", c.price, c.iddealership, d.id, d.\"name\", d.address, d.state, d.country "+
		"FROM cars c join dealerships d on c.iddealership = d.id WHERE c.model like '%s'", model)
	return p.list(stmt)
}

func (p *carPostgreRepo) ListByBrandAndOrModel(brand, model string) ([]domain.Car, error) {
	modelToCompare := strings.TrimSpace(model)
	brandToCompare := strings.TrimSpace(brand)
	if modelToCompare == "" && brandToCompare == "" {
		return nil, errors.New("Brand and Model is empty")
	}

	if modelToCompare == "" {
		return p.listByBrand(brand)
	}
	if brandToCompare == "" {
		return p.listByModel(model)
	}

	stmt := fmt.Sprintf("SELECT c.id, c.brand, c.model, c.fueltype, c.\"year\", c.price, c.iddealership, d.id, d.\"name\", d.address, d.state, d.country "+
		"FROM cars c join dealerships d on c.iddealership = d.id WHERE c.model like '%s' AND c.brand like '%s'", model, brand)
	return p.list(stmt)
}

func (p *carPostgreRepo) Create(newCar *domain.Car) (*domain.Car, error) {
	stmt := "INSERT INTO cars (id, brand, model, fueltype, \"year\", price, iddealership) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?)"

	_, err := p.db.Exec(stmt, newCar.ID, newCar.Brand, newCar.Model, newCar.FuelType, newCar.Year, newCar.Price, newCar.IdDealerShip)

	if err != nil {
		return nil, err
	}

	return newCar, nil
}