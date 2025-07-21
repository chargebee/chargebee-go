package quotedramp

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	quotedRampEnum "github.com/chargebee/chargebee-go/v3/models/quotedramp/enum"
)

type QuotedRamp struct {
	Id                          string                        `json:"id"`
	LineItems                   []*LineItem                   `json:"line_items"`
	Discounts                   []*Discount                   `json:"discounts"`
	ItemTiers                   []*ItemTier                   `json:"item_tiers"`
	CouponApplicabilityMappings []*CouponApplicabilityMapping `json:"coupon_applicability_mappings"`
	Object                      string                        `json:"object"`
}
type LineItem struct {
	ItemPriceId                               string                 `json:"item_price_id"`
	ItemType                                  enum.ItemType          `json:"item_type"`
	Quantity                                  int32                  `json:"quantity"`
	QuantityInDecimal                         string                 `json:"quantity_in_decimal"`
	MeteredQuantity                           string                 `json:"metered_quantity"`
	UnitPrice                                 int64                  `json:"unit_price"`
	UnitPriceInDecimal                        string                 `json:"unit_price_in_decimal"`
	Amount                                    int64                  `json:"amount"`
	AmountInDecimal                           string                 `json:"amount_in_decimal"`
	BillingPeriod                             int32                  `json:"billing_period"`
	BillingPeriodUnit                         enum.BillingPeriodUnit `json:"billing_period_unit"`
	FreeQuantity                              int32                  `json:"free_quantity"`
	FreeQuantityInDecimal                     string                 `json:"free_quantity_in_decimal"`
	BillingCycles                             int32                  `json:"billing_cycles"`
	ServicePeriodDays                         int32                  `json:"service_period_days"`
	ChargeOnEvent                             enum.ChargeOnEvent     `json:"charge_on_event"`
	ChargeOnce                                bool                   `json:"charge_once"`
	ChargeOnOption                            enum.ChargeOnOption    `json:"charge_on_option"`
	StartDate                                 int64                  `json:"start_date"`
	EndDate                                   int64                  `json:"end_date"`
	RampTierId                                string                 `json:"ramp_tier_id"`
	DiscountPerBillingCycle                   int64                  `json:"discount_per_billing_cycle"`
	DiscountPerBillingCycleInDecimal          string                 `json:"discount_per_billing_cycle_in_decimal"`
	ItemLevelDiscountPerBillingCycle          int64                  `json:"item_level_discount_per_billing_cycle"`
	ItemLevelDiscountPerBillingCycleInDecimal string                 `json:"item_level_discount_per_billing_cycle_in_decimal"`
	AmountPerBillingCycle                     int64                  `json:"amount_per_billing_cycle"`
	AmountPerBillingCycleInDecimal            string                 `json:"amount_per_billing_cycle_in_decimal"`
	NetAmountPerBillingCycle                  int64                  `json:"net_amount_per_billing_cycle"`
	NetAmountPerBillingCycleInDecimal         string                 `json:"net_amount_per_billing_cycle_in_decimal"`
	Object                                    string                 `json:"object"`
}
type Discount struct {
	Id            string                            `json:"id"`
	InvoiceName   string                            `json:"invoice_name"`
	Type          quotedRampEnum.DiscountType       `json:"type"`
	Percentage    float64                           `json:"percentage"`
	Amount        int64                             `json:"amount"`
	DurationType  enum.DurationType                 `json:"duration_type"`
	EntityType    quotedRampEnum.DiscountEntityType `json:"entity_type"`
	EntityId      string                            `json:"entity_id"`
	Period        int32                             `json:"period"`
	PeriodUnit    enum.PeriodUnit                   `json:"period_unit"`
	IncludedInMrr bool                              `json:"included_in_mrr"`
	ApplyOn       enum.ApplyOn                      `json:"apply_on"`
	ItemPriceId   string                            `json:"item_price_id"`
	CreatedAt     int64                             `json:"created_at"`
	UpdatedAt     int64                             `json:"updated_at"`
	StartDate     int64                             `json:"start_date"`
	EndDate       int64                             `json:"end_date"`
	Object        string                            `json:"object"`
}
type ItemTier struct {
	ItemPriceId           string `json:"item_price_id"`
	StartingUnit          int32  `json:"starting_unit"`
	EndingUnit            int32  `json:"ending_unit"`
	Price                 int64  `json:"price"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal"`
	PriceInDecimal        string `json:"price_in_decimal"`
	RampTierId            string `json:"ramp_tier_id"`
	Object                string `json:"object"`
}
type CouponApplicabilityMapping struct {
	CouponId               string   `json:"coupon_id"`
	ApplicableItemPriceIds []string `json:"applicable_item_price_ids"`
	Object                 string   `json:"object"`
}
