package carservice

import (
	ports "github.com/brenos/qap/internal/core/ports/car"
)

type service struct {
	usecase ports.CarUseCase
}

// New returns contract implementation of UserService
func New(usecase ports.CarUseCase) ports.CarService {
	return &service{
		usecase: usecase,
	}
}
