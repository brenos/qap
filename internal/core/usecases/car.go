package usecases

import (
	"log"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
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

func (c *carUseCase) Get(id string) (*domain.Car, error) {
	car, err := c.carRepo.Get(id)
	if err != nil {
		log.Panicf("Error getting car by id %s from repo - %s", id, err)
		return nil, err
	}
	return car, nil
}

func (c *carUseCase) ListByDealership(idDealership string) ([]domain.CleanCar, error) {
	cars, err := c.carRepo.ListByDealership(idDealership)
	if err != nil {
		log.Panicf("Error getting cars by dealershipId %s from repo - %s", idDealership, err)
		return nil, err
	}
	return cars, nil
}

func (c *carUseCase) ListByBrandAndOrModel(brand, model string) ([]domain.Car, error) {
	cars, err := c.carRepo.ListByBrandAndOrModel(brand, model)
	if err != nil {
		log.Panicf("Error getting cars by brand %s and model %s from repo - %s", brand, model, err)
		return nil, err
	}
	return cars, nil
}

func (c *carUseCase) Create(carDto *domain.CreateCarRequest) (*domain.Car, error) {
	var carId = helpers.RandomUUIDAsString()
	newCar := domain.NewCar(carId, carDto.Brand, carDto.Model, carDto.FuelType, carDto.IdDealerShip, carDto.Year, carDto.Price)

	_, err := c.carRepo.Create(newCar)
	if err != nil {
		log.Panicf("Error creating car from repo - %s", err)
		return nil, err
	}

	return newCar, nil
}
