package chargebee

import (
	"encoding/json"
)

type CreditNoteType string

const (
	CreditNoteTypeAdjustment CreditNoteType = "adjustment"
	CreditNoteTypeRefundable CreditNoteType = "refundable"
	CreditNoteTypeStore      CreditNoteType = "store"
)

type CreditNoteReasonCode string

const (
	CreditNoteReasonCodeWriteOff                 CreditNoteReasonCode = "write_off"
	CreditNoteReasonCodeSubscriptionChange       CreditNoteReasonCode = "subscription_change"
	CreditNoteReasonCodeSubscriptionCancellation CreditNoteReasonCode = "subscription_cancellation"
	CreditNoteReasonCodeSubscriptionPause        CreditNoteReasonCode = "subscription_pause"
	CreditNoteReasonCodeChargeback               CreditNoteReasonCode = "chargeback"
	CreditNoteReasonCodeProductUnsatisfactory    CreditNoteReasonCode = "product_unsatisfactory"
	CreditNoteReasonCodeServiceUnsatisfactory    CreditNoteReasonCode = "service_unsatisfactory"
	CreditNoteReasonCodeOrderChange              CreditNoteReasonCode = "order_change"
	CreditNoteReasonCodeOrderCancellation        CreditNoteReasonCode = "order_cancellation"
	CreditNoteReasonCodeWaiver                   CreditNoteReasonCode = "waiver"
	CreditNoteReasonCodeOther                    CreditNoteReasonCode = "other"
	CreditNoteReasonCodeFraudulent               CreditNoteReasonCode = "fraudulent"
)

type CreditNoteStatus string

const (
	CreditNoteStatusAdjusted  CreditNoteStatus = "adjusted"
	CreditNoteStatusRefunded  CreditNoteStatus = "refunded"
	CreditNoteStatusRefundDue CreditNoteStatus = "refund_due"
	CreditNoteStatusVoided    CreditNoteStatus = "voided"
)

type CreditNoteLineItemEntityType string

const (
	CreditNoteLineItemEntityTypeAdhoc           CreditNoteLineItemEntityType = "adhoc"
	CreditNoteLineItemEntityTypePlanItemPrice   CreditNoteLineItemEntityType = "plan_item_price"
	CreditNoteLineItemEntityTypeAddonItemPrice  CreditNoteLineItemEntityType = "addon_item_price"
	CreditNoteLineItemEntityTypeChargeItemPrice CreditNoteLineItemEntityType = "charge_item_price"
	CreditNoteLineItemEntityTypePlanSetup       CreditNoteLineItemEntityType = "plan_setup"
	CreditNoteLineItemEntityTypePlan            CreditNoteLineItemEntityType = "plan"
	CreditNoteLineItemEntityTypeAddon           CreditNoteLineItemEntityType = "addon"
)

type CreditNoteLineItemDiscountDiscountType string

const (
	CreditNoteLineItemDiscountDiscountTypeItemLevelCoupon       CreditNoteLineItemDiscountDiscountType = "item_level_coupon"
	CreditNoteLineItemDiscountDiscountTypeDocumentLevelCoupon   CreditNoteLineItemDiscountDiscountType = "document_level_coupon"
	CreditNoteLineItemDiscountDiscountTypePromotionalCredits    CreditNoteLineItemDiscountDiscountType = "promotional_credits"
	CreditNoteLineItemDiscountDiscountTypeProratedCredits       CreditNoteLineItemDiscountDiscountType = "prorated_credits"
	CreditNoteLineItemDiscountDiscountTypeItemLevelDiscount     CreditNoteLineItemDiscountDiscountType = "item_level_discount"
	CreditNoteLineItemDiscountDiscountTypeDocumentLevelDiscount CreditNoteLineItemDiscountDiscountType = "document_level_discount"
)

type CreditNoteDiscountEntityType string

const (
	CreditNoteDiscountEntityTypeItemLevelCoupon       CreditNoteDiscountEntityType = "item_level_coupon"
	CreditNoteDiscountEntityTypeDocumentLevelCoupon   CreditNoteDiscountEntityType = "document_level_coupon"
	CreditNoteDiscountEntityTypePromotionalCredits    CreditNoteDiscountEntityType = "promotional_credits"
	CreditNoteDiscountEntityTypeProratedCredits       CreditNoteDiscountEntityType = "prorated_credits"
	CreditNoteDiscountEntityTypeItemLevelDiscount     CreditNoteDiscountEntityType = "item_level_discount"
	CreditNoteDiscountEntityTypeDocumentLevelDiscount CreditNoteDiscountEntityType = "document_level_discount"
)

type CreditNoteDiscountDiscountType string

const (
	CreditNoteDiscountDiscountTypeFixedAmount CreditNoteDiscountDiscountType = "fixed_amount"
	CreditNoteDiscountDiscountTypePercentage  CreditNoteDiscountDiscountType = "percentage"
)

type CreditNoteLinkedRefundTxnStatus string

const (
	CreditNoteLinkedRefundTxnStatusInProgress     CreditNoteLinkedRefundTxnStatus = "in_progress"
	CreditNoteLinkedRefundTxnStatusSuccess        CreditNoteLinkedRefundTxnStatus = "success"
	CreditNoteLinkedRefundTxnStatusVoided         CreditNoteLinkedRefundTxnStatus = "voided"
	CreditNoteLinkedRefundTxnStatusFailure        CreditNoteLinkedRefundTxnStatus = "failure"
	CreditNoteLinkedRefundTxnStatusTimeout        CreditNoteLinkedRefundTxnStatus = "timeout"
	CreditNoteLinkedRefundTxnStatusNeedsAttention CreditNoteLinkedRefundTxnStatus = "needs_attention"
	CreditNoteLinkedRefundTxnStatusLateFailure    CreditNoteLinkedRefundTxnStatus = "late_failure"
)

type CreditNoteAllocationInvoiceStatus string

const (
	CreditNoteAllocationInvoiceStatusPaid       CreditNoteAllocationInvoiceStatus = "paid"
	CreditNoteAllocationInvoiceStatusPosted     CreditNoteAllocationInvoiceStatus = "posted"
	CreditNoteAllocationInvoiceStatusPaymentDue CreditNoteAllocationInvoiceStatus = "payment_due"
	CreditNoteAllocationInvoiceStatusNotPaid    CreditNoteAllocationInvoiceStatus = "not_paid"
	CreditNoteAllocationInvoiceStatusVoided     CreditNoteAllocationInvoiceStatus = "voided"
	CreditNoteAllocationInvoiceStatusPending    CreditNoteAllocationInvoiceStatus = "pending"
)

type CreditNoteAllocationTaxApplication string

const (
	CreditNoteAllocationTaxApplicationPreTax  CreditNoteAllocationTaxApplication = "pre_tax"
	CreditNoteAllocationTaxApplicationPostTax CreditNoteAllocationTaxApplication = "post_tax"
)

type CreditNoteEinvoiceStatus string

const (
	CreditNoteEinvoiceStatusScheduled  CreditNoteEinvoiceStatus = "scheduled"
	CreditNoteEinvoiceStatusSkipped    CreditNoteEinvoiceStatus = "skipped"
	CreditNoteEinvoiceStatusInProgress CreditNoteEinvoiceStatus = "in_progress"
	CreditNoteEinvoiceStatusSuccess    CreditNoteEinvoiceStatus = "success"
	CreditNoteEinvoiceStatusFailed     CreditNoteEinvoiceStatus = "failed"
	CreditNoteEinvoiceStatusRegistered CreditNoteEinvoiceStatus = "registered"
)

// just struct
type CreditNote struct {
	Id                        string                           `json:"id"`
	CustomerId                string                           `json:"customer_id"`
	SubscriptionId            string                           `json:"subscription_id"`
	ReferenceInvoiceId        string                           `json:"reference_invoice_id"`
	Type                      CreditNoteType                   `json:"type"`
	ReasonCode                CreditNoteReasonCode             `json:"reason_code"`
	Status                    CreditNoteStatus                 `json:"status"`
	VatNumber                 string                           `json:"vat_number"`
	Date                      int64                            `json:"date"`
	PriceType                 PriceType                        `json:"price_type"`
	CurrencyCode              string                           `json:"currency_code"`
	Total                     int64                            `json:"total"`
	AmountAllocated           int64                            `json:"amount_allocated"`
	AmountRefunded            int64                            `json:"amount_refunded"`
	AmountAvailable           int64                            `json:"amount_available"`
	RefundedAt                int64                            `json:"refunded_at"`
	VoidedAt                  int64                            `json:"voided_at"`
	GeneratedAt               int64                            `json:"generated_at"`
	ResourceVersion           int64                            `json:"resource_version"`
	UpdatedAt                 int64                            `json:"updated_at"`
	Channel                   Channel                          `json:"channel"`
	LineItemsNextOffset       string                           `json:"line_items_next_offset"`
	SubTotal                  int64                            `json:"sub_total"`
	SubTotalInLocalCurrency   int64                            `json:"sub_total_in_local_currency"`
	TotalInLocalCurrency      int64                            `json:"total_in_local_currency"`
	LocalCurrencyCode         string                           `json:"local_currency_code"`
	RoundOffAmount            int64                            `json:"round_off_amount"`
	FractionalCorrection      int64                            `json:"fractional_correction"`
	LineItems                 []*CreditNoteLineItem            `json:"line_items"`
	LineItemTiers             []*CreditNoteLineItemTier        `json:"line_item_tiers"`
	LineItemDiscounts         []*CreditNoteLineItemDiscount    `json:"line_item_discounts"`
	LineItemTaxes             []*CreditNoteLineItemTax         `json:"line_item_taxes"`
	LineItemAddresses         []*CreditNoteLineItemAddress     `json:"line_item_addresses"`
	Discounts                 []*CreditNoteDiscount            `json:"discounts"`
	Taxes                     []*CreditNoteTax                 `json:"taxes"`
	TaxOrigin                 *CreditNoteTaxOrigin             `json:"tax_origin"`
	LinkedRefunds             []*CreditNoteLinkedRefund        `json:"linked_refunds"`
	Allocations               []*CreditNoteAllocation          `json:"allocations"`
	Deleted                   bool                             `json:"deleted"`
	TaxCategory               string                           `json:"tax_category"`
	LocalCurrencyExchangeRate float64                          `json:"local_currency_exchange_rate"`
	CreateReasonCode          string                           `json:"create_reason_code"`
	VatNumberPrefix           string                           `json:"vat_number_prefix"`
	BusinessEntityId          string                           `json:"business_entity_id"`
	ShippingAddress           *CreditNoteShippingAddress       `json:"shipping_address"`
	BillingAddress            *CreditNoteBillingAddress        `json:"billing_address"`
	Einvoice                  *CreditNoteEinvoice              `json:"einvoice"`
	SiteDetailsAtCreation     *CreditNoteSiteDetailsAtCreation `json:"site_details_at_creation"`
	CustomField               CustomField                      `json:"custom_field"`
	Object                    string                           `json:"object"`
}

// sub resources
type CreditNoteLineItem struct {
	Id                      string                       `json:"id"`
	SubscriptionId          string                       `json:"subscription_id"`
	DateFrom                int64                        `json:"date_from"`
	DateTo                  int64                        `json:"date_to"`
	UnitAmount              int64                        `json:"unit_amount"`
	Quantity                int32                        `json:"quantity"`
	Amount                  int64                        `json:"amount"`
	PricingModel            PricingModel                 `json:"pricing_model"`
	IsTaxed                 bool                         `json:"is_taxed"`
	TaxAmount               int64                        `json:"tax_amount"`
	TaxRate                 float64                      `json:"tax_rate"`
	UnitAmountInDecimal     string                       `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string                       `json:"quantity_in_decimal"`
	AmountInDecimal         string                       `json:"amount_in_decimal"`
	DiscountAmount          int64                        `json:"discount_amount"`
	ItemLevelDiscountAmount int64                        `json:"item_level_discount_amount"`
	Metered                 bool                         `json:"metered"`
	IsPercentagePricing     bool                         `json:"is_percentage_pricing"`
	ReferenceLineItemId     string                       `json:"reference_line_item_id"`
	Description             string                       `json:"description"`
	EntityDescription       string                       `json:"entity_description"`
	EntityType              CreditNoteLineItemEntityType `json:"entity_type"`
	TaxExemptReason         TaxExemptReason              `json:"tax_exempt_reason"`
	EntityId                string                       `json:"entity_id"`
	CustomerId              string                       `json:"customer_id"`
	Object                  string                       `json:"object"`
}

type CreditNoteLineItemTier struct {
	LineItemId            string      `json:"line_item_id"`
	StartingUnit          int32       `json:"starting_unit"`
	EndingUnit            int32       `json:"ending_unit"`
	QuantityUsed          int32       `json:"quantity_used"`
	UnitAmount            int64       `json:"unit_amount"`
	StartingUnitInDecimal string      `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string      `json:"ending_unit_in_decimal"`
	QuantityUsedInDecimal string      `json:"quantity_used_in_decimal"`
	UnitAmountInDecimal   string      `json:"unit_amount_in_decimal"`
	PricingType           PricingType `json:"pricing_type"`
	PackageSize           int32       `json:"package_size"`
	Object                string      `json:"object"`
}

type CreditNoteLineItemDiscount struct {
	LineItemId     string                                 `json:"line_item_id"`
	DiscountType   CreditNoteLineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                 `json:"coupon_id"`
	EntityId       string                                 `json:"entity_id"`
	DiscountAmount int64                                  `json:"discount_amount"`
	Object         string                                 `json:"object"`
}

type CreditNoteLineItemTax struct {
	LineItemId               string       `json:"line_item_id"`
	TaxName                  string       `json:"tax_name"`
	TaxRate                  float64      `json:"tax_rate"`
	DateTo                   int64        `json:"date_to"`
	DateFrom                 int64        `json:"date_from"`
	ProratedTaxableAmount    float64      `json:"prorated_taxable_amount"`
	IsPartialTaxApplied      bool         `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool         `json:"is_non_compliance_tax"`
	TaxableAmount            int64        `json:"taxable_amount"`
	TaxAmount                int64        `json:"tax_amount"`
	TaxJurisType             TaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string       `json:"tax_juris_name"`
	TaxJurisCode             string       `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int64        `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string       `json:"local_currency_code"`
	Object                   string       `json:"object"`
}

type CreditNoteLineItemAddress struct {
	LineItemId       string           `json:"line_item_id"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Email            string           `json:"email"`
	Company          string           `json:"company"`
	Phone            string           `json:"phone"`
	Line1            string           `json:"line1"`
	Line2            string           `json:"line2"`
	Line3            string           `json:"line3"`
	City             string           `json:"city"`
	StateCode        string           `json:"state_code"`
	State            string           `json:"state"`
	Country          string           `json:"country"`
	Zip              string           `json:"zip"`
	ValidationStatus ValidationStatus `json:"validation_status"`
	Object           string           `json:"object"`
}

type CreditNoteDiscount struct {
	Amount        int64                          `json:"amount"`
	Description   string                         `json:"description"`
	LineItemId    string                         `json:"line_item_id"`
	EntityType    CreditNoteDiscountEntityType   `json:"entity_type"`
	DiscountType  CreditNoteDiscountDiscountType `json:"discount_type"`
	EntityId      string                         `json:"entity_id"`
	CouponSetCode string                         `json:"coupon_set_code"`
	Object        string                         `json:"object"`
}

type CreditNoteTax struct {
	Name        string `json:"name"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}

type CreditNoteTaxOrigin struct {
	Country            string `json:"country"`
	RegistrationNumber string `json:"registration_number"`
	Object             string `json:"object"`
}

type CreditNoteLinkedRefund struct {
	TxnId            string                          `json:"txn_id"`
	AppliedAmount    int64                           `json:"applied_amount"`
	AppliedAt        int64                           `json:"applied_at"`
	TxnStatus        CreditNoteLinkedRefundTxnStatus `json:"txn_status"`
	TxnDate          int64                           `json:"txn_date"`
	TxnAmount        int64                           `json:"txn_amount"`
	RefundReasonCode string                          `json:"refund_reason_code"`
	Object           string                          `json:"object"`
}

type CreditNoteAllocation struct {
	InvoiceId       string                             `json:"invoice_id"`
	AllocatedAmount int64                              `json:"allocated_amount"`
	AllocatedAt     int64                              `json:"allocated_at"`
	InvoiceDate     int64                              `json:"invoice_date"`
	InvoiceStatus   CreditNoteAllocationInvoiceStatus  `json:"invoice_status"`
	TaxApplication  CreditNoteAllocationTaxApplication `json:"tax_application"`
	Object          string                             `json:"object"`
}

type CreditNoteShippingAddress struct {
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Email            string           `json:"email"`
	Company          string           `json:"company"`
	Phone            string           `json:"phone"`
	Line1            string           `json:"line1"`
	Line2            string           `json:"line2"`
	Line3            string           `json:"line3"`
	City             string           `json:"city"`
	StateCode        string           `json:"state_code"`
	State            string           `json:"state"`
	Country          string           `json:"country"`
	Zip              string           `json:"zip"`
	ValidationStatus ValidationStatus `json:"validation_status"`
	Object           string           `json:"object"`
}

type CreditNoteBillingAddress struct {
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Email            string           `json:"email"`
	Company          string           `json:"company"`
	Phone            string           `json:"phone"`
	Line1            string           `json:"line1"`
	Line2            string           `json:"line2"`
	Line3            string           `json:"line3"`
	City             string           `json:"city"`
	StateCode        string           `json:"state_code"`
	State            string           `json:"state"`
	Country          string           `json:"country"`
	Zip              string           `json:"zip"`
	ValidationStatus ValidationStatus `json:"validation_status"`
	Object           string           `json:"object"`
}

type CreditNoteEinvoice struct {
	Id              string                   `json:"id"`
	ReferenceNumber string                   `json:"reference_number"`
	Status          CreditNoteEinvoiceStatus `json:"status"`
	Message         string                   `json:"message"`
	Object          string                   `json:"object"`
}

type CreditNoteSiteDetailsAtCreation struct {
	Timezone            string          `json:"timezone"`
	OrganizationAddress json.RawMessage `json:"organization_address"`
	Object              string          `json:"object"`
}

// operations
// input params
type CreditNoteCreateRequest struct {
	ReferenceInvoiceId string                      `json:"reference_invoice_id,omitempty"`
	CustomerId         string                      `json:"customer_id,omitempty"`
	Total              *int64                      `json:"total,omitempty"`
	Type               CreditNoteType              `json:"type"`
	ReasonCode         CreditNoteReasonCode        `json:"reason_code,omitempty"`
	CreateReasonCode   string                      `json:"create_reason_code,omitempty"`
	Date               *int64                      `json:"date,omitempty"`
	CustomerNotes      string                      `json:"customer_notes,omitempty"`
	CurrencyCode       string                      `json:"currency_code,omitempty"`
	LineItems          []*CreditNoteCreateLineItem `json:"line_items,omitempty"`
	Comment            string                      `json:"comment,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *CreditNoteCreateRequest) payload() any { return r }

// input sub resource params multi
type CreditNoteCreateLineItem struct {
	ReferenceLineItemId string                       `json:"reference_line_item_id,omitempty"`
	UnitAmount          *int64                       `json:"unit_amount,omitempty"`
	UnitAmountInDecimal string                       `json:"unit_amount_in_decimal,omitempty"`
	Quantity            *int32                       `json:"quantity,omitempty"`
	QuantityInDecimal   string                       `json:"quantity_in_decimal,omitempty"`
	Amount              *int64                       `json:"amount,omitempty"`
	DateFrom            *int64                       `json:"date_from,omitempty"`
	DateTo              *int64                       `json:"date_to,omitempty"`
	Description         string                       `json:"description,omitempty"`
	EntityType          CreditNoteLineItemEntityType `json:"entity_type,omitempty"`
	EntityId            string                       `json:"entity_id,omitempty"`
}

type CreditNoteRetrieveRequest struct {
	LineItemsLimit  *int32 `json:"line_items_limit,omitempty"`
	LineItemsOffset string `json:"line_items_offset,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *CreditNoteRetrieveRequest) payload() any { return r }

// input sub resource params single
type CreditNoteRetrieveLineItem struct {
	SubscriptionId *StringFilter `json:"subscription_id,omitempty"`
	CustomerId     *StringFilter `json:"customer_id,omitempty"`
}

type CreditNotePdfRequest struct {
	DispositionType DispositionType `json:"disposition_type,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *CreditNotePdfRequest) payload() any { return r }

type CreditNoteRefundRequest struct {
	RefundAmount     *int64 `json:"refund_amount,omitempty"`
	CustomerNotes    string `json:"customer_notes,omitempty"`
	RefundReasonCode string `json:"refund_reason_code,omitempty"`
	apiRequest       `json:"-" form:"-"`
}

func (r *CreditNoteRefundRequest) payload() any { return r }

type CreditNoteRecordRefundRequest struct {
	Transaction      *CreditNoteRecordRefundTransaction `json:"transaction,omitempty"`
	RefundReasonCode string                             `json:"refund_reason_code,omitempty"`
	Comment          string                             `json:"comment,omitempty"`
	apiRequest       `json:"-" form:"-"`
}

func (r *CreditNoteRecordRefundRequest) payload() any { return r }

// input sub resource params single
type CreditNoteRecordRefundTransaction struct {
	Id                    string        `json:"id,omitempty"`
	Amount                *int64        `json:"amount,omitempty"`
	PaymentMethod         PaymentMethod `json:"payment_method"`
	ReferenceNumber       string        `json:"reference_number,omitempty"`
	CustomPaymentMethodId string        `json:"custom_payment_method_id,omitempty"`
	Date                  *int64        `json:"date"`
}

type CreditNoteVoidCreditNoteRequest struct {
	Comment    string `json:"comment,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CreditNoteVoidCreditNoteRequest) payload() any { return r }

type CreditNoteListRequest struct {
	Limit              *int32                  `json:"limit,omitempty"`
	Offset             string                  `json:"offset,omitempty"`
	Einvoice           *CreditNoteListEinvoice `json:"einvoice,omitempty"`
	IncludeDeleted     *bool                   `json:"include_deleted,omitempty"`
	Id                 *StringFilter           `json:"id,omitempty"`
	CustomerId         *StringFilter           `json:"customer_id,omitempty"`
	SubscriptionId     *StringFilter           `json:"subscription_id,omitempty"`
	ReferenceInvoiceId *StringFilter           `json:"reference_invoice_id,omitempty"`
	Type               *EnumFilter             `json:"type,omitempty"`
	ReasonCode         *EnumFilter             `json:"reason_code,omitempty"`
	CreateReasonCode   *StringFilter           `json:"create_reason_code,omitempty"`
	Status             *EnumFilter             `json:"status,omitempty"`
	Date               *TimestampFilter        `json:"date,omitempty"`
	Total              *NumberFilter           `json:"total,omitempty"`
	PriceType          *EnumFilter             `json:"price_type,omitempty"`
	AmountAllocated    *NumberFilter           `json:"amount_allocated,omitempty"`
	AmountRefunded     *NumberFilter           `json:"amount_refunded,omitempty"`
	AmountAvailable    *NumberFilter           `json:"amount_available,omitempty"`
	VoidedAt           *TimestampFilter        `json:"voided_at,omitempty"`
	UpdatedAt          *TimestampFilter        `json:"updated_at,omitempty"`
	SortBy             *SortFilter             `json:"sort_by,omitempty"`
	Channel            *EnumFilter             `json:"channel,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *CreditNoteListRequest) payload() any { return r }

// input sub resource params single
type CreditNoteListEinvoice struct {
	Status *EnumFilter `json:"status,omitempty"`
}

type CreditNoteCreditNotesForCustomerRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CreditNoteCreditNotesForCustomerRequest) payload() any { return r }

type CreditNoteDeleteRequest struct {
	Comment    string `json:"comment,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CreditNoteDeleteRequest) payload() any { return r }

type CreditNoteRemoveTaxWithheldRefundRequest struct {
	TaxWithheld *CreditNoteRemoveTaxWithheldRefundTaxWithheld `json:"tax_withheld,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *CreditNoteRemoveTaxWithheldRefundRequest) payload() any { return r }

// input sub resource params single
type CreditNoteRemoveTaxWithheldRefundTaxWithheld struct {
	Id string `json:"id"`
}

type CreditNoteImportCreditNoteRequest struct {
	Id                   string                                    `json:"id"`
	CustomerId           string                                    `json:"customer_id,omitempty"`
	SubscriptionId       string                                    `json:"subscription_id,omitempty"`
	ReferenceInvoiceId   string                                    `json:"reference_invoice_id"`
	Type                 CreditNoteType                            `json:"type"`
	CurrencyCode         string                                    `json:"currency_code,omitempty"`
	CreateReasonCode     string                                    `json:"create_reason_code"`
	Date                 *int64                                    `json:"date"`
	Status               CreditNoteStatus                          `json:"status,omitempty"`
	Total                *int64                                    `json:"total,omitempty"`
	RefundedAt           *int64                                    `json:"refunded_at,omitempty"`
	VoidedAt             *int64                                    `json:"voided_at,omitempty"`
	SubTotal             *int64                                    `json:"sub_total,omitempty"`
	RoundOffAmount       *int64                                    `json:"round_off_amount,omitempty"`
	FractionalCorrection *int64                                    `json:"fractional_correction,omitempty"`
	VatNumberPrefix      string                                    `json:"vat_number_prefix,omitempty"`
	LineItems            []*CreditNoteImportCreditNoteLineItem     `json:"line_items,omitempty"`
	LineItemTiers        []*CreditNoteImportCreditNoteLineItemTier `json:"line_item_tiers,omitempty"`
	Discounts            []*CreditNoteImportCreditNoteDiscount     `json:"discounts,omitempty"`
	Taxes                []*CreditNoteImportCreditNoteTax          `json:"taxes,omitempty"`
	Allocations          []*CreditNoteImportCreditNoteAllocation   `json:"allocations,omitempty"`
	LinkedRefunds        []*CreditNoteImportCreditNoteLinkedRefund `json:"linked_refunds,omitempty"`
	apiRequest           `json:"-" form:"-"`
}

func (r *CreditNoteImportCreditNoteRequest) payload() any { return r }

// input sub resource params multi
type CreditNoteImportCreditNoteLineItem struct {
	ReferenceLineItemId        string                       `json:"reference_line_item_id,omitempty"`
	Id                         string                       `json:"id,omitempty"`
	DateFrom                   *int64                       `json:"date_from,omitempty"`
	DateTo                     *int64                       `json:"date_to,omitempty"`
	SubscriptionId             string                       `json:"subscription_id,omitempty"`
	Description                string                       `json:"description"`
	UnitAmount                 *int64                       `json:"unit_amount,omitempty"`
	Quantity                   *int32                       `json:"quantity,omitempty"`
	Amount                     *int64                       `json:"amount,omitempty"`
	UnitAmountInDecimal        string                       `json:"unit_amount_in_decimal,omitempty"`
	QuantityInDecimal          string                       `json:"quantity_in_decimal,omitempty"`
	AmountInDecimal            string                       `json:"amount_in_decimal,omitempty"`
	EntityType                 CreditNoteLineItemEntityType `json:"entity_type,omitempty"`
	EntityId                   string                       `json:"entity_id,omitempty"`
	ItemLevelDiscount1EntityId string                       `json:"item_level_discount1_entity_id,omitempty"`
	ItemLevelDiscount1Amount   *int64                       `json:"item_level_discount1_amount,omitempty"`
	ItemLevelDiscount2EntityId string                       `json:"item_level_discount2_entity_id,omitempty"`
	ItemLevelDiscount2Amount   *int64                       `json:"item_level_discount2_amount,omitempty"`
	Tax1Name                   string                       `json:"tax1_name,omitempty"`
	Tax1Amount                 *int64                       `json:"tax1_amount,omitempty"`
	Tax2Name                   string                       `json:"tax2_name,omitempty"`
	Tax2Amount                 *int64                       `json:"tax2_amount,omitempty"`
	Tax3Name                   string                       `json:"tax3_name,omitempty"`
	Tax3Amount                 *int64                       `json:"tax3_amount,omitempty"`
	Tax4Name                   string                       `json:"tax4_name,omitempty"`
	Tax4Amount                 *int64                       `json:"tax4_amount,omitempty"`
	Tax5Name                   string                       `json:"tax5_name,omitempty"`
	Tax5Amount                 *int64                       `json:"tax5_amount,omitempty"`
	Tax6Name                   string                       `json:"tax6_name,omitempty"`
	Tax6Amount                 *int64                       `json:"tax6_amount,omitempty"`
	Tax7Name                   string                       `json:"tax7_name,omitempty"`
	Tax7Amount                 *int64                       `json:"tax7_amount,omitempty"`
	Tax8Name                   string                       `json:"tax8_name,omitempty"`
	Tax8Amount                 *int64                       `json:"tax8_amount,omitempty"`
	Tax9Name                   string                       `json:"tax9_name,omitempty"`
	Tax9Amount                 *int64                       `json:"tax9_amount,omitempty"`
	Tax10Name                  string                       `json:"tax10_name,omitempty"`
	Tax10Amount                *int64                       `json:"tax10_amount,omitempty"`
}

// input sub resource params multi
type CreditNoteImportCreditNoteLineItemTier struct {
	LineItemId            string `json:"line_item_id"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	QuantityUsed          *int32 `json:"quantity_used,omitempty"`
	UnitAmount            *int64 `json:"unit_amount,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	QuantityUsedInDecimal string `json:"quantity_used_in_decimal,omitempty"`
	UnitAmountInDecimal   string `json:"unit_amount_in_decimal,omitempty"`
}

// input sub resource params multi
type CreditNoteImportCreditNoteDiscount struct {
	LineItemId  string                       `json:"line_item_id,omitempty"`
	EntityType  CreditNoteDiscountEntityType `json:"entity_type"`
	EntityId    string                       `json:"entity_id,omitempty"`
	Description string                       `json:"description,omitempty"`
	Amount      *int64                       `json:"amount"`
}

// input sub resource params multi
type CreditNoteImportCreditNoteTax struct {
	Name        string       `json:"name"`
	Rate        *float64     `json:"rate"`
	Amount      *int64       `json:"amount,omitempty"`
	Description string       `json:"description,omitempty"`
	JurisType   TaxJurisType `json:"juris_type,omitempty"`
	JurisName   string       `json:"juris_name,omitempty"`
	JurisCode   string       `json:"juris_code,omitempty"`
}

// input sub resource params multi
type CreditNoteImportCreditNoteAllocation struct {
	InvoiceId       string `json:"invoice_id"`
	AllocatedAmount *int64 `json:"allocated_amount"`
	AllocatedAt     *int64 `json:"allocated_at"`
}

// input sub resource params multi
type CreditNoteImportCreditNoteLinkedRefund struct {
	Id              string        `json:"id,omitempty"`
	Amount          *int64        `json:"amount"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	Date            *int64        `json:"date"`
	ReferenceNumber string        `json:"reference_number,omitempty"`
}

// operation response
type CreditNoteCreateResponse struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
	Invoice    Invoice     `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type CreditNoteRetrieveResponse struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type CreditNotePdfResponse struct {
	Download Download `json:"download,omitempty"`
	apiResponse
}

// operation response
type CreditNoteDownloadEinvoiceResponse struct {
	Downloads []Download `json:"downloads,omitempty"`
	apiResponse
}

// operation response
type CreditNoteRefundResponse struct {
	CreditNote  *CreditNote `json:"credit_note,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type CreditNoteRecordRefundResponse struct {
	CreditNote  *CreditNote `json:"credit_note,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type CreditNoteVoidCreditNoteResponse struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
	apiResponse
}

// operation sub response
type CreditNoteListCreditNoteResponse struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
}

// operation response
type CreditNoteListResponse struct {
	List       []*CreditNoteListCreditNoteResponse `json:"list,omitempty"`
	NextOffset string                              `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type CreditNoteCreditNotesForCustomerCreditNoteResponse struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
}

// operation response
type CreditNoteCreditNotesForCustomerResponse struct {
	List       []*CreditNoteCreditNotesForCustomerCreditNoteResponse `json:"list,omitempty"`
	NextOffset string                                                `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type CreditNoteDeleteResponse struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type CreditNoteRemoveTaxWithheldRefundResponse struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type CreditNoteResendEinvoiceResponse struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type CreditNoteSendEinvoiceResponse struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type CreditNoteImportCreditNoteResponse struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
	apiResponse
}
