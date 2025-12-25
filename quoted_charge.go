package chargebee

type QuotedChargeChargeAvalaraSaleType string

const (
	QuotedChargeChargeAvalaraSaleTypeWholesale QuotedChargeChargeAvalaraSaleType = "wholesale"
	QuotedChargeChargeAvalaraSaleTypeRetail    QuotedChargeChargeAvalaraSaleType = "retail"
	QuotedChargeChargeAvalaraSaleTypeConsumed  QuotedChargeChargeAvalaraSaleType = "consumed"
	QuotedChargeChargeAvalaraSaleTypeVendorUse QuotedChargeChargeAvalaraSaleType = "vendor_use"
)

type QuotedChargeItemTierPricingType string

const (
	QuotedChargeItemTierPricingTypePerUnit QuotedChargeItemTierPricingType = "per_unit"
	QuotedChargeItemTierPricingTypeFlatFee QuotedChargeItemTierPricingType = "flat_fee"
	QuotedChargeItemTierPricingTypePackage QuotedChargeItemTierPricingType = "package"
)

type QuotedChargeAddonProrationType string

const (
	QuotedChargeAddonProrationTypeFullTerm    QuotedChargeAddonProrationType = "full_term"
	QuotedChargeAddonProrationTypePartialTerm QuotedChargeAddonProrationType = "partial_term"
	QuotedChargeAddonProrationTypeNone        QuotedChargeAddonProrationType = "none"
)

// just struct
type QuotedCharge struct {
	Charges                     []*QuotedChargeCharge                     `json:"charges"`
	Addons                      []*QuotedChargeAddon                      `json:"addons"`
	InvoiceItems                []*QuotedChargeInvoiceItem                `json:"invoice_items"`
	ItemTiers                   []*QuotedChargeItemTier                   `json:"item_tiers"`
	Coupons                     []*QuotedChargeCoupon                     `json:"coupons"`
	CouponApplicabilityMappings []*QuotedChargeCouponApplicabilityMapping `json:"coupon_applicability_mappings"`
	Object                      string                                    `json:"object"`
}

// sub resources
type QuotedChargeCharge struct {
	Amount                 int64                             `json:"amount"`
	AmountInDecimal        string                            `json:"amount_in_decimal"`
	Description            string                            `json:"description"`
	ServicePeriodInDays    int32                             `json:"service_period_in_days"`
	AvalaraSaleType        QuotedChargeChargeAvalaraSaleType `json:"avalara_sale_type"`
	AvalaraTransactionType int32                             `json:"avalara_transaction_type"`
	AvalaraServiceType     int32                             `json:"avalara_service_type"`
	Object                 string                            `json:"object"`
}

type QuotedChargeInvoiceItem struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           int32  `json:"quantity"`
	QuantityInDecimal  string `json:"quantity_in_decimal"`
	UnitPrice          int64  `json:"unit_price"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal"`
	ServicePeriodDays  int32  `json:"service_period_days"`
	Object             string `json:"object"`
}

type QuotedChargeItemTier struct {
	ItemPriceId           string                          `json:"item_price_id"`
	StartingUnit          int32                           `json:"starting_unit"`
	EndingUnit            int32                           `json:"ending_unit"`
	Price                 int64                           `json:"price"`
	StartingUnitInDecimal string                          `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string                          `json:"ending_unit_in_decimal"`
	PriceInDecimal        string                          `json:"price_in_decimal"`
	PricingType           QuotedChargeItemTierPricingType `json:"pricing_type"`
	PackageSize           int32                           `json:"package_size"`
	Index                 int32                           `json:"index"`
	Object                string                          `json:"object"`
}

type QuotedChargeCoupon struct {
	CouponId string `json:"coupon_id"`
	Object   string `json:"object"`
}

type QuotedChargeCouponApplicabilityMapping struct {
	CouponId               string   `json:"coupon_id"`
	ApplicableItemPriceIds []string `json:"applicable_item_price_ids"`
	Object                 string   `json:"object"`
}

type QuotedChargeAddon struct {
	Id                 string                         `json:"id"`
	Quantity           int32                          `json:"quantity"`
	UnitPrice          int64                          `json:"unit_price"`
	QuantityInDecimal  string                         `json:"quantity_in_decimal"`
	UnitPriceInDecimal string                         `json:"unit_price_in_decimal"`
	ProrationType      QuotedChargeAddonProrationType `json:"proration_type"`
	ServicePeriod      int32                          `json:"service_period"`
	Object             string                         `json:"object"`
}

// operations
// input params
