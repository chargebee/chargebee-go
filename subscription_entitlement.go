package chargebee

type SubscriptionEntitlementScheduleStatus string

const (
	SubscriptionEntitlementScheduleStatusActivated SubscriptionEntitlementScheduleStatus = "activated"
	SubscriptionEntitlementScheduleStatusScheduled SubscriptionEntitlementScheduleStatus = "scheduled"
	SubscriptionEntitlementScheduleStatusFailed    SubscriptionEntitlementScheduleStatus = "failed"
)

// just struct
type SubscriptionEntitlement struct {
	SubscriptionId string                            `json:"subscription_id"`
	FeatureId      string                            `json:"feature_id"`
	FeatureName    string                            `json:"feature_name"`
	FeatureUnit    string                            `json:"feature_unit"`
	FeatureType    string                            `json:"feature_type"`
	Value          string                            `json:"value"`
	Name           string                            `json:"name"`
	IsOverridden   bool                              `json:"is_overridden"`
	IsEnabled      bool                              `json:"is_enabled"`
	ExpiresAt      int64                             `json:"expires_at"`
	Components     *SubscriptionEntitlementComponent `json:"components"`
	Object         string                            `json:"object"`
}

// sub resources
type SubscriptionEntitlementComponent struct {
	EntitlementOverrides *EntitlementOverride `json:"entitlement_overrides"`
	Object               string               `json:"object"`
}

// operations
// input params
type SubscriptionEntitlementSubscriptionEntitlementsForSubscriptionRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *SubscriptionEntitlementSubscriptionEntitlementsForSubscriptionRequest) payload() any {
	return r
}

type SubscriptionEntitlementSetSubscriptionEntitlementAvailabilityRequest struct {
	IsEnabled                *bool                                                                                   `json:"is_enabled"`
	SubscriptionEntitlements []*SubscriptionEntitlementSetSubscriptionEntitlementAvailabilitySubscriptionEntitlement `json:"subscription_entitlements,omitempty"`
	apiRequest               `json:"-" form:"-"`
}

func (r *SubscriptionEntitlementSetSubscriptionEntitlementAvailabilityRequest) payload() any {
	return r
}

// input sub resource params multi
type SubscriptionEntitlementSetSubscriptionEntitlementAvailabilitySubscriptionEntitlement struct {
	FeatureId string `json:"feature_id"`
}

// operation sub response
type SubscriptionEntitlementSubscriptionEntitlementsForSubscriptionSubscriptionEntitlementResponse struct {
	SubscriptionEntitlement *SubscriptionEntitlement `json:"subscription_entitlement,omitempty"`
}

// operation response
type SubscriptionEntitlementSubscriptionEntitlementsForSubscriptionResponse struct {
	List       []*SubscriptionEntitlementSubscriptionEntitlementsForSubscriptionSubscriptionEntitlementResponse `json:"list,omitempty"`
	NextOffset string                                                                                           `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type SubscriptionEntitlementSetSubscriptionEntitlementAvailabilitySubscriptionEntitlementResponse struct {
	SubscriptionEntitlement *SubscriptionEntitlement `json:"subscription_entitlement,omitempty"`
}

// operation response
type SubscriptionEntitlementSetSubscriptionEntitlementAvailabilityResponse struct {
	List []*SubscriptionEntitlementSetSubscriptionEntitlementAvailabilitySubscriptionEntitlementResponse `json:"list,omitempty"`
	apiResponse
}
