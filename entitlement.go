package chargebee

type EntityType string

const (
	EntityTypePlan       EntityType = "plan"
	EntityTypeAddon      EntityType = "addon"
	EntityTypeCharge     EntityType = "charge"
	EntityTypePlanPrice  EntityType = "plan_price"
	EntityTypeAddonPrice EntityType = "addon_price"
)

type Entitlement struct {
	Id          string     `json:"id"`
	EntityId    string     `json:"entity_id"`
	EntityType  EntityType `json:"entity_type"`
	FeatureId   string     `json:"feature_id"`
	FeatureName string     `json:"feature_name"`
	Value       string     `json:"value"`
	Name        string     `json:"name"`
	Object      string     `json:"object"`
}
type ListRequest struct {
	Limit         *int32               `json:"limit,omitempty"`
	Offset        string               `json:"offset,omitempty"`
	FeatureId     *filter.StringFilter `json:"feature_id,omitempty"`
	EntityType    *filter.EnumFilter   `json:"entity_type,omitempty"`
	EntityId      *filter.StringFilter `json:"entity_id,omitempty"`
	IncludeDrafts *bool                `json:"include_drafts,omitempty"`
	Embed         string               `json:"embed,omitempty"`
}
type CreateRequest struct {
	Action       enum.Action          `json:"action"`
	ChangeReason string               `json:"change_reason,omitempty"`
	Entitlements []*CreateEntitlement `json:"entitlements,omitempty"`
}
type CreateEntitlement struct {
	EntityId            string                 `json:"entity_id"`
	FeatureId           string                 `json:"feature_id"`
	EntityType          entitlement.EntityType `json:"entity_type,omitempty"`
	Value               string                 `json:"value,omitempty"`
	ApplyGrandfathering *bool                  `json:"apply_grandfathering,omitempty"`
}

type ListEntitlementResponse struct {
	Entitlement *Entitlement `json:"entitlement,omitempty"`
}

type ListResponse struct {
	List       []*ListEntitlementResponse `json:"list,omitempty"`
	NextOffset string                     `json:"next_offset,omitempty"`
}

type CreateEntitlementResponse struct {
	Entitlement *Entitlement `json:"entitlement,omitempty"`
}

type CreateResponse struct {
	List []*CreateEntitlementResponse `json:"list,omitempty"`
}
