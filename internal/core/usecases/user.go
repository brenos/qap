package usecases

import (
	"log"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	emailPorts "github.com/brenos/qap/internal/core/ports/email"
	tokenPorts "github.com/brenos/qap/internal/core/ports/token"
	ports "github.com/brenos/qap/internal/core/ports/user"
)

type userUseCase struct {
	userRepo     ports.UserRepository
	tokenUseCase tokenPorts.TokenUseCase
	emailAdapter emailPorts.EmailAdapter
}

func NewUserUseCase(userRepo ports.UserRepository, tokenUseCase tokenPorts.TokenUseCase, emailAdapter emailPorts.EmailAdapter) ports.UserUseCase {
	return &userUseCase{
		userRepo:     userRepo,
		tokenUseCase: tokenUseCase,
		emailAdapter: emailAdapter,
	}
}

func (u *userUseCase) Create(userRequest *domain.CreateUserRequest) (*domain.User, error) {
	var userId = helpers.RandomUUIDAsString()
	var token = userId
	newUser := domain.NewUser(userId, userRequest.Email, userRequest.IsPaidUser, 0)

	_, err := u.userRepo.Create(newUser)
	if err != nil {
		log.Panicf("Error creating from repo - %s", err)
		return nil, err
	}

	token, errToken := u.tokenUseCase.GenerateToken(newUser)
	if errToken != nil {
		log.Panicf("Error creating token - %s", errToken)
		return nil, errToken
	}

	errEmail := u.emailAdapter.SendEmail(userRequest.Email, token)
	if errEmail != nil {
		return nil, errEmail
	}

	return newUser, nil
}

func (u *userUseCase) GetById(id string) (*domain.User, error) {
	userGetted, err := u.userRepo.GetById(id)
	if err != nil {
		log.Panicf("Error getting from repo - %s", err)
		return nil, err
	}
	return userGetted, nil
}

func (u *userUseCase) GetByEmail(email string) (*domain.User, error) {
	userGetted, err := u.userRepo.GetByEmail(email)
	if err != nil {
		log.Panicf("Error getting from repo - %s", err)
		return nil, err
	}
	return userGetted, nil
}

func (u *userUseCase) UpdateRequestCount(id string) error {
	err := u.userRepo.UpdateRequestCount(id)
	if err != nil {
		log.Panicf("Error on update request count - %s", err)
	}
	return err
}
