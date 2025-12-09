package chargebee

import (
	"fmt"
	"net/url"
)

type RampService struct {
	transport *Transport
}

func (s *RampService) CreateForSubscription(id string, req *CreateForSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/create_ramp", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *RampService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/ramps/%v/update", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *RampService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/ramps/%v", url.PathEscape(id)), nil)
}

func (s *RampService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/ramps/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *RampService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/ramps"), req)
}
