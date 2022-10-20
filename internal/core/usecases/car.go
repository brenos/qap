package usecases

import (
	"fmt"
	"log"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	"github.com/brenos/qap/internal/core/domain/result"
	ports "github.com/brenos/qap/internal/core/ports/car"
)

type carUseCase struct {
	carRepo ports.CarRepository
}

func NewCarUseCase(carRepo ports.CarRepository) ports.CarUseCase {
	return &carUseCase{
		carRepo: carRepo,
	}
}

func (c *carUseCase) Get(id string) *domain.Result {
	car, err := c.carRepo.Get(id)

	if err != nil {
		errTxt := fmt.Sprintf("Error getting car by id %s", id)
		log.Printf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}

	messageTxt := "Car returned!"

	if car == nil && err == nil {
		messageTxt = "Car not returned!"
		return domain.NewResultMessageAndCode(messageTxt, result.CodeOk)
	}

	return domain.NewResultMessageContextCode(messageTxt, car, result.CodeOk)
}

func (c *carUseCase) ListByDealership(idDealership string) *domain.Result {
	cars, err := c.carRepo.ListByDealership(idDealership)
	if err != nil {
		errTxt := fmt.Sprintf("Error getting cars by dealershipId %s", idDealership)
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}
	return domain.NewResultMessageContextCode("Cars returned by dealership!", cars, result.CodeOk)
}

func (c *carUseCase) ListByBrandAndOrModel(brand, model string) *domain.Result {
	cars, err := c.carRepo.ListByBrandAndOrModel(brand, model)
	if err != nil {
		errTxt := fmt.Sprintf("Error getting cars by brand %s and model %s", brand, model)
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}
	return domain.NewResultMessageContextCode("Cars returned by brand or model!", cars, result.CodeOk)
}

func (c *carUseCase) Create(carDto *domain.CreateCarRequest) *domain.Result {
	var carId = helpers.RandomUUIDAsString()
	newCar := domain.NewCar(carId, carDto.Brand, carDto.Model, carDto.FuelType, carDto.IdDealerShip, carDto.Year, carDto.Price)

	rowsInserted, err := c.carRepo.Create(newCar)

	if err != nil {
		errTxt := "Error creating car"
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}

	messageTxt := "Car created!"

	if rowsInserted <= 0 {
		messageTxt = "Car not created!"
	}

	return domain.NewResultMessageAndCode(messageTxt, result.CodeOk)
}

func (c *carUseCase) Update(car *domain.Car) *domain.Result {
	rowsAffected, err := c.carRepo.Update(car)

	if err != nil {
		errTxt := "Error updating car"
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}

	messageTxt := "Car updated!"

	if rowsAffected <= 0 {
		messageTxt = "Car not updated!"
	}

	return domain.NewResultMessageAndCode(messageTxt, result.CodeOk)
}

func (c *carUseCase) Delete(id string) *domain.Result {
	rowsDeleted, err := c.carRepo.Delete(id)

	if err != nil {
		errTxt := "Error deleting car"
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}

	messageTxt := "Car deleted!"

	if rowsDeleted <= 0 {
		messageTxt = "Car not deleted!"
	}

	return domain.NewResultMessageAndCode(messageTxt, result.CodeOk)
}
