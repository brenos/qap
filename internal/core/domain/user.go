package domain

import (
	"fmt"
)

type User struct {
	ID         string `json:"id"`
	Email      string `json:"email,omitempty"`
	IsPaidUser bool   `json:"isPaidUser,omitempty"`
	RequestQtt int32  `json:"requestQtt,omitempty"`
}

type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func NewUser(id string, email string, isPaidUser bool, requestQtt int32) *User {
	return &User{
		ID:         id,
		Email:      email,
		IsPaidUser: isPaidUser,
		RequestQtt: requestQtt,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("%s - %s - Paid: %t", u.ID, u.Email, u.IsPaidUser)
}
