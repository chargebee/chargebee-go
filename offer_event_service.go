package chargebee

import (
	"fmt"
)

type OfferEventService struct {
	transport *Transport
}

func (s *OfferEventService) OfferEvents(req *OfferEventsRequest) Request {
	return s.transport.SendJsonRequest("POST", fmt.Sprintf("/offer_events"), req).SetSubDomain("grow")
}
