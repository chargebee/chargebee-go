package chargebee

import (
	"fmt"
)

type PurchaseService struct {
	config *ClientConfig
}

func (s *PurchaseService) Create(req *PurchaseCreateRequest) (*PurchaseCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/purchases")
	req.isIdempotent = true
	return send[*PurchaseCreateResponse](req, s.config)
}

func (s *PurchaseService) Estimate(req *PurchaseEstimateRequest) (*PurchaseEstimateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/purchases/estimate")
	return send[*PurchaseEstimateResponse](req, s.config)
}
