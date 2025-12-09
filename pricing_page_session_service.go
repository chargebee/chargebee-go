package chargebee

import (
	"fmt"
)

type PricingPageSessionService struct {
	transport *Transport
}

func (s *PricingPageSessionService) CreateForNewSubscription(req *CreateForNewSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/pricing_page_sessions/create_for_new_subscription"), req).SetIdempotency(true)
}

func (s *PricingPageSessionService) CreateForExistingSubscription(req *CreateForExistingSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/pricing_page_sessions/create_for_existing_subscription"), req).SetIdempotency(true)
}
