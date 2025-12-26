package chargebee

type EntitlementOverrideScheduleStatus string

const (
	EntitlementOverrideScheduleStatusActivated EntitlementOverrideScheduleStatus = "activated"
	EntitlementOverrideScheduleStatusScheduled EntitlementOverrideScheduleStatus = "scheduled"
	EntitlementOverrideScheduleStatusFailed    EntitlementOverrideScheduleStatus = "failed"
)

// just struct
type EntitlementOverride struct {
	Id            string `json:"id"`
	EntityId      string `json:"entity_id"`
	EntityType    string `json:"entity_type"`
	FeatureId     string `json:"feature_id"`
	FeatureName   string `json:"feature_name"`
	Value         string `json:"value"`
	Name          string `json:"name"`
	ExpiresAt     int64  `json:"expires_at"`
	EffectiveFrom int64  `json:"effective_from"`
	Object        string `json:"object"`
}

// sub resources
// operations
// input params
type EntitlementOverrideAddEntitlementOverrideForSubscriptionRequest struct {
	Action               Action                                                                         `json:"action,omitempty"`
	EntitlementOverrides []*EntitlementOverrideAddEntitlementOverrideForSubscriptionEntitlementOverride `json:"entitlement_overrides,omitempty"`
	apiRequest           `json:"-" form:"-"`
}

func (r *EntitlementOverrideAddEntitlementOverrideForSubscriptionRequest) payload() any { return r }

// input sub resource params multi
type EntitlementOverrideAddEntitlementOverrideForSubscriptionEntitlementOverride struct {
	FeatureId     string `json:"feature_id"`
	Value         string `json:"value,omitempty"`
	ExpiresAt     *int64 `json:"expires_at,omitempty"`
	EffectiveFrom *int64 `json:"effective_from,omitempty"`
}

type EntitlementOverrideListEntitlementOverrideForSubscriptionRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *EntitlementOverrideListEntitlementOverrideForSubscriptionRequest) payload() any { return r }

// operation sub response
type EntitlementOverrideAddEntitlementOverrideForSubscriptionEntitlementOverrideResponse struct {
	EntitlementOverride *EntitlementOverride `json:"entitlement_override,omitempty"`
}

// operation response
type EntitlementOverrideAddEntitlementOverrideForSubscriptionResponse struct {
	List []*EntitlementOverrideAddEntitlementOverrideForSubscriptionEntitlementOverrideResponse `json:"list,omitempty"`
	apiResponse
}

// operation sub response
type EntitlementOverrideListEntitlementOverrideForSubscriptionEntitlementOverrideResponse struct {
	EntitlementOverride *EntitlementOverride `json:"entitlement_override,omitempty"`
}

// operation response
type EntitlementOverrideListEntitlementOverrideForSubscriptionResponse struct {
	List       []*EntitlementOverrideListEntitlementOverrideForSubscriptionEntitlementOverrideResponse `json:"list,omitempty"`
	NextOffset string                                                                                  `json:"next_offset,omitempty"`
	apiResponse
}
