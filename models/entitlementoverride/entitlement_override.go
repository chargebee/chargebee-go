package entitlementoverride

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	entitlementOverrideEnum "github.com/chargebee/chargebee-go/v3/models/entitlementoverride/enum"
)

type EntitlementOverride struct {
	Id             string                                 `json:"id"`
	EntityId       string                                 `json:"entity_id"`
	EntityType     string                                 `json:"entity_type"`
	FeatureId      string                                 `json:"feature_id"`
	FeatureName    string                                 `json:"feature_name"`
	Value          string                                 `json:"value"`
	Name           string                                 `json:"name"`
	ExpiresAt      int64                                  `json:"expires_at"`
	EffectiveFrom  int64                                  `json:"effective_from"`
	ScheduleStatus entitlementOverrideEnum.ScheduleStatus `json:"schedule_status"`
	Object         string                                 `json:"object"`
}
type AddEntitlementOverrideForSubscriptionRequestParams struct {
	Action               enum.Action                                                       `json:"action,omitempty"`
	EntitlementOverrides []*AddEntitlementOverrideForSubscriptionEntitlementOverrideParams `json:"entitlement_overrides,omitempty"`
}
type AddEntitlementOverrideForSubscriptionEntitlementOverrideParams struct {
	FeatureId     string `json:"feature_id"`
	Value         string `json:"value,omitempty"`
	ExpiresAt     *int64 `json:"expires_at,omitempty"`
	EffectiveFrom *int64 `json:"effective_from,omitempty"`
}
type ListEntitlementOverrideForSubscriptionRequestParams struct {
	Limit                     *int32 `json:"limit,omitempty"`
	Offset                    string `json:"offset,omitempty"`
	Embed                     string `json:"embed,omitempty"`
	IncludeDrafts             *bool  `json:"include_drafts,omitempty"`
	IncludeScheduledOverrides *bool  `json:"include_scheduled_overrides,omitempty"`
}
