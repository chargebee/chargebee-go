package chargebee

type Status string

const (
	StatusActive  Status = "active"
	StatusDeleted Status = "deleted"
)

type ParentPeriodPeriodUnit string

const (
	ParentPeriodPeriodUnitDay   ParentPeriodPeriodUnit = "day"
	ParentPeriodPeriodUnitWeek  ParentPeriodPeriodUnit = "week"
	ParentPeriodPeriodUnitMonth ParentPeriodPeriodUnit = "month"
	ParentPeriodPeriodUnitYear  ParentPeriodPeriodUnit = "year"
)

type DifferentialPrice struct {
	Id               string          `json:"id"`
	ItemPriceId      string          `json:"item_price_id"`
	ParentItemId     string          `json:"parent_item_id"`
	Price            int64           `json:"price"`
	PriceInDecimal   string          `json:"price_in_decimal"`
	Status           Status          `json:"status"`
	ResourceVersion  int64           `json:"resource_version"`
	UpdatedAt        int64           `json:"updated_at"`
	CreatedAt        int64           `json:"created_at"`
	ModifiedAt       int64           `json:"modified_at"`
	Tiers            []*Tier         `json:"tiers"`
	CurrencyCode     string          `json:"currency_code"`
	ParentPeriods    []*ParentPeriod `json:"parent_periods"`
	BusinessEntityId string          `json:"business_entity_id"`
	Deleted          bool            `json:"deleted"`
	Object           string          `json:"object"`
}
type Tier struct {
	StartingUnit          int32            `json:"starting_unit"`
	EndingUnit            int32            `json:"ending_unit"`
	Price                 int64            `json:"price"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal"`
	PriceInDecimal        string           `json:"price_in_decimal"`
	PricingType           enum.PricingType `json:"pricing_type"`
	PackageSize           int32            `json:"package_size"`
	Object                string           `json:"object"`
}
type ParentPeriod struct {
	PeriodUnit ParentPeriodPeriodUnit `json:"period_unit"`
	Period     json.RawMessage        `json:"period"`
	Object     string                 `json:"object"`
}
type CreateRequest struct {
	ParentItemId     string                `json:"parent_item_id"`
	Price            *int64                `json:"price,omitempty"`
	PriceInDecimal   string                `json:"price_in_decimal,omitempty"`
	ParentPeriods    []*CreateParentPeriod `json:"parent_periods,omitempty"`
	Tiers            []*CreateTier         `json:"tiers,omitempty"`
	BusinessEntityId string                `json:"business_entity_id,omitempty"`
}
type CreateParentPeriod struct {
	PeriodUnit differentialPrice.ParentPeriodPeriodUnit `json:"period_unit"`
	Period     []interface{}                            `json:"period,omitempty"`
}
type CreateTier struct {
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type RetrieveRequest struct {
	ItemPriceId string `json:"item_price_id"`
}
type UpdateRequest struct {
	ItemPriceId    string                `json:"item_price_id"`
	Price          *int64                `json:"price,omitempty"`
	PriceInDecimal string                `json:"price_in_decimal,omitempty"`
	ParentPeriods  []*UpdateParentPeriod `json:"parent_periods,omitempty"`
	Tiers          []*UpdateTier         `json:"tiers,omitempty"`
}
type UpdateParentPeriod struct {
	PeriodUnit differentialPrice.ParentPeriodPeriodUnit `json:"period_unit"`
	Period     []interface{}                            `json:"period,omitempty"`
}
type UpdateTier struct {
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type DeleteRequest struct {
	ItemPriceId string `json:"item_price_id"`
}
type ListRequest struct {
	Limit        *int32               `json:"limit,omitempty"`
	Offset       string               `json:"offset,omitempty"`
	ItemPriceId  *filter.StringFilter `json:"item_price_id,omitempty"`
	ItemId       *filter.StringFilter `json:"item_id,omitempty"`
	Id           *filter.StringFilter `json:"id,omitempty"`
	ParentItemId *filter.StringFilter `json:"parent_item_id,omitempty"`
}

type CreateResponse struct {
	DifferentialPrice *DifferentialPrice `json:"differential_price,omitempty"`
}

type RetrieveResponse struct {
	DifferentialPrice *DifferentialPrice `json:"differential_price,omitempty"`
}

type UpdateResponse struct {
	DifferentialPrice *DifferentialPrice `json:"differential_price,omitempty"`
}

type DeleteResponse struct {
	DifferentialPrice *DifferentialPrice `json:"differential_price,omitempty"`
}

type ListDifferentialPriceResponse struct {
	DifferentialPrice *DifferentialPrice `json:"differential_price,omitempty"`
}

type ListResponse struct {
	List       []*ListDifferentialPriceResponse `json:"list,omitempty"`
	NextOffset string                           `json:"next_offset,omitempty"`
}
