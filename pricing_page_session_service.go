package chargebee

import (
	"fmt"
)

type PricingPageSessionService struct {
	config *ClientConfig
}

func (s *PricingPageSessionService) CreateForNewSubscription(req *PricingPageSessionCreateForNewSubscriptionRequest) (*PricingPageSessionCreateForNewSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/pricing_page_sessions/create_for_new_subscription")
	req.isIdempotent = true
	return send[*PricingPageSessionCreateForNewSubscriptionResponse](req, s.config)
}

func (s *PricingPageSessionService) CreateForExistingSubscription(req *PricingPageSessionCreateForExistingSubscriptionRequest) (*PricingPageSessionCreateForExistingSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/pricing_page_sessions/create_for_existing_subscription")
	req.isIdempotent = true
	return send[*PricingPageSessionCreateForExistingSubscriptionResponse](req, s.config)
}
