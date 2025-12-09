package chargebee

import (
	"fmt"
	"net/url"
)

type PaymentIntentService struct {
	transport *Transport
}

func (s *PaymentIntentService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_intents"), req).SetIdempotency(true)
}

func (s *PaymentIntentService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_intents/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *PaymentIntentService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/payment_intents/%v", url.PathEscape(id)), nil)
}
