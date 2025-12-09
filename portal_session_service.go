package chargebee

import (
	"fmt"
	"net/url"
)

type PortalSessionService struct {
	transport *Transport
}

func (s *PortalSessionService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/portal_sessions"), req).SetIdempotency(true)
}

func (s *PortalSessionService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/portal_sessions/%v", url.PathEscape(id)), nil)
}

func (s *PortalSessionService) Logout(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/portal_sessions/%v/logout", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *PortalSessionService) Activate(id string, req *ActivateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/portal_sessions/%v/activate", url.PathEscape(id)), req).SetIdempotency(true)
}
