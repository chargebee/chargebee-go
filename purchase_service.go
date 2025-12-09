package chargebee

import (
	"fmt"
)

type PurchaseService struct {
	transport *Transport
}

func (s *PurchaseService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/purchases"), req).SetIdempotency(true)
}

func (s *PurchaseService) Estimate(req *EstimateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/purchases/estimate"), req)
}
