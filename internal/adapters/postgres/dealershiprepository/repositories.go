package dealershiprepository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/brenos/qap/internal/core/domain"
	ports "github.com/brenos/qap/internal/core/ports/dealership"
)

type dealershipPostgre struct {
	ID      string
	Name    string
	Address string
	State   string
	Country string
}

type dealershipListPostgre []dealershipPostgre

func (p *dealershipPostgre) ToDomain() *domain.Dealership {
	return &domain.Dealership{
		ID:      p.ID,
		Name:    p.Name,
		Address: p.Address,
		State:   p.State,
		Country: p.Country,
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
}

func (p dealershipListPostgre) ToDomain() []domain.Dealership {
	dealerships := make([]domain.Dealership, len(p))
	for k, dealershipIt := range p {
		dealership := dealershipIt.ToDomain()
		dealerships[k] = *dealership
	}

	return dealerships
}

type dealershipPostgreRepo struct {
	db *sql.DB
}

func NewDealershipPostgreRepo(db *sql.DB) ports.DealershipRepository {
	return &dealershipPostgreRepo{
		db: db,
	}
}

func (p *dealershipPostgreRepo) Get(id string) (*domain.Dealership, error) {
	var dealership dealershipPostgre = dealershipPostgre{}
	stmt := fmt.Sprintf("SELECT id, \"name\", address, state, country FROM dealerships WHERE id = '%s'", id)

	result := p.db.QueryRow(stmt)
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Scan(&dealership.ID, &dealership.Name, &dealership.Address, &dealership.State, &dealership.Country)
	if err != nil {
		return nil, err
	}

	//GET DEALERSHIP

	return dealership.ToDomain(), nil
}

func (p *dealershipPostgreRepo) list(stmt string) ([]domain.Dealership, error) {
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

	return dealerships.ToDomain(), nil
}

func (p *dealershipPostgreRepo) listByCountry(country string) ([]domain.Dealership, error) {
	if strings.TrimSpace(country) == "" {
		return nil, errors.New("Country is empty")
	}

	stmt := fmt.Sprintf("SELECT id, \"name\", address, state, country FROM dealerships WHERE country like '%s'", country)
	return p.list(stmt)
}

func (p *dealershipPostgreRepo) listByState(state string) ([]domain.Dealership, error) {
	if strings.TrimSpace(state) == "" {
		return nil, errors.New("State is empty")
	}

	stmt := fmt.Sprintf("SELECT id, \"name\", address, state, country FROM dealerships WHERE state like '%s'", state)
	return p.list(stmt)
}

func (p *dealershipPostgreRepo) List() ([]domain.Dealership, error) {
	stmt := fmt.Sprintf("SELECT id, \"name\", address, state, country FROM dealerships")
	return p.list(stmt)
}

func (p *dealershipPostgreRepo) ListByCountryAndState(country, state string) ([]domain.Dealership, error) {
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

func (p *dealershipPostgreRepo) Create(newDealership *domain.Dealership) (*domain.Dealership, error) {
	stmt := "INSERT INTO dealerships (id, name, address, state, country) " +
		"VALUES(?, ?, ?, ?, ?)"
	_, err := p.db.Exec(stmt, newDealership.ID, newDealership.Name, newDealership.Address, newDealership.State, newDealership.Country)

	if err != nil {
		return nil, err
	}

	return newDealership, nil
}
