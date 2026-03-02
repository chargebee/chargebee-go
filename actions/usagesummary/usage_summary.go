package usagesummary

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/usagesummary"
	"net/url"
)

// Deprecated: This function is deprecated.
func RetrieveUsageSummaryForSubscription(id string, params *usagesummary.RetrieveUsageSummaryForSubscriptionRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/usage_summary", url.PathEscape(id)), params)
}
