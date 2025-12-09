package chargebee

type ItemType string

const (
	ItemTypePlan         ItemType = "plan"
	ItemTypeAddon        ItemType = "addon"
	ItemTypeCharge       ItemType = "charge"
	ItemTypeSubscription ItemType = "subscription"
	ItemTypeItem         ItemType = "item"
)

type ItemEntitlement struct {
	Id          string   `json:"id"`
	ItemId      string   `json:"item_id"`
	ItemType    ItemType `json:"item_type"`
	FeatureId   string   `json:"feature_id"`
	FeatureName string   `json:"feature_name"`
	Value       string   `json:"value"`
	Name        string   `json:"name"`
	Object      string   `json:"object"`
}
type ItemEntitlementsForItemRequest struct {
	Limit         *int32 `json:"limit,omitempty"`
	Offset        string `json:"offset,omitempty"`
	IncludeDrafts *bool  `json:"include_drafts,omitempty"`
	Embed         string `json:"embed,omitempty"`
}
type ItemEntitlementsForFeatureRequest struct {
	Limit         *int32 `json:"limit,omitempty"`
	Offset        string `json:"offset,omitempty"`
	IncludeDrafts *bool  `json:"include_drafts,omitempty"`
}
type AddItemEntitlementsRequest struct {
	Action           enum.Action                           `json:"action"`
	ItemEntitlements []*AddItemEntitlementsItemEntitlement `json:"item_entitlements,omitempty"`
}
type AddItemEntitlementsItemEntitlement struct {
	ItemId   string                   `json:"item_id"`
	ItemType itemEntitlement.ItemType `json:"item_type,omitempty"`
	Value    string                   `json:"value,omitempty"`
}
type UpsertOrRemoveItemEntitlementsForItemRequest struct {
	Action           enum.Action                                             `json:"action"`
	ItemEntitlements []*UpsertOrRemoveItemEntitlementsForItemItemEntitlement `json:"item_entitlements,omitempty"`
}
type UpsertOrRemoveItemEntitlementsForItemItemEntitlement struct {
	FeatureId string `json:"feature_id"`
	Value     string `json:"value,omitempty"`
}

type ItemEntitlementsForItemItemEntitlementResponse struct {
	ItemEntitlement *ItemEntitlement `json:"item_entitlement,omitempty"`
}

type ItemEntitlementsForItemResponse struct {
	List       []*ItemEntitlementsForItemItemEntitlementResponse `json:"list,omitempty"`
	NextOffset string                                            `json:"next_offset,omitempty"`
}

type ItemEntitlementsForFeatureItemEntitlementResponse struct {
	ItemEntitlement *ItemEntitlement `json:"item_entitlement,omitempty"`
}

type ItemEntitlementsForFeatureResponse struct {
	List       []*ItemEntitlementsForFeatureItemEntitlementResponse `json:"list,omitempty"`
	NextOffset string                                               `json:"next_offset,omitempty"`
}

type AddItemEntitlementsResponse struct {
	ItemEntitlement *ItemEntitlement `json:"item_entitlement,omitempty"`
}

type UpsertOrRemoveItemEntitlementsForItemResponse struct {
	ItemEntitlement *ItemEntitlement `json:"item_entitlement,omitempty"`
}
