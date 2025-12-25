package chargebee

type ItemEntitlementItemType string

const (
	ItemEntitlementItemTypePlan         ItemEntitlementItemType = "plan"
	ItemEntitlementItemTypeAddon        ItemEntitlementItemType = "addon"
	ItemEntitlementItemTypeCharge       ItemEntitlementItemType = "charge"
	ItemEntitlementItemTypeSubscription ItemEntitlementItemType = "subscription"
	ItemEntitlementItemTypeItem         ItemEntitlementItemType = "item"
)

type ItemEntitlementAction string

const (
	ItemEntitlementActionUpsert ItemEntitlementAction = "upsert"
	ItemEntitlementActionRemove ItemEntitlementAction = "remove"
)

// just struct
type ItemEntitlement struct {
	Id          string                  `json:"id"`
	ItemId      string                  `json:"item_id"`
	ItemType    ItemEntitlementItemType `json:"item_type"`
	FeatureId   string                  `json:"feature_id"`
	FeatureName string                  `json:"feature_name"`
	Value       string                  `json:"value"`
	Name        string                  `json:"name"`
	Object      string                  `json:"object"`
}

// sub resources
// operations
// input params
type ItemEntitlementItemEntitlementsForItemRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *ItemEntitlementItemEntitlementsForItemRequest) payload() any { return r }

type ItemEntitlementItemEntitlementsForFeatureRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *ItemEntitlementItemEntitlementsForFeatureRequest) payload() any { return r }

type ItemEntitlementAddItemEntitlementsRequest struct {
	Action           ItemEntitlementAction                                `json:"action"`
	ItemEntitlements []*ItemEntitlementAddItemEntitlementsItemEntitlement `json:"item_entitlements,omitempty"`
	apiRequest       `json:"-" form:"-"`
}

func (r *ItemEntitlementAddItemEntitlementsRequest) payload() any { return r }

// input sub resource params multi
type ItemEntitlementAddItemEntitlementsItemEntitlement struct {
	ItemId   string                                 `json:"item_id"`
	ItemType ItemEntitlementItemEntitlementItemType `json:"item_type,omitempty"`
	Value    string                                 `json:"value,omitempty"`
}
type ItemEntitlementUpsertOrRemoveItemEntitlementsForItemRequest struct {
	Action           ItemEntitlementAction                                                  `json:"action"`
	ItemEntitlements []*ItemEntitlementUpsertOrRemoveItemEntitlementsForItemItemEntitlement `json:"item_entitlements,omitempty"`
	apiRequest       `json:"-" form:"-"`
}

func (r *ItemEntitlementUpsertOrRemoveItemEntitlementsForItemRequest) payload() any { return r }

// input sub resource params multi
type ItemEntitlementUpsertOrRemoveItemEntitlementsForItemItemEntitlement struct {
	FeatureId string `json:"feature_id"`
	Value     string `json:"value,omitempty"`
}

// operation sub response
type ItemEntitlementItemEntitlementsForItemItemEntitlementResponse struct {
	ItemEntitlement *ItemEntitlement `json:"item_entitlement,omitempty"`
}

// operation response
type ItemEntitlementItemEntitlementsForItemResponse struct {
	List       []*ItemEntitlementItemEntitlementsForItemItemEntitlementResponse `json:"list,omitempty"`
	NextOffset string                                                           `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type ItemEntitlementItemEntitlementsForFeatureItemEntitlementResponse struct {
	ItemEntitlement *ItemEntitlement `json:"item_entitlement,omitempty"`
}

// operation response
type ItemEntitlementItemEntitlementsForFeatureResponse struct {
	List       []*ItemEntitlementItemEntitlementsForFeatureItemEntitlementResponse `json:"list,omitempty"`
	NextOffset string                                                              `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type ItemEntitlementAddItemEntitlementsResponse struct {
	ItemEntitlement *ItemEntitlement `json:"item_entitlement,omitempty"`
	apiResponse
}

// operation response
type ItemEntitlementUpsertOrRemoveItemEntitlementsForItemResponse struct {
	ItemEntitlement *ItemEntitlement `json:"item_entitlement,omitempty"`
	apiResponse
}
