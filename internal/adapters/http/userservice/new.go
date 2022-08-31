package userservice

import (
	ports "github.com/brenos/qap/internal/core/ports/user"
)

type service struct {
	usecase ports.UserUseCase
}

// New returns contract implementation of UserService
func New(usecase ports.UserUseCase) ports.UserService {
	return &service{
		usecase: usecase,
	}
}
