package quoteentitlement

import (
	quoteEntitlementEnum "github.com/chargebee/chargebee-go/v3/models/quoteentitlement/enum"
)

type QuoteEntitlement struct {
	EntityId   string                          `json:"entity_id"`
	EntityType quoteEntitlementEnum.EntityType `json:"entity_type"`
	FeatureId  string                          `json:"feature_id"`
	Value      string                          `json:"value"`
	IsEnabled  bool                            `json:"is_enabled"`
	StartDate  int64                           `json:"start_date"`
	EndDate    int64                           `json:"end_date"`
	CreatedAt  int64                           `json:"created_at"`
	ModifiedAt int64                           `json:"modified_at"`
	Object     string                          `json:"object"`
}
