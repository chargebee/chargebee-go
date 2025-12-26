package chargebee

type QuotedRampLineItemBillingPeriodUnit string

const (
	QuotedRampLineItemBillingPeriodUnitDay   QuotedRampLineItemBillingPeriodUnit = "day"
	QuotedRampLineItemBillingPeriodUnitWeek  QuotedRampLineItemBillingPeriodUnit = "week"
	QuotedRampLineItemBillingPeriodUnitMonth QuotedRampLineItemBillingPeriodUnit = "month"
	QuotedRampLineItemBillingPeriodUnitYear  QuotedRampLineItemBillingPeriodUnit = "year"
)

type QuotedRampDiscountType string

const (
	QuotedRampDiscountTypeFixedAmount QuotedRampDiscountType = "fixed_amount"
	QuotedRampDiscountTypePercentage  QuotedRampDiscountType = "percentage"
)

type QuotedRampDiscountEntityType string

const (
	QuotedRampDiscountEntityTypeItemLevelCoupon       QuotedRampDiscountEntityType = "item_level_coupon"
	QuotedRampDiscountEntityTypeDocumentLevelCoupon   QuotedRampDiscountEntityType = "document_level_coupon"
	QuotedRampDiscountEntityTypeItemLevelDiscount     QuotedRampDiscountEntityType = "item_level_discount"
	QuotedRampDiscountEntityTypeDocumentLevelDiscount QuotedRampDiscountEntityType = "document_level_discount"
)

// just struct
type QuotedRamp struct {
	Id                          string                                  `json:"id"`
	LineItems                   []*QuotedRampLineItem                   `json:"line_items"`
	Discounts                   []*QuotedRampDiscount                   `json:"discounts"`
	ItemTiers                   []*QuotedRampItemTier                   `json:"item_tiers"`
	CouponApplicabilityMappings []*QuotedRampCouponApplicabilityMapping `json:"coupon_applicability_mappings"`
	Object                      string                                  `json:"object"`
}

// sub resources
type QuotedRampLineItem struct {
	ItemPriceId                               string                              `json:"item_price_id"`
	ItemType                                  ItemType                            `json:"item_type"`
	Quantity                                  int32                               `json:"quantity"`
	QuantityInDecimal                         string                              `json:"quantity_in_decimal"`
	MeteredQuantity                           string                              `json:"metered_quantity"`
	UnitPrice                                 int64                               `json:"unit_price"`
	UnitPriceInDecimal                        string                              `json:"unit_price_in_decimal"`
	Amount                                    int64                               `json:"amount"`
	AmountInDecimal                           string                              `json:"amount_in_decimal"`
	BillingPeriod                             int32                               `json:"billing_period"`
	BillingPeriodUnit                         QuotedRampLineItemBillingPeriodUnit `json:"billing_period_unit"`
	FreeQuantity                              int32                               `json:"free_quantity"`
	FreeQuantityInDecimal                     string                              `json:"free_quantity_in_decimal"`
	BillingCycles                             int32                               `json:"billing_cycles"`
	ServicePeriodDays                         int32                               `json:"service_period_days"`
	ChargeOnEvent                             ChargeOnEvent                       `json:"charge_on_event"`
	ChargeOnce                                bool                                `json:"charge_once"`
	ChargeOnOption                            ChargeOnOption                      `json:"charge_on_option"`
	StartDate                                 int64                               `json:"start_date"`
	EndDate                                   int64                               `json:"end_date"`
	RampTierId                                string                              `json:"ramp_tier_id"`
	DiscountPerBillingCycle                   int64                               `json:"discount_per_billing_cycle"`
	DiscountPerBillingCycleInDecimal          string                              `json:"discount_per_billing_cycle_in_decimal"`
	ItemLevelDiscountPerBillingCycle          int64                               `json:"item_level_discount_per_billing_cycle"`
	ItemLevelDiscountPerBillingCycleInDecimal string                              `json:"item_level_discount_per_billing_cycle_in_decimal"`
	AmountPerBillingCycle                     int64                               `json:"amount_per_billing_cycle"`
	AmountPerBillingCycleInDecimal            string                              `json:"amount_per_billing_cycle_in_decimal"`
	NetAmountPerBillingCycle                  int64                               `json:"net_amount_per_billing_cycle"`
	NetAmountPerBillingCycleInDecimal         string                              `json:"net_amount_per_billing_cycle_in_decimal"`
	Object                                    string                              `json:"object"`
}

type QuotedRampDiscount struct {
	Id            string                       `json:"id"`
	InvoiceName   string                       `json:"invoice_name"`
	Type          QuotedRampDiscountType       `json:"type"`
	Percentage    float64                      `json:"percentage"`
	Amount        int64                        `json:"amount"`
	DurationType  DurationType                 `json:"duration_type"`
	EntityType    QuotedRampDiscountEntityType `json:"entity_type"`
	EntityId      string                       `json:"entity_id"`
	Period        int32                        `json:"period"`
	PeriodUnit    PeriodUnit                   `json:"period_unit"`
	IncludedInMrr bool                         `json:"included_in_mrr"`
	ApplyOn       ApplyOn                      `json:"apply_on"`
	ItemPriceId   string                       `json:"item_price_id"`
	CreatedAt     int64                        `json:"created_at"`
	UpdatedAt     int64                        `json:"updated_at"`
	StartDate     int64                        `json:"start_date"`
	EndDate       int64                        `json:"end_date"`
	Object        string                       `json:"object"`
}

type QuotedRampItemTier struct {
	ItemPriceId           string      `json:"item_price_id"`
	StartingUnit          int32       `json:"starting_unit"`
	EndingUnit            int32       `json:"ending_unit"`
	Price                 int64       `json:"price"`
	StartingUnitInDecimal string      `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string      `json:"ending_unit_in_decimal"`
	PriceInDecimal        string      `json:"price_in_decimal"`
	RampTierId            string      `json:"ramp_tier_id"`
	PricingType           PricingType `json:"pricing_type"`
	PackageSize           int32       `json:"package_size"`
	Object                string      `json:"object"`
}

type QuotedRampCouponApplicabilityMapping struct {
	CouponId               string   `json:"coupon_id"`
	ApplicableItemPriceIds []string `json:"applicable_item_price_ids"`
	Object                 string   `json:"object"`
}

// operations
// input params
