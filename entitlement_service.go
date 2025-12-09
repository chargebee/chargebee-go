package chargebee

import (
	"fmt"
)

type EntitlementService struct {
	transport *Transport
}

func (s *EntitlementService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/entitlements"), req)
}

func (s *EntitlementService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/entitlements"), req).SetIdempotency(true)
}
