package chargebee

import (
	"fmt"
	"net/url"
)

type FeatureService struct {
	transport *Transport
}

func (s *FeatureService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/features"), req)
}

func (s *FeatureService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/features"), req).SetIdempotency(true)
}

func (s *FeatureService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/features/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *FeatureService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/features/%v", url.PathEscape(id)), nil)
}

func (s *FeatureService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/features/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *FeatureService) Activate(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/features/%v/activate_command", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *FeatureService) Archive(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/features/%v/archive_command", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *FeatureService) Reactivate(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/features/%v/reactivate_command", url.PathEscape(id)), nil).SetIdempotency(true)
}
