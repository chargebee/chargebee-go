package itementitlement

import (
	"github.com/chargebee/chargebee-go/enum"
	itemEntitlementEnum "github.com/chargebee/chargebee-go/models/itementitlement/enum"
)

type ItemEntitlement struct {
	Id          string                       `json:"id"`
	ItemId      string                       `json:"item_id"`
	ItemType    itemEntitlementEnum.ItemType `json:"item_type"`
	FeatureId   string                       `json:"feature_id"`
	FeatureName string                       `json:"feature_name"`
	Value       string                       `json:"value"`
	Name        string                       `json:"name"`
	Object      string                       `json:"object"`
}
type ItemEntitlementsForItemRequestParams struct {
	Limit         *int32 `json:"limit,omitempty"`
	Offset        string `json:"offset,omitempty"`
	IncludeDrafts *bool  `json:"include_drafts,omitempty"`
	Embed         string `json:"embed,omitempty"`
}
type ItemEntitlementsForFeatureRequestParams struct {
	Limit         *int32 `json:"limit,omitempty"`
	Offset        string `json:"offset,omitempty"`
	IncludeDrafts *bool  `json:"include_drafts,omitempty"`
}
type AddItemEntitlementsRequestParams struct {
	Action           enum.Action                                 `json:"action"`
	ItemEntitlements []*AddItemEntitlementsItemEntitlementParams `json:"item_entitlements,omitempty"`
}
type AddItemEntitlementsItemEntitlementParams struct {
	ItemId   string                       `json:"item_id"`
	ItemType itemEntitlementEnum.ItemType `json:"item_type,omitempty"`
	Value    string                       `json:"value,omitempty"`
}
type UpsertOrRemoveItemEntitlementsForItemRequestParams struct {
	Action           enum.Action                                                   `json:"action"`
	ItemEntitlements []*UpsertOrRemoveItemEntitlementsForItemItemEntitlementParams `json:"item_entitlements,omitempty"`
}
type UpsertOrRemoveItemEntitlementsForItemItemEntitlementParams struct {
	FeatureId string `json:"feature_id"`
	Value     string `json:"value,omitempty"`
}
