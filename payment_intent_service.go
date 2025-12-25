package chargebee

import (
	"fmt"
	"net/url"
)

type PaymentIntentService struct {
	config *ClientConfig
}

func (s *PaymentIntentService) Create(req *PaymentIntentCreateRequest) (*PaymentIntentCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_intents")
	req.isIdempotent = true
	return send[*PaymentIntentCreateResponse](req, s.config)
}

func (s *PaymentIntentService) Update(id string, req *PaymentIntentUpdateRequest) (*PaymentIntentUpdateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_intents/%v", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PaymentIntentUpdateResponse](req, s.config)
}

func (s *PaymentIntentService) Retrieve(id string) (*PaymentIntentRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/payment_intents/%v", url.PathEscape(id))
	return send[*PaymentIntentRetrieveResponse](req, s.config)
}
