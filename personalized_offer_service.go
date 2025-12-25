package chargebee

import (
	"fmt"
)

type PersonalizedOfferService struct {
	config *ClientConfig
}

func (s *PersonalizedOfferService) PersonalizedOffers(req *PersonalizedOfferPersonalizedOffersRequest) (*PersonalizedOfferPersonalizedOffersResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/personalized_offers")
	req.isJsonRequest = true
	return send[*PersonalizedOfferPersonalizedOffersResponse](req, s.config)
}
