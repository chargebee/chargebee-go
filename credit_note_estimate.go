package chargebee

type CreditNoteEstimateType string

const (
	CreditNoteEstimateTypeAdjustment CreditNoteEstimateType = "adjustment"
	CreditNoteEstimateTypeRefundable CreditNoteEstimateType = "refundable"
	CreditNoteEstimateTypeStore      CreditNoteEstimateType = "store"
)

type CreditNoteEstimatePriceType string

const (
	CreditNoteEstimatePriceTypeTaxExclusive CreditNoteEstimatePriceType = "tax_exclusive"
	CreditNoteEstimatePriceTypeTaxInclusive CreditNoteEstimatePriceType = "tax_inclusive"
)

type CreditNoteEstimateLineItemPricingModel string

const (
	CreditNoteEstimateLineItemPricingModelFlatFee   CreditNoteEstimateLineItemPricingModel = "flat_fee"
	CreditNoteEstimateLineItemPricingModelPerUnit   CreditNoteEstimateLineItemPricingModel = "per_unit"
	CreditNoteEstimateLineItemPricingModelTiered    CreditNoteEstimateLineItemPricingModel = "tiered"
	CreditNoteEstimateLineItemPricingModelVolume    CreditNoteEstimateLineItemPricingModel = "volume"
	CreditNoteEstimateLineItemPricingModelStairstep CreditNoteEstimateLineItemPricingModel = "stairstep"
)

type CreditNoteEstimateLineItemEntityType string

const (
	CreditNoteEstimateLineItemEntityTypeAdhoc           CreditNoteEstimateLineItemEntityType = "adhoc"
	CreditNoteEstimateLineItemEntityTypePlanItemPrice   CreditNoteEstimateLineItemEntityType = "plan_item_price"
	CreditNoteEstimateLineItemEntityTypeAddonItemPrice  CreditNoteEstimateLineItemEntityType = "addon_item_price"
	CreditNoteEstimateLineItemEntityTypeChargeItemPrice CreditNoteEstimateLineItemEntityType = "charge_item_price"
	CreditNoteEstimateLineItemEntityTypePlanSetup       CreditNoteEstimateLineItemEntityType = "plan_setup"
	CreditNoteEstimateLineItemEntityTypePlan            CreditNoteEstimateLineItemEntityType = "plan"
	CreditNoteEstimateLineItemEntityTypeAddon           CreditNoteEstimateLineItemEntityType = "addon"
)

type CreditNoteEstimateLineItemTaxExemptReason string

const (
	CreditNoteEstimateLineItemTaxExemptReasonTaxNotConfigured                 CreditNoteEstimateLineItemTaxExemptReason = "tax_not_configured"
	CreditNoteEstimateLineItemTaxExemptReasonRegionNonTaxable                 CreditNoteEstimateLineItemTaxExemptReason = "region_non_taxable"
	CreditNoteEstimateLineItemTaxExemptReasonExport                           CreditNoteEstimateLineItemTaxExemptReason = "export"
	CreditNoteEstimateLineItemTaxExemptReasonCustomerExempt                   CreditNoteEstimateLineItemTaxExemptReason = "customer_exempt"
	CreditNoteEstimateLineItemTaxExemptReasonProductExempt                    CreditNoteEstimateLineItemTaxExemptReason = "product_exempt"
	CreditNoteEstimateLineItemTaxExemptReasonZeroRated                        CreditNoteEstimateLineItemTaxExemptReason = "zero_rated"
	CreditNoteEstimateLineItemTaxExemptReasonReverseCharge                    CreditNoteEstimateLineItemTaxExemptReason = "reverse_charge"
	CreditNoteEstimateLineItemTaxExemptReasonHighValuePhysicalGoods           CreditNoteEstimateLineItemTaxExemptReason = "high_value_physical_goods"
	CreditNoteEstimateLineItemTaxExemptReasonZeroValueItem                    CreditNoteEstimateLineItemTaxExemptReason = "zero_value_item"
	CreditNoteEstimateLineItemTaxExemptReasonTaxNotConfiguredExternalProvider CreditNoteEstimateLineItemTaxExemptReason = "tax_not_configured_external_provider"
)

type CreditNoteEstimateLineItemTierPricingType string

const (
	CreditNoteEstimateLineItemTierPricingTypePerUnit CreditNoteEstimateLineItemTierPricingType = "per_unit"
	CreditNoteEstimateLineItemTierPricingTypeFlatFee CreditNoteEstimateLineItemTierPricingType = "flat_fee"
	CreditNoteEstimateLineItemTierPricingTypePackage CreditNoteEstimateLineItemTierPricingType = "package"
)

type CreditNoteEstimateLineItemDiscountDiscountType string

const (
	CreditNoteEstimateLineItemDiscountDiscountTypeItemLevelCoupon       CreditNoteEstimateLineItemDiscountDiscountType = "item_level_coupon"
	CreditNoteEstimateLineItemDiscountDiscountTypeDocumentLevelCoupon   CreditNoteEstimateLineItemDiscountDiscountType = "document_level_coupon"
	CreditNoteEstimateLineItemDiscountDiscountTypePromotionalCredits    CreditNoteEstimateLineItemDiscountDiscountType = "promotional_credits"
	CreditNoteEstimateLineItemDiscountDiscountTypeProratedCredits       CreditNoteEstimateLineItemDiscountDiscountType = "prorated_credits"
	CreditNoteEstimateLineItemDiscountDiscountTypeItemLevelDiscount     CreditNoteEstimateLineItemDiscountDiscountType = "item_level_discount"
	CreditNoteEstimateLineItemDiscountDiscountTypeDocumentLevelDiscount CreditNoteEstimateLineItemDiscountDiscountType = "document_level_discount"
)

type CreditNoteEstimateLineItemTaxTaxJurisType string

const (
	CreditNoteEstimateLineItemTaxTaxJurisTypeCountry        CreditNoteEstimateLineItemTaxTaxJurisType = "country"
	CreditNoteEstimateLineItemTaxTaxJurisTypeFederal        CreditNoteEstimateLineItemTaxTaxJurisType = "federal"
	CreditNoteEstimateLineItemTaxTaxJurisTypeState          CreditNoteEstimateLineItemTaxTaxJurisType = "state"
	CreditNoteEstimateLineItemTaxTaxJurisTypeCounty         CreditNoteEstimateLineItemTaxTaxJurisType = "county"
	CreditNoteEstimateLineItemTaxTaxJurisTypeCity           CreditNoteEstimateLineItemTaxTaxJurisType = "city"
	CreditNoteEstimateLineItemTaxTaxJurisTypeSpecial        CreditNoteEstimateLineItemTaxTaxJurisType = "special"
	CreditNoteEstimateLineItemTaxTaxJurisTypeUnincorporated CreditNoteEstimateLineItemTaxTaxJurisType = "unincorporated"
	CreditNoteEstimateLineItemTaxTaxJurisTypeOther          CreditNoteEstimateLineItemTaxTaxJurisType = "other"
)

type CreditNoteEstimateDiscountEntityType string

const (
	CreditNoteEstimateDiscountEntityTypeItemLevelCoupon       CreditNoteEstimateDiscountEntityType = "item_level_coupon"
	CreditNoteEstimateDiscountEntityTypeDocumentLevelCoupon   CreditNoteEstimateDiscountEntityType = "document_level_coupon"
	CreditNoteEstimateDiscountEntityTypePromotionalCredits    CreditNoteEstimateDiscountEntityType = "promotional_credits"
	CreditNoteEstimateDiscountEntityTypeProratedCredits       CreditNoteEstimateDiscountEntityType = "prorated_credits"
	CreditNoteEstimateDiscountEntityTypeItemLevelDiscount     CreditNoteEstimateDiscountEntityType = "item_level_discount"
	CreditNoteEstimateDiscountEntityTypeDocumentLevelDiscount CreditNoteEstimateDiscountEntityType = "document_level_discount"
)

type CreditNoteEstimateDiscountDiscountType string

const (
	CreditNoteEstimateDiscountDiscountTypeFixedAmount CreditNoteEstimateDiscountDiscountType = "fixed_amount"
	CreditNoteEstimateDiscountDiscountTypePercentage  CreditNoteEstimateDiscountDiscountType = "percentage"
)

// just struct
type CreditNoteEstimate struct {
	ReferenceInvoiceId string                                `json:"reference_invoice_id"`
	Type               CreditNoteEstimateType                `json:"type"`
	PriceType          CreditNoteEstimatePriceType           `json:"price_type"`
	CurrencyCode       string                                `json:"currency_code"`
	SubTotal           int64                                 `json:"sub_total"`
	Total              int64                                 `json:"total"`
	AmountAllocated    int64                                 `json:"amount_allocated"`
	AmountAvailable    int64                                 `json:"amount_available"`
	LineItems          []*CreditNoteEstimateLineItem         `json:"line_items"`
	LineItemTiers      []*CreditNoteEstimateLineItemTier     `json:"line_item_tiers"`
	LineItemDiscounts  []*CreditNoteEstimateLineItemDiscount `json:"line_item_discounts"`
	LineItemTaxes      []*CreditNoteEstimateLineItemTax      `json:"line_item_taxes"`
	Discounts          []*CreditNoteEstimateDiscount         `json:"discounts"`
	Taxes              []*CreditNoteEstimateTax              `json:"taxes"`
	RoundOffAmount     int64                                 `json:"round_off_amount"`
	CustomerId         string                                `json:"customer_id"`
	Object             string                                `json:"object"`
}

// sub resources
type CreditNoteEstimateLineItem struct {
	Id                      string                                    `json:"id"`
	SubscriptionId          string                                    `json:"subscription_id"`
	DateFrom                int64                                     `json:"date_from"`
	DateTo                  int64                                     `json:"date_to"`
	UnitAmount              int64                                     `json:"unit_amount"`
	Quantity                int32                                     `json:"quantity"`
	Amount                  int64                                     `json:"amount"`
	PricingModel            CreditNoteEstimateLineItemPricingModel    `json:"pricing_model"`
	IsTaxed                 bool                                      `json:"is_taxed"`
	TaxAmount               int64                                     `json:"tax_amount"`
	TaxRate                 float64                                   `json:"tax_rate"`
	UnitAmountInDecimal     string                                    `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string                                    `json:"quantity_in_decimal"`
	AmountInDecimal         string                                    `json:"amount_in_decimal"`
	DiscountAmount          int64                                     `json:"discount_amount"`
	ItemLevelDiscountAmount int64                                     `json:"item_level_discount_amount"`
	Metered                 bool                                      `json:"metered"`
	IsPercentagePricing     bool                                      `json:"is_percentage_pricing"`
	ReferenceLineItemId     string                                    `json:"reference_line_item_id"`
	Description             string                                    `json:"description"`
	EntityDescription       string                                    `json:"entity_description"`
	EntityType              CreditNoteEstimateLineItemEntityType      `json:"entity_type"`
	TaxExemptReason         CreditNoteEstimateLineItemTaxExemptReason `json:"tax_exempt_reason"`
	EntityId                string                                    `json:"entity_id"`
	CustomerId              string                                    `json:"customer_id"`
	Object                  string                                    `json:"object"`
}
type CreditNoteEstimateLineItemTier struct {
	LineItemId            string                                    `json:"line_item_id"`
	StartingUnit          int32                                     `json:"starting_unit"`
	EndingUnit            int32                                     `json:"ending_unit"`
	QuantityUsed          int32                                     `json:"quantity_used"`
	UnitAmount            int64                                     `json:"unit_amount"`
	StartingUnitInDecimal string                                    `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string                                    `json:"ending_unit_in_decimal"`
	QuantityUsedInDecimal string                                    `json:"quantity_used_in_decimal"`
	UnitAmountInDecimal   string                                    `json:"unit_amount_in_decimal"`
	PricingType           CreditNoteEstimateLineItemTierPricingType `json:"pricing_type"`
	PackageSize           int32                                     `json:"package_size"`
	Object                string                                    `json:"object"`
}
type CreditNoteEstimateLineItemDiscount struct {
	LineItemId     string                                         `json:"line_item_id"`
	DiscountType   CreditNoteEstimateLineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                         `json:"coupon_id"`
	EntityId       string                                         `json:"entity_id"`
	DiscountAmount int64                                          `json:"discount_amount"`
	Object         string                                         `json:"object"`
}
type CreditNoteEstimateLineItemTax struct {
	LineItemId               string                                    `json:"line_item_id"`
	TaxName                  string                                    `json:"tax_name"`
	TaxRate                  float64                                   `json:"tax_rate"`
	DateTo                   int64                                     `json:"date_to"`
	DateFrom                 int64                                     `json:"date_from"`
	ProratedTaxableAmount    float64                                   `json:"prorated_taxable_amount"`
	IsPartialTaxApplied      bool                                      `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool                                      `json:"is_non_compliance_tax"`
	TaxableAmount            int64                                     `json:"taxable_amount"`
	TaxAmount                int64                                     `json:"tax_amount"`
	TaxJurisType             CreditNoteEstimateLineItemTaxTaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string                                    `json:"tax_juris_name"`
	TaxJurisCode             string                                    `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int64                                     `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string                                    `json:"local_currency_code"`
	Object                   string                                    `json:"object"`
}
type CreditNoteEstimateDiscount struct {
	Amount        int64                                  `json:"amount"`
	Description   string                                 `json:"description"`
	LineItemId    string                                 `json:"line_item_id"`
	EntityType    CreditNoteEstimateDiscountEntityType   `json:"entity_type"`
	DiscountType  CreditNoteEstimateDiscountDiscountType `json:"discount_type"`
	EntityId      string                                 `json:"entity_id"`
	CouponSetCode string                                 `json:"coupon_set_code"`
	Object        string                                 `json:"object"`
}
type CreditNoteEstimateTax struct {
	Name        string `json:"name"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}

// operations
// input params
