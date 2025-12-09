package chargebee

import (
	"fmt"
	"net/url"
)

type PaymentScheduleSchemeService struct {
	transport *Transport
}

func (s *PaymentScheduleSchemeService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_schedule_schemes"), req).SetIdempotency(true)
}

func (s *PaymentScheduleSchemeService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/payment_schedule_schemes/%v", url.PathEscape(id)), nil)
}

func (s *PaymentScheduleSchemeService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_schedule_schemes/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
