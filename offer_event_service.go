package chargebee

import (
	"fmt"
)

type OfferEventService struct {
	config *ClientConfig
}

func (s *OfferEventService) OfferEvents(req *OfferEventOfferEventsRequest) (*OfferEventOfferEventsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/offer_events")
	req.isJsonRequest = true
	return send[*OfferEventOfferEventsResponse](req, s.config)
}
