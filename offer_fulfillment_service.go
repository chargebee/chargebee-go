package chargebee

import (
	"fmt"
	"net/url"
)

type OfferFulfillmentService struct {
	transport *Transport
}

func (s *OfferFulfillmentService) OfferFulfillments(req *OfferFulfillmentsRequest) Request {
	return s.transport.SendJsonRequest("POST", fmt.Sprintf("/offer_fulfillments"), req).SetSubDomain("grow")
}

func (s *OfferFulfillmentService) OfferFulfillmentsGet(id string) Request {
	return s.transport.SendJsonRequest("GET", fmt.Sprintf("/offer_fulfillments/%v", url.PathEscape(id)), nil).SetSubDomain("grow")
}

func (s *OfferFulfillmentService) OfferFulfillmentsUpdate(id string, req *OfferFulfillmentsUpdateRequest) Request {
	return s.transport.SendJsonRequest("POST", fmt.Sprintf("/offer_fulfillments/%v", url.PathEscape(id)), req).SetSubDomain("grow")
}
