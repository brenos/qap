package domain

import (
	"encoding/json"
	"io"
)

type DealershipRequest struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func FromJSONDealershipRequest(body io.Reader) (*DealershipRequest, error) {
	dealershipRequest := DealershipRequest{}
	if err := json.NewDecoder(body).Decode(&dealershipRequest); err != nil {
		return nil, err
	}

	return &dealershipRequest, nil
}
