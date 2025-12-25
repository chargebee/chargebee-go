package chargebee

import (
	"fmt"
)

type AddressService struct {
	config *ClientConfig
}

func (s *AddressService) Retrieve(req *AddressRetrieveRequest) (*AddressRetrieveResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/addresses")
	return send[*AddressRetrieveResponse](req, s.config)
}

func (s *AddressService) Update(req *AddressUpdateRequest) (*AddressUpdateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/addresses")
	req.isIdempotent = true
	return send[*AddressUpdateResponse](req, s.config)
}
