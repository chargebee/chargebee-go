package usagecharge

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/usagecharge"
	"net/url"
)

func RetrieveUsageChargesForSubscription(id string, params *usagecharge.RetrieveUsageChargesForSubscriptionRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/usage_charges", url.PathEscape(id)), params)
}
