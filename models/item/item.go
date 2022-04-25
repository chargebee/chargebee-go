package item

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	itemEnum "github.com/chargebee/chargebee-go/models/item/enum"
)

type Item struct {
	Id                   string                     `json:"id"`
	Name                 string                     `json:"name"`
	ExternalName         string                     `json:"external_name"`
	Description          string                     `json:"description"`
	Status               itemEnum.Status            `json:"status"`
	ResourceVersion      int64                      `json:"resource_version"`
	UpdatedAt            int64                      `json:"updated_at"`
	ItemFamilyId         string                     `json:"item_family_id"`
	Type                 itemEnum.Type              `json:"type"`
	IsShippable          bool                       `json:"is_shippable"`
	IsGiftable           bool                       `json:"is_giftable"`
	RedirectUrl          string                     `json:"redirect_url"`
	EnabledForCheckout   bool                       `json:"enabled_for_checkout"`
	EnabledInPortal      bool                       `json:"enabled_in_portal"`
	IncludedInMrr        bool                       `json:"included_in_mrr"`
	ItemApplicability    itemEnum.ItemApplicability `json:"item_applicability"`
	GiftClaimRedirectUrl string                     `json:"gift_claim_redirect_url"`
	Unit                 string                     `json:"unit"`
	Metered              bool                       `json:"metered"`
	UsageCalculation     itemEnum.UsageCalculation  `json:"usage_calculation"`
	ArchivedAt           int64                      `json:"archived_at"`
	Channel              enum.Channel               `json:"channel"`
	ApplicableItems      []*ApplicableItem          `json:"applicable_items"`
	Metadata             json.RawMessage            `json:"metadata"`
	CustomField          map[string]interface{}     `json:"custom_field"`
	Object               string                     `json:"object"`
}
type ApplicableItem struct {
	Id     string `json:"id"`
	Object string `json:"object"`
}
type CreateRequestParams struct {
	Id                   string                     `json:"id"`
	Name                 string                     `json:"name"`
	Type                 itemEnum.Type              `json:"type"`
	Description          string                     `json:"description,omitempty"`
	ItemFamilyId         string                     `json:"item_family_id"`
	IsGiftable           *bool                      `json:"is_giftable,omitempty"`
	IsShippable          *bool                      `json:"is_shippable,omitempty"`
	ExternalName         string                     `json:"external_name,omitempty"`
	EnabledInPortal      *bool                      `json:"enabled_in_portal,omitempty"`
	RedirectUrl          string                     `json:"redirect_url,omitempty"`
	EnabledForCheckout   *bool                      `json:"enabled_for_checkout,omitempty"`
	ItemApplicability    itemEnum.ItemApplicability `json:"item_applicability,omitempty"`
	ApplicableItems      []string                   `json:"applicable_items,omitempty"`
	Unit                 string                     `json:"unit,omitempty"`
	GiftClaimRedirectUrl string                     `json:"gift_claim_redirect_url,omitempty"`
	IncludedInMrr        *bool                      `json:"included_in_mrr,omitempty"`
	Metered              *bool                      `json:"metered,omitempty"`
	UsageCalculation     itemEnum.UsageCalculation  `json:"usage_calculation,omitempty"`
	Metadata             map[string]interface{}     `json:"metadata,omitempty"`
}
type UpdateRequestParams struct {
	Name                 string                     `json:"name,omitempty"`
	Description          string                     `json:"description,omitempty"`
	IsShippable          *bool                      `json:"is_shippable,omitempty"`
	ExternalName         string                     `json:"external_name,omitempty"`
	ItemFamilyId         string                     `json:"item_family_id,omitempty"`
	EnabledInPortal      *bool                      `json:"enabled_in_portal,omitempty"`
	RedirectUrl          string                     `json:"redirect_url,omitempty"`
	EnabledForCheckout   *bool                      `json:"enabled_for_checkout,omitempty"`
	ItemApplicability    itemEnum.ItemApplicability `json:"item_applicability,omitempty"`
	ClearApplicableItems *bool                      `json:"clear_applicable_items,omitempty"`
	ApplicableItems      []string                   `json:"applicable_items,omitempty"`
	Unit                 string                     `json:"unit,omitempty"`
	GiftClaimRedirectUrl string                     `json:"gift_claim_redirect_url,omitempty"`
	Metadata             map[string]interface{}     `json:"metadata,omitempty"`
	IncludedInMrr        *bool                      `json:"included_in_mrr,omitempty"`
	Status               itemEnum.Status            `json:"status,omitempty"`
}
type ListRequestParams struct {
	Limit              *int32                  `json:"limit,omitempty"`
	Offset             string                  `json:"offset,omitempty"`
	Id                 *filter.StringFilter    `json:"id,omitempty"`
	ItemFamilyId       *filter.StringFilter    `json:"item_family_id,omitempty"`
	Type               *filter.EnumFilter      `json:"type,omitempty"`
	Name               *filter.StringFilter    `json:"name,omitempty"`
	ItemApplicability  *filter.EnumFilter      `json:"item_applicability,omitempty"`
	Status             *filter.EnumFilter      `json:"status,omitempty"`
	IsGiftable         *filter.BooleanFilter   `json:"is_giftable,omitempty"`
	UpdatedAt          *filter.TimestampFilter `json:"updated_at,omitempty"`
	EnabledForCheckout *filter.BooleanFilter   `json:"enabled_for_checkout,omitempty"`
	EnabledInPortal    *filter.BooleanFilter   `json:"enabled_in_portal,omitempty"`
	Metered            *filter.BooleanFilter   `json:"metered,omitempty"`
	UsageCalculation   *filter.EnumFilter      `json:"usage_calculation,omitempty"`
	Channel            *filter.EnumFilter      `json:"channel,omitempty"`
	SortBy             *filter.SortFilter      `json:"sort_by,omitempty"`
}
