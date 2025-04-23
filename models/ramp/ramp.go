package ramp

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/filter"
	rampEnum "github.com/chargebee/chargebee-go/v3/models/ramp/enum"
)

type Ramp struct {
	Id                     string                  `json:"id"`
	Description            string                  `json:"description"`
	SubscriptionId         string                  `json:"subscription_id"`
	EffectiveFrom          int64                   `json:"effective_from"`
	Status                 rampEnum.Status         `json:"status"`
	CreatedAt              int64                   `json:"created_at"`
	ResourceVersion        int64                   `json:"resource_version"`
	UpdatedAt              int64                   `json:"updated_at"`
	ItemsToAdd             []*ItemsToAdd           `json:"items_to_add"`
	ItemsToUpdate          []*ItemsToUpdate        `json:"items_to_update"`
	CouponsToAdd           []*CouponsToAdd         `json:"coupons_to_add"`
	DiscountsToAdd         []*DiscountsToAdd       `json:"discounts_to_add"`
	ItemTiers              []*ItemTier             `json:"item_tiers"`
	ItemsToRemove          []string                `json:"items_to_remove"`
	CouponsToRemove        []string                `json:"coupons_to_remove"`
	DiscountsToRemove      []string                `json:"discounts_to_remove"`
	Deleted                bool                    `json:"deleted"`
	StatusTransitionReason *StatusTransitionReason `json:"status_transition_reason"`
	Object                 string                  `json:"object"`
}
type ItemsToAdd struct {
	ItemPriceId           string        `json:"item_price_id"`
	ItemType              enum.ItemType `json:"item_type"`
	Quantity              int32         `json:"quantity"`
	QuantityInDecimal     string        `json:"quantity_in_decimal"`
	UnitPrice             int64         `json:"unit_price"`
	UnitPriceInDecimal    string        `json:"unit_price_in_decimal"`
	Amount                int64         `json:"amount"`
	AmountInDecimal       string        `json:"amount_in_decimal"`
	FreeQuantity          int32         `json:"free_quantity"`
	FreeQuantityInDecimal string        `json:"free_quantity_in_decimal"`
	BillingCycles         int32         `json:"billing_cycles"`
	ServicePeriodDays     int32         `json:"service_period_days"`
	MeteredQuantity       string        `json:"metered_quantity"`
	Object                string        `json:"object"`
}
type ItemsToUpdate struct {
	ItemPriceId           string        `json:"item_price_id"`
	ItemType              enum.ItemType `json:"item_type"`
	Quantity              int32         `json:"quantity"`
	QuantityInDecimal     string        `json:"quantity_in_decimal"`
	UnitPrice             int64         `json:"unit_price"`
	UnitPriceInDecimal    string        `json:"unit_price_in_decimal"`
	Amount                int64         `json:"amount"`
	AmountInDecimal       string        `json:"amount_in_decimal"`
	FreeQuantity          int32         `json:"free_quantity"`
	FreeQuantityInDecimal string        `json:"free_quantity_in_decimal"`
	BillingCycles         int32         `json:"billing_cycles"`
	ServicePeriodDays     int32         `json:"service_period_days"`
	MeteredQuantity       string        `json:"metered_quantity"`
	Object                string        `json:"object"`
}
type CouponsToAdd struct {
	CouponId  string `json:"coupon_id"`
	ApplyTill int64  `json:"apply_till"`
	Object    string `json:"object"`
}
type DiscountsToAdd struct {
	Id            string                      `json:"id"`
	InvoiceName   string                      `json:"invoice_name"`
	Type          rampEnum.DiscountsToAddType `json:"type"`
	Percentage    float64                     `json:"percentage"`
	Amount        int64                       `json:"amount"`
	DurationType  enum.DurationType           `json:"duration_type"`
	Period        int32                       `json:"period"`
	PeriodUnit    enum.PeriodUnit             `json:"period_unit"`
	IncludedInMrr bool                        `json:"included_in_mrr"`
	ApplyOn       enum.ApplyOn                `json:"apply_on"`
	ItemPriceId   string                      `json:"item_price_id"`
	CreatedAt     int64                       `json:"created_at"`
	Object        string                      `json:"object"`
}
type ItemTier struct {
	ItemPriceId           string           `json:"item_price_id"`
	StartingUnit          int32            `json:"starting_unit"`
	EndingUnit            int32            `json:"ending_unit"`
	Price                 int64            `json:"price"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal"`
	PriceInDecimal        string           `json:"price_in_decimal"`
	PricingType           enum.PricingType `json:"pricing_type"`
	PackageSize           int32            `json:"package_size"`
	Index                 int32            `json:"index"`
	Object                string           `json:"object"`
}
type StatusTransitionReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Object  string `json:"object"`
}
type CreateForSubscriptionRequestParams struct {
	EffectiveFrom     *int64                                       `json:"effective_from"`
	Description       string                                       `json:"description,omitempty"`
	CouponsToRemove   []string                                     `json:"coupons_to_remove,omitempty"`
	DiscountsToRemove []string                                     `json:"discounts_to_remove,omitempty"`
	ItemsToRemove     []string                                     `json:"items_to_remove,omitempty"`
	ItemsToAdd        []*CreateForSubscriptionItemsToAddParams     `json:"items_to_add,omitempty"`
	ItemsToUpdate     []*CreateForSubscriptionItemsToUpdateParams  `json:"items_to_update,omitempty"`
	ItemTiers         []*CreateForSubscriptionItemTierParams       `json:"item_tiers,omitempty"`
	CouponsToAdd      []*CreateForSubscriptionCouponsToAddParams   `json:"coupons_to_add,omitempty"`
	DiscountsToAdd    []*CreateForSubscriptionDiscountsToAddParams `json:"discounts_to_add,omitempty"`
}
type CreateForSubscriptionItemsToAddParams struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type CreateForSubscriptionItemsToUpdateParams struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type CreateForSubscriptionItemTierParams struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type CreateForSubscriptionCouponsToAddParams struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type CreateForSubscriptionDiscountsToAddParams struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
}
type UpdateRequestParams struct {
	EffectiveFrom     *int64                        `json:"effective_from"`
	Description       string                        `json:"description,omitempty"`
	CouponsToRemove   []string                      `json:"coupons_to_remove,omitempty"`
	DiscountsToRemove []string                      `json:"discounts_to_remove,omitempty"`
	ItemsToRemove     []string                      `json:"items_to_remove,omitempty"`
	ItemsToAdd        []*UpdateItemsToAddParams     `json:"items_to_add,omitempty"`
	ItemsToUpdate     []*UpdateItemsToUpdateParams  `json:"items_to_update,omitempty"`
	ItemTiers         []*UpdateItemTierParams       `json:"item_tiers,omitempty"`
	CouponsToAdd      []*UpdateCouponsToAddParams   `json:"coupons_to_add,omitempty"`
	DiscountsToAdd    []*UpdateDiscountsToAddParams `json:"discounts_to_add,omitempty"`
}
type UpdateItemsToAddParams struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type UpdateItemsToUpdateParams struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type UpdateItemTierParams struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type UpdateCouponsToAddParams struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type UpdateDiscountsToAddParams struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
}
type ListRequestParams struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	IncludeDeleted *bool                   `json:"include_deleted,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id"`
	EffectiveFrom  *filter.TimestampFilter `json:"effective_from,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
