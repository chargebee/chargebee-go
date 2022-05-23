package subscriptionentitlement

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/subscriptionentitlement"
	"net/url"
)

func SubscriptionEntitlementsForSubscription(id string, params *subscriptionentitlement.SubscriptionEntitlementsForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/subscription_entitlements", url.PathEscape(id)), params)
}
func SetSubscriptionEntitlementAvailability(id string, params *subscriptionentitlement.SetSubscriptionEntitlementAvailabilityRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/subscription_entitlements/set_availability", url.PathEscape(id)), params)
}
