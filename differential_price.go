package chargebee

import (
	"encoding/json"
)

type DifferentialPriceStatus string

const (
	DifferentialPriceStatusActive  DifferentialPriceStatus = "active"
	DifferentialPriceStatusDeleted DifferentialPriceStatus = "deleted"
)

type DifferentialPriceTierPricingType string

const (
	DifferentialPriceTierPricingTypePerUnit DifferentialPriceTierPricingType = "per_unit"
	DifferentialPriceTierPricingTypeFlatFee DifferentialPriceTierPricingType = "flat_fee"
	DifferentialPriceTierPricingTypePackage DifferentialPriceTierPricingType = "package"
)

type DifferentialPriceParentPeriodPeriodUnit string

const (
	DifferentialPriceParentPeriodPeriodUnitDay   DifferentialPriceParentPeriodPeriodUnit = "day"
	DifferentialPriceParentPeriodPeriodUnitWeek  DifferentialPriceParentPeriodPeriodUnit = "week"
	DifferentialPriceParentPeriodPeriodUnitMonth DifferentialPriceParentPeriodPeriodUnit = "month"
	DifferentialPriceParentPeriodPeriodUnitYear  DifferentialPriceParentPeriodPeriodUnit = "year"
)

// just struct
type DifferentialPrice struct {
	Id               string                           `json:"id"`
	ItemPriceId      string                           `json:"item_price_id"`
	ParentItemId     string                           `json:"parent_item_id"`
	Price            int64                            `json:"price"`
	PriceInDecimal   string                           `json:"price_in_decimal"`
	Status           DifferentialPriceStatus          `json:"status"`
	ResourceVersion  int64                            `json:"resource_version"`
	UpdatedAt        int64                            `json:"updated_at"`
	CreatedAt        int64                            `json:"created_at"`
	ModifiedAt       int64                            `json:"modified_at"`
	Tiers            []*DifferentialPriceTier         `json:"tiers"`
	CurrencyCode     string                           `json:"currency_code"`
	ParentPeriods    []*DifferentialPriceParentPeriod `json:"parent_periods"`
	BusinessEntityId string                           `json:"business_entity_id"`
	Deleted          bool                             `json:"deleted"`
	Object           string                           `json:"object"`
}

// sub resources
type DifferentialPriceTier struct {
	StartingUnit          int32                            `json:"starting_unit"`
	EndingUnit            int32                            `json:"ending_unit"`
	Price                 int64                            `json:"price"`
	StartingUnitInDecimal string                           `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string                           `json:"ending_unit_in_decimal"`
	PriceInDecimal        string                           `json:"price_in_decimal"`
	PricingType           DifferentialPriceTierPricingType `json:"pricing_type"`
	PackageSize           int32                            `json:"package_size"`
	Object                string                           `json:"object"`
}
type DifferentialPriceParentPeriod struct {
	PeriodUnit DifferentialPriceParentPeriodPeriodUnit `json:"period_unit"`
	Period     json.RawMessage                         `json:"period"`
	Object     string                                  `json:"object"`
}

// operations
// input params
type DifferentialPriceCreateRequest struct {
	ParentItemId     string                                 `json:"parent_item_id"`
	Price            *int64                                 `json:"price,omitempty"`
	PriceInDecimal   string                                 `json:"price_in_decimal,omitempty"`
	ParentPeriods    []*DifferentialPriceCreateParentPeriod `json:"parent_periods,omitempty"`
	Tiers            []*DifferentialPriceCreateTier         `json:"tiers,omitempty"`
	BusinessEntityId string                                 `json:"business_entity_id,omitempty"`
	apiRequest       `json:"-" form:"-"`
}

func (r *DifferentialPriceCreateRequest) payload() any { return r }

// input sub resource params multi
type DifferentialPriceCreateParentPeriod struct {
	PeriodUnit ParentPeriodPeriodUnit `json:"period_unit"`
	Period     []interface{}          `json:"period,omitempty"`
}

// input sub resource params multi
type DifferentialPriceCreateTier struct {
	StartingUnit          *int32      `json:"starting_unit,omitempty"`
	EndingUnit            *int32      `json:"ending_unit,omitempty"`
	Price                 *int64      `json:"price,omitempty"`
	StartingUnitInDecimal string      `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string      `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string      `json:"price_in_decimal,omitempty"`
	PricingType           PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32      `json:"package_size,omitempty"`
}
type DifferentialPriceRetrieveRequest struct {
	ItemPriceId string `json:"item_price_id"`
	apiRequest  `json:"-" form:"-"`
}

func (r *DifferentialPriceRetrieveRequest) payload() any { return r }

type DifferentialPriceUpdateRequest struct {
	ItemPriceId    string                                 `json:"item_price_id"`
	Price          *int64                                 `json:"price,omitempty"`
	PriceInDecimal string                                 `json:"price_in_decimal,omitempty"`
	ParentPeriods  []*DifferentialPriceUpdateParentPeriod `json:"parent_periods,omitempty"`
	Tiers          []*DifferentialPriceUpdateTier         `json:"tiers,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *DifferentialPriceUpdateRequest) payload() any { return r }

// input sub resource params multi
type DifferentialPriceUpdateParentPeriod struct {
	PeriodUnit ParentPeriodPeriodUnit `json:"period_unit"`
	Period     []interface{}          `json:"period,omitempty"`
}

// input sub resource params multi
type DifferentialPriceUpdateTier struct {
	StartingUnit          *int32      `json:"starting_unit,omitempty"`
	EndingUnit            *int32      `json:"ending_unit,omitempty"`
	Price                 *int64      `json:"price,omitempty"`
	StartingUnitInDecimal string      `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string      `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string      `json:"price_in_decimal,omitempty"`
	PricingType           PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32      `json:"package_size,omitempty"`
}
type DifferentialPriceDeleteRequest struct {
	ItemPriceId string `json:"item_price_id"`
	apiRequest  `json:"-" form:"-"`
}

func (r *DifferentialPriceDeleteRequest) payload() any { return r }

type DifferentialPriceListRequest struct {
	Limit        *int32        `json:"limit,omitempty"`
	Offset       string        `json:"offset,omitempty"`
	ItemPriceId  *StringFilter `json:"item_price_id,omitempty"`
	ItemId       *StringFilter `json:"item_id,omitempty"`
	Id           *StringFilter `json:"id,omitempty"`
	ParentItemId *StringFilter `json:"parent_item_id,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *DifferentialPriceListRequest) payload() any { return r }

// operation response
type DifferentialPriceCreateResponse struct {
	DifferentialPrice *DifferentialPrice `json:"differential_price,omitempty"`
	apiResponse
}

// operation response
type DifferentialPriceRetrieveResponse struct {
	DifferentialPrice *DifferentialPrice `json:"differential_price,omitempty"`
	apiResponse
}

// operation response
type DifferentialPriceUpdateResponse struct {
	DifferentialPrice *DifferentialPrice `json:"differential_price,omitempty"`
	apiResponse
}

// operation response
type DifferentialPriceDeleteResponse struct {
	DifferentialPrice *DifferentialPrice `json:"differential_price,omitempty"`
	apiResponse
}

// operation sub response
type DifferentialPriceListDifferentialPriceResponse struct {
	DifferentialPrice *DifferentialPrice `json:"differential_price,omitempty"`
}

// operation response
type DifferentialPriceListResponse struct {
	List       []*DifferentialPriceListDifferentialPriceResponse `json:"list,omitempty"`
	NextOffset string                                            `json:"next_offset,omitempty"`
	apiResponse
}
