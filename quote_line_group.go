package chargebee

type QuoteLineGroupChargeEvent string

const (
	QuoteLineGroupChargeEventImmediate            QuoteLineGroupChargeEvent = "immediate"
	QuoteLineGroupChargeEventSubscriptionCreation QuoteLineGroupChargeEvent = "subscription_creation"
	QuoteLineGroupChargeEventTrialStart           QuoteLineGroupChargeEvent = "trial_start"
	QuoteLineGroupChargeEventSubscriptionChange   QuoteLineGroupChargeEvent = "subscription_change"
	QuoteLineGroupChargeEventSubscriptionRenewal  QuoteLineGroupChargeEvent = "subscription_renewal"
	QuoteLineGroupChargeEventSubscriptionCancel   QuoteLineGroupChargeEvent = "subscription_cancel"
)

type QuoteLineGroupLineItemPricingModel string

const (
	QuoteLineGroupLineItemPricingModelFlatFee   QuoteLineGroupLineItemPricingModel = "flat_fee"
	QuoteLineGroupLineItemPricingModelPerUnit   QuoteLineGroupLineItemPricingModel = "per_unit"
	QuoteLineGroupLineItemPricingModelTiered    QuoteLineGroupLineItemPricingModel = "tiered"
	QuoteLineGroupLineItemPricingModelVolume    QuoteLineGroupLineItemPricingModel = "volume"
	QuoteLineGroupLineItemPricingModelStairstep QuoteLineGroupLineItemPricingModel = "stairstep"
)

type QuoteLineGroupLineItemEntityType string

const (
	QuoteLineGroupLineItemEntityTypeAdhoc           QuoteLineGroupLineItemEntityType = "adhoc"
	QuoteLineGroupLineItemEntityTypePlanItemPrice   QuoteLineGroupLineItemEntityType = "plan_item_price"
	QuoteLineGroupLineItemEntityTypeAddonItemPrice  QuoteLineGroupLineItemEntityType = "addon_item_price"
	QuoteLineGroupLineItemEntityTypeChargeItemPrice QuoteLineGroupLineItemEntityType = "charge_item_price"
	QuoteLineGroupLineItemEntityTypePlanSetup       QuoteLineGroupLineItemEntityType = "plan_setup"
	QuoteLineGroupLineItemEntityTypePlan            QuoteLineGroupLineItemEntityType = "plan"
	QuoteLineGroupLineItemEntityTypeAddon           QuoteLineGroupLineItemEntityType = "addon"
)

type QuoteLineGroupLineItemTaxExemptReason string

const (
	QuoteLineGroupLineItemTaxExemptReasonTaxNotConfigured                 QuoteLineGroupLineItemTaxExemptReason = "tax_not_configured"
	QuoteLineGroupLineItemTaxExemptReasonRegionNonTaxable                 QuoteLineGroupLineItemTaxExemptReason = "region_non_taxable"
	QuoteLineGroupLineItemTaxExemptReasonExport                           QuoteLineGroupLineItemTaxExemptReason = "export"
	QuoteLineGroupLineItemTaxExemptReasonCustomerExempt                   QuoteLineGroupLineItemTaxExemptReason = "customer_exempt"
	QuoteLineGroupLineItemTaxExemptReasonProductExempt                    QuoteLineGroupLineItemTaxExemptReason = "product_exempt"
	QuoteLineGroupLineItemTaxExemptReasonZeroRated                        QuoteLineGroupLineItemTaxExemptReason = "zero_rated"
	QuoteLineGroupLineItemTaxExemptReasonReverseCharge                    QuoteLineGroupLineItemTaxExemptReason = "reverse_charge"
	QuoteLineGroupLineItemTaxExemptReasonHighValuePhysicalGoods           QuoteLineGroupLineItemTaxExemptReason = "high_value_physical_goods"
	QuoteLineGroupLineItemTaxExemptReasonZeroValueItem                    QuoteLineGroupLineItemTaxExemptReason = "zero_value_item"
	QuoteLineGroupLineItemTaxExemptReasonTaxNotConfiguredExternalProvider QuoteLineGroupLineItemTaxExemptReason = "tax_not_configured_external_provider"
)

type QuoteLineGroupLineItemDiscountDiscountType string

const (
	QuoteLineGroupLineItemDiscountDiscountTypeItemLevelCoupon       QuoteLineGroupLineItemDiscountDiscountType = "item_level_coupon"
	QuoteLineGroupLineItemDiscountDiscountTypeDocumentLevelCoupon   QuoteLineGroupLineItemDiscountDiscountType = "document_level_coupon"
	QuoteLineGroupLineItemDiscountDiscountTypePromotionalCredits    QuoteLineGroupLineItemDiscountDiscountType = "promotional_credits"
	QuoteLineGroupLineItemDiscountDiscountTypeProratedCredits       QuoteLineGroupLineItemDiscountDiscountType = "prorated_credits"
	QuoteLineGroupLineItemDiscountDiscountTypeItemLevelDiscount     QuoteLineGroupLineItemDiscountDiscountType = "item_level_discount"
	QuoteLineGroupLineItemDiscountDiscountTypeDocumentLevelDiscount QuoteLineGroupLineItemDiscountDiscountType = "document_level_discount"
)

type QuoteLineGroupLineItemTaxTaxJurisType string

const (
	QuoteLineGroupLineItemTaxTaxJurisTypeCountry        QuoteLineGroupLineItemTaxTaxJurisType = "country"
	QuoteLineGroupLineItemTaxTaxJurisTypeFederal        QuoteLineGroupLineItemTaxTaxJurisType = "federal"
	QuoteLineGroupLineItemTaxTaxJurisTypeState          QuoteLineGroupLineItemTaxTaxJurisType = "state"
	QuoteLineGroupLineItemTaxTaxJurisTypeCounty         QuoteLineGroupLineItemTaxTaxJurisType = "county"
	QuoteLineGroupLineItemTaxTaxJurisTypeCity           QuoteLineGroupLineItemTaxTaxJurisType = "city"
	QuoteLineGroupLineItemTaxTaxJurisTypeSpecial        QuoteLineGroupLineItemTaxTaxJurisType = "special"
	QuoteLineGroupLineItemTaxTaxJurisTypeUnincorporated QuoteLineGroupLineItemTaxTaxJurisType = "unincorporated"
	QuoteLineGroupLineItemTaxTaxJurisTypeOther          QuoteLineGroupLineItemTaxTaxJurisType = "other"
)

type QuoteLineGroupDiscountEntityType string

const (
	QuoteLineGroupDiscountEntityTypeItemLevelCoupon       QuoteLineGroupDiscountEntityType = "item_level_coupon"
	QuoteLineGroupDiscountEntityTypeDocumentLevelCoupon   QuoteLineGroupDiscountEntityType = "document_level_coupon"
	QuoteLineGroupDiscountEntityTypePromotionalCredits    QuoteLineGroupDiscountEntityType = "promotional_credits"
	QuoteLineGroupDiscountEntityTypeProratedCredits       QuoteLineGroupDiscountEntityType = "prorated_credits"
	QuoteLineGroupDiscountEntityTypeItemLevelDiscount     QuoteLineGroupDiscountEntityType = "item_level_discount"
	QuoteLineGroupDiscountEntityTypeDocumentLevelDiscount QuoteLineGroupDiscountEntityType = "document_level_discount"
)

type QuoteLineGroupDiscountDiscountType string

const (
	QuoteLineGroupDiscountDiscountTypeFixedAmount QuoteLineGroupDiscountDiscountType = "fixed_amount"
	QuoteLineGroupDiscountDiscountTypePercentage  QuoteLineGroupDiscountDiscountType = "percentage"
)

// just struct
type QuoteLineGroup struct {
	Version            int32                             `json:"version"`
	Id                 string                            `json:"id"`
	SubTotal           int64                             `json:"sub_total"`
	Total              int64                             `json:"total"`
	CreditsApplied     int64                             `json:"credits_applied"`
	AmountPaid         int64                             `json:"amount_paid"`
	AmountDue          int64                             `json:"amount_due"`
	ChargeEvent        QuoteLineGroupChargeEvent         `json:"charge_event"`
	BillingCycleNumber int32                             `json:"billing_cycle_number"`
	LineItems          []*QuoteLineGroupLineItem         `json:"line_items"`
	LineItemDiscounts  []*QuoteLineGroupLineItemDiscount `json:"line_item_discounts"`
	LineItemTaxes      []*QuoteLineGroupLineItemTax      `json:"line_item_taxes"`
	Discounts          []*QuoteLineGroupDiscount         `json:"discounts"`
	Taxes              []*QuoteLineGroupTax              `json:"taxes"`
	Object             string                            `json:"object"`
}

// sub resources
type QuoteLineGroupLineItem struct {
	Id                      string                                `json:"id"`
	SubscriptionId          string                                `json:"subscription_id"`
	DateFrom                int64                                 `json:"date_from"`
	DateTo                  int64                                 `json:"date_to"`
	UnitAmount              int64                                 `json:"unit_amount"`
	Quantity                int32                                 `json:"quantity"`
	Amount                  int64                                 `json:"amount"`
	PricingModel            QuoteLineGroupLineItemPricingModel    `json:"pricing_model"`
	IsTaxed                 bool                                  `json:"is_taxed"`
	TaxAmount               int64                                 `json:"tax_amount"`
	TaxRate                 float64                               `json:"tax_rate"`
	UnitAmountInDecimal     string                                `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string                                `json:"quantity_in_decimal"`
	AmountInDecimal         string                                `json:"amount_in_decimal"`
	DiscountAmount          int64                                 `json:"discount_amount"`
	ItemLevelDiscountAmount int64                                 `json:"item_level_discount_amount"`
	Metered                 bool                                  `json:"metered"`
	IsPercentagePricing     bool                                  `json:"is_percentage_pricing"`
	ReferenceLineItemId     string                                `json:"reference_line_item_id"`
	Description             string                                `json:"description"`
	EntityDescription       string                                `json:"entity_description"`
	EntityType              QuoteLineGroupLineItemEntityType      `json:"entity_type"`
	TaxExemptReason         QuoteLineGroupLineItemTaxExemptReason `json:"tax_exempt_reason"`
	EntityId                string                                `json:"entity_id"`
	CustomerId              string                                `json:"customer_id"`
	Object                  string                                `json:"object"`
}

type QuoteLineGroupLineItemDiscount struct {
	LineItemId     string                                     `json:"line_item_id"`
	DiscountType   QuoteLineGroupLineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                     `json:"coupon_id"`
	EntityId       string                                     `json:"entity_id"`
	DiscountAmount int64                                      `json:"discount_amount"`
	Object         string                                     `json:"object"`
}

type QuoteLineGroupLineItemTax struct {
	LineItemId               string                                `json:"line_item_id"`
	TaxName                  string                                `json:"tax_name"`
	TaxRate                  float64                               `json:"tax_rate"`
	DateTo                   int64                                 `json:"date_to"`
	DateFrom                 int64                                 `json:"date_from"`
	ProratedTaxableAmount    float64                               `json:"prorated_taxable_amount"`
	IsPartialTaxApplied      bool                                  `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool                                  `json:"is_non_compliance_tax"`
	TaxableAmount            int64                                 `json:"taxable_amount"`
	TaxAmount                int64                                 `json:"tax_amount"`
	TaxJurisType             QuoteLineGroupLineItemTaxTaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string                                `json:"tax_juris_name"`
	TaxJurisCode             string                                `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int64                                 `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string                                `json:"local_currency_code"`
	Object                   string                                `json:"object"`
}

type QuoteLineGroupDiscount struct {
	Amount        int64                              `json:"amount"`
	Description   string                             `json:"description"`
	LineItemId    string                             `json:"line_item_id"`
	EntityType    QuoteLineGroupDiscountEntityType   `json:"entity_type"`
	DiscountType  QuoteLineGroupDiscountDiscountType `json:"discount_type"`
	EntityId      string                             `json:"entity_id"`
	CouponSetCode string                             `json:"coupon_set_code"`
	Object        string                             `json:"object"`
}

type QuoteLineGroupTax struct {
	Name        string `json:"name"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}

// operations
// input params
