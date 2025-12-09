package chargebee

type ScheduleStatus string

const (
	ScheduleStatusActivated ScheduleStatus = "activated"
	ScheduleStatusScheduled ScheduleStatus = "scheduled"
	ScheduleStatusFailed    ScheduleStatus = "failed"
)

type EntitlementOverride struct {
	Id             string         `json:"id"`
	EntityId       string         `json:"entity_id"`
	EntityType     string         `json:"entity_type"`
	FeatureId      string         `json:"feature_id"`
	FeatureName    string         `json:"feature_name"`
	Value          string         `json:"value"`
	Name           string         `json:"name"`
	ExpiresAt      int64          `json:"expires_at"`
	EffectiveFrom  int64          `json:"effective_from"`
	ScheduleStatus ScheduleStatus `json:"schedule_status"`
	Object         string         `json:"object"`
}
type AddEntitlementOverrideForSubscriptionRequest struct {
	Action               enum.Action                                                 `json:"action,omitempty"`
	EntitlementOverrides []*AddEntitlementOverrideForSubscriptionEntitlementOverride `json:"entitlement_overrides,omitempty"`
}
type AddEntitlementOverrideForSubscriptionEntitlementOverride struct {
	FeatureId     string `json:"feature_id"`
	Value         string `json:"value,omitempty"`
	ExpiresAt     *int64 `json:"expires_at,omitempty"`
	EffectiveFrom *int64 `json:"effective_from,omitempty"`
}
type ListEntitlementOverrideForSubscriptionRequest struct {
	Limit                     *int32 `json:"limit,omitempty"`
	Offset                    string `json:"offset,omitempty"`
	Embed                     string `json:"embed,omitempty"`
	IncludeDrafts             *bool  `json:"include_drafts,omitempty"`
	IncludeScheduledOverrides *bool  `json:"include_scheduled_overrides,omitempty"`
}

type AddEntitlementOverrideForSubscriptionEntitlementOverrideResponse struct {
	EntitlementOverride *EntitlementOverride `json:"entitlement_override,omitempty"`
}

type AddEntitlementOverrideForSubscriptionResponse struct {
	List []*AddEntitlementOverrideForSubscriptionEntitlementOverrideResponse `json:"list,omitempty"`
}

type ListEntitlementOverrideForSubscriptionEntitlementOverrideResponse struct {
	EntitlementOverride *EntitlementOverride `json:"entitlement_override,omitempty"`
}

type ListEntitlementOverrideForSubscriptionResponse struct {
	List       []*ListEntitlementOverrideForSubscriptionEntitlementOverrideResponse `json:"list,omitempty"`
	NextOffset string                                                               `json:"next_offset,omitempty"`
}
