package chargebee

import (
	"fmt"
	"net/url"
)

type EntitlementOverrideService struct {
	transport *Transport
}

func (s *EntitlementOverrideService) AddEntitlementOverrideForSubscription(id string, req *AddEntitlementOverrideForSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/entitlement_overrides", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *EntitlementOverrideService) ListEntitlementOverrideForSubscription(id string, req *ListEntitlementOverrideForSubscriptionRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/subscriptions/%v/entitlement_overrides", url.PathEscape(id)), req)
}
