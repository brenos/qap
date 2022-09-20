package usecases

import (
	"log"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	ports "github.com/brenos/qap/internal/core/ports/user"
)

type userUserCase struct {
	userRepo ports.UserRepository
}

func NewUserUseCase(userRepo ports.UserRepository) ports.UserUseCase {
	return &userUserCase{
		userRepo: userRepo,
	}
}

func (u *userUserCase) Create(userRequest *domain.CreateUserRequest) (*domain.User, error) {
	var userId = helpers.RandomUUIDAsString()
	var token = userId
	newUser := domain.NewUser(userId, userRequest.Email, token, userRequest.IsPaidUser, 0)

	_, err := u.userRepo.Create(newUser)
	if err != nil {
		log.Panicf("Error creating from repo - %s", err)
		return nil, err
	}

	return newUser, nil
}

func (u *userUserCase) GetByEmail(email string) (*domain.User, error) {
	userGetted, err := u.userRepo.GetByEmail(email)
	if err != nil {
		log.Panicf("Error getting from repo - %s", err)
		return nil, err
	}
	return userGetted, nil
}

func (u *userUserCase) GetByToken(token string) (*domain.User, error) {
	userGetted, err := u.userRepo.GetByToken(token)
	if err != nil {
		log.Panicf("Error getting from repo - %s", err)
		return nil, err
	}
	return userGetted, nil
}

func (u *userUserCase) UpdateRequestCount(token string) error {
	return nil
}
