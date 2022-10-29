package healthPorts

import "github.com/brenos/qap/internal/core/domain"

type HealthUseCase interface {
	Liveness() *domain.Result
	Readiness() *domain.Result
}
