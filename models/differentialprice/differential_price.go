package differentialprice

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/filter"
	differentialPriceEnum "github.com/chargebee/chargebee-go/models/differentialprice/enum"
)

type DifferentialPrice struct {
	Id              string                       `json:"id"`
	ItemPriceId     string                       `json:"item_price_id"`
	ParentItemId    string                       `json:"parent_item_id"`
	Price           int32                        `json:"price"`
	PriceInDecimal  string                       `json:"price_in_decimal"`
	Status          differentialPriceEnum.Status `json:"status"`
	ResourceVersion int64                        `json:"resource_version"`
	UpdatedAt       int64                        `json:"updated_at"`
	CreatedAt       int64                        `json:"created_at"`
	ModifiedAt      int64                        `json:"modified_at"`
	Tiers           []*Tier                      `json:"tiers"`
	CurrencyCode    string                       `json:"currency_code"`
	ParentPeriods   []*ParentPeriod              `json:"parent_periods"`
	Object          string                       `json:"object"`
}
type Tier struct {
	StartingUnit          int32  `json:"starting_unit"`
	EndingUnit            int32  `json:"ending_unit"`
	Price                 int32  `json:"price"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal"`
	PriceInDecimal        string `json:"price_in_decimal"`
	Object                string `json:"object"`
}
type ParentPeriod struct {
	PeriodUnit differentialPriceEnum.ParentPeriodPeriodUnit `json:"period_unit"`
	Period     json.RawMessage                              `json:"period"`
	Object     string                                       `json:"object"`
}
type CreateRequestParams struct {
	ParentItemId   string                      `json:"parent_item_id"`
	Price          *int32                      `json:"price,omitempty"`
	PriceInDecimal string                      `json:"price_in_decimal,omitempty"`
	ParentPeriods  []*CreateParentPeriodParams `json:"parent_periods,omitempty"`
	Tiers          []*CreateTierParams         `json:"tiers,omitempty"`
}
type CreateParentPeriodParams struct {
	PeriodUnit differentialPriceEnum.ParentPeriodPeriodUnit `json:"period_unit"`
	Period     []interface{}                                `json:"period,omitempty"`
}
type CreateTierParams struct {
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type RetrieveRequestParams struct {
	ItemPriceId string `json:"item_price_id"`
}
type UpdateRequestParams struct {
	ItemPriceId    string                      `json:"item_price_id"`
	Price          *int32                      `json:"price,omitempty"`
	PriceInDecimal string                      `json:"price_in_decimal,omitempty"`
	ParentPeriods  []*UpdateParentPeriodParams `json:"parent_periods,omitempty"`
	Tiers          []*UpdateTierParams         `json:"tiers,omitempty"`
}
type UpdateParentPeriodParams struct {
	PeriodUnit differentialPriceEnum.ParentPeriodPeriodUnit `json:"period_unit"`
	Period     []interface{}                                `json:"period,omitempty"`
}
type UpdateTierParams struct {
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type DeleteRequestParams struct {
	ItemPriceId string `json:"item_price_id"`
}
type ListRequestParams struct {
	Limit        *int32               `json:"limit,omitempty"`
	Offset       string               `json:"offset,omitempty"`
	ItemPriceId  *filter.StringFilter `json:"item_price_id,omitempty"`
	ItemId       *filter.StringFilter `json:"item_id,omitempty"`
	Id           *filter.StringFilter `json:"id,omitempty"`
	ParentItemId *filter.StringFilter `json:"parent_item_id,omitempty"`
}
