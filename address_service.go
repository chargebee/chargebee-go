package chargebee

import (
	"fmt"
)

type AddressService struct {
	transport *Transport
}

func (s *AddressService) Retrieve(req *RetrieveRequest) Request {
	return s.transport.Send("GET", fmt.Sprintf("/addresses"), req)
}

func (s *AddressService) Update(req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/addresses"), req).SetIdempotency(true)
}
