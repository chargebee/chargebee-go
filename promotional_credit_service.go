package chargebee

import (
	"fmt"
	"net/url"
)

type PromotionalCreditService struct {
	transport *Transport
}

func (s *PromotionalCreditService) Add(req *AddRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/promotional_credits/add"), req).SetIdempotency(true)
}

func (s *PromotionalCreditService) Deduct(req *DeductRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/promotional_credits/deduct"), req).SetIdempotency(true)
}

func (s *PromotionalCreditService) Set(req *SetRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/promotional_credits/set"), req).SetIdempotency(true)
}

func (s *PromotionalCreditService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/promotional_credits"), req)
}

func (s *PromotionalCreditService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/promotional_credits/%v", url.PathEscape(id)), nil)
}
