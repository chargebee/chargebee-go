package chargebee

import (
	"fmt"
	"net/url"
)

type SubscriptionEntitlementService struct {
	transport *Transport
}

func (s *SubscriptionEntitlementService) SubscriptionEntitlementsForSubscription(id string, req *SubscriptionEntitlementsForSubscriptionRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/subscriptions/%v/subscription_entitlements", url.PathEscape(id)), req)
}

func (s *SubscriptionEntitlementService) SetSubscriptionEntitlementAvailability(id string, req *SetSubscriptionEntitlementAvailabilityRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/subscription_entitlements/set_availability", url.PathEscape(id)), req).SetIdempotency(true)
}
