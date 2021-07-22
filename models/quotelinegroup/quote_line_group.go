package quotelinegroup

import (
	"github.com/chargebee/chargebee-go/enum"
	quoteLineGroupEnum "github.com/chargebee/chargebee-go/models/quotelinegroup/enum"
)

type QuoteLineGroup struct {
	Version            int32                          `json:"version"`
	Id                 string                         `json:"id"`
	SubTotal           int32                          `json:"sub_total"`
	Total              int32                          `json:"total"`
	CreditsApplied     int32                          `json:"credits_applied"`
	AmountPaid         int32                          `json:"amount_paid"`
	AmountDue          int32                          `json:"amount_due"`
	ChargeEvent        quoteLineGroupEnum.ChargeEvent `json:"charge_event"`
	BillingCycleNumber int32                          `json:"billing_cycle_number"`
	LineItems          []*LineItem                    `json:"line_items"`
	Discounts          []*Discount                    `json:"discounts"`
	LineItemDiscounts  []*LineItemDiscount            `json:"line_item_discounts"`
	Taxes              []*Tax                         `json:"taxes"`
	LineItemTaxes      []*LineItemTax                 `json:"line_item_taxes"`
	Object             string                         `json:"object"`
}
type LineItem struct {
	Id                      string                                `json:"id"`
	SubscriptionId          string                                `json:"subscription_id"`
	DateFrom                int64                                 `json:"date_from"`
	DateTo                  int64                                 `json:"date_to"`
	UnitAmount              int32                                 `json:"unit_amount"`
	Quantity                int32                                 `json:"quantity"`
	Amount                  int32                                 `json:"amount"`
	PricingModel            enum.PricingModel                     `json:"pricing_model"`
	IsTaxed                 bool                                  `json:"is_taxed"`
	TaxAmount               int32                                 `json:"tax_amount"`
	TaxRate                 float64                               `json:"tax_rate"`
	UnitAmountInDecimal     string                                `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string                                `json:"quantity_in_decimal"`
	AmountInDecimal         string                                `json:"amount_in_decimal"`
	DiscountAmount          int32                                 `json:"discount_amount"`
	ItemLevelDiscountAmount int32                                 `json:"item_level_discount_amount"`
	Description             string                                `json:"description"`
	EntityDescription       string                                `json:"entity_description"`
	EntityType              quoteLineGroupEnum.LineItemEntityType `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason                  `json:"tax_exempt_reason"`
	EntityId                string                                `json:"entity_id"`
	CustomerId              string                                `json:"customer_id"`
	Object                  string                                `json:"object"`
}
type Discount struct {
	Amount      int32                                 `json:"amount"`
	Description string                                `json:"description"`
	EntityType  quoteLineGroupEnum.DiscountEntityType `json:"entity_type"`
	EntityId    string                                `json:"entity_id"`
	Object      string                                `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                                          `json:"line_item_id"`
	DiscountType   quoteLineGroupEnum.LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                          `json:"coupon_id"`
	EntityId       string                                          `json:"entity_id"`
	DiscountAmount int32                                           `json:"discount_amount"`
	Object         string                                          `json:"object"`
}
type Tax struct {
	Name        string `json:"name"`
	Amount      int32  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}
type LineItemTax struct {
	LineItemId               string            `json:"line_item_id"`
	TaxName                  string            `json:"tax_name"`
	TaxRate                  float64           `json:"tax_rate"`
	IsPartialTaxApplied      bool              `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool              `json:"is_non_compliance_tax"`
	TaxableAmount            int32             `json:"taxable_amount"`
	TaxAmount                int32             `json:"tax_amount"`
	TaxJurisType             enum.TaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string            `json:"tax_juris_name"`
	TaxJurisCode             string            `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int32             `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string            `json:"local_currency_code"`
	Object                   string            `json:"object"`
}
