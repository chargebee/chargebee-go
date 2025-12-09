package chargebee

import (
	"fmt"
	"net/url"
)

type RecordedPurchaseService struct {
	transport *Transport
}

func (s *RecordedPurchaseService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/recorded_purchases"), req).SetIdempotency(true)
}

func (s *RecordedPurchaseService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/recorded_purchases/%v", url.PathEscape(id)), nil)
}
