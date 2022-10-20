package usecases

import (
	"fmt"
	"log"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	"github.com/brenos/qap/internal/core/domain/result"
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

func (d *dealershipUseCase) Get(id string) *domain.Result {
	dealership, err := d.dealershipRepo.Get(id)

	if err != nil {
		errTxt := fmt.Sprintf("Error getting dealership by id %s", id)
		log.Printf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}

	messageTxt := "Dealership returned!"

	if dealership == nil && err == nil {
		messageTxt = "Dealership not returned!"
		return domain.NewResultMessageAndCode(messageTxt, result.CodeOk)
	}

	return domain.NewResultMessageContextCode(messageTxt, dealership, result.CodeOk)
}

func (d *dealershipUseCase) List() *domain.Result {
	dealerships, err := d.dealershipRepo.List()
	if err != nil {
		errTxt := fmt.Sprintf("Error getting dealership")
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}
	return domain.NewResultMessageContextCode("Dealerships returned!", dealerships, result.CodeOk)
}

func (d *dealershipUseCase) ListByCountryAndOrState(country, state string) *domain.Result {
	dealerships, err := d.dealershipRepo.ListByCountryAndState(country, state)
	if err != nil {
		errTxt := fmt.Sprintf("Error getting dealerships by country %s and state %s", country, state)
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}
	return domain.NewResultMessageContextCode("Dealerships by country or state returned!", dealerships, result.CodeOk)
}

func (d *dealershipUseCase) Create(dealershipRequest *domain.CreateDealershipRequest) *domain.Result {
	var dealershipId = helpers.RandomUUIDAsString()
	newDealership := domain.NewDealership(dealershipId, dealershipRequest.Name, dealershipRequest.Address, dealershipRequest.State, dealershipRequest.Country)

	rowsInserted, err := d.dealershipRepo.Create(newDealership)
	if err != nil {
		errTxt := "Error creating dealership"
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}

	messageTxt := "Dealership created!"

	if rowsInserted <= 0 {
		messageTxt = "Dealership not created!"
	}

	return domain.NewResultMessageAndCode(messageTxt, result.CodeOk)
}

func (d *dealershipUseCase) Update(dealership *domain.Dealership) *domain.Result {
	rowsAffected, err := d.dealershipRepo.Update(dealership)

	if err != nil {
		errTxt := "Error updating dealership"
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}

	messageTxt := "Dealership updated!"

	if rowsAffected <= 0 {
		messageTxt = "Dealership not updated!"
	}

	return domain.NewResultMessageAndCode(messageTxt, result.CodeOk)
}

func (d *dealershipUseCase) Delete(id string) *domain.Result {
	rowsDeleted, err := d.dealershipRepo.Delete(id)

	if err != nil {
		errTxt := "Error deleting dealership"
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}

	messageTxt := "Dealership deleted!"

	if rowsDeleted <= 0 {
		messageTxt = "Dealership not deleted!"
	}

	return domain.NewResultMessageAndCode(messageTxt, result.CodeOk)
}
