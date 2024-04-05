package subscriptionentitlement

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/subscriptionentitlement"
)

func SubscriptionEntitlementsForSubscription(id string, params *subscriptionentitlement.SubscriptionEntitlementsForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/subscription_entitlements", chargebee.IDEscape(id)), params)
}
func SetSubscriptionEntitlementAvailability(id string, params *subscriptionentitlement.SetSubscriptionEntitlementAvailabilityRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/subscription_entitlements/set_availability", chargebee.IDEscape(id)), params)
}
