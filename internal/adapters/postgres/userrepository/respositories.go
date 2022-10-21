package userrepository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/brenos/qap/internal/core/domain"
	ports "github.com/brenos/qap/internal/core/ports/user"
)

type userPostgre struct {
	ID         string
	Email      string
	IsPaidUser bool
	RequestQtt int32
}

type userListPostgre []userPostgre

func (p *userPostgre) ToDomain() *domain.User {
	return &domain.User{
		ID:         p.ID,
		Email:      p.Email,
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

func (p *userPostgreRepo) Create(newUser *domain.User) (*domain.User, error) {
	sqlS := "INSERT INTO users (id, email, ispaiduser, requestsqtt) VALUES ($1, $2, $3, $4)"

	_, err := p.db.Exec(sqlS, newUser.ID, newUser.Email, newUser.IsPaidUser, newUser.RequestQtt)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (p *userPostgreRepo) GetById(id string) (*domain.User, error) {
	var user userPostgre = userPostgre{}
	sqsS := fmt.Sprintf("SELECT id, email, ispaiduser, requestsqtt FROM users WHERE id = '%s'", id)

	result := p.db.QueryRow(sqsS)
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Scan(&user.ID, &user.Email, &user.IsPaidUser, &user.RequestQtt); err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (p *userPostgreRepo) GetByEmail(email string) (*domain.User, error) {
	var user userPostgre = userPostgre{}
	sqsS := fmt.Sprintf("SELECT id, email, ispaiduser, requestsqtt FROM users WHERE email = '%s'", email)

	result := p.db.QueryRow(sqsS)
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Scan(&user.ID, &user.Email, &user.IsPaidUser, &user.RequestQtt); err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (p *userPostgreRepo) UpdateRequestCount(id string) error {
	sqsS := "UPDATE public.users SET requestsqtt=(select requestsqtt+1 from users where id=$1) WHERE id=$2"

	_, err := p.db.Exec(sqsS, id, id)

	return err
}

func (p *userPostgreRepo) Delete(id string) (int64, error) {
	stmt := "DELETE FROM users WHERE id=$1"
	result, err := p.db.Exec(stmt, id)

	rowsDeleted, errResult := result.RowsAffected()

	if errResult != nil {
		log.Panicf("Error on get rows deleted - %s", errResult.Error())
	}
	return rowsDeleted, err
}
