package domain

import "fmt"

type User struct {
	ID         string `json:"id"`
	Email      string `json:"email,omitempty"`
	Token      string `json:"token,omitempty"`
	IsPaidUser bool   `json:"isPaidUser,omitempty"`
	RequestQtt int32  `json:"requestQtt,omitempty"`
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
