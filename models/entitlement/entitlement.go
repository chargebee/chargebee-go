package entitlement

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/filter"
	entitlementEnum "github.com/chargebee/chargebee-go/v3/models/entitlement/enum"
)

type Entitlement struct {
	Id          string                     `json:"id"`
	EntityId    string                     `json:"entity_id"`
	EntityType  entitlementEnum.EntityType `json:"entity_type"`
	FeatureId   string                     `json:"feature_id"`
	FeatureName string                     `json:"feature_name"`
	Value       string                     `json:"value"`
	Name        string                     `json:"name"`
	Object      string                     `json:"object"`
}
type ListRequestParams struct {
	Limit         *int32               `json:"limit,omitempty"`
	Offset        string               `json:"offset,omitempty"`
	FeatureId     *filter.StringFilter `json:"feature_id,omitempty"`
	EntityType    *filter.EnumFilter   `json:"entity_type,omitempty"`
	EntityId      *filter.StringFilter `json:"entity_id,omitempty"`
	IncludeDrafts *bool                `json:"include_drafts,omitempty"`
	Embed         string               `json:"embed,omitempty"`
}
type CreateRequestParams struct {
	Action       enum.Action                `json:"action"`
	ChangeReason string                     `json:"change_reason,omitempty"`
	Entitlements []*CreateEntitlementParams `json:"entitlements,omitempty"`
}
type CreateEntitlementParams struct {
	EntityId            string                     `json:"entity_id"`
	FeatureId           string                     `json:"feature_id"`
	EntityType          entitlementEnum.EntityType `json:"entity_type,omitempty"`
	Value               string                     `json:"value,omitempty"`
	ApplyGrandfathering *bool                      `json:"apply_grandfathering,omitempty"`
}
