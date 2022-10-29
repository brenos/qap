package usecases

import (
	"database/sql"
	"log"

	"github.com/brenos/qap/internal/core/domain"
	"github.com/brenos/qap/internal/core/domain/result"
	ports "github.com/brenos/qap/internal/core/ports/health"
)

type healthUseCase struct {
	db *sql.DB
}

func NewHealthUseCase(db *sql.DB) ports.HealthUseCase {
	return &healthUseCase{
		db: db,
	}
}

func (h *healthUseCase) Liveness() *domain.Result {
	return domain.NewResultMessage("Healthcheck OK")
}

func (h *healthUseCase) Readiness() *domain.Result {
	err := h.db.Ping()
	if err != nil {
		log.Println("DB not connected")
		return domain.NewResultMessageAndCode("DB not connected", result.CodeInternalError)
	}
	return domain.NewResultMessageAndCode("OK", result.CodeOk)
}
