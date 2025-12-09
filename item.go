package chargebee

type Status string

const (
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
	StatusDeleted  Status = "deleted"
)

type Type string

const (
	TypePlan   Type = "plan"
	TypeAddon  Type = "addon"
	TypeCharge Type = "charge"
)

type ItemApplicability string

const (
	ItemApplicabilityAll        ItemApplicability = "all"
	ItemApplicabilityRestricted ItemApplicability = "restricted"
)

type UsageCalculation string

const (
	UsageCalculationSumOfUsages UsageCalculation = "sum_of_usages"
	UsageCalculationLastUsage   UsageCalculation = "last_usage"
	UsageCalculationMaxUsage    UsageCalculation = "max_usage"
)

type BundleConfigurationType string

const (
	BundleConfigurationTypeFixed BundleConfigurationType = "fixed"
)

type Item struct {
	Id                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	ExternalName         string                 `json:"external_name"`
	Description          string                 `json:"description"`
	Status               Status                 `json:"status"`
	ResourceVersion      int64                  `json:"resource_version"`
	UpdatedAt            int64                  `json:"updated_at"`
	ItemFamilyId         string                 `json:"item_family_id"`
	Type                 Type                   `json:"type"`
	IsShippable          bool                   `json:"is_shippable"`
	IsGiftable           bool                   `json:"is_giftable"`
	RedirectUrl          string                 `json:"redirect_url"`
	EnabledForCheckout   bool                   `json:"enabled_for_checkout"`
	EnabledInPortal      bool                   `json:"enabled_in_portal"`
	IncludedInMrr        bool                   `json:"included_in_mrr"`
	ItemApplicability    ItemApplicability      `json:"item_applicability"`
	GiftClaimRedirectUrl string                 `json:"gift_claim_redirect_url"`
	Unit                 string                 `json:"unit"`
	Metered              bool                   `json:"metered"`
	UsageCalculation     UsageCalculation       `json:"usage_calculation"`
	IsPercentagePricing  bool                   `json:"is_percentage_pricing"`
	ArchivedAt           int64                  `json:"archived_at"`
	Channel              enum.Channel           `json:"channel"`
	ApplicableItems      []*ApplicableItem      `json:"applicable_items"`
	BundleItems          []*BundleItem          `json:"bundle_items"`
	BundleConfiguration  *BundleConfiguration   `json:"bundle_configuration"`
	Metadata             json.RawMessage        `json:"metadata"`
	Deleted              bool                   `json:"deleted"`
	BusinessEntityId     string                 `json:"business_entity_id"`
	CustomField          map[string]interface{} `json:"custom_field"`
	Object               string                 `json:"object"`
}
type ApplicableItem struct {
	Id     string `json:"id"`
	Object string `json:"object"`
}
type BundleItem struct {
	ItemId          string        `json:"item_id"`
	ItemType        enum.ItemType `json:"item_type"`
	Quantity        int32         `json:"quantity"`
	PriceAllocation float64       `json:"price_allocation"`
	Object          string        `json:"object"`
}
type BundleConfiguration struct {
	Type   BundleConfigurationType `json:"type"`
	Object string                  `json:"object"`
}
type CreateRequest struct {
	Id                   string                     `json:"id"`
	Name                 string                     `json:"name"`
	Type                 Type                       `json:"type"`
	Description          string                     `json:"description,omitempty"`
	ItemFamilyId         string                     `json:"item_family_id"`
	IsGiftable           *bool                      `json:"is_giftable,omitempty"`
	IsShippable          *bool                      `json:"is_shippable,omitempty"`
	ExternalName         string                     `json:"external_name,omitempty"`
	EnabledInPortal      *bool                      `json:"enabled_in_portal,omitempty"`
	RedirectUrl          string                     `json:"redirect_url,omitempty"`
	EnabledForCheckout   *bool                      `json:"enabled_for_checkout,omitempty"`
	ItemApplicability    ItemApplicability          `json:"item_applicability,omitempty"`
	ApplicableItems      []string                   `json:"applicable_items,omitempty"`
	BundleConfiguration  *CreateBundleConfiguration `json:"bundle_configuration,omitempty"`
	Unit                 string                     `json:"unit,omitempty"`
	GiftClaimRedirectUrl string                     `json:"gift_claim_redirect_url,omitempty"`
	IncludedInMrr        *bool                      `json:"included_in_mrr,omitempty"`
	Metered              *bool                      `json:"metered,omitempty"`
	UsageCalculation     UsageCalculation           `json:"usage_calculation,omitempty"`
	IsPercentagePricing  *bool                      `json:"is_percentage_pricing,omitempty"`
	Metadata             map[string]interface{}     `json:"metadata,omitempty"`
	BusinessEntityId     string                     `json:"business_entity_id,omitempty"`
	BundleItemsToAdd     []*CreateBundleItemsToAdd  `json:"bundle_items_to_add,omitempty"`
}
type CreateBundleConfiguration struct {
	Type item.BundleConfigurationType `json:"type,omitempty"`
}
type CreateBundleItemsToAdd struct {
	ItemId          string        `json:"item_id,omitempty"`
	ItemType        enum.ItemType `json:"item_type,omitempty"`
	Quantity        *int32        `json:"quantity,omitempty"`
	PriceAllocation *float64      `json:"price_allocation,omitempty"`
}
type UpdateRequest struct {
	Name                 string                       `json:"name,omitempty"`
	Description          string                       `json:"description,omitempty"`
	IsShippable          *bool                        `json:"is_shippable,omitempty"`
	ExternalName         string                       `json:"external_name,omitempty"`
	ItemFamilyId         string                       `json:"item_family_id,omitempty"`
	EnabledInPortal      *bool                        `json:"enabled_in_portal,omitempty"`
	RedirectUrl          string                       `json:"redirect_url,omitempty"`
	EnabledForCheckout   *bool                        `json:"enabled_for_checkout,omitempty"`
	ItemApplicability    ItemApplicability            `json:"item_applicability,omitempty"`
	ClearApplicableItems *bool                        `json:"clear_applicable_items,omitempty"`
	ApplicableItems      []string                     `json:"applicable_items,omitempty"`
	BundleConfiguration  *UpdateBundleConfiguration   `json:"bundle_configuration,omitempty"`
	Unit                 string                       `json:"unit,omitempty"`
	GiftClaimRedirectUrl string                       `json:"gift_claim_redirect_url,omitempty"`
	Metadata             map[string]interface{}       `json:"metadata,omitempty"`
	IncludedInMrr        *bool                        `json:"included_in_mrr,omitempty"`
	Status               Status                       `json:"status,omitempty"`
	IsPercentagePricing  *bool                        `json:"is_percentage_pricing,omitempty"`
	BundleItemsToAdd     []*UpdateBundleItemsToAdd    `json:"bundle_items_to_add,omitempty"`
	BundleItemsToUpdate  []*UpdateBundleItemsToUpdate `json:"bundle_items_to_update,omitempty"`
	BundleItemsToRemove  []*UpdateBundleItemsToRemove `json:"bundle_items_to_remove,omitempty"`
}
type UpdateBundleConfiguration struct {
	Type item.BundleConfigurationType `json:"type,omitempty"`
}
type UpdateBundleItemsToAdd struct {
	ItemId          string        `json:"item_id,omitempty"`
	ItemType        enum.ItemType `json:"item_type,omitempty"`
	Quantity        *int32        `json:"quantity,omitempty"`
	PriceAllocation *float64      `json:"price_allocation,omitempty"`
}
type UpdateBundleItemsToUpdate struct {
	ItemId          string        `json:"item_id,omitempty"`
	ItemType        enum.ItemType `json:"item_type,omitempty"`
	Quantity        *int32        `json:"quantity,omitempty"`
	PriceAllocation *float64      `json:"price_allocation,omitempty"`
}
type UpdateBundleItemsToRemove struct {
	ItemId   string        `json:"item_id,omitempty"`
	ItemType enum.ItemType `json:"item_type,omitempty"`
}
type ListRequest struct {
	Limit                     *int32                   `json:"limit,omitempty"`
	Offset                    string                   `json:"offset,omitempty"`
	BundleConfiguration       *ListBundleConfiguration `json:"bundle_configuration,omitempty"`
	Id                        *filter.StringFilter     `json:"id,omitempty"`
	ItemFamilyId              *filter.StringFilter     `json:"item_family_id,omitempty"`
	Type                      *filter.EnumFilter       `json:"type,omitempty"`
	Name                      *filter.StringFilter     `json:"name,omitempty"`
	ItemApplicability         *filter.EnumFilter       `json:"item_applicability,omitempty"`
	Status                    *filter.EnumFilter       `json:"status,omitempty"`
	IsGiftable                *filter.BooleanFilter    `json:"is_giftable,omitempty"`
	UpdatedAt                 *filter.TimestampFilter  `json:"updated_at,omitempty"`
	EnabledForCheckout        *filter.BooleanFilter    `json:"enabled_for_checkout,omitempty"`
	EnabledInPortal           *filter.BooleanFilter    `json:"enabled_in_portal,omitempty"`
	Metered                   *filter.BooleanFilter    `json:"metered,omitempty"`
	UsageCalculation          *filter.EnumFilter       `json:"usage_calculation,omitempty"`
	Channel                   *filter.EnumFilter       `json:"channel,omitempty"`
	BusinessEntityId          *filter.StringFilter     `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *filter.BooleanFilter    `json:"include_site_level_resources,omitempty"`
	SortBy                    *filter.SortFilter       `json:"sort_by,omitempty"`
}
type ListBundleConfiguration struct {
	Type *filter.EnumFilter `json:"type,omitempty"`
}

type CreateResponse struct {
	Item *Item `json:"item,omitempty"`
}

type RetrieveResponse struct {
	Item *Item `json:"item,omitempty"`
}

type UpdateResponse struct {
	Item *Item `json:"item,omitempty"`
}

type ListItemResponse struct {
	Item *Item `json:"item,omitempty"`
}

type ListResponse struct {
	List       []*ListItemResponse `json:"list,omitempty"`
	NextOffset string              `json:"next_offset,omitempty"`
}

type DeleteResponse struct {
	Item *Item `json:"item,omitempty"`
}
