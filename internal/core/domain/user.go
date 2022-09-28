package domain

import (
	"encoding/json"
	"fmt"
	"io"
)

type User struct {
	ID         string `json:"id"`
	Email      string `json:"email,omitempty"`
	Token      string `json:"token,omitempty"`
	IsPaidUser bool   `json:"isPaidUser,omitempty"`
	RequestQtt int32  `json:"requestQtt,omitempty"`
}

type CreateUserRequest struct {
	Email      string `json:"email"`
	IsPaidUser bool   `json:"isPaidUser"`
}

func NewUser(id string, email string, token string, isPaidUser bool, requestQtt int32) *User {
	return &User{
		ID:         id,
		Email:      email,
		Token:      token,
		IsPaidUser: isPaidUser,
		RequestQtt: requestQtt,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("%s - %s - Paid: %t", u.ID, u.Email, u.IsPaidUser)
}

func FromJSONCreateUserRequest(body io.Reader) (*CreateUserRequest, error) {
	createUserRequest := CreateUserRequest{}
	if err := json.NewDecoder(body).Decode(&createUserRequest); err != nil {
		return nil, err
	}

	return &createUserRequest, nil
}
