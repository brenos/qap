package tokenPorts

import "github.com/brenos/qap/internal/core/domain"

type TokenUseCase interface {
	GenerateToken(*domain.User) (string, error)
	VerifyToken(string) bool
	GetUserIdByToken(tokenString string) (string, error)
}
