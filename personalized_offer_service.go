package chargebee

import (
	"fmt"
)

type PersonalizedOfferService struct {
	transport *Transport
}

func (s *PersonalizedOfferService) PersonalizedOffers(req *PersonalizedOffersRequest) Request {
	return s.transport.SendJsonRequest("POST", fmt.Sprintf("/personalized_offers"), req).SetSubDomain("grow")
}
