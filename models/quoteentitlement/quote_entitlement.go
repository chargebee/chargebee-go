package quoteentitlement

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	quoteEntitlementEnum "github.com/chargebee/chargebee-go/v3/models/quoteentitlement/enum"
)

type QuoteEntitlement struct {
	EntityId   string                          `json:"entity_id"`
	EntityType quoteEntitlementEnum.EntityType `json:"entity_type"`
	//Deprecated: this field is deprecated
	ActionType   quoteEntitlementEnum.ActionType `json:"action_type"`
	FeatureId    string                          `json:"feature_id"`
	Value        string                          `json:"value"`
	IsEnabled    bool                            `json:"is_enabled"`
	StartDate    int64                           `json:"start_date"`
	EndDate      int64                           `json:"end_date"`
	CreatedAt    int64                           `json:"created_at"`
	ModifiedAt   int64                           `json:"modified_at"`
	IsOverridden bool                            `json:"is_overridden"`
	FeatureName  string                          `json:"feature_name"`
	FeatureUnit  string                          `json:"feature_unit"`
	FeatureType  string                          `json:"feature_type"`
	Name         string                          `json:"name"`
	Metered      bool                            `json:"metered"`
	Object       string                          `json:"object"`
}
type ListQuoteEntitlementsRequestParams struct {
	Limit     *int32                  `json:"limit,omitempty"`
	Offset    string                  `json:"offset,omitempty"`
	EntityId  *filter.StringFilter    `json:"entity_id,omitempty"`
	StartDate *filter.TimestampFilter `json:"start_date,omitempty"`
	EndDate   *filter.TimestampFilter `json:"end_date,omitempty"`
}
