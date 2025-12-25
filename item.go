package chargebee

import (
	"encoding/json"
)

type ItemStatus string

const (
	ItemStatusActive   ItemStatus = "active"
	ItemStatusArchived ItemStatus = "archived"
	ItemStatusDeleted  ItemStatus = "deleted"
)

type ItemType string

const (
	ItemTypePlan   ItemType = "plan"
	ItemTypeAddon  ItemType = "addon"
	ItemTypeCharge ItemType = "charge"
)

type ItemItemApplicability string

const (
	ItemItemApplicabilityAll        ItemItemApplicability = "all"
	ItemItemApplicabilityRestricted ItemItemApplicability = "restricted"
)

type ItemUsageCalculation string

const (
	ItemUsageCalculationSumOfUsages ItemUsageCalculation = "sum_of_usages"
	ItemUsageCalculationLastUsage   ItemUsageCalculation = "last_usage"
	ItemUsageCalculationMaxUsage    ItemUsageCalculation = "max_usage"
)

type ItemChannel string

const (
	ItemChannelWeb       ItemChannel = "web"
	ItemChannelAppStore  ItemChannel = "app_store"
	ItemChannelPlayStore ItemChannel = "play_store"
)

type ItemBundleItemItemType string

const (
	ItemBundleItemItemTypePlan   ItemBundleItemItemType = "plan"
	ItemBundleItemItemTypeAddon  ItemBundleItemItemType = "addon"
	ItemBundleItemItemTypeCharge ItemBundleItemItemType = "charge"
)

type ItemBundleConfigurationType string

const (
	ItemBundleConfigurationTypeFixed ItemBundleConfigurationType = "fixed"
)

// just struct
type Item struct {
	Id                   string                   `json:"id"`
	Name                 string                   `json:"name"`
	ExternalName         string                   `json:"external_name"`
	Description          string                   `json:"description"`
	Status               ItemStatus               `json:"status"`
	ResourceVersion      int64                    `json:"resource_version"`
	UpdatedAt            int64                    `json:"updated_at"`
	ItemFamilyId         string                   `json:"item_family_id"`
	Type                 ItemType                 `json:"type"`
	IsShippable          bool                     `json:"is_shippable"`
	IsGiftable           bool                     `json:"is_giftable"`
	RedirectUrl          string                   `json:"redirect_url"`
	EnabledForCheckout   bool                     `json:"enabled_for_checkout"`
	EnabledInPortal      bool                     `json:"enabled_in_portal"`
	IncludedInMrr        bool                     `json:"included_in_mrr"`
	ItemApplicability    ItemItemApplicability    `json:"item_applicability"`
	GiftClaimRedirectUrl string                   `json:"gift_claim_redirect_url"`
	Unit                 string                   `json:"unit"`
	Metered              bool                     `json:"metered"`
	UsageCalculation     ItemUsageCalculation     `json:"usage_calculation"`
	IsPercentagePricing  bool                     `json:"is_percentage_pricing"`
	ArchivedAt           int64                    `json:"archived_at"`
	Channel              ItemChannel              `json:"channel"`
	ApplicableItems      []*ItemApplicableItem    `json:"applicable_items"`
	BundleItems          []*ItemBundleItem        `json:"bundle_items"`
	BundleConfiguration  *ItemBundleConfiguration `json:"bundle_configuration"`
	Metadata             json.RawMessage          `json:"metadata"`
	Deleted              bool                     `json:"deleted"`
	BusinessEntityId     string                   `json:"business_entity_id"`
	CustomField          CustomField              `json:"custom_field"`
	Object               string                   `json:"object"`
}

// sub resources
type ItemApplicableItem struct {
	Id     string `json:"id"`
	Object string `json:"object"`
}

type ItemBundleItem struct {
	ItemId          string                 `json:"item_id"`
	ItemType        ItemBundleItemItemType `json:"item_type"`
	Quantity        int32                  `json:"quantity"`
	PriceAllocation float64                `json:"price_allocation"`
	Object          string                 `json:"object"`
}

type ItemBundleConfiguration struct {
	Type   ItemBundleConfigurationType `json:"type"`
	Object string                      `json:"object"`
}

// operations
// input params
type ItemCreateRequest struct {
	Id                   string                         `json:"id"`
	Name                 string                         `json:"name"`
	Type                 ItemType                       `json:"type"`
	Description          string                         `json:"description,omitempty"`
	ItemFamilyId         string                         `json:"item_family_id"`
	IsGiftable           *bool                          `json:"is_giftable,omitempty"`
	IsShippable          *bool                          `json:"is_shippable,omitempty"`
	ExternalName         string                         `json:"external_name,omitempty"`
	EnabledInPortal      *bool                          `json:"enabled_in_portal,omitempty"`
	RedirectUrl          string                         `json:"redirect_url,omitempty"`
	EnabledForCheckout   *bool                          `json:"enabled_for_checkout,omitempty"`
	ItemApplicability    ItemItemApplicability          `json:"item_applicability,omitempty"`
	ApplicableItems      []string                       `json:"applicable_items,omitempty"`
	BundleConfiguration  *ItemCreateBundleConfiguration `json:"bundle_configuration,omitempty"`
	Unit                 string                         `json:"unit,omitempty"`
	GiftClaimRedirectUrl string                         `json:"gift_claim_redirect_url,omitempty"`
	IncludedInMrr        *bool                          `json:"included_in_mrr,omitempty"`
	Metered              *bool                          `json:"metered,omitempty"`
	UsageCalculation     ItemUsageCalculation           `json:"usage_calculation,omitempty"`
	IsPercentagePricing  *bool                          `json:"is_percentage_pricing,omitempty"`
	Metadata             map[string]interface{}         `json:"metadata,omitempty"`
	BusinessEntityId     string                         `json:"business_entity_id,omitempty"`
	BundleItemsToAdd     []*ItemCreateBundleItemsToAdd  `json:"bundle_items_to_add,omitempty"`
	apiRequest           `json:"-" form:"-"`
}

func (r *ItemCreateRequest) payload() any { return r }

// input sub resource params single
type ItemCreateBundleConfiguration struct {
	Type ItemBundleConfigurationType `json:"type,omitempty"`
}

// input sub resource params multi
type ItemCreateBundleItemsToAdd struct {
	ItemId          string                       `json:"item_id,omitempty"`
	ItemType        ItemBundleItemsToAddItemType `json:"item_type,omitempty"`
	Quantity        *int32                       `json:"quantity,omitempty"`
	PriceAllocation *float64                     `json:"price_allocation,omitempty"`
}

type ItemUpdateRequest struct {
	Name                 string                           `json:"name,omitempty"`
	Description          string                           `json:"description,omitempty"`
	IsShippable          *bool                            `json:"is_shippable,omitempty"`
	ExternalName         string                           `json:"external_name,omitempty"`
	ItemFamilyId         string                           `json:"item_family_id,omitempty"`
	EnabledInPortal      *bool                            `json:"enabled_in_portal,omitempty"`
	RedirectUrl          string                           `json:"redirect_url,omitempty"`
	EnabledForCheckout   *bool                            `json:"enabled_for_checkout,omitempty"`
	ItemApplicability    ItemItemApplicability            `json:"item_applicability,omitempty"`
	ApplicableItems      []string                         `json:"applicable_items,omitempty"`
	BundleConfiguration  *ItemUpdateBundleConfiguration   `json:"bundle_configuration,omitempty"`
	Unit                 string                           `json:"unit,omitempty"`
	GiftClaimRedirectUrl string                           `json:"gift_claim_redirect_url,omitempty"`
	Metadata             map[string]interface{}           `json:"metadata,omitempty"`
	IncludedInMrr        *bool                            `json:"included_in_mrr,omitempty"`
	Status               ItemStatus                       `json:"status,omitempty"`
	IsPercentagePricing  *bool                            `json:"is_percentage_pricing,omitempty"`
	BundleItemsToAdd     []*ItemUpdateBundleItemsToAdd    `json:"bundle_items_to_add,omitempty"`
	BundleItemsToUpdate  []*ItemUpdateBundleItemsToUpdate `json:"bundle_items_to_update,omitempty"`
	BundleItemsToRemove  []*ItemUpdateBundleItemsToRemove `json:"bundle_items_to_remove,omitempty"`
	apiRequest           `json:"-" form:"-"`
}

func (r *ItemUpdateRequest) payload() any { return r }

// input sub resource params single
type ItemUpdateBundleConfiguration struct {
	Type ItemBundleConfigurationType `json:"type,omitempty"`
}

// input sub resource params multi
type ItemUpdateBundleItemsToAdd struct {
	ItemId          string                       `json:"item_id,omitempty"`
	ItemType        ItemBundleItemsToAddItemType `json:"item_type,omitempty"`
	Quantity        *int32                       `json:"quantity,omitempty"`
	PriceAllocation *float64                     `json:"price_allocation,omitempty"`
}

// input sub resource params multi
type ItemUpdateBundleItemsToUpdate struct {
	ItemId          string                          `json:"item_id,omitempty"`
	ItemType        ItemBundleItemsToUpdateItemType `json:"item_type,omitempty"`
	Quantity        *int32                          `json:"quantity,omitempty"`
	PriceAllocation *float64                        `json:"price_allocation,omitempty"`
}

// input sub resource params multi
type ItemUpdateBundleItemsToRemove struct {
	ItemId   string                          `json:"item_id,omitempty"`
	ItemType ItemBundleItemsToRemoveItemType `json:"item_type,omitempty"`
}

type ItemListRequest struct {
	Limit                     *int32                   `json:"limit,omitempty"`
	Offset                    string                   `json:"offset,omitempty"`
	BundleConfiguration       *ListBundleConfiguration `json:"bundle_configuration,omitempty"`
	Id                        *StringFilter            `json:"id,omitempty"`
	ItemFamilyId              *StringFilter            `json:"item_family_id,omitempty"`
	Type                      *EnumFilter              `json:"type,omitempty"`
	Name                      *StringFilter            `json:"name,omitempty"`
	ItemApplicability         *EnumFilter              `json:"item_applicability,omitempty"`
	Status                    *EnumFilter              `json:"status,omitempty"`
	IsGiftable                *BooleanFilter           `json:"is_giftable,omitempty"`
	UpdatedAt                 *TimestampFilter         `json:"updated_at,omitempty"`
	EnabledForCheckout        *BooleanFilter           `json:"enabled_for_checkout,omitempty"`
	EnabledInPortal           *BooleanFilter           `json:"enabled_in_portal,omitempty"`
	Metered                   *BooleanFilter           `json:"metered,omitempty"`
	UsageCalculation          *EnumFilter              `json:"usage_calculation,omitempty"`
	Channel                   *EnumFilter              `json:"channel,omitempty"`
	BusinessEntityId          *StringFilter            `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *BooleanFilter           `json:"include_site_level_resources,omitempty"`
	SortBy                    *SortFilter              `json:"sort_by,omitempty"`
	apiRequest                `json:"-" form:"-"`
}

func (r *ItemListRequest) payload() any { return r }

// input sub resource params single
type ItemListBundleConfiguration struct {
	Type *EnumFilter `json:"type,omitempty"`
}

// operation response
type ItemCreateResponse struct {
	Item *Item `json:"item,omitempty"`
	apiResponse
}

// operation response
type ItemRetrieveResponse struct {
	Item *Item `json:"item,omitempty"`
	apiResponse
}

// operation response
type ItemUpdateResponse struct {
	Item *Item `json:"item,omitempty"`
	apiResponse
}

// operation sub response
type ItemListItemResponse struct {
	Item *Item `json:"item,omitempty"`
}

// operation response
type ItemListResponse struct {
	List       []*ItemListItemResponse `json:"list,omitempty"`
	NextOffset string                  `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type ItemDeleteResponse struct {
	Item *Item `json:"item,omitempty"`
	apiResponse
}
