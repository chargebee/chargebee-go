package chargebee

type ChargeEvent string

const (
	ChargeEventImmediate            ChargeEvent = "immediate"
	ChargeEventSubscriptionCreation ChargeEvent = "subscription_creation"
	ChargeEventTrialStart           ChargeEvent = "trial_start"
	ChargeEventSubscriptionChange   ChargeEvent = "subscription_change"
	ChargeEventSubscriptionRenewal  ChargeEvent = "subscription_renewal"
	ChargeEventSubscriptionCancel   ChargeEvent = "subscription_cancel"
)

type LineItemEntityType string

const (
	LineItemEntityTypeAdhoc           LineItemEntityType = "adhoc"
	LineItemEntityTypePlanItemPrice   LineItemEntityType = "plan_item_price"
	LineItemEntityTypeAddonItemPrice  LineItemEntityType = "addon_item_price"
	LineItemEntityTypeChargeItemPrice LineItemEntityType = "charge_item_price"
	LineItemEntityTypePlanSetup       LineItemEntityType = "plan_setup"
	LineItemEntityTypePlan            LineItemEntityType = "plan"
	LineItemEntityTypeAddon           LineItemEntityType = "addon"
)

type LineItemDiscountDiscountType string

const (
	LineItemDiscountDiscountTypeItemLevelCoupon       LineItemDiscountDiscountType = "item_level_coupon"
	LineItemDiscountDiscountTypeDocumentLevelCoupon   LineItemDiscountDiscountType = "document_level_coupon"
	LineItemDiscountDiscountTypePromotionalCredits    LineItemDiscountDiscountType = "promotional_credits"
	LineItemDiscountDiscountTypeProratedCredits       LineItemDiscountDiscountType = "prorated_credits"
	LineItemDiscountDiscountTypeItemLevelDiscount     LineItemDiscountDiscountType = "item_level_discount"
	LineItemDiscountDiscountTypeDocumentLevelDiscount LineItemDiscountDiscountType = "document_level_discount"
)

type DiscountEntityType string

const (
	DiscountEntityTypeItemLevelCoupon       DiscountEntityType = "item_level_coupon"
	DiscountEntityTypeDocumentLevelCoupon   DiscountEntityType = "document_level_coupon"
	DiscountEntityTypePromotionalCredits    DiscountEntityType = "promotional_credits"
	DiscountEntityTypeProratedCredits       DiscountEntityType = "prorated_credits"
	DiscountEntityTypeItemLevelDiscount     DiscountEntityType = "item_level_discount"
	DiscountEntityTypeDocumentLevelDiscount DiscountEntityType = "document_level_discount"
)

type DiscountDiscountType string

const (
	DiscountDiscountTypeFixedAmount DiscountDiscountType = "fixed_amount"
	DiscountDiscountTypePercentage  DiscountDiscountType = "percentage"
)

type QuoteLineGroup struct {
	Version            int32               `json:"version"`
	Id                 string              `json:"id"`
	SubTotal           int64               `json:"sub_total"`
	Total              int64               `json:"total"`
	CreditsApplied     int64               `json:"credits_applied"`
	AmountPaid         int64               `json:"amount_paid"`
	AmountDue          int64               `json:"amount_due"`
	ChargeEvent        ChargeEvent         `json:"charge_event"`
	BillingCycleNumber int32               `json:"billing_cycle_number"`
	LineItems          []*LineItem         `json:"line_items"`
	LineItemDiscounts  []*LineItemDiscount `json:"line_item_discounts"`
	LineItemTaxes      []*LineItemTax      `json:"line_item_taxes"`
	Discounts          []*Discount         `json:"discounts"`
	Taxes              []*Tax              `json:"taxes"`
	Object             string              `json:"object"`
}
type LineItem struct {
	Id                      string               `json:"id"`
	SubscriptionId          string               `json:"subscription_id"`
	DateFrom                int64                `json:"date_from"`
	DateTo                  int64                `json:"date_to"`
	UnitAmount              int64                `json:"unit_amount"`
	Quantity                int32                `json:"quantity"`
	Amount                  int64                `json:"amount"`
	PricingModel            enum.PricingModel    `json:"pricing_model"`
	IsTaxed                 bool                 `json:"is_taxed"`
	TaxAmount               int64                `json:"tax_amount"`
	TaxRate                 float64              `json:"tax_rate"`
	UnitAmountInDecimal     string               `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string               `json:"quantity_in_decimal"`
	AmountInDecimal         string               `json:"amount_in_decimal"`
	DiscountAmount          int64                `json:"discount_amount"`
	ItemLevelDiscountAmount int64                `json:"item_level_discount_amount"`
	Metered                 bool                 `json:"metered"`
	IsPercentagePricing     bool                 `json:"is_percentage_pricing"`
	ReferenceLineItemId     string               `json:"reference_line_item_id"`
	Description             string               `json:"description"`
	EntityDescription       string               `json:"entity_description"`
	EntityType              LineItemEntityType   `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason `json:"tax_exempt_reason"`
	EntityId                string               `json:"entity_id"`
	CustomerId              string               `json:"customer_id"`
	Object                  string               `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                       `json:"line_item_id"`
	DiscountType   LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                       `json:"coupon_id"`
	EntityId       string                       `json:"entity_id"`
	DiscountAmount int64                        `json:"discount_amount"`
	Object         string                       `json:"object"`
}
type LineItemTax struct {
	LineItemId               string            `json:"line_item_id"`
	TaxName                  string            `json:"tax_name"`
	TaxRate                  float64           `json:"tax_rate"`
	DateTo                   int64             `json:"date_to"`
	DateFrom                 int64             `json:"date_from"`
	ProratedTaxableAmount    float64           `json:"prorated_taxable_amount"`
	IsPartialTaxApplied      bool              `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool              `json:"is_non_compliance_tax"`
	TaxableAmount            int64             `json:"taxable_amount"`
	TaxAmount                int64             `json:"tax_amount"`
	TaxJurisType             enum.TaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string            `json:"tax_juris_name"`
	TaxJurisCode             string            `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int64             `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string            `json:"local_currency_code"`
	Object                   string            `json:"object"`
}
type Discount struct {
	Amount        int64                `json:"amount"`
	Description   string               `json:"description"`
	LineItemId    string               `json:"line_item_id"`
	EntityType    DiscountEntityType   `json:"entity_type"`
	DiscountType  DiscountDiscountType `json:"discount_type"`
	EntityId      string               `json:"entity_id"`
	CouponSetCode string               `json:"coupon_set_code"`
	Object        string               `json:"object"`
}
type Tax struct {
	Name        string `json:"name"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}
