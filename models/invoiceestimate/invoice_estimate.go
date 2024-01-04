package invoiceestimate

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	invoiceEstimateEnum "github.com/chargebee/chargebee-go/v3/models/invoiceestimate/enum"
)

type InvoiceEstimate struct {
	Recurring         bool                `json:"recurring"`
	PriceType         enum.PriceType      `json:"price_type"`
	CurrencyCode      string              `json:"currency_code"`
	SubTotal          int64               `json:"sub_total"`
	Total             int64               `json:"total"`
	CreditsApplied    int64               `json:"credits_applied"`
	AmountPaid        int64               `json:"amount_paid"`
	AmountDue         int64               `json:"amount_due"`
	LineItems         []*LineItem         `json:"line_items"`
	Discounts         []*Discount         `json:"discounts"`
	Taxes             []*Tax              `json:"taxes"`
	LineItemTaxes     []*LineItemTax      `json:"line_item_taxes"`
	LineItemTiers     []*LineItemTier     `json:"line_item_tiers"`
	LineItemDiscounts []*LineItemDiscount `json:"line_item_discounts"`
	RoundOffAmount    int64               `json:"round_off_amount"`
	CustomerId        string              `json:"customer_id"`
	Object            string              `json:"object"`
}
type LineItem struct {
	Id                      string                                 `json:"id"`
	SubscriptionId          string                                 `json:"subscription_id"`
	DateFrom                int64                                  `json:"date_from"`
	DateTo                  int64                                  `json:"date_to"`
	UnitAmount              int64                                  `json:"unit_amount"`
	Quantity                int32                                  `json:"quantity"`
	Amount                  int64                                  `json:"amount"`
	PricingModel            enum.PricingModel                      `json:"pricing_model"`
	IsTaxed                 bool                                   `json:"is_taxed"`
	TaxAmount               int64                                  `json:"tax_amount"`
	TaxRate                 float64                                `json:"tax_rate"`
	UnitAmountInDecimal     string                                 `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string                                 `json:"quantity_in_decimal"`
	AmountInDecimal         string                                 `json:"amount_in_decimal"`
	DiscountAmount          int64                                  `json:"discount_amount"`
	ItemLevelDiscountAmount int64                                  `json:"item_level_discount_amount"`
	ReferenceLineItemId     string                                 `json:"reference_line_item_id"`
	Description             string                                 `json:"description"`
	EntityDescription       string                                 `json:"entity_description"`
	EntityType              invoiceEstimateEnum.LineItemEntityType `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason                   `json:"tax_exempt_reason"`
	EntityId                string                                 `json:"entity_id"`
	CustomerId              string                                 `json:"customer_id"`
	Object                  string                                 `json:"object"`
}
type Discount struct {
	Amount        int64                                  `json:"amount"`
	Description   string                                 `json:"description"`
	EntityType    invoiceEstimateEnum.DiscountEntityType `json:"entity_type"`
	EntityId      string                                 `json:"entity_id"`
	CouponSetCode string                                 `json:"coupon_set_code"`
	Object        string                                 `json:"object"`
}
type Tax struct {
	Name        string `json:"name"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
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
type LineItemTier struct {
	LineItemId            string `json:"line_item_id"`
	StartingUnit          int32  `json:"starting_unit"`
	EndingUnit            int32  `json:"ending_unit"`
	QuantityUsed          int32  `json:"quantity_used"`
	UnitAmount            int64  `json:"unit_amount"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal"`
	QuantityUsedInDecimal string `json:"quantity_used_in_decimal"`
	UnitAmountInDecimal   string `json:"unit_amount_in_decimal"`
	Object                string `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                                           `json:"line_item_id"`
	DiscountType   invoiceEstimateEnum.LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                           `json:"coupon_id"`
	EntityId       string                                           `json:"entity_id"`
	DiscountAmount int64                                            `json:"discount_amount"`
	Object         string                                           `json:"object"`
}
