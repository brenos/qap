package ports

import "github.com/brenos/qap/internal/core/domain"

type UserRepository interface {
	Get(id string) (*domain.User, error)
	List() ([]domain.User, error)
	Create(newUser *domain.User) (*domain.User, error)
}
