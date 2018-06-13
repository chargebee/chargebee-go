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
	Object             string                      `json:"object"`
}
type LineItem struct {
	Id                      string                                    `json:"id"`
	SubscriptionId          string                                    `json:"subscription_id"`
	DateFrom                int64                                     `json:"date_from"`
	DateTo                  int64                                     `json:"date_to"`
	UnitAmount              int32                                     `json:"unit_amount"`
	Quantity                int32                                     `json:"quantity"`
	IsTaxed                 bool                                      `json:"is_taxed"`
	TaxAmount               int32                                     `json:"tax_amount"`
	TaxRate                 float64                                   `json:"tax_rate"`
	Amount                  int32                                     `json:"amount"`
	DiscountAmount          int32                                     `json:"discount_amount"`
	ItemLevelDiscountAmount int32                                     `json:"item_level_discount_amount"`
	Description             string                                    `json:"description"`
	EntityType              creditNoteEstimateEnum.LineItemEntityType `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason                      `json:"tax_exempt_reason"`
	EntityId                string                                    `json:"entity_id"`
	Object                  string                                    `json:"object"`
}
type Discount struct {
	Amount      int32                                     `json:"amount"`
	Description string                                    `json:"description"`
	EntityType  creditNoteEstimateEnum.DiscountEntityType `json:"entity_type"`
	EntityId    string                                    `json:"entity_id"`
	Object      string                                    `json:"object"`
}
type Tax struct {
	Name        string `json:"name"`
	Amount      int32  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}
type LineItemTax struct {
	LineItemId   string            `json:"line_item_id"`
	TaxName      string            `json:"tax_name"`
	TaxRate      float64           `json:"tax_rate"`
	TaxAmount    int32             `json:"tax_amount"`
	TaxJurisType enum.TaxJurisType `json:"tax_juris_type"`
	TaxJurisName string            `json:"tax_juris_name"`
	TaxJurisCode string            `json:"tax_juris_code"`
	Object       string            `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                                              `json:"line_item_id"`
	DiscountType   creditNoteEstimateEnum.LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                              `json:"coupon_id"`
	DiscountAmount int32                                               `json:"discount_amount"`
	Object         string                                              `json:"object"`
}
