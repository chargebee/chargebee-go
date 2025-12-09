package chargebee

import (
	"fmt"
	"net/url"
)

type PlanService struct {
	transport *Transport
}

func (s *PlanService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/plans"), req).SetIdempotency(true)
}

func (s *PlanService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/plans/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *PlanService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/plans"), req)
}

func (s *PlanService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/plans/%v", url.PathEscape(id)), nil)
}

func (s *PlanService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/plans/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *PlanService) Copy(req *CopyRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/plans/copy"), req).SetIdempotency(true)
}

func (s *PlanService) Unarchive(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/plans/%v/unarchive", url.PathEscape(id)), nil).SetIdempotency(true)
}
