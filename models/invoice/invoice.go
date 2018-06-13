package invoice

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	creditNoteEnum "github.com/chargebee/chargebee-go/models/creditnote/enum"
	invoiceEnum "github.com/chargebee/chargebee-go/models/invoice/enum"
	transactionEnum "github.com/chargebee/chargebee-go/models/transaction/enum"
)

type Invoice struct {
	Id                    string                    `json:"id"`
	PoNumber              string                    `json:"po_number"`
	CustomerId            string                    `json:"customer_id"`
	SubscriptionId        string                    `json:"subscription_id"`
	Recurring             bool                      `json:"recurring"`
	Status                invoiceEnum.Status        `json:"status"`
	VatNumber             string                    `json:"vat_number"`
	PriceType             enum.PriceType            `json:"price_type"`
	Date                  int64                     `json:"date"`
	DueDate               int64                     `json:"due_date"`
	NetTermDays           int32                     `json:"net_term_days"`
	ExchangeRate          float64                   `json:"exchange_rate"`
	CurrencyCode          string                    `json:"currency_code"`
	Total                 int32                     `json:"total"`
	AmountPaid            int32                     `json:"amount_paid"`
	AmountAdjusted        int32                     `json:"amount_adjusted"`
	WriteOffAmount        int32                     `json:"write_off_amount"`
	CreditsApplied        int32                     `json:"credits_applied"`
	AmountDue             int32                     `json:"amount_due"`
	PaidAt                int64                     `json:"paid_at"`
	DunningStatus         invoiceEnum.DunningStatus `json:"dunning_status"`
	NextRetryAt           int64                     `json:"next_retry_at"`
	VoidedAt              int64                     `json:"voided_at"`
	ResourceVersion       int64                     `json:"resource_version"`
	UpdatedAt             int64                     `json:"updated_at"`
	SubTotal              int32                     `json:"sub_total"`
	Tax                   int32                     `json:"tax"`
	FirstInvoice          bool                      `json:"first_invoice"`
	NewSalesAmount        int32                     `json:"new_sales_amount"`
	HasAdvanceCharges     bool                      `json:"has_advance_charges"`
	BaseCurrencyCode      string                    `json:"base_currency_code"`
	ExpectedPaymentDate   int64                     `json:"expected_payment_date"`
	AmountToCollect       int32                     `json:"amount_to_collect"`
	RoundOffAmount        int32                     `json:"round_off_amount"`
	LineItems             []*LineItem               `json:"line_items"`
	Discounts             []*Discount               `json:"discounts"`
	LineItemDiscounts     []*LineItemDiscount       `json:"line_item_discounts"`
	Taxes                 []*Tax                    `json:"taxes"`
	LineItemTaxes         []*LineItemTax            `json:"line_item_taxes"`
	LinkedPayments        []*LinkedPayment          `json:"linked_payments"`
	AppliedCredits        []*AppliedCredit          `json:"applied_credits"`
	AdjustmentCreditNotes []*AdjustmentCreditNote   `json:"adjustment_credit_notes"`
	IssuedCreditNotes     []*IssuedCreditNote       `json:"issued_credit_notes"`
	LinkedOrders          []*LinkedOrder            `json:"linked_orders"`
	Notes                 []*Note                   `json:"notes"`
	ShippingAddress       *ShippingAddress          `json:"shipping_address"`
	BillingAddress        *BillingAddress           `json:"billing_address"`
	Deleted               bool                      `json:"deleted"`
	IsVatMossRegistered   bool                      `json:"is_vat_moss_registered"`
	IsDigital             bool                      `json:"is_digital"`
	Object                string                    `json:"object"`
}
type LineItem struct {
	Id                      string                         `json:"id"`
	SubscriptionId          string                         `json:"subscription_id"`
	DateFrom                int64                          `json:"date_from"`
	DateTo                  int64                          `json:"date_to"`
	UnitAmount              int32                          `json:"unit_amount"`
	Quantity                int32                          `json:"quantity"`
	IsTaxed                 bool                           `json:"is_taxed"`
	TaxAmount               int32                          `json:"tax_amount"`
	TaxRate                 float64                        `json:"tax_rate"`
	Amount                  int32                          `json:"amount"`
	DiscountAmount          int32                          `json:"discount_amount"`
	ItemLevelDiscountAmount int32                          `json:"item_level_discount_amount"`
	Description             string                         `json:"description"`
	EntityType              invoiceEnum.LineItemEntityType `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason           `json:"tax_exempt_reason"`
	EntityId                string                         `json:"entity_id"`
	Object                  string                         `json:"object"`
}
type Discount struct {
	Amount      int32                          `json:"amount"`
	Description string                         `json:"description"`
	EntityType  invoiceEnum.DiscountEntityType `json:"entity_type"`
	EntityId    string                         `json:"entity_id"`
	Object      string                         `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                                   `json:"line_item_id"`
	DiscountType   invoiceEnum.LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                   `json:"coupon_id"`
	DiscountAmount int32                                    `json:"discount_amount"`
	Object         string                                   `json:"object"`
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
type LinkedPayment struct {
	TxnId         string                 `json:"txn_id"`
	AppliedAmount int32                  `json:"applied_amount"`
	AppliedAt     int64                  `json:"applied_at"`
	TxnStatus     transactionEnum.Status `json:"txn_status"`
	TxnDate       int64                  `json:"txn_date"`
	TxnAmount     int32                  `json:"txn_amount"`
	Object        string                 `json:"object"`
}
type AppliedCredit struct {
	CnId          string                    `json:"cn_id"`
	AppliedAmount int32                     `json:"applied_amount"`
	AppliedAt     int64                     `json:"applied_at"`
	CnReasonCode  creditNoteEnum.ReasonCode `json:"cn_reason_code"`
	CnDate        int64                     `json:"cn_date"`
	CnStatus      creditNoteEnum.Status     `json:"cn_status"`
	Object        string                    `json:"object"`
}
type AdjustmentCreditNote struct {
	CnId         string                    `json:"cn_id"`
	CnReasonCode creditNoteEnum.ReasonCode `json:"cn_reason_code"`
	CnDate       int64                     `json:"cn_date"`
	CnTotal      int32                     `json:"cn_total"`
	CnStatus     creditNoteEnum.Status     `json:"cn_status"`
	Object       string                    `json:"object"`
}
type IssuedCreditNote struct {
	CnId         string                    `json:"cn_id"`
	CnReasonCode creditNoteEnum.ReasonCode `json:"cn_reason_code"`
	CnDate       int64                     `json:"cn_date"`
	CnTotal      int32                     `json:"cn_total"`
	CnStatus     creditNoteEnum.Status     `json:"cn_status"`
	Object       string                    `json:"object"`
}
type LinkedOrder struct {
	Id                string                  `json:"id"`
	Status            invoiceEnum.OrderStatus `json:"status"`
	ReferenceId       string                  `json:"reference_id"`
	FulfillmentStatus string                  `json:"fulfillment_status"`
	BatchId           string                  `json:"batch_id"`
	CreatedAt         int64                   `json:"created_at"`
	Object            string                  `json:"object"`
}
type Note struct {
	EntityType invoiceEnum.NoteEntityType `json:"entity_type"`
	Note       string                     `json:"note"`
	EntityId   string                     `json:"entity_id"`
	Object     string                     `json:"object"`
}
type ShippingAddress struct {
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Email            string                `json:"email"`
	Company          string                `json:"company"`
	Phone            string                `json:"phone"`
	Line1            string                `json:"line1"`
	Line2            string                `json:"line2"`
	Line3            string                `json:"line3"`
	City             string                `json:"city"`
	StateCode        string                `json:"state_code"`
	State            string                `json:"state"`
	Country          string                `json:"country"`
	Zip              string                `json:"zip"`
	ValidationStatus enum.ValidationStatus `json:"validation_status"`
	Object           string                `json:"object"`
}
type BillingAddress struct {
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Email            string                `json:"email"`
	Company          string                `json:"company"`
	Phone            string                `json:"phone"`
	Line1            string                `json:"line1"`
	Line2            string                `json:"line2"`
	Line3            string                `json:"line3"`
	City             string                `json:"city"`
	StateCode        string                `json:"state_code"`
	State            string                `json:"state"`
	Country          string                `json:"country"`
	Zip              string                `json:"zip"`
	ValidationStatus enum.ValidationStatus `json:"validation_status"`
	Object           string                `json:"object"`
}
type CreateRequestParams struct {
	CustomerId      string                       `json:"customer_id"`
	CurrencyCode    string                       `json:"currency_code,omitempty"`
	Addons          []*CreateAddonParams         `json:"addons,omitempty"`
	Charges         []*CreateChargeParams        `json:"charges,omitempty"`
	Coupon          string                       `json:"coupon,omitempty"`
	PoNumber        string                       `json:"po_number,omitempty"`
	PaymentSourceId string                       `json:"payment_source_id,omitempty"`
	ShippingAddress *CreateShippingAddressParams `json:"shipping_address,omitempty"`
}
type CreateAddonParams struct {
	Id        string `json:"id,omitempty"`
	Quantity  *int32 `json:"quantity,omitempty"`
	UnitPrice *int32 `json:"unit_price,omitempty"`
}
type CreateChargeParams struct {
	Amount      *int32 `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}
type CreateShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}

type ChargeRequestParams struct {
	CustomerId      string `json:"customer_id,omitempty"`
	SubscriptionId  string `json:"subscription_id,omitempty"`
	CurrencyCode    string `json:"currency_code,omitempty"`
	Amount          *int32 `json:"amount"`
	Description     string `json:"description"`
	Coupon          string `json:"coupon,omitempty"`
	PoNumber        string `json:"po_number,omitempty"`
	PaymentSourceId string `json:"payment_source_id,omitempty"`
}

type ChargeAddonRequestParams struct {
	CustomerId      string `json:"customer_id,omitempty"`
	SubscriptionId  string `json:"subscription_id,omitempty"`
	AddonId         string `json:"addon_id"`
	AddonQuantity   *int32 `json:"addon_quantity,omitempty"`
	AddonUnitPrice  *int32 `json:"addon_unit_price,omitempty"`
	Coupon          string `json:"coupon,omitempty"`
	PoNumber        string `json:"po_number,omitempty"`
	PaymentSourceId string `json:"payment_source_id,omitempty"`
}

type ImportInvoiceRequestParams struct {
	Id                string                              `json:"id"`
	CurrencyCode      string                              `json:"currency_code,omitempty"`
	CustomerId        string                              `json:"customer_id,omitempty"`
	SubscriptionId    string                              `json:"subscription_id,omitempty"`
	PoNumber          string                              `json:"po_number,omitempty"`
	PriceType         enum.PriceType                      `json:"price_type,omitempty"`
	TaxOverrideReason enum.TaxOverrideReason              `json:"tax_override_reason,omitempty"`
	VatNumber         string                              `json:"vat_number,omitempty"`
	Date              *int64                              `json:"date"`
	Total             *int32                              `json:"total"`
	RoundOff          *int32                              `json:"round_off,omitempty"`
	Status            invoiceEnum.Status                  `json:"status,omitempty"`
	DueDate           *int64                              `json:"due_date,omitempty"`
	NetTermDays       *int32                              `json:"net_term_days,omitempty"`
	UseForProration   *bool                               `json:"use_for_proration,omitempty"`
	LineItems         []*ImportInvoiceLineItemParams      `json:"line_items,omitempty"`
	Discounts         []*ImportInvoiceDiscountParams      `json:"discounts,omitempty"`
	Taxes             []*ImportInvoiceTaxParams           `json:"taxes,omitempty"`
	Payments          []*ImportInvoicePaymentParams       `json:"payments,omitempty"`
	Notes             []*ImportInvoiceNoteParams          `json:"notes,omitempty"`
	BillingAddress    *ImportInvoiceBillingAddressParams  `json:"billing_address,omitempty"`
	ShippingAddress   *ImportInvoiceShippingAddressParams `json:"shipping_address,omitempty"`
}
type ImportInvoiceLineItemParams struct {
	DateFrom                   *int64                         `json:"date_from,omitempty"`
	DateTo                     *int64                         `json:"date_to,omitempty"`
	Description                string                         `json:"description"`
	UnitAmount                 *int32                         `json:"unit_amount,omitempty"`
	Quantity                   *int32                         `json:"quantity,omitempty"`
	Amount                     *int32                         `json:"amount,omitempty"`
	EntityType                 invoiceEnum.LineItemEntityType `json:"entity_type,omitempty"`
	EntityId                   string                         `json:"entity_id,omitempty"`
	ItemLevelDiscount1EntityId string                         `json:"item_level_discount1_entity_id,omitempty"`
	ItemLevelDiscount1Amount   *int32                         `json:"item_level_discount1_amount,omitempty"`
	ItemLevelDiscount2EntityId string                         `json:"item_level_discount2_entity_id,omitempty"`
	ItemLevelDiscount2Amount   *int32                         `json:"item_level_discount2_amount,omitempty"`
	Tax1Name                   string                         `json:"tax1_name,omitempty"`
	Tax1Amount                 *int32                         `json:"tax1_amount,omitempty"`
	Tax2Name                   string                         `json:"tax2_name,omitempty"`
	Tax2Amount                 *int32                         `json:"tax2_amount,omitempty"`
	Tax3Name                   string                         `json:"tax3_name,omitempty"`
	Tax3Amount                 *int32                         `json:"tax3_amount,omitempty"`
	Tax4Name                   string                         `json:"tax4_name,omitempty"`
	Tax4Amount                 *int32                         `json:"tax4_amount,omitempty"`
}
type ImportInvoiceDiscountParams struct {
	EntityType  invoiceEnum.DiscountEntityType `json:"entity_type"`
	EntityId    string                         `json:"entity_id,omitempty"`
	Description string                         `json:"description,omitempty"`
	Amount      *int32                         `json:"amount"`
}
type ImportInvoiceTaxParams struct {
	Name        string            `json:"name"`
	Rate        *float64          `json:"rate"`
	Amount      *int32            `json:"amount,omitempty"`
	Description string            `json:"description,omitempty"`
	JurisType   enum.TaxJurisType `json:"juris_type,omitempty"`
	JurisName   string            `json:"juris_name,omitempty"`
	JurisCode   string            `json:"juris_code,omitempty"`
}
type ImportInvoicePaymentParams struct {
	Amount          *int32             `json:"amount"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method"`
	Date            *int64             `json:"date,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
}
type ImportInvoiceNoteParams struct {
	EntityType invoiceEnum.NoteEntityType `json:"entity_type,omitempty"`
	EntityId   string                     `json:"entity_id,omitempty"`
	Note       string                     `json:"note,omitempty"`
}
type ImportInvoiceBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type ImportInvoiceShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}

type ApplyPaymentsRequestParams struct {
	Transactions []*ApplyPaymentsTransactionParams `json:"transactions,omitempty"`
}
type ApplyPaymentsTransactionParams struct {
	Id string `json:"id,omitempty"`
}

type ApplyCreditsRequestParams struct {
	CreditNotes []*ApplyCreditsCreditNoteParams `json:"credit_notes,omitempty"`
}
type ApplyCreditsCreditNoteParams struct {
	Id string `json:"id,omitempty"`
}

type ListRequestParams struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	PaidOnAfter    *int64                  `json:"paid_on_after,omitempty"`
	IncludeDeleted *bool                   `json:"include_deleted,omitempty"`
	Id             *filter.StringFilter    `json:"id,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id,omitempty"`
	CustomerId     *filter.StringFilter    `json:"customer_id,omitempty"`
	Recurring      *filter.BooleanFilter   `json:"recurring,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	PriceType      *filter.EnumFilter      `json:"price_type,omitempty"`
	Date           *filter.TimestampFilter `json:"date,omitempty"`
	PaidAt         *filter.TimestampFilter `json:"paid_at,omitempty"`
	Total          *filter.NumberFilter    `json:"total,omitempty"`
	AmountPaid     *filter.NumberFilter    `json:"amount_paid,omitempty"`
	AmountAdjusted *filter.NumberFilter    `json:"amount_adjusted,omitempty"`
	CreditsApplied *filter.NumberFilter    `json:"credits_applied,omitempty"`
	AmountDue      *filter.NumberFilter    `json:"amount_due,omitempty"`
	DunningStatus  *filter.EnumFilter      `json:"dunning_status,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	VoidedAt       *filter.TimestampFilter `json:"voided_at,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}

type InvoicesForCustomerRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}

type InvoicesForSubscriptionRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}

type PdfRequestParams struct {
	DispositionType enum.DispositionType `json:"disposition_type,omitempty"`
}

type AddChargeRequestParams struct {
	Amount      *int32                   `json:"amount"`
	Description string                   `json:"description"`
	LineItem    *AddChargeLineItemParams `json:"line_item,omitempty"`
}
type AddChargeLineItemParams struct {
	DateFrom *int64 `json:"date_from,omitempty"`
	DateTo   *int64 `json:"date_to,omitempty"`
}

type AddAddonChargeRequestParams struct {
	AddonId        string                        `json:"addon_id"`
	AddonQuantity  *int32                        `json:"addon_quantity,omitempty"`
	AddonUnitPrice *int32                        `json:"addon_unit_price,omitempty"`
	LineItem       *AddAddonChargeLineItemParams `json:"line_item,omitempty"`
}
type AddAddonChargeLineItemParams struct {
	DateFrom *int64 `json:"date_from,omitempty"`
	DateTo   *int64 `json:"date_to,omitempty"`
}

type CollectPaymentRequestParams struct {
	Amount          *int32 `json:"amount,omitempty"`
	PaymentSourceId string `json:"payment_source_id,omitempty"`
}

type RecordPaymentRequestParams struct {
	Transaction *RecordPaymentTransactionParams `json:"transaction,omitempty"`
	Comment     string                          `json:"comment,omitempty"`
}
type RecordPaymentTransactionParams struct {
	Amount          *int32                 `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod     `json:"payment_method"`
	ReferenceNumber string                 `json:"reference_number,omitempty"`
	IdAtGateway     string                 `json:"id_at_gateway,omitempty"`
	Status          transactionEnum.Status `json:"status,omitempty"`
	Date            *int64                 `json:"date,omitempty"`
	ErrorCode       string                 `json:"error_code,omitempty"`
	ErrorText       string                 `json:"error_text,omitempty"`
}

type RefundRequestParams struct {
	RefundAmount  *int32                  `json:"refund_amount,omitempty"`
	CreditNote    *RefundCreditNoteParams `json:"credit_note,omitempty"`
	Comment       string                  `json:"comment,omitempty"`
	CustomerNotes string                  `json:"customer_notes,omitempty"`
}
type RefundCreditNoteParams struct {
	ReasonCode creditNoteEnum.ReasonCode `json:"reason_code,omitempty"`
}

type RecordRefundRequestParams struct {
	Transaction   *RecordRefundTransactionParams `json:"transaction,omitempty"`
	CreditNote    *RecordRefundCreditNoteParams  `json:"credit_note,omitempty"`
	Comment       string                         `json:"comment,omitempty"`
	CustomerNotes string                         `json:"customer_notes,omitempty"`
}
type RecordRefundTransactionParams struct {
	Amount          *int32             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date"`
}
type RecordRefundCreditNoteParams struct {
	ReasonCode creditNoteEnum.ReasonCode `json:"reason_code,omitempty"`
}

type RemovePaymentRequestParams struct {
	Transaction *RemovePaymentTransactionParams `json:"transaction,omitempty"`
}
type RemovePaymentTransactionParams struct {
	Id string `json:"id"`
}

type RemoveCreditNoteRequestParams struct {
	CreditNote *RemoveCreditNoteCreditNoteParams `json:"credit_note,omitempty"`
}
type RemoveCreditNoteCreditNoteParams struct {
	Id string `json:"id"`
}

type VoidInvoiceRequestParams struct {
	Comment string `json:"comment,omitempty"`
}

type WriteOffRequestParams struct {
	Comment string `json:"comment,omitempty"`
}

type DeleteRequestParams struct {
	Comment string `json:"comment,omitempty"`
}

type UpdateDetailsRequestParams struct {
	BillingAddress  *UpdateDetailsBillingAddressParams  `json:"billing_address,omitempty"`
	ShippingAddress *UpdateDetailsShippingAddressParams `json:"shipping_address,omitempty"`
	VatNumber       string                              `json:"vat_number,omitempty"`
	PoNumber        string                              `json:"po_number,omitempty"`
}
type UpdateDetailsBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateDetailsShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
