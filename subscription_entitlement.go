package chargebee

type ScheduleStatus string

const (
	ScheduleStatusActivated ScheduleStatus = "activated"
	ScheduleStatusScheduled ScheduleStatus = "scheduled"
	ScheduleStatusFailed    ScheduleStatus = "failed"
)

type SubscriptionEntitlement struct {
	SubscriptionId string         `json:"subscription_id"`
	FeatureId      string         `json:"feature_id"`
	FeatureName    string         `json:"feature_name"`
	FeatureUnit    string         `json:"feature_unit"`
	FeatureType    string         `json:"feature_type"`
	Value          string         `json:"value"`
	Name           string         `json:"name"`
	IsOverridden   bool           `json:"is_overridden"`
	IsEnabled      bool           `json:"is_enabled"`
	EffectiveFrom  int64          `json:"effective_from"`
	ScheduleStatus ScheduleStatus `json:"schedule_status"`
	ExpiresAt      int64          `json:"expires_at"`
	Components     *Component     `json:"components"`
	Object         string         `json:"object"`
}
type Component struct {
	EntitlementOverrides *entitlementoverride.EntitlementOverride `json:"entitlement_overrides"`
	Object               string                                   `json:"object"`
}
type SubscriptionEntitlementsForSubscriptionRequest struct {
	Limit                     *int32 `json:"limit,omitempty"`
	Offset                    string `json:"offset,omitempty"`
	IncludeDrafts             *bool  `json:"include_drafts,omitempty"`
	Embed                     string `json:"embed,omitempty"`
	IncludeScheduledOverrides *bool  `json:"include_scheduled_overrides,omitempty"`
}
type SetSubscriptionEntitlementAvailabilityRequest struct {
	IsEnabled                *bool                                                            `json:"is_enabled"`
	SubscriptionEntitlements []*SetSubscriptionEntitlementAvailabilitySubscriptionEntitlement `json:"subscription_entitlements,omitempty"`
}
type SetSubscriptionEntitlementAvailabilitySubscriptionEntitlement struct {
	FeatureId string `json:"feature_id"`
}

type SubscriptionEntitlementsForSubscriptionSubscriptionEntitlementResponse struct {
	SubscriptionEntitlement *SubscriptionEntitlement `json:"subscription_entitlement,omitempty"`
}

type SubscriptionEntitlementsForSubscriptionResponse struct {
	List       []*SubscriptionEntitlementsForSubscriptionSubscriptionEntitlementResponse `json:"list,omitempty"`
	NextOffset string                                                                    `json:"next_offset,omitempty"`
}

type SetSubscriptionEntitlementAvailabilitySubscriptionEntitlementResponse struct {
	SubscriptionEntitlement *SubscriptionEntitlement `json:"subscription_entitlement,omitempty"`
}

type SetSubscriptionEntitlementAvailabilityResponse struct {
	List []*SetSubscriptionEntitlementAvailabilitySubscriptionEntitlementResponse `json:"list,omitempty"`
}
