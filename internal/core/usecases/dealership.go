package usecases

import (
	"log"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	ports "github.com/brenos/qap/internal/core/ports/dealership"
)

type dealershipUseCase struct {
	dealershipRepo ports.DealershipRepository
}

func NewDealershipUseCase(dealershipRepo ports.DealershipRepository) ports.DealershipUseCase {
	return &dealershipUseCase{
		dealershipRepo: dealershipRepo,
	}
}

func (d *dealershipUseCase) Get(id string) (*domain.Dealership, error) {
	dealership, err := d.dealershipRepo.Get(id)
	if err != nil {
		log.Panicf("Error getting dealership by id %s from repo - %s", id, err)
		return nil, err
	}
	return dealership, nil
}

func (d *dealershipUseCase) List() ([]domain.Dealership, error) {
	dealerships, err := d.dealershipRepo.List()
	if err != nil {
		log.Panicf("Error getting dealerships from repo - %s", err)
		return nil, err
	}
	return dealerships, nil
}

func (d *dealershipUseCase) ListByState(state string) ([]domain.Dealership, error) {
	dealerships, err := d.dealershipRepo.ListByState(state)
	if err != nil {
		log.Panicf("Error getting dealerships by state %s from repo - %s", state, err)
		return nil, err
	}
	return dealerships, nil
}

func (d *dealershipUseCase) ListByCountry(country string) ([]domain.Dealership, error) {
	dealerships, err := d.dealershipRepo.ListByCountry(country)
	if err != nil {
		log.Panicf("Error getting dealerships by state %s from repo - %s", country, err)
		return nil, err
	}
	return dealerships, nil
}

func (d *dealershipUseCase) ListByCountryAndOrState(country, state string) ([]domain.Dealership, error) {
	if country == "" {
		return this.ListByState(state)
	}
	if state == "" {
		return this.ListByCountry(country)
	}

	dealerships, err := d.dealershipRepo.ListByCountryAndState(country, state)
	if err != nil {
		log.Panicf("Error getting dealerships by country %s and state %s from repo - %s", country, state, err)
		return nil, err
	}
	return dealerships, nil
}

func (d *dealershipUseCase) Create(dealershipRequest *domain.CreateDealershipRequest) (*domain.Dealership, error) {
	var dealershipId = helpers.RandomUUIDAsString()
	newDealership := domain.NewDealership(dealershipId, dealershipRequest.Name, dealershipRequest.Address, dealershipRequest.State, dealershipRequest.Country)

	_, err := d.dealershipRepo.Create(newDealership)
	if err != nil {
		log.Panicf("Error creating dealership from repo - %s", err)
		return nil, err
	}

	return newDealership, nil
}
