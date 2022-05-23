package creditnote

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	creditNoteEnum "github.com/chargebee/chargebee-go/models/creditnote/enum"
	invoiceEnum "github.com/chargebee/chargebee-go/models/invoice/enum"
	transactionEnum "github.com/chargebee/chargebee-go/models/transaction/enum"
)

type CreditNote struct {
	Id                      string                    `json:"id"`
	CustomerId              string                    `json:"customer_id"`
	SubscriptionId          string                    `json:"subscription_id"`
	ReferenceInvoiceId      string                    `json:"reference_invoice_id"`
	Type                    creditNoteEnum.Type       `json:"type"`
	ReasonCode              creditNoteEnum.ReasonCode `json:"reason_code"`
	Status                  creditNoteEnum.Status     `json:"status"`
	VatNumber               string                    `json:"vat_number"`
	Date                    int64                     `json:"date"`
	PriceType               enum.PriceType            `json:"price_type"`
	CurrencyCode            string                    `json:"currency_code"`
	Total                   int32                     `json:"total"`
	AmountAllocated         int32                     `json:"amount_allocated"`
	AmountRefunded          int32                     `json:"amount_refunded"`
	AmountAvailable         int32                     `json:"amount_available"`
	RefundedAt              int64                     `json:"refunded_at"`
	VoidedAt                int64                     `json:"voided_at"`
	GeneratedAt             int64                     `json:"generated_at"`
	ResourceVersion         int64                     `json:"resource_version"`
	UpdatedAt               int64                     `json:"updated_at"`
	Channel                 enum.Channel              `json:"channel"`
	Einvoice                *Einvoice                 `json:"einvoice"`
	SubTotal                int32                     `json:"sub_total"`
	SubTotalInLocalCurrency int32                     `json:"sub_total_in_local_currency"`
	TotalInLocalCurrency    int32                     `json:"total_in_local_currency"`
	LocalCurrencyCode       string                    `json:"local_currency_code"`
	RoundOffAmount          int32                     `json:"round_off_amount"`
	FractionalCorrection    int32                     `json:"fractional_correction"`
	LineItems               []*LineItem               `json:"line_items"`
	Discounts               []*Discount               `json:"discounts"`
	LineItemDiscounts       []*LineItemDiscount       `json:"line_item_discounts"`
	LineItemTiers           []*LineItemTier           `json:"line_item_tiers"`
	Taxes                   []*Tax                    `json:"taxes"`
	LineItemTaxes           []*LineItemTax            `json:"line_item_taxes"`
	LinkedRefunds           []*LinkedRefund           `json:"linked_refunds"`
	Allocations             []*Allocation             `json:"allocations"`
	Deleted                 bool                      `json:"deleted"`
	CreateReasonCode        string                    `json:"create_reason_code"`
	VatNumberPrefix         string                    `json:"vat_number_prefix"`
	Object                  string                    `json:"object"`
}
type Einvoice struct {
	Id      string                        `json:"id"`
	Status  creditNoteEnum.EinvoiceStatus `json:"status"`
	Message string                        `json:"message"`
	Object  string                        `json:"object"`
}
type LineItem struct {
	Id                      string                            `json:"id"`
	SubscriptionId          string                            `json:"subscription_id"`
	DateFrom                int64                             `json:"date_from"`
	DateTo                  int64                             `json:"date_to"`
	UnitAmount              int32                             `json:"unit_amount"`
	Quantity                int32                             `json:"quantity"`
	Amount                  int32                             `json:"amount"`
	PricingModel            enum.PricingModel                 `json:"pricing_model"`
	IsTaxed                 bool                              `json:"is_taxed"`
	TaxAmount               int32                             `json:"tax_amount"`
	TaxRate                 float64                           `json:"tax_rate"`
	UnitAmountInDecimal     string                            `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string                            `json:"quantity_in_decimal"`
	AmountInDecimal         string                            `json:"amount_in_decimal"`
	DiscountAmount          int32                             `json:"discount_amount"`
	ItemLevelDiscountAmount int32                             `json:"item_level_discount_amount"`
	Description             string                            `json:"description"`
	EntityDescription       string                            `json:"entity_description"`
	EntityType              creditNoteEnum.LineItemEntityType `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason              `json:"tax_exempt_reason"`
	EntityId                string                            `json:"entity_id"`
	CustomerId              string                            `json:"customer_id"`
	Object                  string                            `json:"object"`
}
type Discount struct {
	Amount      int32                             `json:"amount"`
	Description string                            `json:"description"`
	EntityType  creditNoteEnum.DiscountEntityType `json:"entity_type"`
	EntityId    string                            `json:"entity_id"`
	Object      string                            `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                                      `json:"line_item_id"`
	DiscountType   creditNoteEnum.LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                      `json:"coupon_id"`
	EntityId       string                                      `json:"entity_id"`
	DiscountAmount int32                                       `json:"discount_amount"`
	Object         string                                      `json:"object"`
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
type LinkedRefund struct {
	TxnId            string                 `json:"txn_id"`
	AppliedAmount    int32                  `json:"applied_amount"`
	AppliedAt        int64                  `json:"applied_at"`
	TxnStatus        transactionEnum.Status `json:"txn_status"`
	TxnDate          int64                  `json:"txn_date"`
	TxnAmount        int32                  `json:"txn_amount"`
	RefundReasonCode string                 `json:"refund_reason_code"`
	Object           string                 `json:"object"`
}
type Allocation struct {
	InvoiceId       string             `json:"invoice_id"`
	AllocatedAmount int32              `json:"allocated_amount"`
	AllocatedAt     int64              `json:"allocated_at"`
	InvoiceDate     int64              `json:"invoice_date"`
	InvoiceStatus   invoiceEnum.Status `json:"invoice_status"`
	Object          string             `json:"object"`
}
type CreateRequestParams struct {
	ReferenceInvoiceId string                    `json:"reference_invoice_id"`
	Total              *int32                    `json:"total,omitempty"`
	Type               creditNoteEnum.Type       `json:"type"`
	ReasonCode         creditNoteEnum.ReasonCode `json:"reason_code,omitempty"`
	CreateReasonCode   string                    `json:"create_reason_code,omitempty"`
	Date               *int64                    `json:"date,omitempty"`
	CustomerNotes      string                    `json:"customer_notes,omitempty"`
	LineItems          []*CreateLineItemParams   `json:"line_items,omitempty"`
	Comment            string                    `json:"comment,omitempty"`
}
type CreateLineItemParams struct {
	ReferenceLineItemId string `json:"reference_line_item_id"`
	UnitAmount          *int32 `json:"unit_amount,omitempty"`
	UnitAmountInDecimal string `json:"unit_amount_in_decimal,omitempty"`
	Quantity            *int32 `json:"quantity,omitempty"`
	QuantityInDecimal   string `json:"quantity_in_decimal,omitempty"`
	Amount              *int32 `json:"amount,omitempty"`
	DateFrom            *int64 `json:"date_from,omitempty"`
	DateTo              *int64 `json:"date_to,omitempty"`
	Description         string `json:"description,omitempty"`
}
type PdfRequestParams struct {
	DispositionType enum.DispositionType `json:"disposition_type,omitempty"`
}
type RefundRequestParams struct {
	RefundAmount     *int32 `json:"refund_amount,omitempty"`
	CustomerNotes    string `json:"customer_notes,omitempty"`
	RefundReasonCode string `json:"refund_reason_code,omitempty"`
}
type RecordRefundRequestParams struct {
	Transaction      *RecordRefundTransactionParams `json:"transaction,omitempty"`
	RefundReasonCode string                         `json:"refund_reason_code,omitempty"`
	Comment          string                         `json:"comment,omitempty"`
}
type RecordRefundTransactionParams struct {
	Amount          *int32             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date"`
}
type VoidCreditNoteRequestParams struct {
	Comment string `json:"comment,omitempty"`
}
type ListRequestParams struct {
	Limit              *int32                  `json:"limit,omitempty"`
	Offset             string                  `json:"offset,omitempty"`
	IncludeDeleted     *bool                   `json:"include_deleted,omitempty"`
	Id                 *filter.StringFilter    `json:"id,omitempty"`
	CustomerId         *filter.StringFilter    `json:"customer_id,omitempty"`
	SubscriptionId     *filter.StringFilter    `json:"subscription_id,omitempty"`
	ReferenceInvoiceId *filter.StringFilter    `json:"reference_invoice_id,omitempty"`
	Type               *filter.EnumFilter      `json:"type,omitempty"`
	ReasonCode         *filter.EnumFilter      `json:"reason_code,omitempty"`
	CreateReasonCode   *filter.StringFilter    `json:"create_reason_code,omitempty"`
	Status             *filter.EnumFilter      `json:"status,omitempty"`
	Date               *filter.TimestampFilter `json:"date,omitempty"`
	Total              *filter.NumberFilter    `json:"total,omitempty"`
	PriceType          *filter.EnumFilter      `json:"price_type,omitempty"`
	AmountAllocated    *filter.NumberFilter    `json:"amount_allocated,omitempty"`
	AmountRefunded     *filter.NumberFilter    `json:"amount_refunded,omitempty"`
	AmountAvailable    *filter.NumberFilter    `json:"amount_available,omitempty"`
	VoidedAt           *filter.TimestampFilter `json:"voided_at,omitempty"`
	UpdatedAt          *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy             *filter.SortFilter      `json:"sort_by,omitempty"`
	Channel            *filter.EnumFilter      `json:"channel,omitempty"`
	Einvoice           *ListEinvoiceParams     `json:"einvoice,omitempty"`
}
type ListEinvoiceParams struct {
	Status *filter.EnumFilter `json:"status,omitempty"`
}
type CreditNotesForCustomerRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type DeleteRequestParams struct {
	Comment string `json:"comment,omitempty"`
}
