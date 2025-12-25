package chargebee

import (
	"fmt"
	"net/url"
)

type PortalSessionService struct {
	config *ClientConfig
}

func (s *PortalSessionService) Create(req *PortalSessionCreateRequest) (*PortalSessionCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/portal_sessions")
	req.isIdempotent = true
	return send[*PortalSessionCreateResponse](req, s.config)
}

func (s *PortalSessionService) Retrieve(id string) (*PortalSessionRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/portal_sessions/%v", url.PathEscape(id))
	return send[*PortalSessionRetrieveResponse](req, s.config)
}

func (s *PortalSessionService) Logout(id string) (*PortalSessionLogoutResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/portal_sessions/%v/logout", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PortalSessionLogoutResponse](req, s.config)
}

func (s *PortalSessionService) Activate(id string, req *PortalSessionActivateRequest) (*PortalSessionActivateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/portal_sessions/%v/activate", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PortalSessionActivateResponse](req, s.config)
}
