package usagesummary

import (
	"github.com/chargebee/chargebee-go/v3/enum"
)

type UsageSummary struct {
	SubscriptionId  string `json:"subscription_id"`
	FeatureId       string `json:"feature_id"`
	AggregatedValue string `json:"aggregated_value"`
	AggregatedFrom  int64  `json:"aggregated_from"`
	AggregatedTo    int64  `json:"aggregated_to"`
	Object          string `json:"object"`
}
type RetrieveUsageSummaryForSubscriptionRequestParams struct {
	Limit          *int32          `json:"limit,omitempty"`
	Offset         string          `json:"offset,omitempty"`
	FeatureId      string          `json:"feature_id"`
	WindowSize     enum.WindowSize `json:"window_size,omitempty"`
	TimeframeStart *int64          `json:"timeframe_start,omitempty"`
	TimeframeEnd   *int64          `json:"timeframe_end,omitempty"`
}
