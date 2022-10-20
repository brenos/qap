package dealershiprepository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/brenos/qap/internal/core/domain"
	portsCar "github.com/brenos/qap/internal/core/ports/car"
	ports "github.com/brenos/qap/internal/core/ports/dealership"
)

type dealershipPostgre struct {
	ID      string
	Name    string
	Address string
	State   string
	Country string
	Cars    []domain.CleanCar
}

type dealershipListPostgre []dealershipPostgre

func (p *dealershipPostgre) ToDomain() *domain.Dealership {
	return &domain.Dealership{
		ID:      p.ID,
		Name:    p.Name,
		Address: p.Address,
		State:   p.State,
		Country: p.Country,
		Cars:    p.Cars,
	}
}

func (p *dealershipPostgre) FromDomain(dealership *domain.Dealership) {
	if p == nil {
		p = &dealershipPostgre{}
	}

	p.ID = dealership.ID
	p.Name = dealership.Name
	p.Address = dealership.Name
	p.State = dealership.State
	p.Country = dealership.Country
	p.Cars = dealership.Cars
}

func (p dealershipListPostgre) ToDomain() []domain.Dealership {
	dealerships := make([]domain.Dealership, len(p))
	for k, dealershipIt := range p {
		dealership := dealershipIt.ToDomain()
		dealerships[k] = *dealership
	}

	return dealerships
}

func (p *dealershipPostgre) ToCleanDomain() *domain.CleanDealership {
	return &domain.CleanDealership{
		ID:      p.ID,
		Name:    p.Name,
		Address: p.Address,
		State:   p.State,
		Country: p.Country,
	}
}

func (p *dealershipPostgre) FromCleanDomain(dealership *domain.CleanDealership) {
	if p == nil {
		p = &dealershipPostgre{}
	}

	p.ID = dealership.ID
	p.Name = dealership.Name
	p.Address = dealership.Name
	p.State = dealership.State
	p.Country = dealership.Country
}

func (p dealershipListPostgre) ToCleanDomain() []domain.CleanDealership {
	dealerships := make([]domain.CleanDealership, len(p))
	for k, dealershipIt := range p {
		dealership := dealershipIt.ToCleanDomain()
		dealerships[k] = *dealership
	}

	return dealerships
}

// Create class

type dealershipPostgreRepo struct {
	db            *sql.DB
	carRepository portsCar.CarRepository
}

func NewDealershipPostgreRepo(db *sql.DB, carRepository *portsCar.CarRepository) ports.DealershipRepository {
	return &dealershipPostgreRepo{
		db:            db,
		carRepository: *carRepository,
	}
}

// Methods

func (p *dealershipPostgreRepo) Get(id string) (*domain.Dealership, error) {
	var dealership dealershipPostgre = dealershipPostgre{}
	stmt := fmt.Sprintf("SELECT id, \"name\", address, state, country FROM dealerships WHERE id = '%s'", id)

	result := p.db.QueryRow(stmt)
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Scan(&dealership.ID, &dealership.Name, &dealership.Address, &dealership.State, &dealership.Country)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		dealership.Cars, _ = p.carRepository.ListByDealership(dealership.ID)
		return dealership.ToDomain(), nil
	default:
		return nil, err
	}
}

func (p *dealershipPostgreRepo) list(stmt string) ([]domain.CleanDealership, error) {
	var dealerships dealershipListPostgre

	result, err := p.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		dealership := dealershipPostgre{}

		err := result.Scan(&dealership.ID, &dealership.Name, &dealership.Address, &dealership.State, &dealership.Country)
		if err != nil {
			return nil, err
		}

		dealerships = append(dealerships, dealership)
	}

	return dealerships.ToCleanDomain(), nil
}

func (p *dealershipPostgreRepo) listByCountry(country string) ([]domain.CleanDealership, error) {
	if strings.TrimSpace(country) == "" {
		return nil, errors.New("Country is empty")
	}

	stmt := fmt.Sprintf("SELECT id, \"name\", address, state, country FROM dealerships WHERE country like '%s'", country)
	return p.list(stmt)
}

func (p *dealershipPostgreRepo) listByState(state string) ([]domain.CleanDealership, error) {
	if strings.TrimSpace(state) == "" {
		return nil, errors.New("State is empty")
	}

	stmt := fmt.Sprintf("SELECT id, \"name\", address, state, country FROM dealerships WHERE state like '%s'", state)
	return p.list(stmt)
}

func (p *dealershipPostgreRepo) List() ([]domain.CleanDealership, error) {
	stmt := fmt.Sprintf("SELECT id, \"name\", address, state, country FROM dealerships")
	return p.list(stmt)
}

func (p *dealershipPostgreRepo) ListByCountryAndState(country, state string) ([]domain.CleanDealership, error) {
	countryToCompate := strings.TrimSpace(country)
	stateToCompare := strings.TrimSpace(state)
	if countryToCompate == "" && stateToCompare == "" {
		return nil, errors.New("Country and State is empty")
	}

	if countryToCompate == "" {
		return p.listByState(state)
	}
	if stateToCompare == "" {
		return p.listByCountry(country)
	}

	stmt := fmt.Sprintf("SELECT id, \"name\", address, state, country FROM dealerships WHERE country like '%s' AND state like '%s'", country, state)
	return p.list(stmt)
}

func (p *dealershipPostgreRepo) Create(newDealership *domain.Dealership) (int64, error) {
	stmt := "INSERT INTO dealerships (id, name, address, state, country) VALUES($1, $2, $3, $4, $5)"
	result, err := p.db.Exec(stmt, newDealership.ID, newDealership.Name, newDealership.Address, newDealership.State, newDealership.Country)

	rowsInserted, errResult := result.RowsAffected()

	if errResult != nil {
		log.Panicf("Error on get rows affected - %s", errResult.Error())
	}

	return rowsInserted, err
}

func (p *dealershipPostgreRepo) Update(dealership *domain.Dealership) (int64, error) {
	stmt := "UPDATE dealerships SET name=$1, address=$2, state=$3, country=$4 WHERE id=$5"
	result, err := p.db.Exec(stmt, dealership.Name, dealership.Address, dealership.State, dealership.Country, dealership.ID)

	rowsAffected, errResult := result.RowsAffected()

	if errResult != nil {
		log.Panicf("Error on get rows affected - %s", errResult.Error())
	}
	return rowsAffected, err
}

func (p *dealershipPostgreRepo) Delete(id string) (int64, error) {
	stmt := "DELETE FROM dealerships WHERE id=$1"
	result, err := p.db.Exec(stmt, id)

	rowsDeleted, errResult := result.RowsAffected()

	if errResult != nil {
		log.Panicf("Error on get rows affected - %s", errResult.Error())
	}
	return rowsDeleted, err
}
