package userrepository

import (
	"database/sql"
	"fmt"

	"github.com/brenos/qap/internal/core/domain"
	ports "github.com/brenos/qap/internal/core/ports/user"
)

type userPostgre struct {
	ID         string
	Email      string
	Token      string
	IsPaidUser bool
	RequestQtt int32
}

type userListPostgre []userPostgre

func (p *userPostgre) ToDomain() *domain.User {
	return &domain.User{
		ID:         p.ID,
		Email:      p.Email,
		Token:      p.Token,
		IsPaidUser: p.IsPaidUser,
		RequestQtt: p.RequestQtt,
	}
}
func (p *userPostgre) FromDomain(user *domain.User) {
	if p == nil {
		p = &userPostgre{}
	}

	p.ID = user.ID
	p.Email = user.Email
	p.Token = user.Token
	p.IsPaidUser = user.IsPaidUser
	p.RequestQtt = user.RequestQtt
}

func (p userListPostgre) ToDomain() []domain.User {
	users := make([]domain.User, len(p))
	for k, usr := range p {
		user := usr.ToDomain()
		users[k] = *user
	}

	return users
}

type userPostgreRepo struct {
	db *sql.DB
}

func NewUserPostgreRepo(db *sql.DB) ports.UserRepository {
	return &userPostgreRepo{
		db: db,
	}
}

func (p *userPostgreRepo) Get(id string) (*domain.User, error) {
	var user userPostgre = userPostgre{}
	sqsS := fmt.Sprintf("SELECT id, email, token, ispaiduser, requestsqtt FROM users WHERE id = '%s'", id)

	result := p.db.QueryRow(sqsS)
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Scan(&user.ID, &user.Email, &user.Token, &user.IsPaidUser, &user.RequestQtt); err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (p *userPostgreRepo) List() ([]domain.User, error) {
	var users userListPostgre
	sqsS := "SELECT id, email, token, ispaiduser, requestsqtt FROM users"

	result, err := p.db.Query(sqsS)
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		user := userPostgre{}

		if err := result.Scan(&user.ID, &user.Email, &user.Token, &user.IsPaidUser, &user.RequestQtt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users.ToDomain(), nil
}

func (p *userPostgreRepo) Create(newUser *domain.User) (*domain.User, error) {
	sqlS := "INSERT INTO users (id, email, token, ispaiduser, requestsqtt) VALUES (?, ?, ?, ?, ?)"

	_, err := p.db.Exec(sqlS, newUser.ID, newUser.Email, newUser.Token, newUser.IsPaidUser, newUser.RequestQtt)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}