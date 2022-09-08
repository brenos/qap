package dealershipservice

import (
	ports "github.com/brenos/qap/internal/core/ports/dealership"
)

type service struct {
	usecase ports.DealershipUseCase
}

// New returns contract implementation of UserService
func New(usecase ports.DealershipUseCase) ports.DealershipService {
	return &service{
		usecase: usecase,
	}
}
