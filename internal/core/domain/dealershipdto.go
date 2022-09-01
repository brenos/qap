package domain

import (
	"encoding/json"
	"io"
)

type CreateDealershipRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func FromJSONCreateDealershipRequest(body io.Reader) (*CreateDealershipRequest, error) {
	createDealershipRequest := CreateDealershipRequest{}
	if err := json.NewDecoder(body).Decode(&createDealershipRequest); err != nil {
		return nil, err
	}

	return &createDealershipRequest, nil
}
