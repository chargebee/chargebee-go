package entitlementoverride

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/entitlementoverride"
)

func AddEntitlementOverrideForSubscription(id string, params *entitlementoverride.AddEntitlementOverrideForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/entitlement_overrides", chargebee.IDEscape(id)), params)
}
func ListEntitlementOverrideForSubscription(id string, params *entitlementoverride.ListEntitlementOverrideForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/entitlement_overrides", chargebee.IDEscape(id)), params)
}
