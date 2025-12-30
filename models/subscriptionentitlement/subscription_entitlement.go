package subscriptionentitlement

import (
	"github.com/chargebee/chargebee-go/v3/models/entitlementoverride"
	subscriptionEntitlementEnum "github.com/chargebee/chargebee-go/v3/models/subscriptionentitlement/enum"
)

type SubscriptionEntitlement struct {
	SubscriptionId string `json:"subscription_id"`
	FeatureId      string `json:"feature_id"`
	FeatureName    string `json:"feature_name"`
	FeatureUnit    string `json:"feature_unit"`
	FeatureType    string `json:"feature_type"`
	Value          string `json:"value"`
	Name           string `json:"name"`
	IsOverridden   bool   `json:"is_overridden"`
	IsEnabled      bool   `json:"is_enabled"`
	//Deprecated: this field is deprecated
	EffectiveFrom int64 `json:"effective_from"`
	//Deprecated: this field is deprecated
	ScheduleStatus subscriptionEntitlementEnum.ScheduleStatus `json:"schedule_status"`
	ExpiresAt      int64                                      `json:"expires_at"`
	Components     *Component                                 `json:"components"`
	Object         string                                     `json:"object"`
}
type Component struct {
	EntitlementOverrides *entitlementoverride.EntitlementOverride `json:"entitlement_overrides"`
	Object               string                                   `json:"object"`
}
type SubscriptionEntitlementsForSubscriptionRequestParams struct {
	Limit                     *int32 `json:"limit,omitempty"`
	Offset                    string `json:"offset,omitempty"`
	IncludeDrafts             *bool  `json:"include_drafts,omitempty"`
	Embed                     string `json:"embed,omitempty"`
	IncludeScheduledOverrides *bool  `json:"include_scheduled_overrides,omitempty"`
}
type SetSubscriptionEntitlementAvailabilityRequestParams struct {
	IsEnabled                *bool                                                                  `json:"is_enabled"`
	SubscriptionEntitlements []*SetSubscriptionEntitlementAvailabilitySubscriptionEntitlementParams `json:"subscription_entitlements,omitempty"`
}
type SetSubscriptionEntitlementAvailabilitySubscriptionEntitlementParams struct {
	FeatureId string `json:"feature_id"`
}
