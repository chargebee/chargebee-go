package chargebee

import (
	"fmt"
	"net/url"
)

type CurrencyService struct {
	transport *Transport
}

func (s *CurrencyService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/currencies/list"), req)
}

func (s *CurrencyService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/currencies/%v", url.PathEscape(id)), nil)
}

func (s *CurrencyService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/currencies"), req).SetIdempotency(true)
}

func (s *CurrencyService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/currencies/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CurrencyService) AddSchedule(id string, req *AddScheduleRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/currencies/%v/add_schedule", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CurrencyService) RemoveSchedule(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/currencies/%v/remove_schedule", url.PathEscape(id)), nil).SetIdempotency(true)
}
