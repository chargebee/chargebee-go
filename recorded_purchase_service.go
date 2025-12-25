package chargebee

import (
	"fmt"
	"net/url"
)

type RecordedPurchaseService struct {
	config *ClientConfig
}

func (s *RecordedPurchaseService) Create(req *RecordedPurchaseCreateRequest) (*RecordedPurchaseCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/recorded_purchases")
	req.isIdempotent = true
	return send[*RecordedPurchaseCreateResponse](req, s.config)
}

func (s *RecordedPurchaseService) Retrieve(id string) (*RecordedPurchaseRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/recorded_purchases/%v", url.PathEscape(id))
	return send[*RecordedPurchaseRetrieveResponse](req, s.config)
}
