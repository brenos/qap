package carrepository

import (
	"database/sql"
	"errors"
	"log"
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

func (p *carPostgre) ToCleanDomain() *domain.CleanCar {
	return &domain.CleanCar{
		ID:       p.ID,
		Brand:    p.Brand,
		Model:    p.Model,
		FuelType: p.FuelType,
		Year:     p.Year,
		Price:    p.Price,
	}
}

func (p *carPostgre) FromCleanDomain(car *domain.CleanCar) {
	if p == nil {
		p = &carPostgre{}
	}

	p.ID = car.ID
	p.Brand = car.Brand
	p.Model = car.Model
	p.FuelType = car.FuelType
	p.Year = car.Year
	p.Price = car.Price
}

func (p carListPostgre) ToCleanDomain() []domain.CleanCar {
	cars := make([]domain.CleanCar, len(p))
	for k, carIt := range p {
		car := carIt.ToCleanDomain()
		cars[k] = *car
	}

	return cars
}

// Create repository

type carPostgreRepo struct {
	db *sql.DB
}

func NewCarPostgreRepo(db *sql.DB) ports.CarRepository {
	return &carPostgreRepo{
		db: db,
	}
}

// Methods

func (p *carPostgreRepo) Get(id string) (*domain.Car, error) {
	var car carPostgre = carPostgre{}

	result := p.db.QueryRow("SELECT c.id, c.brand, c.model, c.fueltype, c.\"year\", c.price, c.iddealership, d.id, d.\"name\", d.address, d.state, d.country "+
		"FROM cars c join dealerships d on c.iddealership = d.id WHERE c.id = $1", id)
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Scan(&car.ID, &car.Brand, &car.Model, &car.FuelType, &car.Year, &car.Price,
		&car.IdDealerShip, &car.Dealership.ID, &car.Dealership.Name, &car.Dealership.Address,
		&car.Dealership.State, &car.Dealership.Country)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return car.ToDomain(), nil
	default:
		return nil, err
	}
}

func (p *carPostgreRepo) list(stmt string, args ...any) ([]domain.Car, error) {
	var cars carListPostgre

	result, err := p.db.Query(stmt, args...)
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

func (p *carPostgreRepo) listClean(stmt string, args ...any) ([]domain.CleanCar, error) {
	var cars carListPostgre

	result, err := p.db.Query(stmt, args...)
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		car := carPostgre{}

		err := result.Scan(&car.ID, &car.Brand, &car.Model, &car.FuelType, &car.Year, &car.Price)

		if err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	return cars.ToCleanDomain(), nil
}

func (p *carPostgreRepo) ListByDealership(idDealership string) ([]domain.CleanCar, error) {
	if strings.TrimSpace(idDealership) == "" {
		return nil, errors.New("dealership ID is empty")
	}

	stmt := "SELECT id, brand, model, fueltype, \"year\", price FROM cars WHERE iddealership like $1"
	return p.listClean(stmt, idDealership)
}

func (p *carPostgreRepo) listByBrand(brand string) ([]domain.Car, error) {
	if strings.TrimSpace(brand) == "" {
		return nil, errors.New("brand is empty")
	}

	stmt := "SELECT c.id, c.brand, c.model, c.fueltype, c.\"year\", c.price, c.iddealership, d.id, d.\"name\", d.address, d.state, d.country " +
		"FROM cars c join dealerships d on c.iddealership = d.id WHERE LOWER(c.brand) like $1"
	return p.list(stmt, brand)
}

func (p *carPostgreRepo) listByModel(model string) ([]domain.Car, error) {
	if strings.TrimSpace(model) == "" {
		return nil, errors.New("model is empty")
	}

	stmt := "SELECT c.id, c.brand, c.model, c.fueltype, c.\"year\", c.price, c.iddealership, d.id, d.\"name\", d.address, d.state, d.country " +
		"FROM cars c join dealerships d on c.iddealership = d.id WHERE LOWER(c.model) like $1"
	return p.list(stmt, model)
}

func (p *carPostgreRepo) ListByBrandAndOrModel(brand, model string) ([]domain.Car, error) {
	modelToCompare := strings.TrimSpace(model)
	brandToCompare := strings.TrimSpace(brand)
	if modelToCompare == "" && brandToCompare == "" {
		return nil, errors.New("brand and Model is empty")
	}

	model = strings.ToLower(model)
	brand = strings.ToLower(brand)

	if modelToCompare == "" {
		return p.listByBrand(brand)
	}
	if brandToCompare == "" {
		return p.listByModel(model)
	}

	stmt := "SELECT c.id, c.brand, c.model, c.fueltype, c.\"year\", c.price, c.iddealership, d.id, d.\"name\", d.address, d.state, d.country " +
		"FROM cars c join dealerships d on c.iddealership = d.id WHERE LOWER(c.model) like $1 AND LOWER(c.brand) like $2"
	return p.list(stmt, model, brand)
}

func (p *carPostgreRepo) Create(newCar *domain.Car) (int64, error) {
	stmt := "INSERT INTO cars (id, brand, model, fueltype, \"year\", price, iddealership) VALUES($1, $2, $3, $4, $5, $6, $7)"

	result, err := p.db.Exec(stmt, newCar.ID, newCar.Brand, newCar.Model, newCar.FuelType, newCar.Year, newCar.Price, newCar.IdDealerShip)

	rowsInserted, errResult := result.RowsAffected()

	if errResult != nil {
		log.Panicf("Error on get rows affected - %s", errResult.Error())
	}

	return rowsInserted, err
}

func (p *carPostgreRepo) Update(car *domain.Car) (int64, error) {
	stmt := "UPDATE cars SET brand=$1, model=$2, fueltype=$3, \"year\"=$4, price=$5, iddealership=$6 WHERE id=$7"

	result, err := p.db.Exec(stmt, car.Brand, car.Model, car.FuelType, car.Year, car.Price, car.IdDealerShip, car.ID)

	rowsAffected, errResult := result.RowsAffected()

	if errResult != nil {
		log.Panicf("Error on get rows affected - %s", errResult.Error())
	}
	return rowsAffected, err
}

func (p *carPostgreRepo) Delete(id string) (int64, error) {
	stmt := "DELETE from cars WHERE id=$1"

	result, err := p.db.Exec(stmt, id)

	rowsDeleted, errResult := result.RowsAffected()

	if errResult != nil {
		log.Panicf("Error on get rows affected - %s", errResult.Error())
	}
	return rowsDeleted, err
}
