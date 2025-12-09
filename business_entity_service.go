package chargebee

import (
	"fmt"
)

type BusinessEntityService struct {
	transport *Transport
}

func (s *BusinessEntityService) CreateTransfers(req *CreateTransfersRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/business_entities/transfers"), req).SetIdempotency(true)
}

func (s *BusinessEntityService) GetTransfers(req *GetTransfersRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/business_entities/transfers"), req)
}
