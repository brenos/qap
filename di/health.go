package di

import (
	"database/sql"

	healthservice "github.com/brenos/qap/internal/adapters/http/health"
	ports "github.com/brenos/qap/internal/core/ports/health"
	"github.com/brenos/qap/internal/core/usecases"
)

func ConfigHealthDI(conn *sql.DB) ports.HealthService {
	healthUseCase := usecases.NewHealthUseCase(conn)
	healthService := healthservice.New(healthUseCase)

	return healthService
}
