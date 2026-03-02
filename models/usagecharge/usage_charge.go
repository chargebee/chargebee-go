package usagecharge

import (
	"github.com/chargebee/chargebee-go/v3/filter"
)

type UsageCharge struct {
	SubscriptionId     string `json:"subscription_id"`
	FeatureId          string `json:"feature_id"`
	IncludedUsage      string `json:"included_usage"`
	TotalUsage         string `json:"total_usage"`
	OnDemandUsage      string `json:"on_demand_usage"`
	MeteredItemPriceId string `json:"metered_item_price_id"`
	Amount             string `json:"amount"`
	CurrencyCode       string `json:"currency_code"`
	UsageFrom          int64  `json:"usage_from"`
	UsageTo            int64  `json:"usage_to"`
	Object             string `json:"object"`
}
type RetrieveUsageChargesForSubscriptionRequestParams struct {
	Limit     *int32               `json:"limit,omitempty"`
	Offset    string               `json:"offset,omitempty"`
	FeatureId *filter.StringFilter `json:"feature_id,omitempty"`
}
