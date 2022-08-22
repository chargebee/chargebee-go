package creditnoteestimate

import (
	"github.com/chargebee/chargebee-go/enum"
	creditNoteEstimateEnum "github.com/chargebee/chargebee-go/models/creditnoteestimate/enum"
)

type CreditNoteEstimate struct {
	ReferenceInvoiceId string                      `json:"reference_invoice_id"`
	Type               creditNoteEstimateEnum.Type `json:"type"`
	PriceType          enum.PriceType              `json:"price_type"`
	CurrencyCode       string                      `json:"currency_code"`
	SubTotal           int32                       `json:"sub_total"`
	Total              int32                       `json:"total"`
	AmountAllocated    int32                       `json:"amount_allocated"`
	AmountAvailable    int32                       `json:"amount_available"`
	LineItems          []*LineItem                 `json:"line_items"`
	Discounts          []*Discount                 `json:"discounts"`
	Taxes              []*Tax                      `json:"taxes"`
	LineItemTaxes      []*LineItemTax              `json:"line_item_taxes"`
	LineItemDiscounts  []*LineItemDiscount         `json:"line_item_discounts"`
	LineItemTiers      []*LineItemTier             `json:"line_item_tiers"`
	RoundOffAmount     int32                       `json:"round_off_amount"`
	CustomerId         string                      `json:"customer_id"`
	Object             string                      `json:"object"`
}
type LineItem struct {
	Id                      string                                    `json:"id"`
	SubscriptionId          string                                    `json:"subscription_id"`
	DateFrom                int64                                     `json:"date_from"`
	DateTo                  int64                                     `json:"date_to"`
	UnitAmount              int32                                     `json:"unit_amount"`
	Quantity                int32                                     `json:"quantity"`
	Amount                  int32                                     `json:"amount"`
	PricingModel            enum.PricingModel                         `json:"pricing_model"`
	IsTaxed                 bool                                      `json:"is_taxed"`
	TaxAmount               int32                                     `json:"tax_amount"`
	TaxRate                 float64                                   `json:"tax_rate"`
	UnitAmountInDecimal     string                                    `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string                                    `json:"quantity_in_decimal"`
	AmountInDecimal         string                                    `json:"amount_in_decimal"`
	DiscountAmount          int32                                     `json:"discount_amount"`
	ItemLevelDiscountAmount int32                                     `json:"item_level_discount_amount"`
	ReferenceLineItemId     string                                    `json:"reference_line_item_id"`
	Description             string                                    `json:"description"`
	EntityDescription       string                                    `json:"entity_description"`
	EntityType              creditNoteEstimateEnum.LineItemEntityType `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason                      `json:"tax_exempt_reason"`
	EntityId                string                                    `json:"entity_id"`
	CustomerId              string                                    `json:"customer_id"`
	Object                  string                                    `json:"object"`
}
type Discount struct {
	Amount        int32                                     `json:"amount"`
	Description   string                                    `json:"description"`
	EntityType    creditNoteEstimateEnum.DiscountEntityType `json:"entity_type"`
	EntityId      string                                    `json:"entity_id"`
	CouponSetCode string                                    `json:"coupon_set_code"`
	Object        string                                    `json:"object"`
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
type LineItemDiscount struct {
	LineItemId     string                                              `json:"line_item_id"`
	DiscountType   creditNoteEstimateEnum.LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                              `json:"coupon_id"`
	EntityId       string                                              `json:"entity_id"`
	DiscountAmount int32                                               `json:"discount_amount"`
	Object         string                                              `json:"object"`
}
type LineItemTier struct {
	LineItemId            string `json:"line_item_id"`
	StartingUnit          int32  `json:"starting_unit"`
	EndingUnit            int32  `json:"ending_unit"`
	QuantityUsed          int32  `json:"quantity_used"`
	UnitAmount            int32  `json:"unit_amount"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal"`
	QuantityUsedInDecimal string `json:"quantity_used_in_decimal"`
	UnitAmountInDecimal   string `json:"unit_amount_in_decimal"`
	Object                string `json:"object"`
}
