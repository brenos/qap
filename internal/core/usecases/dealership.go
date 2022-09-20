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

func (d *dealershipUseCase) List() ([]domain.CleanDealership, error) {
	dealerships, err := d.dealershipRepo.List()
	if err != nil {
		log.Panicf("Error getting dealerships from repo - %s", err)
		return nil, err
	}
	return dealerships, nil
}

func (d *dealershipUseCase) ListByCountryAndOrState(country, state string) ([]domain.CleanDealership, error) {
	dealerships, err := d.dealershipRepo.ListByCountryAndState(country, state)
	if err != nil {
		log.Panicf("Error getting dealerships by country %s and state %s from repo - %s", country, state, err)
		return nil, err
	}
	return dealerships, nil
}

func (d *dealershipUseCase) Create(dealershipRequest *domain.CreateDealershipRequest) error {
	var dealershipId = helpers.RandomUUIDAsString()
	newDealership := domain.NewDealership(dealershipId, dealershipRequest.Name, dealershipRequest.Address, dealershipRequest.State, dealershipRequest.Country)

	err := d.dealershipRepo.Create(newDealership)
	if err != nil {
		log.Panicf("Error creating dealership from repo - %s", err)
	}

	return err
}

func (d *dealershipUseCase) Update(dealership *domain.Dealership) error {
	err := d.dealershipRepo.Update(dealership)

	if err != nil {
		log.Panicf("Error updating dealership from repo - %s", err)
	}
	return err
}

func (d *dealershipUseCase) Delete(id string) error {
	err := d.dealershipRepo.Delete(id)

	if err != nil {
		log.Panicf("Error deleting dealership from repo - %s", err)
	}
	return err
}
