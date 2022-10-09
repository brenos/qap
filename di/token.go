package di

import (
	ports "github.com/brenos/qap/internal/core/ports/token"
	"github.com/brenos/qap/internal/core/usecases"
)

func ConfigTokenDi() ports.TokenUseCase {
	tokenUseCase := usecases.NewTokenUseCase()
	return tokenUseCase
}
