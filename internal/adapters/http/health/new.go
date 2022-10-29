package healthservice

import (
	ports "github.com/brenos/qap/internal/core/ports/health"
)

type service struct {
	usecase ports.HealthUseCase
}

func New(usecase ports.HealthUseCase) ports.HealthService {
	return &service{
		usecase: usecase,
	}
}
