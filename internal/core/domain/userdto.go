package domain

import (
	"encoding/json"
	"io"
)

// CreateUserRequest is an representation request body to create a new User
type CreateUserRequest struct {
	Email      string `json:"email"`
	IsPaidUser bool   `json:"isPaidUser"`
}

// FromJSONCreateUserRequest converts json body request to a CreateProductRequest struct
func FromJSONCreateUserRequest(body io.Reader) (*CreateUserRequest, error) {
	createUserRequest := CreateUserRequest{}
	if err := json.NewDecoder(body).Decode(&createUserRequest); err != nil {
		return nil, err
	}

	return &createUserRequest, nil
}
