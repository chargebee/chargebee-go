package chargebee

import (
	"fmt"
	"net/url"
)

type CardService struct {
	transport *Transport
}

func (s *CardService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/cards/%v", url.PathEscape(id)), nil)
}

func (s *CardService) UpdateCardForCustomer(id string, req *UpdateCardForCustomerRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/credit_card", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CardService) SwitchGatewayForCustomer(id string, req *SwitchGatewayForCustomerRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/switch_gateway", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CardService) CopyCardForCustomer(id string, req *CopyCardForCustomerRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/copy_card", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CardService) DeleteCardForCustomer(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/delete_card", url.PathEscape(id)), nil).SetIdempotency(true)
}
