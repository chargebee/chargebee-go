package chargebee

import (
	"fmt"
)

type EntitlementService struct {
	config *ClientConfig
}

func (s *EntitlementService) List(req *EntitlementListRequest) (*EntitlementListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/entitlements")
	req.isListRequest = true
	return send[*EntitlementListResponse](req, s.config)
}

func (s *EntitlementService) Create(req *EntitlementCreateRequest) (*EntitlementCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/entitlements")
	req.isIdempotent = true
	return send[*EntitlementCreateResponse](req, s.config)
}
