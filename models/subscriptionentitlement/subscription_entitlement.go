package subscriptionentitlement

import (
	"github.com/chargebee/chargebee-go/models/entitlementoverride"
)

type SubscriptionEntitlement struct {
	SubscriptionId string     `json:"subscription_id"`
	FeatureId      string     `json:"feature_id"`
	FeatureName    string     `json:"feature_name"`
	FeatureUnit    string     `json:"feature_unit"`
	Value          string     `json:"value"`
	Name           string     `json:"name"`
	IsOverridden   bool       `json:"is_overridden"`
	IsEnabled      bool       `json:"is_enabled"`
	ExpiresAt      int64      `json:"expires_at"`
	Components     *Component `json:"components"`
	Object         string     `json:"object"`
}
type Component struct {
	EntitlementOverrides *entitlementoverride.EntitlementOverride `json:"entitlement_overrides"`
	Object               string                                   `json:"object"`
}
type SubscriptionEntitlementsForSubscriptionRequestParams struct {
	Limit         *int32 `json:"limit,omitempty"`
	Offset        string `json:"offset,omitempty"`
	IncludeDrafts *bool  `json:"include_drafts,omitempty"`
	Embed         string `json:"embed,omitempty"`
}
type SetSubscriptionEntitlementAvailabilityRequestParams struct {
	IsEnabled                *bool                                                                  `json:"is_enabled"`
	SubscriptionEntitlements []*SetSubscriptionEntitlementAvailabilitySubscriptionEntitlementParams `json:"subscription_entitlements,omitempty"`
}
type SetSubscriptionEntitlementAvailabilitySubscriptionEntitlementParams struct {
	FeatureId string `json:"feature_id"`
}
