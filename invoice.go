package chargebee

import (
	"encoding/json"
)

type InvoiceStatus string

const (
	InvoiceStatusPaid       InvoiceStatus = "paid"
	InvoiceStatusPosted     InvoiceStatus = "posted"
	InvoiceStatusPaymentDue InvoiceStatus = "payment_due"
	InvoiceStatusNotPaid    InvoiceStatus = "not_paid"
	InvoiceStatusVoided     InvoiceStatus = "voided"
	InvoiceStatusPending    InvoiceStatus = "pending"
)

type InvoicePriceType string

const (
	InvoicePriceTypeTaxExclusive InvoicePriceType = "tax_exclusive"
	InvoicePriceTypeTaxInclusive InvoicePriceType = "tax_inclusive"
)

type InvoiceDunningStatus string

const (
	InvoiceDunningStatusInProgress InvoiceDunningStatus = "in_progress"
	InvoiceDunningStatusExhausted  InvoiceDunningStatus = "exhausted"
	InvoiceDunningStatusStopped    InvoiceDunningStatus = "stopped"
	InvoiceDunningStatusSuccess    InvoiceDunningStatus = "success"
)

type InvoiceChannel string

const (
	InvoiceChannelWeb       InvoiceChannel = "web"
	InvoiceChannelAppStore  InvoiceChannel = "app_store"
	InvoiceChannelPlayStore InvoiceChannel = "play_store"
)

type InvoiceLineItemPricingModel string

const (
	InvoiceLineItemPricingModelFlatFee   InvoiceLineItemPricingModel = "flat_fee"
	InvoiceLineItemPricingModelPerUnit   InvoiceLineItemPricingModel = "per_unit"
	InvoiceLineItemPricingModelTiered    InvoiceLineItemPricingModel = "tiered"
	InvoiceLineItemPricingModelVolume    InvoiceLineItemPricingModel = "volume"
	InvoiceLineItemPricingModelStairstep InvoiceLineItemPricingModel = "stairstep"
)

type InvoiceLineItemEntityType string

const (
	InvoiceLineItemEntityTypeAdhoc           InvoiceLineItemEntityType = "adhoc"
	InvoiceLineItemEntityTypePlanItemPrice   InvoiceLineItemEntityType = "plan_item_price"
	InvoiceLineItemEntityTypeAddonItemPrice  InvoiceLineItemEntityType = "addon_item_price"
	InvoiceLineItemEntityTypeChargeItemPrice InvoiceLineItemEntityType = "charge_item_price"
	InvoiceLineItemEntityTypePlanSetup       InvoiceLineItemEntityType = "plan_setup"
	InvoiceLineItemEntityTypePlan            InvoiceLineItemEntityType = "plan"
	InvoiceLineItemEntityTypeAddon           InvoiceLineItemEntityType = "addon"
)

type InvoiceLineItemTaxExemptReason string

const (
	InvoiceLineItemTaxExemptReasonTaxNotConfigured                 InvoiceLineItemTaxExemptReason = "tax_not_configured"
	InvoiceLineItemTaxExemptReasonRegionNonTaxable                 InvoiceLineItemTaxExemptReason = "region_non_taxable"
	InvoiceLineItemTaxExemptReasonExport                           InvoiceLineItemTaxExemptReason = "export"
	InvoiceLineItemTaxExemptReasonCustomerExempt                   InvoiceLineItemTaxExemptReason = "customer_exempt"
	InvoiceLineItemTaxExemptReasonProductExempt                    InvoiceLineItemTaxExemptReason = "product_exempt"
	InvoiceLineItemTaxExemptReasonZeroRated                        InvoiceLineItemTaxExemptReason = "zero_rated"
	InvoiceLineItemTaxExemptReasonReverseCharge                    InvoiceLineItemTaxExemptReason = "reverse_charge"
	InvoiceLineItemTaxExemptReasonHighValuePhysicalGoods           InvoiceLineItemTaxExemptReason = "high_value_physical_goods"
	InvoiceLineItemTaxExemptReasonZeroValueItem                    InvoiceLineItemTaxExemptReason = "zero_value_item"
	InvoiceLineItemTaxExemptReasonTaxNotConfiguredExternalProvider InvoiceLineItemTaxExemptReason = "tax_not_configured_external_provider"
)

type InvoiceLineItemTierPricingType string

const (
	InvoiceLineItemTierPricingTypePerUnit InvoiceLineItemTierPricingType = "per_unit"
	InvoiceLineItemTierPricingTypeFlatFee InvoiceLineItemTierPricingType = "flat_fee"
	InvoiceLineItemTierPricingTypePackage InvoiceLineItemTierPricingType = "package"
)

type InvoiceLineItemDiscountDiscountType string

const (
	InvoiceLineItemDiscountDiscountTypeItemLevelCoupon       InvoiceLineItemDiscountDiscountType = "item_level_coupon"
	InvoiceLineItemDiscountDiscountTypeDocumentLevelCoupon   InvoiceLineItemDiscountDiscountType = "document_level_coupon"
	InvoiceLineItemDiscountDiscountTypePromotionalCredits    InvoiceLineItemDiscountDiscountType = "promotional_credits"
	InvoiceLineItemDiscountDiscountTypeProratedCredits       InvoiceLineItemDiscountDiscountType = "prorated_credits"
	InvoiceLineItemDiscountDiscountTypeItemLevelDiscount     InvoiceLineItemDiscountDiscountType = "item_level_discount"
	InvoiceLineItemDiscountDiscountTypeDocumentLevelDiscount InvoiceLineItemDiscountDiscountType = "document_level_discount"
)

type InvoiceLineItemTaxTaxJurisType string

const (
	InvoiceLineItemTaxTaxJurisTypeCountry        InvoiceLineItemTaxTaxJurisType = "country"
	InvoiceLineItemTaxTaxJurisTypeFederal        InvoiceLineItemTaxTaxJurisType = "federal"
	InvoiceLineItemTaxTaxJurisTypeState          InvoiceLineItemTaxTaxJurisType = "state"
	InvoiceLineItemTaxTaxJurisTypeCounty         InvoiceLineItemTaxTaxJurisType = "county"
	InvoiceLineItemTaxTaxJurisTypeCity           InvoiceLineItemTaxTaxJurisType = "city"
	InvoiceLineItemTaxTaxJurisTypeSpecial        InvoiceLineItemTaxTaxJurisType = "special"
	InvoiceLineItemTaxTaxJurisTypeUnincorporated InvoiceLineItemTaxTaxJurisType = "unincorporated"
	InvoiceLineItemTaxTaxJurisTypeOther          InvoiceLineItemTaxTaxJurisType = "other"
)

type InvoiceLineItemAddressValidationStatus string

const (
	InvoiceLineItemAddressValidationStatusNotValidated   InvoiceLineItemAddressValidationStatus = "not_validated"
	InvoiceLineItemAddressValidationStatusValid          InvoiceLineItemAddressValidationStatus = "valid"
	InvoiceLineItemAddressValidationStatusPartiallyValid InvoiceLineItemAddressValidationStatus = "partially_valid"
	InvoiceLineItemAddressValidationStatusInvalid        InvoiceLineItemAddressValidationStatus = "invalid"
)

type InvoiceDiscountEntityType string

const (
	InvoiceDiscountEntityTypeItemLevelCoupon       InvoiceDiscountEntityType = "item_level_coupon"
	InvoiceDiscountEntityTypeDocumentLevelCoupon   InvoiceDiscountEntityType = "document_level_coupon"
	InvoiceDiscountEntityTypePromotionalCredits    InvoiceDiscountEntityType = "promotional_credits"
	InvoiceDiscountEntityTypeProratedCredits       InvoiceDiscountEntityType = "prorated_credits"
	InvoiceDiscountEntityTypeItemLevelDiscount     InvoiceDiscountEntityType = "item_level_discount"
	InvoiceDiscountEntityTypeDocumentLevelDiscount InvoiceDiscountEntityType = "document_level_discount"
)

type InvoiceDiscountDiscountType string

const (
	InvoiceDiscountDiscountTypeFixedAmount InvoiceDiscountDiscountType = "fixed_amount"
	InvoiceDiscountDiscountTypePercentage  InvoiceDiscountDiscountType = "percentage"
)

type InvoiceInvoiceTransactionTxnStatus string

const (
	InvoiceInvoiceTransactionTxnStatusInProgress     InvoiceInvoiceTransactionTxnStatus = "in_progress"
	InvoiceInvoiceTransactionTxnStatusSuccess        InvoiceInvoiceTransactionTxnStatus = "success"
	InvoiceInvoiceTransactionTxnStatusVoided         InvoiceInvoiceTransactionTxnStatus = "voided"
	InvoiceInvoiceTransactionTxnStatusFailure        InvoiceInvoiceTransactionTxnStatus = "failure"
	InvoiceInvoiceTransactionTxnStatusTimeout        InvoiceInvoiceTransactionTxnStatus = "timeout"
	InvoiceInvoiceTransactionTxnStatusNeedsAttention InvoiceInvoiceTransactionTxnStatus = "needs_attention"
	InvoiceInvoiceTransactionTxnStatusLateFailure    InvoiceInvoiceTransactionTxnStatus = "late_failure"
)

type InvoiceReferenceTransactionTxnStatus string

const (
	InvoiceReferenceTransactionTxnStatusInProgress     InvoiceReferenceTransactionTxnStatus = "in_progress"
	InvoiceReferenceTransactionTxnStatusSuccess        InvoiceReferenceTransactionTxnStatus = "success"
	InvoiceReferenceTransactionTxnStatusVoided         InvoiceReferenceTransactionTxnStatus = "voided"
	InvoiceReferenceTransactionTxnStatusFailure        InvoiceReferenceTransactionTxnStatus = "failure"
	InvoiceReferenceTransactionTxnStatusTimeout        InvoiceReferenceTransactionTxnStatus = "timeout"
	InvoiceReferenceTransactionTxnStatusNeedsAttention InvoiceReferenceTransactionTxnStatus = "needs_attention"
	InvoiceReferenceTransactionTxnStatusLateFailure    InvoiceReferenceTransactionTxnStatus = "late_failure"
)

type InvoiceReferenceTransactionTxnType string

const (
	InvoiceReferenceTransactionTxnTypeAuthorization   InvoiceReferenceTransactionTxnType = "authorization"
	InvoiceReferenceTransactionTxnTypePayment         InvoiceReferenceTransactionTxnType = "payment"
	InvoiceReferenceTransactionTxnTypeRefund          InvoiceReferenceTransactionTxnType = "refund"
	InvoiceReferenceTransactionTxnTypePaymentReversal InvoiceReferenceTransactionTxnType = "payment_reversal"
)

type InvoiceReferenceTransactionAuthorizationReason string

const (
	InvoiceReferenceTransactionAuthorizationReasonVerification     InvoiceReferenceTransactionAuthorizationReason = "verification"
	InvoiceReferenceTransactionAuthorizationReasonBlockingFunds    InvoiceReferenceTransactionAuthorizationReason = "blocking_funds"
	InvoiceReferenceTransactionAuthorizationReasonScheduledCapture InvoiceReferenceTransactionAuthorizationReason = "scheduled_capture"
)

type InvoiceDunningAttemptDunningType string

const (
	InvoiceDunningAttemptDunningTypeAutoCollect InvoiceDunningAttemptDunningType = "auto_collect"
	InvoiceDunningAttemptDunningTypeOffline     InvoiceDunningAttemptDunningType = "offline"
	InvoiceDunningAttemptDunningTypeDirectDebit InvoiceDunningAttemptDunningType = "direct_debit"
)

type InvoiceDunningAttemptTxnStatus string

const (
	InvoiceDunningAttemptTxnStatusInProgress     InvoiceDunningAttemptTxnStatus = "in_progress"
	InvoiceDunningAttemptTxnStatusSuccess        InvoiceDunningAttemptTxnStatus = "success"
	InvoiceDunningAttemptTxnStatusVoided         InvoiceDunningAttemptTxnStatus = "voided"
	InvoiceDunningAttemptTxnStatusFailure        InvoiceDunningAttemptTxnStatus = "failure"
	InvoiceDunningAttemptTxnStatusTimeout        InvoiceDunningAttemptTxnStatus = "timeout"
	InvoiceDunningAttemptTxnStatusNeedsAttention InvoiceDunningAttemptTxnStatus = "needs_attention"
	InvoiceDunningAttemptTxnStatusLateFailure    InvoiceDunningAttemptTxnStatus = "late_failure"
)

type InvoiceAppliedCreditCnReasonCode string

const (
	InvoiceAppliedCreditCnReasonCodeWriteOff                 InvoiceAppliedCreditCnReasonCode = "write_off"
	InvoiceAppliedCreditCnReasonCodeSubscriptionChange       InvoiceAppliedCreditCnReasonCode = "subscription_change"
	InvoiceAppliedCreditCnReasonCodeSubscriptionCancellation InvoiceAppliedCreditCnReasonCode = "subscription_cancellation"
	InvoiceAppliedCreditCnReasonCodeSubscriptionPause        InvoiceAppliedCreditCnReasonCode = "subscription_pause"
	InvoiceAppliedCreditCnReasonCodeChargeback               InvoiceAppliedCreditCnReasonCode = "chargeback"
	InvoiceAppliedCreditCnReasonCodeProductUnsatisfactory    InvoiceAppliedCreditCnReasonCode = "product_unsatisfactory"
	InvoiceAppliedCreditCnReasonCodeServiceUnsatisfactory    InvoiceAppliedCreditCnReasonCode = "service_unsatisfactory"
	InvoiceAppliedCreditCnReasonCodeOrderChange              InvoiceAppliedCreditCnReasonCode = "order_change"
	InvoiceAppliedCreditCnReasonCodeOrderCancellation        InvoiceAppliedCreditCnReasonCode = "order_cancellation"
	InvoiceAppliedCreditCnReasonCodeWaiver                   InvoiceAppliedCreditCnReasonCode = "waiver"
	InvoiceAppliedCreditCnReasonCodeOther                    InvoiceAppliedCreditCnReasonCode = "other"
	InvoiceAppliedCreditCnReasonCodeFraudulent               InvoiceAppliedCreditCnReasonCode = "fraudulent"
)

type InvoiceAppliedCreditCnStatus string

const (
	InvoiceAppliedCreditCnStatusAdjusted  InvoiceAppliedCreditCnStatus = "adjusted"
	InvoiceAppliedCreditCnStatusRefunded  InvoiceAppliedCreditCnStatus = "refunded"
	InvoiceAppliedCreditCnStatusRefundDue InvoiceAppliedCreditCnStatus = "refund_due"
	InvoiceAppliedCreditCnStatusVoided    InvoiceAppliedCreditCnStatus = "voided"
)

type InvoiceAppliedCreditTaxApplication string

const (
	InvoiceAppliedCreditTaxApplicationPreTax  InvoiceAppliedCreditTaxApplication = "pre_tax"
	InvoiceAppliedCreditTaxApplicationPostTax InvoiceAppliedCreditTaxApplication = "post_tax"
)

type InvoiceCreatedCreditNoteCnReasonCode string

const (
	InvoiceCreatedCreditNoteCnReasonCodeWriteOff                 InvoiceCreatedCreditNoteCnReasonCode = "write_off"
	InvoiceCreatedCreditNoteCnReasonCodeSubscriptionChange       InvoiceCreatedCreditNoteCnReasonCode = "subscription_change"
	InvoiceCreatedCreditNoteCnReasonCodeSubscriptionCancellation InvoiceCreatedCreditNoteCnReasonCode = "subscription_cancellation"
	InvoiceCreatedCreditNoteCnReasonCodeSubscriptionPause        InvoiceCreatedCreditNoteCnReasonCode = "subscription_pause"
	InvoiceCreatedCreditNoteCnReasonCodeChargeback               InvoiceCreatedCreditNoteCnReasonCode = "chargeback"
	InvoiceCreatedCreditNoteCnReasonCodeProductUnsatisfactory    InvoiceCreatedCreditNoteCnReasonCode = "product_unsatisfactory"
	InvoiceCreatedCreditNoteCnReasonCodeServiceUnsatisfactory    InvoiceCreatedCreditNoteCnReasonCode = "service_unsatisfactory"
	InvoiceCreatedCreditNoteCnReasonCodeOrderChange              InvoiceCreatedCreditNoteCnReasonCode = "order_change"
	InvoiceCreatedCreditNoteCnReasonCodeOrderCancellation        InvoiceCreatedCreditNoteCnReasonCode = "order_cancellation"
	InvoiceCreatedCreditNoteCnReasonCodeWaiver                   InvoiceCreatedCreditNoteCnReasonCode = "waiver"
	InvoiceCreatedCreditNoteCnReasonCodeOther                    InvoiceCreatedCreditNoteCnReasonCode = "other"
	InvoiceCreatedCreditNoteCnReasonCodeFraudulent               InvoiceCreatedCreditNoteCnReasonCode = "fraudulent"
)

type InvoiceCreatedCreditNoteCnStatus string

const (
	InvoiceCreatedCreditNoteCnStatusAdjusted  InvoiceCreatedCreditNoteCnStatus = "adjusted"
	InvoiceCreatedCreditNoteCnStatusRefunded  InvoiceCreatedCreditNoteCnStatus = "refunded"
	InvoiceCreatedCreditNoteCnStatusRefundDue InvoiceCreatedCreditNoteCnStatus = "refund_due"
	InvoiceCreatedCreditNoteCnStatusVoided    InvoiceCreatedCreditNoteCnStatus = "voided"
)

type InvoiceLinkedOrderStatus string

const (
	InvoiceLinkedOrderStatusNew                InvoiceLinkedOrderStatus = "new"
	InvoiceLinkedOrderStatusProcessing         InvoiceLinkedOrderStatus = "processing"
	InvoiceLinkedOrderStatusComplete           InvoiceLinkedOrderStatus = "complete"
	InvoiceLinkedOrderStatusCancelled          InvoiceLinkedOrderStatus = "cancelled"
	InvoiceLinkedOrderStatusVoided             InvoiceLinkedOrderStatus = "voided"
	InvoiceLinkedOrderStatusQueued             InvoiceLinkedOrderStatus = "queued"
	InvoiceLinkedOrderStatusAwaitingShipment   InvoiceLinkedOrderStatus = "awaiting_shipment"
	InvoiceLinkedOrderStatusOnHold             InvoiceLinkedOrderStatus = "on_hold"
	InvoiceLinkedOrderStatusDelivered          InvoiceLinkedOrderStatus = "delivered"
	InvoiceLinkedOrderStatusShipped            InvoiceLinkedOrderStatus = "shipped"
	InvoiceLinkedOrderStatusPartiallyDelivered InvoiceLinkedOrderStatus = "partially_delivered"
	InvoiceLinkedOrderStatusReturned           InvoiceLinkedOrderStatus = "returned"
)

type InvoiceLinkedOrderOrderType string

const (
	InvoiceLinkedOrderOrderTypeManual          InvoiceLinkedOrderOrderType = "manual"
	InvoiceLinkedOrderOrderTypeSystemGenerated InvoiceLinkedOrderOrderType = "system_generated"
)

type InvoiceNoteEntityType string

const (
	InvoiceNoteEntityTypeCoupon          InvoiceNoteEntityType = "coupon"
	InvoiceNoteEntityTypeSubscription    InvoiceNoteEntityType = "subscription"
	InvoiceNoteEntityTypeCustomer        InvoiceNoteEntityType = "customer"
	InvoiceNoteEntityTypePlanItemPrice   InvoiceNoteEntityType = "plan_item_price"
	InvoiceNoteEntityTypeAddonItemPrice  InvoiceNoteEntityType = "addon_item_price"
	InvoiceNoteEntityTypeChargeItemPrice InvoiceNoteEntityType = "charge_item_price"
	InvoiceNoteEntityTypeTax             InvoiceNoteEntityType = "tax"
	InvoiceNoteEntityTypePlan            InvoiceNoteEntityType = "plan"
	InvoiceNoteEntityTypeAddon           InvoiceNoteEntityType = "addon"
)

type InvoiceShippingAddressValidationStatus string

const (
	InvoiceShippingAddressValidationStatusNotValidated   InvoiceShippingAddressValidationStatus = "not_validated"
	InvoiceShippingAddressValidationStatusValid          InvoiceShippingAddressValidationStatus = "valid"
	InvoiceShippingAddressValidationStatusPartiallyValid InvoiceShippingAddressValidationStatus = "partially_valid"
	InvoiceShippingAddressValidationStatusInvalid        InvoiceShippingAddressValidationStatus = "invalid"
)

type InvoiceBillingAddressValidationStatus string

const (
	InvoiceBillingAddressValidationStatusNotValidated   InvoiceBillingAddressValidationStatus = "not_validated"
	InvoiceBillingAddressValidationStatusValid          InvoiceBillingAddressValidationStatus = "valid"
	InvoiceBillingAddressValidationStatusPartiallyValid InvoiceBillingAddressValidationStatus = "partially_valid"
	InvoiceBillingAddressValidationStatusInvalid        InvoiceBillingAddressValidationStatus = "invalid"
)

type InvoiceEinvoiceStatus string

const (
	InvoiceEinvoiceStatusScheduled  InvoiceEinvoiceStatus = "scheduled"
	InvoiceEinvoiceStatusSkipped    InvoiceEinvoiceStatus = "skipped"
	InvoiceEinvoiceStatusInProgress InvoiceEinvoiceStatus = "in_progress"
	InvoiceEinvoiceStatusSuccess    InvoiceEinvoiceStatus = "success"
	InvoiceEinvoiceStatusFailed     InvoiceEinvoiceStatus = "failed"
	InvoiceEinvoiceStatusRegistered InvoiceEinvoiceStatus = "registered"
)

type InvoiceAutoCollection string

const (
	InvoiceAutoCollectionOn  InvoiceAutoCollection = "on"
	InvoiceAutoCollectionOff InvoiceAutoCollection = "off"
)

type InvoicePaymentInitiator string

const (
	InvoicePaymentInitiatorCustomer InvoicePaymentInitiator = "customer"
	InvoicePaymentInitiatorMerchant InvoicePaymentInitiator = "merchant"
)

type InvoiceAvalaraSaleType string

const (
	InvoiceAvalaraSaleTypeWholesale InvoiceAvalaraSaleType = "wholesale"
	InvoiceAvalaraSaleTypeRetail    InvoiceAvalaraSaleType = "retail"
	InvoiceAvalaraSaleTypeConsumed  InvoiceAvalaraSaleType = "consumed"
	InvoiceAvalaraSaleTypeVendorUse InvoiceAvalaraSaleType = "vendor_use"
)

type InvoiceTaxOverrideReason string

const (
	InvoiceTaxOverrideReasonIdExempt       InvoiceTaxOverrideReason = "id_exempt"
	InvoiceTaxOverrideReasonCustomerExempt InvoiceTaxOverrideReason = "customer_exempt"
	InvoiceTaxOverrideReasonExport         InvoiceTaxOverrideReason = "export"
)

type InvoiceDispositionType string

const (
	InvoiceDispositionTypeAttachment InvoiceDispositionType = "attachment"
	InvoiceDispositionTypeInline     InvoiceDispositionType = "inline"
)

// just struct
type Invoice struct {
	Id                        string                         `json:"id"`
	CustomerId                string                         `json:"customer_id"`
	PaymentOwner              string                         `json:"payment_owner"`
	SubscriptionId            string                         `json:"subscription_id"`
	Recurring                 bool                           `json:"recurring"`
	Status                    InvoiceStatus                  `json:"status"`
	Date                      int64                          `json:"date"`
	DueDate                   int64                          `json:"due_date"`
	NetTermDays               int32                          `json:"net_term_days"`
	PoNumber                  string                         `json:"po_number"`
	VatNumber                 string                         `json:"vat_number"`
	PriceType                 InvoicePriceType               `json:"price_type"`
	ExchangeRate              float64                        `json:"exchange_rate"`
	LocalCurrencyExchangeRate float64                        `json:"local_currency_exchange_rate"`
	CurrencyCode              string                         `json:"currency_code"`
	LocalCurrencyCode         string                         `json:"local_currency_code"`
	Tax                       int64                          `json:"tax"`
	SubTotal                  int64                          `json:"sub_total"`
	SubTotalInLocalCurrency   int64                          `json:"sub_total_in_local_currency"`
	Total                     int64                          `json:"total"`
	TotalInLocalCurrency      int64                          `json:"total_in_local_currency"`
	AmountDue                 int64                          `json:"amount_due"`
	AmountAdjusted            int64                          `json:"amount_adjusted"`
	AmountPaid                int64                          `json:"amount_paid"`
	PaidAt                    int64                          `json:"paid_at"`
	WriteOffAmount            int64                          `json:"write_off_amount"`
	CreditsApplied            int64                          `json:"credits_applied"`
	DunningStatus             InvoiceDunningStatus           `json:"dunning_status"`
	NextRetryAt               int64                          `json:"next_retry_at"`
	VoidedAt                  int64                          `json:"voided_at"`
	ResourceVersion           int64                          `json:"resource_version"`
	UpdatedAt                 int64                          `json:"updated_at"`
	LineItemsNextOffset       string                         `json:"line_items_next_offset"`
	FirstInvoice              bool                           `json:"first_invoice"`
	NewSalesAmount            int64                          `json:"new_sales_amount"`
	HasAdvanceCharges         bool                           `json:"has_advance_charges"`
	TermFinalized             bool                           `json:"term_finalized"`
	IsGifted                  bool                           `json:"is_gifted"`
	GeneratedAt               int64                          `json:"generated_at"`
	ExpectedPaymentDate       int64                          `json:"expected_payment_date"`
	AmountToCollect           int64                          `json:"amount_to_collect"`
	RoundOffAmount            int64                          `json:"round_off_amount"`
	LineItems                 []*InvoiceLineItem             `json:"line_items"`
	LineItemTiers             []*InvoiceLineItemTier         `json:"line_item_tiers"`
	LineItemDiscounts         []*InvoiceLineItemDiscount     `json:"line_item_discounts"`
	LineItemTaxes             []*InvoiceLineItemTax          `json:"line_item_taxes"`
	LineItemCredits           []*InvoiceLineItemCredit       `json:"line_item_credits"`
	LineItemAddresses         []*InvoiceLineItemAddress      `json:"line_item_addresses"`
	Discounts                 []*InvoiceDiscount             `json:"discounts"`
	Taxes                     []*InvoiceTax                  `json:"taxes"`
	TaxOrigin                 *InvoiceTaxOrigin              `json:"tax_origin"`
	LinkedPayments            []*InvoiceLinkedPayment        `json:"linked_payments"`
	ReferenceTransactions     []*InvoiceReferenceTransaction `json:"reference_transactions"`
	DunningAttempts           []*InvoiceDunningAttempt       `json:"dunning_attempts"`
	AppliedCredits            []*InvoiceAppliedCredit        `json:"applied_credits"`
	AdjustmentCreditNotes     []*InvoiceAdjustmentCreditNote `json:"adjustment_credit_notes"`
	IssuedCreditNotes         []*InvoiceIssuedCreditNote     `json:"issued_credit_notes"`
	LinkedOrders              []*InvoiceLinkedOrder          `json:"linked_orders"`
	Notes                     []*InvoiceNote                 `json:"notes"`
	ShippingAddress           *InvoiceShippingAddress        `json:"shipping_address"`
	BillingAddress            *InvoiceBillingAddress         `json:"billing_address"`
	StatementDescriptor       *InvoiceStatementDescriptor    `json:"statement_descriptor"`
	Einvoice                  *InvoiceEinvoice               `json:"einvoice"`
	VoidReasonCode            string                         `json:"void_reason_code"`
	Deleted                   bool                           `json:"deleted"`
	TaxCategory               string                         `json:"tax_category"`
	VatNumberPrefix           string                         `json:"vat_number_prefix"`
	Channel                   InvoiceChannel                 `json:"channel"`
	BusinessEntityId          string                         `json:"business_entity_id"`
	SiteDetailsAtCreation     *InvoiceSiteDetailsAtCreation  `json:"site_details_at_creation"`
	CustomField               CustomField                    `json:"custom_field"`
	Object                    string                         `json:"object"`
}

// sub resources
type InvoiceLineItem struct {
	Id                      string                         `json:"id"`
	SubscriptionId          string                         `json:"subscription_id"`
	DateFrom                int64                          `json:"date_from"`
	DateTo                  int64                          `json:"date_to"`
	UnitAmount              int64                          `json:"unit_amount"`
	Quantity                int32                          `json:"quantity"`
	Amount                  int64                          `json:"amount"`
	PricingModel            InvoiceLineItemPricingModel    `json:"pricing_model"`
	IsTaxed                 bool                           `json:"is_taxed"`
	TaxAmount               int64                          `json:"tax_amount"`
	TaxRate                 float64                        `json:"tax_rate"`
	UnitAmountInDecimal     string                         `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string                         `json:"quantity_in_decimal"`
	AmountInDecimal         string                         `json:"amount_in_decimal"`
	DiscountAmount          int64                          `json:"discount_amount"`
	ItemLevelDiscountAmount int64                          `json:"item_level_discount_amount"`
	Metered                 bool                           `json:"metered"`
	IsPercentagePricing     bool                           `json:"is_percentage_pricing"`
	ReferenceLineItemId     string                         `json:"reference_line_item_id"`
	Description             string                         `json:"description"`
	EntityDescription       string                         `json:"entity_description"`
	EntityType              InvoiceLineItemEntityType      `json:"entity_type"`
	TaxExemptReason         InvoiceLineItemTaxExemptReason `json:"tax_exempt_reason"`
	EntityId                string                         `json:"entity_id"`
	CustomerId              string                         `json:"customer_id"`
	Object                  string                         `json:"object"`
}

type InvoiceLineItemTier struct {
	LineItemId            string                         `json:"line_item_id"`
	StartingUnit          int32                          `json:"starting_unit"`
	EndingUnit            int32                          `json:"ending_unit"`
	QuantityUsed          int32                          `json:"quantity_used"`
	UnitAmount            int64                          `json:"unit_amount"`
	StartingUnitInDecimal string                         `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string                         `json:"ending_unit_in_decimal"`
	QuantityUsedInDecimal string                         `json:"quantity_used_in_decimal"`
	UnitAmountInDecimal   string                         `json:"unit_amount_in_decimal"`
	PricingType           InvoiceLineItemTierPricingType `json:"pricing_type"`
	PackageSize           int32                          `json:"package_size"`
	Object                string                         `json:"object"`
}

type InvoiceLineItemDiscount struct {
	LineItemId     string                              `json:"line_item_id"`
	DiscountType   InvoiceLineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                              `json:"coupon_id"`
	EntityId       string                              `json:"entity_id"`
	DiscountAmount int64                               `json:"discount_amount"`
	Object         string                              `json:"object"`
}

type InvoiceLineItemTax struct {
	LineItemId               string                         `json:"line_item_id"`
	TaxName                  string                         `json:"tax_name"`
	TaxRate                  float64                        `json:"tax_rate"`
	DateTo                   int64                          `json:"date_to"`
	DateFrom                 int64                          `json:"date_from"`
	ProratedTaxableAmount    float64                        `json:"prorated_taxable_amount"`
	IsPartialTaxApplied      bool                           `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool                           `json:"is_non_compliance_tax"`
	TaxableAmount            int64                          `json:"taxable_amount"`
	TaxAmount                int64                          `json:"tax_amount"`
	TaxJurisType             InvoiceLineItemTaxTaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string                         `json:"tax_juris_name"`
	TaxJurisCode             string                         `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int64                          `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string                         `json:"local_currency_code"`
	Object                   string                         `json:"object"`
}

type InvoiceLineItemCredit struct {
	CnId          string  `json:"cn_id"`
	AppliedAmount float64 `json:"applied_amount"`
	LineItemId    string  `json:"line_item_id"`
	Object        string  `json:"object"`
}

type InvoiceLineItemAddress struct {
	LineItemId       string                                 `json:"line_item_id"`
	FirstName        string                                 `json:"first_name"`
	LastName         string                                 `json:"last_name"`
	Email            string                                 `json:"email"`
	Company          string                                 `json:"company"`
	Phone            string                                 `json:"phone"`
	Line1            string                                 `json:"line1"`
	Line2            string                                 `json:"line2"`
	Line3            string                                 `json:"line3"`
	City             string                                 `json:"city"`
	StateCode        string                                 `json:"state_code"`
	State            string                                 `json:"state"`
	Country          string                                 `json:"country"`
	Zip              string                                 `json:"zip"`
	ValidationStatus InvoiceLineItemAddressValidationStatus `json:"validation_status"`
	Object           string                                 `json:"object"`
}

type InvoiceDiscount struct {
	Amount        int64                       `json:"amount"`
	Description   string                      `json:"description"`
	LineItemId    string                      `json:"line_item_id"`
	EntityType    InvoiceDiscountEntityType   `json:"entity_type"`
	DiscountType  InvoiceDiscountDiscountType `json:"discount_type"`
	EntityId      string                      `json:"entity_id"`
	CouponSetCode string                      `json:"coupon_set_code"`
	Object        string                      `json:"object"`
}

type InvoiceTax struct {
	Name        string `json:"name"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}

type InvoiceTaxOrigin struct {
	Country            string `json:"country"`
	RegistrationNumber string `json:"registration_number"`
	Object             string `json:"object"`
}

type InvoiceLinkedPayment struct {
	TxnId         string            `json:"txn_id"`
	AppliedAmount int64             `json:"applied_amount"`
	AppliedAt     int64             `json:"applied_at"`
	TxnStatus     TransactionStatus `json:"txn_status"`
	TxnDate       int64             `json:"txn_date"`
	TxnAmount     int64             `json:"txn_amount"`
	Object        string            `json:"object"`
}

type InvoiceReferenceTransaction struct {
	AppliedAmount       int64                          `json:"applied_amount"`
	AppliedAt           int64                          `json:"applied_at"`
	TxnId               string                         `json:"txn_id"`
	TxnStatus           TransactionStatus              `json:"txn_status"`
	TxnDate             int64                          `json:"txn_date"`
	TxnAmount           int64                          `json:"txn_amount"`
	TxnType             TransactionType                `json:"txn_type"`
	AmountCapturable    int64                          `json:"amount_capturable"`
	AuthorizationReason TransactionAuthorizationReason `json:"authorization_reason"`
	Object              string                         `json:"object"`
}

type InvoiceDunningAttempt struct {
	Attempt       int32                            `json:"attempt"`
	TransactionId string                           `json:"transaction_id"`
	DunningType   InvoiceDunningAttemptDunningType `json:"dunning_type"`
	CreatedAt     int64                            `json:"created_at"`
	TxnStatus     TransactionStatus                `json:"txn_status"`
	TxnAmount     int64                            `json:"txn_amount"`
	Object        string                           `json:"object"`
}

type InvoiceAppliedCredit struct {
	CnId               string                             `json:"cn_id"`
	AppliedAmount      int64                              `json:"applied_amount"`
	AppliedAt          int64                              `json:"applied_at"`
	CnReasonCode       CreditNoteReasonCode               `json:"cn_reason_code"`
	CnCreateReasonCode string                             `json:"cn_create_reason_code"`
	CnDate             int64                              `json:"cn_date"`
	CnStatus           CreditNoteStatus                   `json:"cn_status"`
	TaxApplication     InvoiceAppliedCreditTaxApplication `json:"tax_application"`
	Object             string                             `json:"object"`
}

type InvoiceAdjustmentCreditNote struct {
	CnId               string               `json:"cn_id"`
	CnReasonCode       CreditNoteReasonCode `json:"cn_reason_code"`
	CnCreateReasonCode string               `json:"cn_create_reason_code"`
	CnDate             int64                `json:"cn_date"`
	CnTotal            int64                `json:"cn_total"`
	CnStatus           CreditNoteStatus     `json:"cn_status"`
	Object             string               `json:"object"`
}

type InvoiceIssuedCreditNote struct {
	CnId               string               `json:"cn_id"`
	CnReasonCode       CreditNoteReasonCode `json:"cn_reason_code"`
	CnCreateReasonCode string               `json:"cn_create_reason_code"`
	CnDate             int64                `json:"cn_date"`
	CnTotal            int64                `json:"cn_total"`
	CnStatus           CreditNoteStatus     `json:"cn_status"`
	Object             string               `json:"object"`
}

type InvoiceLinkedOrder struct {
	Id                string                      `json:"id"`
	DocumentNumber    string                      `json:"document_number"`
	Status            InvoiceLinkedOrderStatus    `json:"status"`
	OrderType         InvoiceLinkedOrderOrderType `json:"order_type"`
	ReferenceId       string                      `json:"reference_id"`
	FulfillmentStatus string                      `json:"fulfillment_status"`
	BatchId           string                      `json:"batch_id"`
	CreatedAt         int64                       `json:"created_at"`
	Object            string                      `json:"object"`
}

type InvoiceNote struct {
	Note       string                `json:"note"`
	EntityId   string                `json:"entity_id"`
	EntityType InvoiceNoteEntityType `json:"entity_type"`
	Object     string                `json:"object"`
}

type InvoiceShippingAddress struct {
	FirstName        string                                 `json:"first_name"`
	LastName         string                                 `json:"last_name"`
	Email            string                                 `json:"email"`
	Company          string                                 `json:"company"`
	Phone            string                                 `json:"phone"`
	Line1            string                                 `json:"line1"`
	Line2            string                                 `json:"line2"`
	Line3            string                                 `json:"line3"`
	City             string                                 `json:"city"`
	StateCode        string                                 `json:"state_code"`
	State            string                                 `json:"state"`
	Country          string                                 `json:"country"`
	Zip              string                                 `json:"zip"`
	ValidationStatus InvoiceShippingAddressValidationStatus `json:"validation_status"`
	Object           string                                 `json:"object"`
}

type InvoiceBillingAddress struct {
	FirstName        string                                `json:"first_name"`
	LastName         string                                `json:"last_name"`
	Email            string                                `json:"email"`
	Company          string                                `json:"company"`
	Phone            string                                `json:"phone"`
	Line1            string                                `json:"line1"`
	Line2            string                                `json:"line2"`
	Line3            string                                `json:"line3"`
	City             string                                `json:"city"`
	StateCode        string                                `json:"state_code"`
	State            string                                `json:"state"`
	Country          string                                `json:"country"`
	Zip              string                                `json:"zip"`
	ValidationStatus InvoiceBillingAddressValidationStatus `json:"validation_status"`
	Object           string                                `json:"object"`
}

type InvoiceStatementDescriptor struct {
	Id         string `json:"id"`
	Descriptor string `json:"descriptor"`
	Object     string `json:"object"`
}

type InvoiceEinvoice struct {
	Id              string                `json:"id"`
	ReferenceNumber string                `json:"reference_number"`
	Status          InvoiceEinvoiceStatus `json:"status"`
	Message         string                `json:"message"`
	Object          string                `json:"object"`
}

type InvoiceSiteDetailsAtCreation struct {
	Timezone            string          `json:"timezone"`
	OrganizationAddress json.RawMessage `json:"organization_address"`
	Object              string          `json:"object"`
}

// operations
// input params
type InvoiceCreateRequest struct {
	CustomerId                  string                            `json:"customer_id,omitempty"`
	SubscriptionId              string                            `json:"subscription_id,omitempty"`
	CurrencyCode                string                            `json:"currency_code,omitempty"`
	Addons                      []*InvoiceCreateAddon             `json:"addons,omitempty"`
	InvoiceDate                 *int64                            `json:"invoice_date,omitempty"`
	Charges                     []*InvoiceCreateCharge            `json:"charges,omitempty"`
	TaxProvidersFields          []*InvoiceCreateTaxProvidersField `json:"tax_providers_fields,omitempty"`
	InvoiceNote                 string                            `json:"invoice_note,omitempty"`
	RemoveGeneralNote           *bool                             `json:"remove_general_note,omitempty"`
	NotesToRemove               []*InvoiceCreateNotesToRemove     `json:"notes_to_remove,omitempty"`
	PoNumber                    string                            `json:"po_number,omitempty"`
	CouponIds                   []string                          `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId  string                            `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId             string                            `json:"payment_source_id,omitempty"`
	AutoCollection              InvoiceAutoCollection             `json:"auto_collection,omitempty"`
	ShippingAddress             *InvoiceCreateShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor         *InvoiceCreateStatementDescriptor `json:"statement_descriptor,omitempty"`
	Card                        *InvoiceCreateCard                `json:"card,omitempty"`
	BankAccount                 *InvoiceCreateBankAccount         `json:"bank_account,omitempty"`
	TokenId                     string                            `json:"token_id,omitempty"`
	PaymentMethod               *InvoiceCreatePaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent               *InvoiceCreatePaymentIntent       `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                             `json:"replace_primary_payment_source,omitempty"`
	RetainPaymentSource         *bool                             `json:"retain_payment_source,omitempty"`
	PaymentInitiator            InvoicePaymentInitiator           `json:"payment_initiator,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *InvoiceCreateRequest) payload() any { return r }

// input sub resource params multi
type InvoiceCreateAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}

// input sub resource params multi
type InvoiceCreateCharge struct {
	Amount                 *int64                       `json:"amount,omitempty"`
	AmountInDecimal        string                       `json:"amount_in_decimal,omitempty"`
	Description            string                       `json:"description,omitempty"`
	Taxable                *bool                        `json:"taxable,omitempty"`
	TaxProfileId           string                       `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string                       `json:"avalara_tax_code,omitempty"`
	HsnCode                string                       `json:"hsn_code,omitempty"`
	TaxjarProductCode      string                       `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        InvoiceChargeAvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                       `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                       `json:"avalara_service_type,omitempty"`
	DateFrom               *int64                       `json:"date_from,omitempty"`
	DateTo                 *int64                       `json:"date_to,omitempty"`
}

// input sub resource params multi
type InvoiceCreateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}

// input sub resource params multi
type InvoiceCreateNotesToRemove struct {
	EntityType InvoiceNotesToRemoveEntityType `json:"entity_type,omitempty"`
	EntityId   string                         `json:"entity_id,omitempty"`
}

// input sub resource params single
type InvoiceCreateShippingAddress struct {
	FirstName        string                                 `json:"first_name,omitempty"`
	LastName         string                                 `json:"last_name,omitempty"`
	Email            string                                 `json:"email,omitempty"`
	Company          string                                 `json:"company,omitempty"`
	Phone            string                                 `json:"phone,omitempty"`
	Line1            string                                 `json:"line1,omitempty"`
	Line2            string                                 `json:"line2,omitempty"`
	Line3            string                                 `json:"line3,omitempty"`
	City             string                                 `json:"city,omitempty"`
	StateCode        string                                 `json:"state_code,omitempty"`
	State            string                                 `json:"state,omitempty"`
	Zip              string                                 `json:"zip,omitempty"`
	Country          string                                 `json:"country,omitempty"`
	ValidationStatus InvoiceShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type InvoiceCreateStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}

// input sub resource params single
type InvoiceCreateCard struct {
	Gateway               InvoiceCardGateway         `json:"gateway,omitempty"`
	GatewayAccountId      string                     `json:"gateway_account_id,omitempty"`
	TmpToken              string                     `json:"tmp_token,omitempty"`
	FirstName             string                     `json:"first_name,omitempty"`
	LastName              string                     `json:"last_name,omitempty"`
	Number                string                     `json:"number,omitempty"`
	ExpiryMonth           *int32                     `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                     `json:"expiry_year,omitempty"`
	Cvv                   string                     `json:"cvv,omitempty"`
	PreferredScheme       InvoiceCardPreferredScheme `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                     `json:"billing_addr1,omitempty"`
	BillingAddr2          string                     `json:"billing_addr2,omitempty"`
	BillingCity           string                     `json:"billing_city,omitempty"`
	BillingStateCode      string                     `json:"billing_state_code,omitempty"`
	BillingState          string                     `json:"billing_state,omitempty"`
	BillingZip            string                     `json:"billing_zip,omitempty"`
	BillingCountry        string                     `json:"billing_country,omitempty"`
	IpAddress             string                     `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{}     `json:"additional_information,omitempty"`
}

// input sub resource params single
type InvoiceCreateBankAccount struct {
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	Iban                  string                              `json:"iban,omitempty"`
	FirstName             string                              `json:"first_name,omitempty"`
	LastName              string                              `json:"last_name,omitempty"`
	Company               string                              `json:"company,omitempty"`
	Email                 string                              `json:"email,omitempty"`
	Phone                 string                              `json:"phone,omitempty"`
	BankName              string                              `json:"bank_name,omitempty"`
	AccountNumber         string                              `json:"account_number,omitempty"`
	RoutingNumber         string                              `json:"routing_number,omitempty"`
	BankCode              string                              `json:"bank_code,omitempty"`
	AccountType           InvoiceBankAccountAccountType       `json:"account_type,omitempty"`
	AccountHolderType     InvoiceBankAccountAccountHolderType `json:"account_holder_type,omitempty"`
	EcheckType            InvoiceBankAccountEcheckType        `json:"echeck_type,omitempty"`
	IssuingCountry        string                              `json:"issuing_country,omitempty"`
	SwedishIdentityNumber string                              `json:"swedish_identity_number,omitempty"`
	BillingAddress        map[string]interface{}              `json:"billing_address,omitempty"`
}

// input sub resource params single
type InvoiceCreatePaymentMethod struct {
	Type                  InvoicePaymentMethodType    `json:"type,omitempty"`
	Gateway               InvoicePaymentMethodGateway `json:"gateway,omitempty"`
	GatewayAccountId      string                      `json:"gateway_account_id,omitempty"`
	ReferenceId           string                      `json:"reference_id,omitempty"`
	TmpToken              string                      `json:"tmp_token,omitempty"`
	IssuingCountry        string                      `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{}      `json:"additional_information,omitempty"`
}

// input sub resource params single
type InvoiceCreatePaymentIntent struct {
	Id                    string                                `json:"id,omitempty"`
	GatewayAccountId      string                                `json:"gateway_account_id,omitempty"`
	GwToken               string                                `json:"gw_token,omitempty"`
	PaymentMethodType     InvoicePaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                                `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                                `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}                `json:"additional_information,omitempty"`
}

type InvoiceCreateForChargeItemsAndChargesRequest struct {
	CustomerId                  string                                                    `json:"customer_id,omitempty"`
	SubscriptionId              string                                                    `json:"subscription_id,omitempty"`
	CurrencyCode                string                                                    `json:"currency_code,omitempty"`
	ItemPrices                  []*InvoiceCreateForChargeItemsAndChargesItemPrice         `json:"item_prices,omitempty"`
	ItemTiers                   []*InvoiceCreateForChargeItemsAndChargesItemTier          `json:"item_tiers,omitempty"`
	Charges                     []*InvoiceCreateForChargeItemsAndChargesCharge            `json:"charges,omitempty"`
	InvoiceNote                 string                                                    `json:"invoice_note,omitempty"`
	RemoveGeneralNote           *bool                                                     `json:"remove_general_note,omitempty"`
	NotesToRemove               []*InvoiceCreateForChargeItemsAndChargesNotesToRemove     `json:"notes_to_remove,omitempty"`
	PoNumber                    string                                                    `json:"po_number,omitempty"`
	CouponIds                   []string                                                  `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId  string                                                    `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId             string                                                    `json:"payment_source_id,omitempty"`
	AutoCollection              InvoiceAutoCollection                                     `json:"auto_collection,omitempty"`
	TaxProvidersFields          []*InvoiceCreateForChargeItemsAndChargesTaxProvidersField `json:"tax_providers_fields,omitempty"`
	Discounts                   []*InvoiceCreateForChargeItemsAndChargesDiscount          `json:"discounts,omitempty"`
	InvoiceDate                 *int64                                                    `json:"invoice_date,omitempty"`
	ShippingAddress             *InvoiceCreateForChargeItemsAndChargesShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor         *InvoiceCreateForChargeItemsAndChargesStatementDescriptor `json:"statement_descriptor,omitempty"`
	Card                        *InvoiceCreateForChargeItemsAndChargesCard                `json:"card,omitempty"`
	BankAccount                 *InvoiceCreateForChargeItemsAndChargesBankAccount         `json:"bank_account,omitempty"`
	TokenId                     string                                                    `json:"token_id,omitempty"`
	PaymentMethod               *InvoiceCreateForChargeItemsAndChargesPaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent               *InvoiceCreateForChargeItemsAndChargesPaymentIntent       `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                                                     `json:"replace_primary_payment_source,omitempty"`
	RetainPaymentSource         *bool                                                     `json:"retain_payment_source,omitempty"`
	PaymentInitiator            InvoicePaymentInitiator                                   `json:"payment_initiator,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *InvoiceCreateForChargeItemsAndChargesRequest) payload() any { return r }

// input sub resource params multi
type InvoiceCreateForChargeItemsAndChargesItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}

// input sub resource params multi
type InvoiceCreateForChargeItemsAndChargesItemTier struct {
	ItemPriceId           string                     `json:"item_price_id,omitempty"`
	StartingUnit          *int32                     `json:"starting_unit,omitempty"`
	EndingUnit            *int32                     `json:"ending_unit,omitempty"`
	Price                 *int64                     `json:"price,omitempty"`
	StartingUnitInDecimal string                     `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                     `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                     `json:"price_in_decimal,omitempty"`
	PricingType           InvoiceItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                     `json:"package_size,omitempty"`
}

// input sub resource params multi
type InvoiceCreateForChargeItemsAndChargesCharge struct {
	Amount                 *int64                       `json:"amount,omitempty"`
	AmountInDecimal        string                       `json:"amount_in_decimal,omitempty"`
	Description            string                       `json:"description,omitempty"`
	Taxable                *bool                        `json:"taxable,omitempty"`
	TaxProfileId           string                       `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string                       `json:"avalara_tax_code,omitempty"`
	HsnCode                string                       `json:"hsn_code,omitempty"`
	TaxjarProductCode      string                       `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        InvoiceChargeAvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                       `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                       `json:"avalara_service_type,omitempty"`
	DateFrom               *int64                       `json:"date_from,omitempty"`
	DateTo                 *int64                       `json:"date_to,omitempty"`
}

// input sub resource params multi
type InvoiceCreateForChargeItemsAndChargesNotesToRemove struct {
	EntityType InvoiceNotesToRemoveEntityType `json:"entity_type,omitempty"`
	EntityId   string                         `json:"entity_id,omitempty"`
}

// input sub resource params multi
type InvoiceCreateForChargeItemsAndChargesTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}

// input sub resource params multi
type InvoiceCreateForChargeItemsAndChargesDiscount struct {
	Percentage  *float64               `json:"percentage,omitempty"`
	Amount      *int64                 `json:"amount,omitempty"`
	Quantity    *int32                 `json:"quantity,omitempty"`
	ApplyOn     InvoiceDiscountApplyOn `json:"apply_on"`
	ItemPriceId string                 `json:"item_price_id,omitempty"`
}

// input sub resource params single
type InvoiceCreateForChargeItemsAndChargesShippingAddress struct {
	FirstName        string                                 `json:"first_name,omitempty"`
	LastName         string                                 `json:"last_name,omitempty"`
	Email            string                                 `json:"email,omitempty"`
	Company          string                                 `json:"company,omitempty"`
	Phone            string                                 `json:"phone,omitempty"`
	Line1            string                                 `json:"line1,omitempty"`
	Line2            string                                 `json:"line2,omitempty"`
	Line3            string                                 `json:"line3,omitempty"`
	City             string                                 `json:"city,omitempty"`
	StateCode        string                                 `json:"state_code,omitempty"`
	State            string                                 `json:"state,omitempty"`
	Zip              string                                 `json:"zip,omitempty"`
	Country          string                                 `json:"country,omitempty"`
	ValidationStatus InvoiceShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type InvoiceCreateForChargeItemsAndChargesStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}

// input sub resource params single
type InvoiceCreateForChargeItemsAndChargesCard struct {
	Gateway               InvoiceCardGateway         `json:"gateway,omitempty"`
	GatewayAccountId      string                     `json:"gateway_account_id,omitempty"`
	TmpToken              string                     `json:"tmp_token,omitempty"`
	FirstName             string                     `json:"first_name,omitempty"`
	LastName              string                     `json:"last_name,omitempty"`
	Number                string                     `json:"number,omitempty"`
	ExpiryMonth           *int32                     `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                     `json:"expiry_year,omitempty"`
	Cvv                   string                     `json:"cvv,omitempty"`
	PreferredScheme       InvoiceCardPreferredScheme `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                     `json:"billing_addr1,omitempty"`
	BillingAddr2          string                     `json:"billing_addr2,omitempty"`
	BillingCity           string                     `json:"billing_city,omitempty"`
	BillingStateCode      string                     `json:"billing_state_code,omitempty"`
	BillingState          string                     `json:"billing_state,omitempty"`
	BillingZip            string                     `json:"billing_zip,omitempty"`
	BillingCountry        string                     `json:"billing_country,omitempty"`
	IpAddress             string                     `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{}     `json:"additional_information,omitempty"`
}

// input sub resource params single
type InvoiceCreateForChargeItemsAndChargesBankAccount struct {
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	Iban                  string                              `json:"iban,omitempty"`
	FirstName             string                              `json:"first_name,omitempty"`
	LastName              string                              `json:"last_name,omitempty"`
	Company               string                              `json:"company,omitempty"`
	Email                 string                              `json:"email,omitempty"`
	Phone                 string                              `json:"phone,omitempty"`
	BankName              string                              `json:"bank_name,omitempty"`
	AccountNumber         string                              `json:"account_number,omitempty"`
	RoutingNumber         string                              `json:"routing_number,omitempty"`
	BankCode              string                              `json:"bank_code,omitempty"`
	AccountType           InvoiceBankAccountAccountType       `json:"account_type,omitempty"`
	AccountHolderType     InvoiceBankAccountAccountHolderType `json:"account_holder_type,omitempty"`
	EcheckType            InvoiceBankAccountEcheckType        `json:"echeck_type,omitempty"`
	IssuingCountry        string                              `json:"issuing_country,omitempty"`
	SwedishIdentityNumber string                              `json:"swedish_identity_number,omitempty"`
	BillingAddress        map[string]interface{}              `json:"billing_address,omitempty"`
}

// input sub resource params single
type InvoiceCreateForChargeItemsAndChargesPaymentMethod struct {
	Type                  InvoicePaymentMethodType    `json:"type,omitempty"`
	Gateway               InvoicePaymentMethodGateway `json:"gateway,omitempty"`
	GatewayAccountId      string                      `json:"gateway_account_id,omitempty"`
	ReferenceId           string                      `json:"reference_id,omitempty"`
	TmpToken              string                      `json:"tmp_token,omitempty"`
	IssuingCountry        string                      `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{}      `json:"additional_information,omitempty"`
}

// input sub resource params single
type InvoiceCreateForChargeItemsAndChargesPaymentIntent struct {
	Id                    string                                `json:"id,omitempty"`
	GatewayAccountId      string                                `json:"gateway_account_id,omitempty"`
	GwToken               string                                `json:"gw_token,omitempty"`
	PaymentMethodType     InvoicePaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                                `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                                `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}                `json:"additional_information,omitempty"`
}

type InvoiceChargeRequest struct {
	CustomerId             string                            `json:"customer_id,omitempty"`
	SubscriptionId         string                            `json:"subscription_id,omitempty"`
	CurrencyCode           string                            `json:"currency_code,omitempty"`
	Amount                 *int64                            `json:"amount,omitempty"`
	AmountInDecimal        string                            `json:"amount_in_decimal,omitempty"`
	Description            string                            `json:"description"`
	DateFrom               *int64                            `json:"date_from,omitempty"`
	DateTo                 *int64                            `json:"date_to,omitempty"`
	CouponIds              []string                          `json:"coupon_ids,omitempty"`
	AvalaraSaleType        InvoiceAvalaraSaleType            `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                            `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                            `json:"avalara_service_type,omitempty"`
	PoNumber               string                            `json:"po_number,omitempty"`
	InvoiceDate            *int64                            `json:"invoice_date,omitempty"`
	TaxProvidersFields     []*InvoiceChargeTaxProvidersField `json:"tax_providers_fields,omitempty"`
	PaymentSourceId        string                            `json:"payment_source_id,omitempty"`
	PaymentInitiator       InvoicePaymentInitiator           `json:"payment_initiator,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *InvoiceChargeRequest) payload() any { return r }

// input sub resource params multi
type InvoiceChargeTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}

type InvoiceChargeAddonRequest struct {
	CustomerId              string                  `json:"customer_id,omitempty"`
	SubscriptionId          string                  `json:"subscription_id,omitempty"`
	AddonId                 string                  `json:"addon_id"`
	AddonQuantity           *int32                  `json:"addon_quantity,omitempty"`
	AddonUnitPrice          *int64                  `json:"addon_unit_price,omitempty"`
	AddonQuantityInDecimal  string                  `json:"addon_quantity_in_decimal,omitempty"`
	AddonUnitPriceInDecimal string                  `json:"addon_unit_price_in_decimal,omitempty"`
	DateFrom                *int64                  `json:"date_from,omitempty"`
	DateTo                  *int64                  `json:"date_to,omitempty"`
	CouponIds               []string                `json:"coupon_ids,omitempty"`
	PoNumber                string                  `json:"po_number,omitempty"`
	InvoiceDate             *int64                  `json:"invoice_date,omitempty"`
	PaymentSourceId         string                  `json:"payment_source_id,omitempty"`
	PaymentInitiator        InvoicePaymentInitiator `json:"payment_initiator,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *InvoiceChargeAddonRequest) payload() any { return r }

type InvoiceCreateForChargeItemRequest struct {
	CustomerId       string                                `json:"customer_id,omitempty"`
	SubscriptionId   string                                `json:"subscription_id,omitempty"`
	ItemPrice        *InvoiceCreateForChargeItemItemPrice  `json:"item_price,omitempty"`
	ItemTiers        []*InvoiceCreateForChargeItemItemTier `json:"item_tiers,omitempty"`
	PoNumber         string                                `json:"po_number,omitempty"`
	Coupon           string                                `json:"coupon,omitempty"`
	PaymentSourceId  string                                `json:"payment_source_id,omitempty"`
	PaymentInitiator InvoicePaymentInitiator               `json:"payment_initiator,omitempty"`
	InvoiceDate      *int64                                `json:"invoice_date,omitempty"`
	apiRequest       `json:"-" form:"-"`
}

func (r *InvoiceCreateForChargeItemRequest) payload() any { return r }

// input sub resource params single
type InvoiceCreateForChargeItemItemPrice struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}

// input sub resource params multi
type InvoiceCreateForChargeItemItemTier struct {
	StartingUnit          *int32                     `json:"starting_unit,omitempty"`
	EndingUnit            *int32                     `json:"ending_unit,omitempty"`
	Price                 *int64                     `json:"price,omitempty"`
	StartingUnitInDecimal string                     `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                     `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                     `json:"price_in_decimal,omitempty"`
	PricingType           InvoiceItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                     `json:"package_size,omitempty"`
}

type InvoiceStopDunningRequest struct {
	Comment    string `json:"comment,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InvoiceStopDunningRequest) payload() any { return r }

type InvoicePauseDunningRequest struct {
	ExpectedPaymentDate *int64 `json:"expected_payment_date"`
	Comment             string `json:"comment,omitempty"`
	apiRequest          `json:"-" form:"-"`
}

func (r *InvoicePauseDunningRequest) payload() any { return r }

type InvoiceResumeDunningRequest struct {
	Comment    string `json:"comment,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InvoiceResumeDunningRequest) payload() any { return r }

type InvoiceImportInvoiceRequest struct {
	Id                      string                                        `json:"id"`
	CurrencyCode            string                                        `json:"currency_code,omitempty"`
	CustomerId              string                                        `json:"customer_id,omitempty"`
	SubscriptionId          string                                        `json:"subscription_id,omitempty"`
	PoNumber                string                                        `json:"po_number,omitempty"`
	PriceType               InvoicePriceType                              `json:"price_type,omitempty"`
	TaxOverrideReason       InvoiceTaxOverrideReason                      `json:"tax_override_reason,omitempty"`
	VatNumber               string                                        `json:"vat_number,omitempty"`
	VatNumberPrefix         string                                        `json:"vat_number_prefix,omitempty"`
	Date                    *int64                                        `json:"date"`
	Total                   *int64                                        `json:"total"`
	RoundOff                *int64                                        `json:"round_off,omitempty"`
	Status                  InvoiceStatus                                 `json:"status,omitempty"`
	VoidedAt                *int64                                        `json:"voided_at,omitempty"`
	VoidReasonCode          string                                        `json:"void_reason_code,omitempty"`
	IsWrittenOff            *bool                                         `json:"is_written_off,omitempty"`
	WriteOffAmount          *int64                                        `json:"write_off_amount,omitempty"`
	WriteOffDate            *int64                                        `json:"write_off_date,omitempty"`
	DueDate                 *int64                                        `json:"due_date,omitempty"`
	NetTermDays             *int32                                        `json:"net_term_days,omitempty"`
	HasAdvanceCharges       *bool                                         `json:"has_advance_charges,omitempty"`
	UseForProration         *bool                                         `json:"use_for_proration,omitempty"`
	LineItems               []*InvoiceImportInvoiceLineItem               `json:"line_items,omitempty"`
	PaymentReferenceNumbers []*InvoiceImportInvoicePaymentReferenceNumber `json:"payment_reference_numbers,omitempty"`
	LineItemTiers           []*InvoiceImportInvoiceLineItemTier           `json:"line_item_tiers,omitempty"`
	Discounts               []*InvoiceImportInvoiceDiscount               `json:"discounts,omitempty"`
	Taxes                   []*InvoiceImportInvoiceTax                    `json:"taxes,omitempty"`
	CreditNote              *InvoiceImportInvoiceCreditNote               `json:"credit_note,omitempty"`
	Payments                []*InvoiceImportInvoicePayment                `json:"payments,omitempty"`
	Notes                   []*InvoiceImportInvoiceNote                   `json:"notes,omitempty"`
	BillingAddress          *InvoiceImportInvoiceBillingAddress           `json:"billing_address,omitempty"`
	ShippingAddress         *InvoiceImportInvoiceShippingAddress          `json:"shipping_address,omitempty"`
	LineItemAddresses       []*InvoiceImportInvoiceLineItemAddress        `json:"line_item_addresses,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *InvoiceImportInvoiceRequest) payload() any { return r }

// input sub resource params multi
type InvoiceImportInvoiceLineItem struct {
	Id                         string                    `json:"id,omitempty"`
	DateFrom                   *int64                    `json:"date_from,omitempty"`
	DateTo                     *int64                    `json:"date_to,omitempty"`
	SubscriptionId             string                    `json:"subscription_id,omitempty"`
	Description                string                    `json:"description"`
	UnitAmount                 *int64                    `json:"unit_amount,omitempty"`
	Quantity                   *int32                    `json:"quantity,omitempty"`
	Amount                     *int64                    `json:"amount,omitempty"`
	UnitAmountInDecimal        string                    `json:"unit_amount_in_decimal,omitempty"`
	QuantityInDecimal          string                    `json:"quantity_in_decimal,omitempty"`
	AmountInDecimal            string                    `json:"amount_in_decimal,omitempty"`
	EntityType                 InvoiceLineItemEntityType `json:"entity_type,omitempty"`
	EntityId                   string                    `json:"entity_id,omitempty"`
	ItemLevelDiscount1EntityId string                    `json:"item_level_discount1_entity_id,omitempty"`
	ItemLevelDiscount1Amount   *int64                    `json:"item_level_discount1_amount,omitempty"`
	ItemLevelDiscount2EntityId string                    `json:"item_level_discount2_entity_id,omitempty"`
	ItemLevelDiscount2Amount   *int64                    `json:"item_level_discount2_amount,omitempty"`
	Tax1Name                   string                    `json:"tax1_name,omitempty"`
	Tax1Amount                 *int64                    `json:"tax1_amount,omitempty"`
	Tax2Name                   string                    `json:"tax2_name,omitempty"`
	Tax2Amount                 *int64                    `json:"tax2_amount,omitempty"`
	Tax3Name                   string                    `json:"tax3_name,omitempty"`
	Tax3Amount                 *int64                    `json:"tax3_amount,omitempty"`
	Tax4Name                   string                    `json:"tax4_name,omitempty"`
	Tax4Amount                 *int64                    `json:"tax4_amount,omitempty"`
	Tax5Name                   string                    `json:"tax5_name,omitempty"`
	Tax5Amount                 *int64                    `json:"tax5_amount,omitempty"`
	Tax6Name                   string                    `json:"tax6_name,omitempty"`
	Tax6Amount                 *int64                    `json:"tax6_amount,omitempty"`
	Tax7Name                   string                    `json:"tax7_name,omitempty"`
	Tax7Amount                 *int64                    `json:"tax7_amount,omitempty"`
	Tax8Name                   string                    `json:"tax8_name,omitempty"`
	Tax8Amount                 *int64                    `json:"tax8_amount,omitempty"`
	Tax9Name                   string                    `json:"tax9_name,omitempty"`
	Tax9Amount                 *int64                    `json:"tax9_amount,omitempty"`
	Tax10Name                  string                    `json:"tax10_name,omitempty"`
	Tax10Amount                *int64                    `json:"tax10_amount,omitempty"`
	CreatedAt                  *int64                    `json:"created_at,omitempty"`
}

// input sub resource params multi
type InvoiceImportInvoicePaymentReferenceNumber struct {
	Id     string                            `json:"id,omitempty"`
	Type   InvoicePaymentReferenceNumberType `json:"type"`
	Number string                            `json:"number"`
}

// input sub resource params multi
type InvoiceImportInvoiceLineItemTier struct {
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
type InvoiceImportInvoiceDiscount struct {
	LineItemId  string                    `json:"line_item_id,omitempty"`
	EntityType  InvoiceDiscountEntityType `json:"entity_type"`
	EntityId    string                    `json:"entity_id,omitempty"`
	Description string                    `json:"description,omitempty"`
	Amount      *int64                    `json:"amount"`
}

// input sub resource params multi
type InvoiceImportInvoiceTax struct {
	Name        string              `json:"name"`
	Rate        *float64            `json:"rate"`
	Amount      *int64              `json:"amount,omitempty"`
	Description string              `json:"description,omitempty"`
	JurisType   InvoiceTaxJurisType `json:"juris_type,omitempty"`
	JurisName   string              `json:"juris_name,omitempty"`
	JurisCode   string              `json:"juris_code,omitempty"`
}

// input sub resource params single
type InvoiceImportInvoiceCreditNote struct {
	Id string `json:"id,omitempty"`
}

// input sub resource params multi
type InvoiceImportInvoicePayment struct {
	Id              string                      `json:"id,omitempty"`
	Amount          *int64                      `json:"amount"`
	PaymentMethod   InvoicePaymentPaymentMethod `json:"payment_method"`
	Date            *int64                      `json:"date,omitempty"`
	ReferenceNumber string                      `json:"reference_number,omitempty"`
}

// input sub resource params multi
type InvoiceImportInvoiceNote struct {
	EntityType InvoiceNoteEntityType `json:"entity_type,omitempty"`
	EntityId   string                `json:"entity_id,omitempty"`
	Note       string                `json:"note,omitempty"`
}

// input sub resource params single
type InvoiceImportInvoiceBillingAddress struct {
	FirstName        string                                `json:"first_name,omitempty"`
	LastName         string                                `json:"last_name,omitempty"`
	Email            string                                `json:"email,omitempty"`
	Company          string                                `json:"company,omitempty"`
	Phone            string                                `json:"phone,omitempty"`
	Line1            string                                `json:"line1,omitempty"`
	Line2            string                                `json:"line2,omitempty"`
	Line3            string                                `json:"line3,omitempty"`
	City             string                                `json:"city,omitempty"`
	StateCode        string                                `json:"state_code,omitempty"`
	State            string                                `json:"state,omitempty"`
	Zip              string                                `json:"zip,omitempty"`
	Country          string                                `json:"country,omitempty"`
	ValidationStatus InvoiceBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type InvoiceImportInvoiceShippingAddress struct {
	FirstName        string                                 `json:"first_name,omitempty"`
	LastName         string                                 `json:"last_name,omitempty"`
	Email            string                                 `json:"email,omitempty"`
	Company          string                                 `json:"company,omitempty"`
	Phone            string                                 `json:"phone,omitempty"`
	Line1            string                                 `json:"line1,omitempty"`
	Line2            string                                 `json:"line2,omitempty"`
	Line3            string                                 `json:"line3,omitempty"`
	City             string                                 `json:"city,omitempty"`
	StateCode        string                                 `json:"state_code,omitempty"`
	State            string                                 `json:"state,omitempty"`
	Zip              string                                 `json:"zip,omitempty"`
	Country          string                                 `json:"country,omitempty"`
	ValidationStatus InvoiceShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type InvoiceImportInvoiceLineItemAddress struct {
	LineItemId       string                                 `json:"line_item_id,omitempty"`
	FirstName        string                                 `json:"first_name,omitempty"`
	LastName         string                                 `json:"last_name,omitempty"`
	Email            string                                 `json:"email,omitempty"`
	Company          string                                 `json:"company,omitempty"`
	Phone            string                                 `json:"phone,omitempty"`
	Line1            string                                 `json:"line1,omitempty"`
	Line2            string                                 `json:"line2,omitempty"`
	Line3            string                                 `json:"line3,omitempty"`
	City             string                                 `json:"city,omitempty"`
	StateCode        string                                 `json:"state_code,omitempty"`
	State            string                                 `json:"state,omitempty"`
	Zip              string                                 `json:"zip,omitempty"`
	Country          string                                 `json:"country,omitempty"`
	ValidationStatus InvoiceLineItemAddressValidationStatus `json:"validation_status,omitempty"`
}

type InvoiceApplyPaymentsRequest struct {
	Transactions []*InvoiceApplyPaymentsTransaction `json:"transactions,omitempty"`
	Comment      string                             `json:"comment,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *InvoiceApplyPaymentsRequest) payload() any { return r }

// input sub resource params multi
type InvoiceApplyPaymentsTransaction struct {
	Id     string `json:"id,omitempty"`
	Amount *int64 `json:"amount,omitempty"`
}

type InvoiceDeleteLineItemsRequest struct {
	LineItems  []*InvoiceDeleteLineItemsLineItem `json:"line_items,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InvoiceDeleteLineItemsRequest) payload() any { return r }

// input sub resource params multi
type InvoiceDeleteLineItemsLineItem struct {
	Id string `json:"id,omitempty"`
}

type InvoiceApplyCreditsRequest struct {
	CreditNotes []*InvoiceApplyCreditsCreditNote `json:"credit_notes,omitempty"`
	Comment     string                           `json:"comment,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *InvoiceApplyCreditsRequest) payload() any { return r }

// input sub resource params multi
type InvoiceApplyCreditsCreditNote struct {
	Id string `json:"id,omitempty"`
}

type InvoiceListRequest struct {
	Limit          *int32           `json:"limit,omitempty"`
	Offset         string           `json:"offset,omitempty"`
	Einvoice       *ListEinvoice    `json:"einvoice,omitempty"`
	IncludeDeleted *bool            `json:"include_deleted,omitempty"`
	Id             *StringFilter    `json:"id,omitempty"`
	SubscriptionId *StringFilter    `json:"subscription_id,omitempty"`
	CustomerId     *StringFilter    `json:"customer_id,omitempty"`
	Recurring      *BooleanFilter   `json:"recurring,omitempty"`
	Status         *EnumFilter      `json:"status,omitempty"`
	PriceType      *EnumFilter      `json:"price_type,omitempty"`
	Date           *TimestampFilter `json:"date,omitempty"`
	PaidAt         *TimestampFilter `json:"paid_at,omitempty"`
	Total          *NumberFilter    `json:"total,omitempty"`
	AmountPaid     *NumberFilter    `json:"amount_paid,omitempty"`
	AmountAdjusted *NumberFilter    `json:"amount_adjusted,omitempty"`
	CreditsApplied *NumberFilter    `json:"credits_applied,omitempty"`
	AmountDue      *NumberFilter    `json:"amount_due,omitempty"`
	DunningStatus  *EnumFilter      `json:"dunning_status,omitempty"`
	PaymentOwner   *StringFilter    `json:"payment_owner,omitempty"`
	UpdatedAt      *TimestampFilter `json:"updated_at,omitempty"`
	Channel        *EnumFilter      `json:"channel,omitempty"`
	VoidedAt       *TimestampFilter `json:"voided_at,omitempty"`
	VoidReasonCode *StringFilter    `json:"void_reason_code,omitempty"`
	SortBy         *SortFilter      `json:"sort_by,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *InvoiceListRequest) payload() any { return r }

// input sub resource params single
type InvoiceListEinvoice struct {
	Status *EnumFilter `json:"status,omitempty"`
}

type InvoiceInvoicesForCustomerRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InvoiceInvoicesForCustomerRequest) payload() any { return r }

type InvoiceInvoicesForSubscriptionRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InvoiceInvoicesForSubscriptionRequest) payload() any { return r }

type InvoiceRetrieveRequest struct {
	LineItemsLimit  *int32 `json:"line_items_limit,omitempty"`
	LineItemsOffset string `json:"line_items_offset,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *InvoiceRetrieveRequest) payload() any { return r }

// input sub resource params single
type InvoiceRetrieveLineItem struct {
	SubscriptionId *StringFilter `json:"subscription_id,omitempty"`
	CustomerId     *StringFilter `json:"customer_id,omitempty"`
}

type InvoicePdfRequest struct {
	DispositionType InvoiceDispositionType `json:"disposition_type,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *InvoicePdfRequest) payload() any { return r }

type InvoiceListPaymentReferenceNumbersRequest struct {
	Limit                  *int32                                             `json:"limit,omitempty"`
	Offset                 string                                             `json:"offset,omitempty"`
	PaymentReferenceNumber *ListPaymentReferenceNumbersPaymentReferenceNumber `json:"payment_reference_number,omitempty"`
	Id                     *StringFilter                                      `json:"id,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *InvoiceListPaymentReferenceNumbersRequest) payload() any { return r }

// input sub resource params single
type InvoiceListPaymentReferenceNumbersPaymentReferenceNumber struct {
	Number *StringFilter `json:"number,omitempty"`
}

type InvoiceAddChargeRequest struct {
	Amount                 *int64                    `json:"amount"`
	Description            string                    `json:"description"`
	AvalaraSaleType        InvoiceAvalaraSaleType    `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                    `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                    `json:"avalara_service_type,omitempty"`
	AvalaraTaxCode         string                    `json:"avalara_tax_code,omitempty"`
	HsnCode                string                    `json:"hsn_code,omitempty"`
	TaxjarProductCode      string                    `json:"taxjar_product_code,omitempty"`
	LineItem               *InvoiceAddChargeLineItem `json:"line_item,omitempty"`
	Comment                string                    `json:"comment,omitempty"`
	SubscriptionId         string                    `json:"subscription_id,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *InvoiceAddChargeRequest) payload() any { return r }

// input sub resource params single
type InvoiceAddChargeLineItem struct {
	DateFrom *int64 `json:"date_from,omitempty"`
	DateTo   *int64 `json:"date_to,omitempty"`
}

type InvoiceAddAddonChargeRequest struct {
	AddonId                 string                         `json:"addon_id"`
	AddonQuantity           *int32                         `json:"addon_quantity,omitempty"`
	AddonUnitPrice          *int64                         `json:"addon_unit_price,omitempty"`
	AddonQuantityInDecimal  string                         `json:"addon_quantity_in_decimal,omitempty"`
	AddonUnitPriceInDecimal string                         `json:"addon_unit_price_in_decimal,omitempty"`
	LineItem                *InvoiceAddAddonChargeLineItem `json:"line_item,omitempty"`
	Comment                 string                         `json:"comment,omitempty"`
	SubscriptionId          string                         `json:"subscription_id,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *InvoiceAddAddonChargeRequest) payload() any { return r }

// input sub resource params single
type InvoiceAddAddonChargeLineItem struct {
	DateFrom *int64 `json:"date_from,omitempty"`
	DateTo   *int64 `json:"date_to,omitempty"`
}

type InvoiceAddChargeItemRequest struct {
	ItemPrice      *InvoiceAddChargeItemItemPrice  `json:"item_price,omitempty"`
	ItemTiers      []*InvoiceAddChargeItemItemTier `json:"item_tiers,omitempty"`
	Comment        string                          `json:"comment,omitempty"`
	SubscriptionId string                          `json:"subscription_id,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *InvoiceAddChargeItemRequest) payload() any { return r }

// input sub resource params single
type InvoiceAddChargeItemItemPrice struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}

// input sub resource params multi
type InvoiceAddChargeItemItemTier struct {
	StartingUnit          *int32                     `json:"starting_unit,omitempty"`
	EndingUnit            *int32                     `json:"ending_unit,omitempty"`
	Price                 *int64                     `json:"price,omitempty"`
	StartingUnitInDecimal string                     `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                     `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                     `json:"price_in_decimal,omitempty"`
	PricingType           InvoiceItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                     `json:"package_size,omitempty"`
}

type InvoiceCloseRequest struct {
	Comment           string                       `json:"comment,omitempty"`
	InvoiceNote       string                       `json:"invoice_note,omitempty"`
	RemoveGeneralNote *bool                        `json:"remove_general_note,omitempty"`
	NotesToRemove     []*InvoiceCloseNotesToRemove `json:"notes_to_remove,omitempty"`
	InvoiceDate       *int64                       `json:"invoice_date,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *InvoiceCloseRequest) payload() any { return r }

// input sub resource params multi
type InvoiceCloseNotesToRemove struct {
	EntityType InvoiceNotesToRemoveEntityType `json:"entity_type,omitempty"`
	EntityId   string                         `json:"entity_id,omitempty"`
}

type InvoiceCollectPaymentRequest struct {
	Amount                     *int64                  `json:"amount,omitempty"`
	AuthorizationTransactionId string                  `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId            string                  `json:"payment_source_id,omitempty"`
	Comment                    string                  `json:"comment,omitempty"`
	PaymentInitiator           InvoicePaymentInitiator `json:"payment_initiator,omitempty"`
	apiRequest                 `json:"-" form:"-"`
}

func (r *InvoiceCollectPaymentRequest) payload() any { return r }

type InvoiceRecordPaymentRequest struct {
	Transaction *InvoiceRecordPaymentTransaction `json:"transaction,omitempty"`
	Comment     string                           `json:"comment,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *InvoiceRecordPaymentRequest) payload() any { return r }

// input sub resource params single
type InvoiceRecordPaymentTransaction struct {
	Amount                *int64                          `json:"amount,omitempty"`
	PaymentMethod         InvoiceTransactionPaymentMethod `json:"payment_method"`
	ReferenceNumber       string                          `json:"reference_number,omitempty"`
	CustomPaymentMethodId string                          `json:"custom_payment_method_id,omitempty"`
	IdAtGateway           string                          `json:"id_at_gateway,omitempty"`
	Status                InvoiceTransactionStatus        `json:"status,omitempty"`
	Date                  *int64                          `json:"date,omitempty"`
	ErrorCode             string                          `json:"error_code,omitempty"`
	ErrorText             string                          `json:"error_text,omitempty"`
}

type InvoiceRecordTaxWithheldRequest struct {
	TaxWithheld *InvoiceRecordTaxWithheldTaxWithheld `json:"tax_withheld,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *InvoiceRecordTaxWithheldRequest) payload() any { return r }

// input sub resource params single
type InvoiceRecordTaxWithheldTaxWithheld struct {
	Amount          *int64 `json:"amount"`
	ReferenceNumber string `json:"reference_number,omitempty"`
	Date            *int64 `json:"date,omitempty"`
	Description     string `json:"description,omitempty"`
}

type InvoiceRemoveTaxWithheldRequest struct {
	TaxWithheld *InvoiceRemoveTaxWithheldTaxWithheld `json:"tax_withheld,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *InvoiceRemoveTaxWithheldRequest) payload() any { return r }

// input sub resource params single
type InvoiceRemoveTaxWithheldTaxWithheld struct {
	Id string `json:"id"`
}

type InvoiceRefundRequest struct {
	RefundAmount  *int64                   `json:"refund_amount,omitempty"`
	CreditNote    *InvoiceRefundCreditNote `json:"credit_note,omitempty"`
	Comment       string                   `json:"comment,omitempty"`
	CustomerNotes string                   `json:"customer_notes,omitempty"`
	apiRequest    `json:"-" form:"-"`
}

func (r *InvoiceRefundRequest) payload() any { return r }

// input sub resource params single
type InvoiceRefundCreditNote struct {
	ReasonCode       InvoiceCreditNoteReasonCode `json:"reason_code,omitempty"`
	CreateReasonCode string                      `json:"create_reason_code,omitempty"`
}

type InvoiceRecordRefundRequest struct {
	Transaction   *InvoiceRecordRefundTransaction `json:"transaction,omitempty"`
	CreditNote    *InvoiceRecordRefundCreditNote  `json:"credit_note,omitempty"`
	Comment       string                          `json:"comment,omitempty"`
	CustomerNotes string                          `json:"customer_notes,omitempty"`
	apiRequest    `json:"-" form:"-"`
}

func (r *InvoiceRecordRefundRequest) payload() any { return r }

// input sub resource params single
type InvoiceRecordRefundTransaction struct {
	Amount                *int64                          `json:"amount,omitempty"`
	PaymentMethod         InvoiceTransactionPaymentMethod `json:"payment_method"`
	ReferenceNumber       string                          `json:"reference_number,omitempty"`
	CustomPaymentMethodId string                          `json:"custom_payment_method_id,omitempty"`
	Date                  *int64                          `json:"date"`
}

// input sub resource params single
type InvoiceRecordRefundCreditNote struct {
	ReasonCode       InvoiceCreditNoteReasonCode `json:"reason_code,omitempty"`
	CreateReasonCode string                      `json:"create_reason_code,omitempty"`
}

type InvoiceRemovePaymentRequest struct {
	Transaction *InvoiceRemovePaymentTransaction `json:"transaction,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *InvoiceRemovePaymentRequest) payload() any { return r }

// input sub resource params single
type InvoiceRemovePaymentTransaction struct {
	Id string `json:"id"`
}

type InvoiceRemoveCreditNoteRequest struct {
	CreditNote *InvoiceRemoveCreditNoteCreditNote `json:"credit_note,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InvoiceRemoveCreditNoteRequest) payload() any { return r }

// input sub resource params single
type InvoiceRemoveCreditNoteCreditNote struct {
	Id string `json:"id"`
}

type InvoiceVoidInvoiceRequest struct {
	Comment        string `json:"comment,omitempty"`
	VoidReasonCode string `json:"void_reason_code,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *InvoiceVoidInvoiceRequest) payload() any { return r }

type InvoiceWriteOffRequest struct {
	Comment    string `json:"comment,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InvoiceWriteOffRequest) payload() any { return r }

type InvoiceDeleteRequest struct {
	Comment    string `json:"comment,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InvoiceDeleteRequest) payload() any { return r }

type InvoiceUpdateDetailsRequest struct {
	BillingAddress      *InvoiceUpdateDetailsBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress     *InvoiceUpdateDetailsShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor *InvoiceUpdateDetailsStatementDescriptor `json:"statement_descriptor,omitempty"`
	VatNumber           string                                   `json:"vat_number,omitempty"`
	VatNumberPrefix     string                                   `json:"vat_number_prefix,omitempty"`
	PoNumber            string                                   `json:"po_number,omitempty"`
	Comment             string                                   `json:"comment,omitempty"`
	apiRequest          `json:"-" form:"-"`
}

func (r *InvoiceUpdateDetailsRequest) payload() any { return r }

// input sub resource params single
type InvoiceUpdateDetailsBillingAddress struct {
	FirstName        string                                `json:"first_name,omitempty"`
	LastName         string                                `json:"last_name,omitempty"`
	Email            string                                `json:"email,omitempty"`
	Company          string                                `json:"company,omitempty"`
	Phone            string                                `json:"phone,omitempty"`
	Line1            string                                `json:"line1,omitempty"`
	Line2            string                                `json:"line2,omitempty"`
	Line3            string                                `json:"line3,omitempty"`
	City             string                                `json:"city,omitempty"`
	StateCode        string                                `json:"state_code,omitempty"`
	State            string                                `json:"state,omitempty"`
	Zip              string                                `json:"zip,omitempty"`
	Country          string                                `json:"country,omitempty"`
	ValidationStatus InvoiceBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type InvoiceUpdateDetailsShippingAddress struct {
	FirstName        string                                 `json:"first_name,omitempty"`
	LastName         string                                 `json:"last_name,omitempty"`
	Email            string                                 `json:"email,omitempty"`
	Company          string                                 `json:"company,omitempty"`
	Phone            string                                 `json:"phone,omitempty"`
	Line1            string                                 `json:"line1,omitempty"`
	Line2            string                                 `json:"line2,omitempty"`
	Line3            string                                 `json:"line3,omitempty"`
	City             string                                 `json:"city,omitempty"`
	StateCode        string                                 `json:"state_code,omitempty"`
	State            string                                 `json:"state,omitempty"`
	Zip              string                                 `json:"zip,omitempty"`
	Country          string                                 `json:"country,omitempty"`
	ValidationStatus InvoiceShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type InvoiceUpdateDetailsStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}

type InvoiceApplyPaymentScheduleSchemeRequest struct {
	SchemeId   string `json:"scheme_id"`
	Amount     *int64 `json:"amount,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InvoiceApplyPaymentScheduleSchemeRequest) payload() any { return r }

// operation response
type InvoiceCreateResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceCreateForChargeItemsAndChargesResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceChargeResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceChargeAddonResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceCreateForChargeItemResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceStopDunningResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoicePauseDunningResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceResumeDunningResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceImportInvoiceResponse struct {
	Invoice    *Invoice   `json:"invoice,omitempty"`
	CreditNote CreditNote `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type InvoiceApplyPaymentsResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceSyncUsagesResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceDeleteLineItemsResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceApplyCreditsResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation sub response
type InvoiceListInvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

// operation response
type InvoiceListResponse struct {
	List       []*InvoiceListInvoiceResponse `json:"list,omitempty"`
	NextOffset string                        `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type InvoiceInvoicesForCustomerInvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

// operation response
type InvoiceInvoicesForCustomerResponse struct {
	List       []*InvoiceInvoicesForCustomerInvoiceResponse `json:"list,omitempty"`
	NextOffset string                                       `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type InvoiceInvoicesForSubscriptionInvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

// operation response
type InvoiceInvoicesForSubscriptionResponse struct {
	List       []*InvoiceInvoicesForSubscriptionInvoiceResponse `json:"list,omitempty"`
	NextOffset string                                           `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type InvoiceRetrieveResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoicePdfResponse struct {
	Download Download `json:"download,omitempty"`
	apiResponse
}

// operation response
type InvoiceDownloadEinvoiceResponse struct {
	Downloads []Download `json:"downloads,omitempty"`
	apiResponse
}

// operation sub response
type InvoiceListPaymentReferenceNumbersInvoiceResponse struct {
	PaymentReferenceNumber PaymentReferenceNumber `json:"payment_reference_number,omitempty"`
}

// operation response
type InvoiceListPaymentReferenceNumbersResponse struct {
	List       []*InvoiceListPaymentReferenceNumbersInvoiceResponse `json:"list,omitempty"`
	NextOffset string                                               `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type InvoiceAddChargeResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceAddAddonChargeResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceAddChargeItemResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceCloseResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceCollectPaymentResponse struct {
	Invoice     *Invoice    `json:"invoice,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type InvoiceRecordPaymentResponse struct {
	Invoice     *Invoice    `json:"invoice,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type InvoiceRecordTaxWithheldResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceRemoveTaxWithheldResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceRefundResponse struct {
	Invoice     *Invoice    `json:"invoice,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
	CreditNote  CreditNote  `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type InvoiceRecordRefundResponse struct {
	Invoice     *Invoice    `json:"invoice,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
	CreditNote  CreditNote  `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type InvoiceRemovePaymentResponse struct {
	Invoice     *Invoice    `json:"invoice,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type InvoiceRemoveCreditNoteResponse struct {
	Invoice    *Invoice   `json:"invoice,omitempty"`
	CreditNote CreditNote `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type InvoiceVoidInvoiceResponse struct {
	Invoice    *Invoice   `json:"invoice,omitempty"`
	CreditNote CreditNote `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type InvoiceWriteOffResponse struct {
	Invoice    *Invoice   `json:"invoice,omitempty"`
	CreditNote CreditNote `json:"credit_note,omitempty"`
	apiResponse
}

// operation response
type InvoiceDeleteResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceUpdateDetailsResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceApplyPaymentScheduleSchemeResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoicePaymentSchedulesResponse struct {
	PaymentSchedules []PaymentSchedule `json:"payment_schedules,omitempty"`
	apiResponse
}

// operation response
type InvoiceResendEinvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type InvoiceSendEinvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	apiResponse
}
