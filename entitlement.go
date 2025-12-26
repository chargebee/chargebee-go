package chargebee

type EntitlementEntityType string

const (
	EntitlementEntityTypePlan       EntitlementEntityType = "plan"
	EntitlementEntityTypeAddon      EntitlementEntityType = "addon"
	EntitlementEntityTypeCharge     EntitlementEntityType = "charge"
	EntitlementEntityTypePlanPrice  EntitlementEntityType = "plan_price"
	EntitlementEntityTypeAddonPrice EntitlementEntityType = "addon_price"
)

// just struct
type Entitlement struct {
	Id          string                `json:"id"`
	EntityId    string                `json:"entity_id"`
	EntityType  EntitlementEntityType `json:"entity_type"`
	FeatureId   string                `json:"feature_id"`
	FeatureName string                `json:"feature_name"`
	Value       string                `json:"value"`
	Name        string                `json:"name"`
	Object      string                `json:"object"`
}

// sub resources
// operations
// input params
type EntitlementListRequest struct {
	Limit      *int32        `json:"limit,omitempty"`
	Offset     string        `json:"offset,omitempty"`
	FeatureId  *StringFilter `json:"feature_id,omitempty"`
	EntityType *EnumFilter   `json:"entity_type,omitempty"`
	EntityId   *StringFilter `json:"entity_id,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *EntitlementListRequest) payload() any { return r }

type EntitlementCreateRequest struct {
	Action       Action                          `json:"action"`
	ChangeReason string                          `json:"change_reason,omitempty"`
	Entitlements []*EntitlementCreateEntitlement `json:"entitlements,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *EntitlementCreateRequest) payload() any { return r }

// input sub resource params multi
type EntitlementCreateEntitlement struct {
	EntityId            string                `json:"entity_id"`
	FeatureId           string                `json:"feature_id"`
	EntityType          EntitlementEntityType `json:"entity_type,omitempty"`
	Value               string                `json:"value,omitempty"`
	ApplyGrandfathering *bool                 `json:"apply_grandfathering,omitempty"`
}

// operation sub response
type EntitlementListEntitlementResponse struct {
	Entitlement *Entitlement `json:"entitlement,omitempty"`
}

// operation response
type EntitlementListResponse struct {
	List       []*EntitlementListEntitlementResponse `json:"list,omitempty"`
	NextOffset string                                `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type EntitlementCreateEntitlementResponse struct {
	Entitlement *Entitlement `json:"entitlement,omitempty"`
}

// operation response
type EntitlementCreateResponse struct {
	List []*EntitlementCreateEntitlementResponse `json:"list,omitempty"`
	apiResponse
}
