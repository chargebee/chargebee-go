package priceramp

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	priceRampEnum "github.com/chargebee/chargebee-go/v3/models/priceramp/enum"
)

type PriceRamp struct {
	Id               string               `json:"id"`
	ItemPriceId      string               `json:"item_price_id"`
	Status           priceRampEnum.Status `json:"status"`
	CreatedAt        int64                `json:"created_at"`
	ResourceVersion  int64                `json:"resource_version"`
	UpdatedAt        int64                `json:"updated_at"`
	ArchivedAt       int64                `json:"archived_at"`
	RampIntervals    []*RampInterval      `json:"ramp_intervals"`
	BusinessEntityId string               `json:"business_entity_id"`
	Object           string               `json:"object"`
}
type RampInterval struct {
	StartingBillingCycle int32  `json:"starting_billing_cycle"`
	EndingBillingCycle   int32  `json:"ending_billing_cycle"`
	ApplyCatalogPrice    bool   `json:"apply_catalog_price"`
	Price                int64  `json:"price"`
	PriceInDecimal       string `json:"price_in_decimal"`
	Object               string `json:"object"`
}
type CreateRequestParams struct {
	ItemPriceId      string                      `json:"item_price_id"`
	BusinessEntityId string                      `json:"business_entity_id,omitempty"`
	RampIntervals    []*CreateRampIntervalParams `json:"ramp_intervals,omitempty"`
}
type CreateRampIntervalParams struct {
	StartingBillingCycle *int32 `json:"starting_billing_cycle"`
	EndingBillingCycle   *int32 `json:"ending_billing_cycle,omitempty"`
	ApplyCatalogPrice    *bool  `json:"apply_catalog_price,omitempty"`
	Price                *int64 `json:"price,omitempty"`
	PriceInDecimal       string `json:"price_in_decimal,omitempty"`
}
type UpdateRequestParams struct {
	Status        priceRampEnum.Status        `json:"status,omitempty"`
	RampIntervals []*UpdateRampIntervalParams `json:"ramp_intervals,omitempty"`
}
type UpdateRampIntervalParams struct {
	StartingBillingCycle *int32 `json:"starting_billing_cycle"`
	EndingBillingCycle   *int32 `json:"ending_billing_cycle,omitempty"`
	ApplyCatalogPrice    *bool  `json:"apply_catalog_price,omitempty"`
	Price                *int64 `json:"price,omitempty"`
	PriceInDecimal       string `json:"price_in_decimal,omitempty"`
}
type ListRequestParams struct {
	Limit                     *int32                  `json:"limit,omitempty"`
	Offset                    string                  `json:"offset,omitempty"`
	Id                        *filter.StringFilter    `json:"id,omitempty"`
	ItemPriceId               *filter.StringFilter    `json:"item_price_id,omitempty"`
	Status                    *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt                 *filter.TimestampFilter `json:"updated_at,omitempty"`
	BusinessEntityId          *filter.StringFilter    `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *filter.BooleanFilter   `json:"include_site_level_resources,omitempty"`
	SortBy                    *filter.SortFilter      `json:"sort_by,omitempty"`
}
