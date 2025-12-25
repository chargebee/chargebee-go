package chargebee

import (
	"encoding/json"
)

type QuoteStatus string

const (
	QuoteStatusOpen             QuoteStatus = "open"
	QuoteStatusAccepted         QuoteStatus = "accepted"
	QuoteStatusDeclined         QuoteStatus = "declined"
	QuoteStatusInvoiced         QuoteStatus = "invoiced"
	QuoteStatusClosed           QuoteStatus = "closed"
	QuoteStatusPendingApproval  QuoteStatus = "pending_approval"
	QuoteStatusApprovalRejected QuoteStatus = "approval_rejected"
	QuoteStatusProposed         QuoteStatus = "proposed"
	QuoteStatusVoided           QuoteStatus = "voided"
	QuoteStatusExpired          QuoteStatus = "expired"
)

type QuoteOperationType string

const (
	QuoteOperationTypeCreateSubscriptionForCustomer QuoteOperationType = "create_subscription_for_customer"
	QuoteOperationTypeChangeSubscription            QuoteOperationType = "change_subscription"
	QuoteOperationTypeOnetimeInvoice                QuoteOperationType = "onetime_invoice"
	QuoteOperationTypeRenewSubscription             QuoteOperationType = "renew_subscription"
)

type QuotePriceType string

const (
	QuotePriceTypeTaxExclusive QuotePriceType = "tax_exclusive"
	QuotePriceTypeTaxInclusive QuotePriceType = "tax_inclusive"
)

type QuoteLineItemPricingModel string

const (
	QuoteLineItemPricingModelFlatFee   QuoteLineItemPricingModel = "flat_fee"
	QuoteLineItemPricingModelPerUnit   QuoteLineItemPricingModel = "per_unit"
	QuoteLineItemPricingModelTiered    QuoteLineItemPricingModel = "tiered"
	QuoteLineItemPricingModelVolume    QuoteLineItemPricingModel = "volume"
	QuoteLineItemPricingModelStairstep QuoteLineItemPricingModel = "stairstep"
)

type QuoteLineItemEntityType string

const (
	QuoteLineItemEntityTypeAdhoc           QuoteLineItemEntityType = "adhoc"
	QuoteLineItemEntityTypePlanItemPrice   QuoteLineItemEntityType = "plan_item_price"
	QuoteLineItemEntityTypeAddonItemPrice  QuoteLineItemEntityType = "addon_item_price"
	QuoteLineItemEntityTypeChargeItemPrice QuoteLineItemEntityType = "charge_item_price"
	QuoteLineItemEntityTypePlanSetup       QuoteLineItemEntityType = "plan_setup"
	QuoteLineItemEntityTypePlan            QuoteLineItemEntityType = "plan"
	QuoteLineItemEntityTypeAddon           QuoteLineItemEntityType = "addon"
)

type QuoteLineItemTaxExemptReason string

const (
	QuoteLineItemTaxExemptReasonTaxNotConfigured                 QuoteLineItemTaxExemptReason = "tax_not_configured"
	QuoteLineItemTaxExemptReasonRegionNonTaxable                 QuoteLineItemTaxExemptReason = "region_non_taxable"
	QuoteLineItemTaxExemptReasonExport                           QuoteLineItemTaxExemptReason = "export"
	QuoteLineItemTaxExemptReasonCustomerExempt                   QuoteLineItemTaxExemptReason = "customer_exempt"
	QuoteLineItemTaxExemptReasonProductExempt                    QuoteLineItemTaxExemptReason = "product_exempt"
	QuoteLineItemTaxExemptReasonZeroRated                        QuoteLineItemTaxExemptReason = "zero_rated"
	QuoteLineItemTaxExemptReasonReverseCharge                    QuoteLineItemTaxExemptReason = "reverse_charge"
	QuoteLineItemTaxExemptReasonHighValuePhysicalGoods           QuoteLineItemTaxExemptReason = "high_value_physical_goods"
	QuoteLineItemTaxExemptReasonZeroValueItem                    QuoteLineItemTaxExemptReason = "zero_value_item"
	QuoteLineItemTaxExemptReasonTaxNotConfiguredExternalProvider QuoteLineItemTaxExemptReason = "tax_not_configured_external_provider"
)

type QuoteLineItemTierPricingType string

const (
	QuoteLineItemTierPricingTypePerUnit QuoteLineItemTierPricingType = "per_unit"
	QuoteLineItemTierPricingTypeFlatFee QuoteLineItemTierPricingType = "flat_fee"
	QuoteLineItemTierPricingTypePackage QuoteLineItemTierPricingType = "package"
)

type QuoteLineItemDiscountDiscountType string

const (
	QuoteLineItemDiscountDiscountTypeItemLevelCoupon       QuoteLineItemDiscountDiscountType = "item_level_coupon"
	QuoteLineItemDiscountDiscountTypeDocumentLevelCoupon   QuoteLineItemDiscountDiscountType = "document_level_coupon"
	QuoteLineItemDiscountDiscountTypePromotionalCredits    QuoteLineItemDiscountDiscountType = "promotional_credits"
	QuoteLineItemDiscountDiscountTypeProratedCredits       QuoteLineItemDiscountDiscountType = "prorated_credits"
	QuoteLineItemDiscountDiscountTypeItemLevelDiscount     QuoteLineItemDiscountDiscountType = "item_level_discount"
	QuoteLineItemDiscountDiscountTypeDocumentLevelDiscount QuoteLineItemDiscountDiscountType = "document_level_discount"
)

type QuoteLineItemTaxTaxJurisType string

const (
	QuoteLineItemTaxTaxJurisTypeCountry        QuoteLineItemTaxTaxJurisType = "country"
	QuoteLineItemTaxTaxJurisTypeFederal        QuoteLineItemTaxTaxJurisType = "federal"
	QuoteLineItemTaxTaxJurisTypeState          QuoteLineItemTaxTaxJurisType = "state"
	QuoteLineItemTaxTaxJurisTypeCounty         QuoteLineItemTaxTaxJurisType = "county"
	QuoteLineItemTaxTaxJurisTypeCity           QuoteLineItemTaxTaxJurisType = "city"
	QuoteLineItemTaxTaxJurisTypeSpecial        QuoteLineItemTaxTaxJurisType = "special"
	QuoteLineItemTaxTaxJurisTypeUnincorporated QuoteLineItemTaxTaxJurisType = "unincorporated"
	QuoteLineItemTaxTaxJurisTypeOther          QuoteLineItemTaxTaxJurisType = "other"
)

type QuoteDiscountEntityType string

const (
	QuoteDiscountEntityTypeItemLevelCoupon       QuoteDiscountEntityType = "item_level_coupon"
	QuoteDiscountEntityTypeDocumentLevelCoupon   QuoteDiscountEntityType = "document_level_coupon"
	QuoteDiscountEntityTypePromotionalCredits    QuoteDiscountEntityType = "promotional_credits"
	QuoteDiscountEntityTypeProratedCredits       QuoteDiscountEntityType = "prorated_credits"
	QuoteDiscountEntityTypeItemLevelDiscount     QuoteDiscountEntityType = "item_level_discount"
	QuoteDiscountEntityTypeDocumentLevelDiscount QuoteDiscountEntityType = "document_level_discount"
)

type QuoteDiscountDiscountType string

const (
	QuoteDiscountDiscountTypeFixedAmount QuoteDiscountDiscountType = "fixed_amount"
	QuoteDiscountDiscountTypePercentage  QuoteDiscountDiscountType = "percentage"
)

type QuoteShippingAddressValidationStatus string

const (
	QuoteShippingAddressValidationStatusNotValidated   QuoteShippingAddressValidationStatus = "not_validated"
	QuoteShippingAddressValidationStatusValid          QuoteShippingAddressValidationStatus = "valid"
	QuoteShippingAddressValidationStatusPartiallyValid QuoteShippingAddressValidationStatus = "partially_valid"
	QuoteShippingAddressValidationStatusInvalid        QuoteShippingAddressValidationStatus = "invalid"
)

type QuoteBillingAddressValidationStatus string

const (
	QuoteBillingAddressValidationStatusNotValidated   QuoteBillingAddressValidationStatus = "not_validated"
	QuoteBillingAddressValidationStatusValid          QuoteBillingAddressValidationStatus = "valid"
	QuoteBillingAddressValidationStatusPartiallyValid QuoteBillingAddressValidationStatus = "partially_valid"
	QuoteBillingAddressValidationStatusInvalid        QuoteBillingAddressValidationStatus = "invalid"
)

type QuoteBillingAlignmentMode string

const (
	QuoteBillingAlignmentModeImmediate QuoteBillingAlignmentMode = "immediate"
	QuoteBillingAlignmentModeDelayed   QuoteBillingAlignmentMode = "delayed"
)

type QuoteChangeOption string

const (
	QuoteChangeOptionImmediately  QuoteChangeOption = "immediately"
	QuoteChangeOptionSpecificDate QuoteChangeOption = "specific_date"
)

type QuoteBillingStartOption string

const (
	QuoteBillingStartOptionImmediately    QuoteBillingStartOption = "immediately"
	QuoteBillingStartOptionOnSpecificDate QuoteBillingStartOption = "on_specific_date"
)

type QuoteDispositionType string

const (
	QuoteDispositionTypeAttachment QuoteDispositionType = "attachment"
	QuoteDispositionTypeInline     QuoteDispositionType = "inline"
)

// just struct
type Quote struct {
	Id                         string                   `json:"id"`
	Name                       string                   `json:"name"`
	PoNumber                   string                   `json:"po_number"`
	CustomerId                 string                   `json:"customer_id"`
	SubscriptionId             string                   `json:"subscription_id"`
	InvoiceId                  string                   `json:"invoice_id"`
	Status                     QuoteStatus              `json:"status"`
	OperationType              QuoteOperationType       `json:"operation_type"`
	VatNumber                  string                   `json:"vat_number"`
	PriceType                  QuotePriceType           `json:"price_type"`
	ValidTill                  int64                    `json:"valid_till"`
	Date                       int64                    `json:"date"`
	TotalPayable               int64                    `json:"total_payable"`
	ChargeOnAcceptance         int64                    `json:"charge_on_acceptance"`
	SubTotal                   int64                    `json:"sub_total"`
	Total                      int64                    `json:"total"`
	CreditsApplied             int64                    `json:"credits_applied"`
	AmountPaid                 int64                    `json:"amount_paid"`
	AmountDue                  int64                    `json:"amount_due"`
	Version                    int32                    `json:"version"`
	ResourceVersion            int64                    `json:"resource_version"`
	UpdatedAt                  int64                    `json:"updated_at"`
	VatNumberPrefix            string                   `json:"vat_number_prefix"`
	LineItems                  []*QuoteLineItem         `json:"line_items"`
	LineItemTiers              []*QuoteLineItemTier     `json:"line_item_tiers"`
	LineItemDiscounts          []*QuoteLineItemDiscount `json:"line_item_discounts"`
	LineItemTaxes              []*QuoteLineItemTax      `json:"line_item_taxes"`
	Discounts                  []*QuoteDiscount         `json:"discounts"`
	Taxes                      []*QuoteTax              `json:"taxes"`
	TaxCategory                string                   `json:"tax_category"`
	CurrencyCode               string                   `json:"currency_code"`
	Notes                      json.RawMessage          `json:"notes"`
	ShippingAddress            *QuoteShippingAddress    `json:"shipping_address"`
	BillingAddress             *QuoteBillingAddress     `json:"billing_address"`
	ContractTermStart          int64                    `json:"contract_term_start"`
	ContractTermEnd            int64                    `json:"contract_term_end"`
	ContractTermTerminationFee int64                    `json:"contract_term_termination_fee"`
	BusinessEntityId           string                   `json:"business_entity_id"`
	Deleted                    bool                     `json:"deleted"`
	TotalContractValue         int64                    `json:"total_contract_value"`
	TotalDiscount              int64                    `json:"total_discount"`
	CustomField                CustomField              `json:"custom_field"`
	Object                     string                   `json:"object"`
}

// sub resources
type QuoteLineItem struct {
	Id                      string                       `json:"id"`
	SubscriptionId          string                       `json:"subscription_id"`
	DateFrom                int64                        `json:"date_from"`
	DateTo                  int64                        `json:"date_to"`
	UnitAmount              int64                        `json:"unit_amount"`
	Quantity                int32                        `json:"quantity"`
	Amount                  int64                        `json:"amount"`
	PricingModel            QuoteLineItemPricingModel    `json:"pricing_model"`
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
	EntityType              QuoteLineItemEntityType      `json:"entity_type"`
	TaxExemptReason         QuoteLineItemTaxExemptReason `json:"tax_exempt_reason"`
	EntityId                string                       `json:"entity_id"`
	CustomerId              string                       `json:"customer_id"`
	Object                  string                       `json:"object"`
}
type QuoteLineItemTier struct {
	LineItemId            string                       `json:"line_item_id"`
	StartingUnit          int32                        `json:"starting_unit"`
	EndingUnit            int32                        `json:"ending_unit"`
	QuantityUsed          int32                        `json:"quantity_used"`
	UnitAmount            int64                        `json:"unit_amount"`
	StartingUnitInDecimal string                       `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string                       `json:"ending_unit_in_decimal"`
	QuantityUsedInDecimal string                       `json:"quantity_used_in_decimal"`
	UnitAmountInDecimal   string                       `json:"unit_amount_in_decimal"`
	PricingType           QuoteLineItemTierPricingType `json:"pricing_type"`
	PackageSize           int32                        `json:"package_size"`
	Object                string                       `json:"object"`
}
type QuoteLineItemDiscount struct {
	LineItemId     string                            `json:"line_item_id"`
	DiscountType   QuoteLineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                            `json:"coupon_id"`
	EntityId       string                            `json:"entity_id"`
	DiscountAmount int64                             `json:"discount_amount"`
	Object         string                            `json:"object"`
}
type QuoteLineItemTax struct {
	LineItemId               string                       `json:"line_item_id"`
	TaxName                  string                       `json:"tax_name"`
	TaxRate                  float64                      `json:"tax_rate"`
	DateTo                   int64                        `json:"date_to"`
	DateFrom                 int64                        `json:"date_from"`
	ProratedTaxableAmount    float64                      `json:"prorated_taxable_amount"`
	IsPartialTaxApplied      bool                         `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool                         `json:"is_non_compliance_tax"`
	TaxableAmount            int64                        `json:"taxable_amount"`
	TaxAmount                int64                        `json:"tax_amount"`
	TaxJurisType             QuoteLineItemTaxTaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string                       `json:"tax_juris_name"`
	TaxJurisCode             string                       `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int64                        `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string                       `json:"local_currency_code"`
	Object                   string                       `json:"object"`
}
type QuoteDiscount struct {
	Amount        int64                     `json:"amount"`
	Description   string                    `json:"description"`
	LineItemId    string                    `json:"line_item_id"`
	EntityType    QuoteDiscountEntityType   `json:"entity_type"`
	DiscountType  QuoteDiscountDiscountType `json:"discount_type"`
	EntityId      string                    `json:"entity_id"`
	CouponSetCode string                    `json:"coupon_set_code"`
	Object        string                    `json:"object"`
}
type QuoteTax struct {
	Name        string `json:"name"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}
type QuoteShippingAddress struct {
	FirstName        string                               `json:"first_name"`
	LastName         string                               `json:"last_name"`
	Email            string                               `json:"email"`
	Company          string                               `json:"company"`
	Phone            string                               `json:"phone"`
	Line1            string                               `json:"line1"`
	Line2            string                               `json:"line2"`
	Line3            string                               `json:"line3"`
	City             string                               `json:"city"`
	StateCode        string                               `json:"state_code"`
	State            string                               `json:"state"`
	Country          string                               `json:"country"`
	Zip              string                               `json:"zip"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status"`
	Object           string                               `json:"object"`
}
type QuoteBillingAddress struct {
	FirstName        string                              `json:"first_name"`
	LastName         string                              `json:"last_name"`
	Email            string                              `json:"email"`
	Company          string                              `json:"company"`
	Phone            string                              `json:"phone"`
	Line1            string                              `json:"line1"`
	Line2            string                              `json:"line2"`
	Line3            string                              `json:"line3"`
	City             string                              `json:"city"`
	StateCode        string                              `json:"state_code"`
	State            string                              `json:"state"`
	Country          string                              `json:"country"`
	Zip              string                              `json:"zip"`
	ValidationStatus QuoteBillingAddressValidationStatus `json:"validation_status"`
	Object           string                              `json:"object"`
}

// operations
// input params
type QuoteCreateSubForCustomerQuoteRequest struct {
	Name                    string                                           `json:"name,omitempty"`
	Notes                   string                                           `json:"notes,omitempty"`
	ExpiresAt               *int64                                           `json:"expires_at,omitempty"`
	Subscription            *QuoteCreateSubForCustomerQuoteSubscription      `json:"subscription,omitempty"`
	BillingCycles           *int32                                           `json:"billing_cycles,omitempty"`
	Addons                  []*QuoteCreateSubForCustomerQuoteAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*QuoteCreateSubForCustomerQuoteEventBasedAddon `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                                         `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                           `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    QuoteBillingAlignmentMode                        `json:"billing_alignment_mode,omitempty"`
	ShippingAddress         *QuoteCreateSubForCustomerQuoteShippingAddress   `json:"shipping_address,omitempty"`
	ContractTerm            *QuoteCreateSubForCustomerQuoteContractTerm      `json:"contract_term,omitempty"`
	CouponIds               []string                                         `json:"coupon_ids,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *QuoteCreateSubForCustomerQuoteRequest) payload() any { return r }

// input sub resource params single
type QuoteCreateSubForCustomerQuoteSubscription struct {
	Id                                string                                `json:"id,omitempty"`
	PoNumber                          string                                `json:"po_number,omitempty"`
	PlanId                            string                                `json:"plan_id"`
	PlanQuantity                      *int32                                `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                                `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                                `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                `json:"trial_end,omitempty"`
	StartDate                         *int64                                `json:"start_date,omitempty"`
	OfflinePaymentMethod              QuoteSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params multi
type QuoteCreateSubForCustomerQuoteAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}

// input sub resource params multi
type QuoteCreateSubForCustomerQuoteEventBasedAddon struct {
	Id                  string                       `json:"id,omitempty"`
	Quantity            *int32                       `json:"quantity,omitempty"`
	UnitPrice           *int64                       `json:"unit_price,omitempty"`
	QuantityInDecimal   string                       `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                       `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32                       `json:"service_period_in_days,omitempty"`
	OnEvent             QuoteEventBasedAddonOnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool                        `json:"charge_once,omitempty"`
	ChargeOn            QuoteEventBasedAddonChargeOn `json:"charge_on,omitempty"`
}

// input sub resource params single
type QuoteCreateSubForCustomerQuoteShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteCreateSubForCustomerQuoteContractTerm struct {
	ActionAtTermEnd          QuoteContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type QuoteEditCreateSubForCustomerQuoteRequest struct {
	Notes                   string                                               `json:"notes,omitempty"`
	ExpiresAt               *int64                                               `json:"expires_at,omitempty"`
	Subscription            *QuoteEditCreateSubForCustomerQuoteSubscription      `json:"subscription,omitempty"`
	BillingCycles           *int32                                               `json:"billing_cycles,omitempty"`
	Addons                  []*QuoteEditCreateSubForCustomerQuoteAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*QuoteEditCreateSubForCustomerQuoteEventBasedAddon `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                                             `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                               `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    QuoteBillingAlignmentMode                            `json:"billing_alignment_mode,omitempty"`
	ShippingAddress         *QuoteEditCreateSubForCustomerQuoteShippingAddress   `json:"shipping_address,omitempty"`
	ContractTerm            *QuoteEditCreateSubForCustomerQuoteContractTerm      `json:"contract_term,omitempty"`
	CouponIds               []string                                             `json:"coupon_ids,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *QuoteEditCreateSubForCustomerQuoteRequest) payload() any { return r }

// input sub resource params single
type QuoteEditCreateSubForCustomerQuoteSubscription struct {
	Id                                string                                `json:"id,omitempty"`
	PoNumber                          string                                `json:"po_number,omitempty"`
	PlanId                            string                                `json:"plan_id"`
	PlanQuantity                      *int32                                `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                                `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                                `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                `json:"trial_end,omitempty"`
	StartDate                         *int64                                `json:"start_date,omitempty"`
	OfflinePaymentMethod              QuoteSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params multi
type QuoteEditCreateSubForCustomerQuoteAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}

// input sub resource params multi
type QuoteEditCreateSubForCustomerQuoteEventBasedAddon struct {
	Id                  string                       `json:"id,omitempty"`
	Quantity            *int32                       `json:"quantity,omitempty"`
	UnitPrice           *int64                       `json:"unit_price,omitempty"`
	QuantityInDecimal   string                       `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                       `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32                       `json:"service_period_in_days,omitempty"`
	OnEvent             QuoteEventBasedAddonOnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool                        `json:"charge_once,omitempty"`
	ChargeOn            QuoteEventBasedAddonChargeOn `json:"charge_on,omitempty"`
}

// input sub resource params single
type QuoteEditCreateSubForCustomerQuoteShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteEditCreateSubForCustomerQuoteContractTerm struct {
	ActionAtTermEnd          QuoteContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type QuoteUpdateSubscriptionQuoteRequest struct {
	Name                    string                                         `json:"name,omitempty"`
	Notes                   string                                         `json:"notes,omitempty"`
	ExpiresAt               *int64                                         `json:"expires_at,omitempty"`
	Subscription            *QuoteUpdateSubscriptionQuoteSubscription      `json:"subscription,omitempty"`
	Addons                  []*QuoteUpdateSubscriptionQuoteAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*QuoteUpdateSubscriptionQuoteEventBasedAddon `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                          `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                       `json:"mandatory_addons_to_remove,omitempty"`
	BillingCycles           *int32                                         `json:"billing_cycles,omitempty"`
	TermsToCharge           *int32                                         `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                         `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    QuoteBillingAlignmentMode                      `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                       `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                          `json:"replace_coupon_list,omitempty"`
	ChangeOption            QuoteChangeOption                              `json:"change_option,omitempty"`
	ChangesScheduledAt      *int64                                         `json:"changes_scheduled_at,omitempty"`
	ForceTermReset          *bool                                          `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                          `json:"reactivate,omitempty"`
	BillingAddress          *QuoteUpdateSubscriptionQuoteBillingAddress    `json:"billing_address,omitempty"`
	ShippingAddress         *QuoteUpdateSubscriptionQuoteShippingAddress   `json:"shipping_address,omitempty"`
	Customer                *QuoteUpdateSubscriptionQuoteCustomer          `json:"customer,omitempty"`
	ContractTerm            *QuoteUpdateSubscriptionQuoteContractTerm      `json:"contract_term,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *QuoteUpdateSubscriptionQuoteRequest) payload() any { return r }

// input sub resource params single
type QuoteUpdateSubscriptionQuoteSubscription struct {
	Id                                string                                `json:"id"`
	PlanId                            string                                `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                                `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int64                                `json:"plan_unit_price,omitempty"`
	SetupFee                          *int64                                `json:"setup_fee,omitempty"`
	PlanQuantityInDecimal             string                                `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                                `json:"plan_unit_price_in_decimal,omitempty"`
	StartDate                         *int64                                `json:"start_date,omitempty"`
	TrialEnd                          *int64                                `json:"trial_end,omitempty"`
	Coupon                            string                                `json:"coupon,omitempty"`
	AutoCollection                    QuoteSubscriptionAutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              QuoteSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params multi
type QuoteUpdateSubscriptionQuoteAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}

// input sub resource params multi
type QuoteUpdateSubscriptionQuoteEventBasedAddon struct {
	Id                  string                       `json:"id,omitempty"`
	Quantity            *int32                       `json:"quantity,omitempty"`
	UnitPrice           *int64                       `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32                       `json:"service_period_in_days,omitempty"`
	ChargeOn            QuoteEventBasedAddonChargeOn `json:"charge_on,omitempty"`
	OnEvent             QuoteEventBasedAddonOnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool                        `json:"charge_once,omitempty"`
	QuantityInDecimal   string                       `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                       `json:"unit_price_in_decimal,omitempty"`
}

// input sub resource params single
type QuoteUpdateSubscriptionQuoteBillingAddress struct {
	FirstName        string                              `json:"first_name,omitempty"`
	LastName         string                              `json:"last_name,omitempty"`
	Email            string                              `json:"email,omitempty"`
	Company          string                              `json:"company,omitempty"`
	Phone            string                              `json:"phone,omitempty"`
	Line1            string                              `json:"line1,omitempty"`
	Line2            string                              `json:"line2,omitempty"`
	Line3            string                              `json:"line3,omitempty"`
	City             string                              `json:"city,omitempty"`
	StateCode        string                              `json:"state_code,omitempty"`
	State            string                              `json:"state,omitempty"`
	Zip              string                              `json:"zip,omitempty"`
	Country          string                              `json:"country,omitempty"`
	ValidationStatus QuoteBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteUpdateSubscriptionQuoteShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteUpdateSubscriptionQuoteCustomer struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}

// input sub resource params single
type QuoteUpdateSubscriptionQuoteContractTerm struct {
	ActionAtTermEnd          QuoteContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type QuoteEditUpdateSubscriptionQuoteRequest struct {
	Notes                   string                                             `json:"notes,omitempty"`
	ExpiresAt               *int64                                             `json:"expires_at,omitempty"`
	Subscription            *QuoteEditUpdateSubscriptionQuoteSubscription      `json:"subscription,omitempty"`
	Addons                  []*QuoteEditUpdateSubscriptionQuoteAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*QuoteEditUpdateSubscriptionQuoteEventBasedAddon `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                              `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                           `json:"mandatory_addons_to_remove,omitempty"`
	BillingCycles           *int32                                             `json:"billing_cycles,omitempty"`
	TermsToCharge           *int32                                             `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                             `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    QuoteBillingAlignmentMode                          `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                           `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                              `json:"replace_coupon_list,omitempty"`
	ChangeOption            QuoteChangeOption                                  `json:"change_option,omitempty"`
	ChangesScheduledAt      *int64                                             `json:"changes_scheduled_at,omitempty"`
	ForceTermReset          *bool                                              `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                              `json:"reactivate,omitempty"`
	BillingAddress          *QuoteEditUpdateSubscriptionQuoteBillingAddress    `json:"billing_address,omitempty"`
	ShippingAddress         *QuoteEditUpdateSubscriptionQuoteShippingAddress   `json:"shipping_address,omitempty"`
	Customer                *QuoteEditUpdateSubscriptionQuoteCustomer          `json:"customer,omitempty"`
	ContractTerm            *QuoteEditUpdateSubscriptionQuoteContractTerm      `json:"contract_term,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *QuoteEditUpdateSubscriptionQuoteRequest) payload() any { return r }

// input sub resource params single
type QuoteEditUpdateSubscriptionQuoteSubscription struct {
	PlanId                            string                                `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                                `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int64                                `json:"plan_unit_price,omitempty"`
	SetupFee                          *int64                                `json:"setup_fee,omitempty"`
	PlanQuantityInDecimal             string                                `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                                `json:"plan_unit_price_in_decimal,omitempty"`
	StartDate                         *int64                                `json:"start_date,omitempty"`
	TrialEnd                          *int64                                `json:"trial_end,omitempty"`
	Coupon                            string                                `json:"coupon,omitempty"`
	AutoCollection                    QuoteSubscriptionAutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              QuoteSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params multi
type QuoteEditUpdateSubscriptionQuoteAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}

// input sub resource params multi
type QuoteEditUpdateSubscriptionQuoteEventBasedAddon struct {
	Id                  string                       `json:"id,omitempty"`
	Quantity            *int32                       `json:"quantity,omitempty"`
	UnitPrice           *int64                       `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32                       `json:"service_period_in_days,omitempty"`
	ChargeOn            QuoteEventBasedAddonChargeOn `json:"charge_on,omitempty"`
	OnEvent             QuoteEventBasedAddonOnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool                        `json:"charge_once,omitempty"`
	QuantityInDecimal   string                       `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                       `json:"unit_price_in_decimal,omitempty"`
}

// input sub resource params single
type QuoteEditUpdateSubscriptionQuoteBillingAddress struct {
	FirstName        string                              `json:"first_name,omitempty"`
	LastName         string                              `json:"last_name,omitempty"`
	Email            string                              `json:"email,omitempty"`
	Company          string                              `json:"company,omitempty"`
	Phone            string                              `json:"phone,omitempty"`
	Line1            string                              `json:"line1,omitempty"`
	Line2            string                              `json:"line2,omitempty"`
	Line3            string                              `json:"line3,omitempty"`
	City             string                              `json:"city,omitempty"`
	StateCode        string                              `json:"state_code,omitempty"`
	State            string                              `json:"state,omitempty"`
	Zip              string                              `json:"zip,omitempty"`
	Country          string                              `json:"country,omitempty"`
	ValidationStatus QuoteBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteEditUpdateSubscriptionQuoteShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteEditUpdateSubscriptionQuoteCustomer struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}

// input sub resource params single
type QuoteEditUpdateSubscriptionQuoteContractTerm struct {
	ActionAtTermEnd          QuoteContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type QuoteCreateForOnetimeChargesRequest struct {
	Name               string                                           `json:"name,omitempty"`
	CustomerId         string                                           `json:"customer_id"`
	PoNumber           string                                           `json:"po_number,omitempty"`
	Notes              string                                           `json:"notes,omitempty"`
	ExpiresAt          *int64                                           `json:"expires_at,omitempty"`
	CurrencyCode       string                                           `json:"currency_code,omitempty"`
	Addons             []*QuoteCreateForOnetimeChargesAddon             `json:"addons,omitempty"`
	Charges            []*QuoteCreateForOnetimeChargesCharge            `json:"charges,omitempty"`
	Coupon             string                                           `json:"coupon,omitempty"`
	CouponIds          []string                                         `json:"coupon_ids,omitempty"`
	ShippingAddress    *QuoteCreateForOnetimeChargesShippingAddress     `json:"shipping_address,omitempty"`
	TaxProvidersFields []*QuoteCreateForOnetimeChargesTaxProvidersField `json:"tax_providers_fields,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *QuoteCreateForOnetimeChargesRequest) payload() any { return r }

// input sub resource params multi
type QuoteCreateForOnetimeChargesAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriod      *int32 `json:"service_period,omitempty"`
}

// input sub resource params multi
type QuoteCreateForOnetimeChargesCharge struct {
	Amount                 *int64                     `json:"amount,omitempty"`
	AmountInDecimal        string                     `json:"amount_in_decimal,omitempty"`
	Description            string                     `json:"description,omitempty"`
	AvalaraSaleType        QuoteChargeAvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                     `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                     `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32                     `json:"service_period,omitempty"`
}

// input sub resource params single
type QuoteCreateForOnetimeChargesShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type QuoteCreateForOnetimeChargesTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type QuoteEditOneTimeQuoteRequest struct {
	PoNumber           string                                    `json:"po_number,omitempty"`
	Notes              string                                    `json:"notes,omitempty"`
	ExpiresAt          *int64                                    `json:"expires_at,omitempty"`
	CurrencyCode       string                                    `json:"currency_code,omitempty"`
	Addons             []*QuoteEditOneTimeQuoteAddon             `json:"addons,omitempty"`
	Charges            []*QuoteEditOneTimeQuoteCharge            `json:"charges,omitempty"`
	Coupon             string                                    `json:"coupon,omitempty"`
	CouponIds          []string                                  `json:"coupon_ids,omitempty"`
	ShippingAddress    *QuoteEditOneTimeQuoteShippingAddress     `json:"shipping_address,omitempty"`
	TaxProvidersFields []*QuoteEditOneTimeQuoteTaxProvidersField `json:"tax_providers_fields,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *QuoteEditOneTimeQuoteRequest) payload() any { return r }

// input sub resource params multi
type QuoteEditOneTimeQuoteAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriod      *int32 `json:"service_period,omitempty"`
}

// input sub resource params multi
type QuoteEditOneTimeQuoteCharge struct {
	Amount                 *int64                     `json:"amount,omitempty"`
	AmountInDecimal        string                     `json:"amount_in_decimal,omitempty"`
	Description            string                     `json:"description,omitempty"`
	AvalaraSaleType        QuoteChargeAvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                     `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                     `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32                     `json:"service_period,omitempty"`
}

// input sub resource params single
type QuoteEditOneTimeQuoteShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type QuoteEditOneTimeQuoteTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type QuoteCreateSubItemsForCustomerQuoteRequest struct {
	Name                   string                                                 `json:"name,omitempty"`
	Notes                  string                                                 `json:"notes,omitempty"`
	ExpiresAt              *int64                                                 `json:"expires_at,omitempty"`
	Subscription           *QuoteCreateSubItemsForCustomerQuoteSubscription       `json:"subscription,omitempty"`
	BillingCycles          *int32                                                 `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*QuoteCreateSubItemsForCustomerQuoteSubscriptionItem `json:"subscription_items,omitempty"`
	Discounts              []*QuoteCreateSubItemsForCustomerQuoteDiscount         `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                               `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*QuoteCreateSubItemsForCustomerQuoteItemTier         `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                                 `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   QuoteBillingAlignmentMode                              `json:"billing_alignment_mode,omitempty"`
	ShippingAddress        *QuoteCreateSubItemsForCustomerQuoteShippingAddress    `json:"shipping_address,omitempty"`
	ContractTerm           *QuoteCreateSubItemsForCustomerQuoteContractTerm       `json:"contract_term,omitempty"`
	CouponIds              []string                                               `json:"coupon_ids,omitempty"`
	BillingStartOption     QuoteBillingStartOption                                `json:"billing_start_option,omitempty"`
	BillingAddress         *QuoteCreateSubItemsForCustomerQuoteBillingAddress     `json:"billing_address,omitempty"`
	NetTermDays            *int32                                                 `json:"net_term_days,omitempty"`
	Coupons                []*QuoteCreateSubItemsForCustomerQuoteCoupon           `json:"coupons,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *QuoteCreateSubItemsForCustomerQuoteRequest) payload() any { return r }

// input sub resource params single
type QuoteCreateSubItemsForCustomerQuoteSubscription struct {
	Id                                string                                `json:"id,omitempty"`
	PoNumber                          string                                `json:"po_number,omitempty"`
	TrialEnd                          *int64                                `json:"trial_end,omitempty"`
	SetupFee                          *int64                                `json:"setup_fee,omitempty"`
	StartDate                         *int64                                `json:"start_date,omitempty"`
	OfflinePaymentMethod              QuoteSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params multi
type QuoteCreateSubItemsForCustomerQuoteSubscriptionItem struct {
	ItemPriceId        string                              `json:"item_price_id"`
	Quantity           *int32                              `json:"quantity,omitempty"`
	QuantityInDecimal  string                              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64                              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string                              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32                              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64                              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32                              `json:"service_period_days,omitempty"`
	ChargeOnEvent      QuoteSubscriptionItemChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool                               `json:"charge_once,omitempty"`
	ItemType           QuoteSubscriptionItemItemType       `json:"item_type,omitempty"`
	ChargeOnOption     QuoteSubscriptionItemChargeOnOption `json:"charge_on_option,omitempty"`
	StartDate          *int64                              `json:"start_date,omitempty"`
	EndDate            *int64                              `json:"end_date,omitempty"`
	RampTierId         string                              `json:"ramp_tier_id,omitempty"`
}

// input sub resource params multi
type QuoteCreateSubItemsForCustomerQuoteDiscount struct {
	ApplyOn       QuoteDiscountApplyOn      `json:"apply_on,omitempty"`
	DurationType  QuoteDiscountDurationType `json:"duration_type"`
	Percentage    *float64                  `json:"percentage,omitempty"`
	Amount        *int64                    `json:"amount,omitempty"`
	Period        *int32                    `json:"period,omitempty"`
	PeriodUnit    QuoteDiscountPeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool                     `json:"included_in_mrr,omitempty"`
	ItemPriceId   string                    `json:"item_price_id,omitempty"`
	Quantity      *int32                    `json:"quantity,omitempty"`
	StartDate     *int64                    `json:"start_date,omitempty"`
	EndDate       *int64                    `json:"end_date,omitempty"`
}

// input sub resource params multi
type QuoteCreateSubItemsForCustomerQuoteItemTier struct {
	ItemPriceId           string                   `json:"item_price_id,omitempty"`
	StartingUnit          *int32                   `json:"starting_unit,omitempty"`
	EndingUnit            *int32                   `json:"ending_unit,omitempty"`
	Price                 *int64                   `json:"price,omitempty"`
	StartingUnitInDecimal string                   `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                   `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                   `json:"price_in_decimal,omitempty"`
	PricingType           QuoteItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                   `json:"package_size,omitempty"`
	RampTierId            string                   `json:"ramp_tier_id,omitempty"`
}

// input sub resource params single
type QuoteCreateSubItemsForCustomerQuoteShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteCreateSubItemsForCustomerQuoteContractTerm struct {
	ActionAtTermEnd          QuoteContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params single
type QuoteCreateSubItemsForCustomerQuoteBillingAddress struct {
	FirstName        string                              `json:"first_name,omitempty"`
	LastName         string                              `json:"last_name,omitempty"`
	Email            string                              `json:"email,omitempty"`
	Company          string                              `json:"company,omitempty"`
	Phone            string                              `json:"phone,omitempty"`
	Line1            string                              `json:"line1,omitempty"`
	Line2            string                              `json:"line2,omitempty"`
	Line3            string                              `json:"line3,omitempty"`
	City             string                              `json:"city,omitempty"`
	StateCode        string                              `json:"state_code,omitempty"`
	State            string                              `json:"state,omitempty"`
	Zip              string                              `json:"zip,omitempty"`
	Country          string                              `json:"country,omitempty"`
	ValidationStatus QuoteBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type QuoteCreateSubItemsForCustomerQuoteCoupon struct {
	Id        string `json:"id,omitempty"`
	StartDate *int64 `json:"start_date,omitempty"`
	EndDate   *int64 `json:"end_date,omitempty"`
}
type QuoteEditCreateSubCustomerQuoteForItemsRequest struct {
	Notes                  string                                                     `json:"notes,omitempty"`
	ExpiresAt              *int64                                                     `json:"expires_at,omitempty"`
	Subscription           *QuoteEditCreateSubCustomerQuoteForItemsSubscription       `json:"subscription,omitempty"`
	BillingCycles          *int32                                                     `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*QuoteEditCreateSubCustomerQuoteForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	Discounts              []*QuoteEditCreateSubCustomerQuoteForItemsDiscount         `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                                   `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*QuoteEditCreateSubCustomerQuoteForItemsItemTier         `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                                     `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   QuoteBillingAlignmentMode                                  `json:"billing_alignment_mode,omitempty"`
	ShippingAddress        *QuoteEditCreateSubCustomerQuoteForItemsShippingAddress    `json:"shipping_address,omitempty"`
	ContractTerm           *QuoteEditCreateSubCustomerQuoteForItemsContractTerm       `json:"contract_term,omitempty"`
	CouponIds              []string                                                   `json:"coupon_ids,omitempty"`
	BillingStartOption     QuoteBillingStartOption                                    `json:"billing_start_option,omitempty"`
	BillingAddress         *QuoteEditCreateSubCustomerQuoteForItemsBillingAddress     `json:"billing_address,omitempty"`
	NetTermDays            *int32                                                     `json:"net_term_days,omitempty"`
	Coupons                []*QuoteEditCreateSubCustomerQuoteForItemsCoupon           `json:"coupons,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *QuoteEditCreateSubCustomerQuoteForItemsRequest) payload() any { return r }

// input sub resource params single
type QuoteEditCreateSubCustomerQuoteForItemsSubscription struct {
	Id                                string                                `json:"id,omitempty"`
	PoNumber                          string                                `json:"po_number,omitempty"`
	TrialEnd                          *int64                                `json:"trial_end,omitempty"`
	SetupFee                          *int64                                `json:"setup_fee,omitempty"`
	StartDate                         *int64                                `json:"start_date,omitempty"`
	OfflinePaymentMethod              QuoteSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params multi
type QuoteEditCreateSubCustomerQuoteForItemsSubscriptionItem struct {
	ItemPriceId        string                              `json:"item_price_id"`
	Quantity           *int32                              `json:"quantity,omitempty"`
	QuantityInDecimal  string                              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64                              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string                              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32                              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64                              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32                              `json:"service_period_days,omitempty"`
	ChargeOnEvent      QuoteSubscriptionItemChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool                               `json:"charge_once,omitempty"`
	ItemType           QuoteSubscriptionItemItemType       `json:"item_type,omitempty"`
	ChargeOnOption     QuoteSubscriptionItemChargeOnOption `json:"charge_on_option,omitempty"`
	StartDate          *int64                              `json:"start_date,omitempty"`
	EndDate            *int64                              `json:"end_date,omitempty"`
	RampTierId         string                              `json:"ramp_tier_id,omitempty"`
}

// input sub resource params multi
type QuoteEditCreateSubCustomerQuoteForItemsDiscount struct {
	ApplyOn       QuoteDiscountApplyOn      `json:"apply_on,omitempty"`
	DurationType  QuoteDiscountDurationType `json:"duration_type"`
	Percentage    *float64                  `json:"percentage,omitempty"`
	Amount        *int64                    `json:"amount,omitempty"`
	Period        *int32                    `json:"period,omitempty"`
	PeriodUnit    QuoteDiscountPeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool                     `json:"included_in_mrr,omitempty"`
	ItemPriceId   string                    `json:"item_price_id,omitempty"`
	Quantity      *int32                    `json:"quantity,omitempty"`
	StartDate     *int64                    `json:"start_date,omitempty"`
	EndDate       *int64                    `json:"end_date,omitempty"`
}

// input sub resource params multi
type QuoteEditCreateSubCustomerQuoteForItemsItemTier struct {
	ItemPriceId           string                   `json:"item_price_id,omitempty"`
	StartingUnit          *int32                   `json:"starting_unit,omitempty"`
	EndingUnit            *int32                   `json:"ending_unit,omitempty"`
	Price                 *int64                   `json:"price,omitempty"`
	StartingUnitInDecimal string                   `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                   `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                   `json:"price_in_decimal,omitempty"`
	PricingType           QuoteItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                   `json:"package_size,omitempty"`
	RampTierId            string                   `json:"ramp_tier_id,omitempty"`
}

// input sub resource params single
type QuoteEditCreateSubCustomerQuoteForItemsShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteEditCreateSubCustomerQuoteForItemsContractTerm struct {
	ActionAtTermEnd          QuoteContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params single
type QuoteEditCreateSubCustomerQuoteForItemsBillingAddress struct {
	FirstName        string                              `json:"first_name,omitempty"`
	LastName         string                              `json:"last_name,omitempty"`
	Email            string                              `json:"email,omitempty"`
	Company          string                              `json:"company,omitempty"`
	Phone            string                              `json:"phone,omitempty"`
	Line1            string                              `json:"line1,omitempty"`
	Line2            string                              `json:"line2,omitempty"`
	Line3            string                              `json:"line3,omitempty"`
	City             string                              `json:"city,omitempty"`
	StateCode        string                              `json:"state_code,omitempty"`
	State            string                              `json:"state,omitempty"`
	Zip              string                              `json:"zip,omitempty"`
	Country          string                              `json:"country,omitempty"`
	ValidationStatus QuoteBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type QuoteEditCreateSubCustomerQuoteForItemsCoupon struct {
	Id        string `json:"id,omitempty"`
	StartDate *int64 `json:"start_date,omitempty"`
	EndDate   *int64 `json:"end_date,omitempty"`
}
type QuoteUpdateSubscriptionQuoteForItemsRequest struct {
	Name                   string                                                  `json:"name,omitempty"`
	Notes                  string                                                  `json:"notes,omitempty"`
	ExpiresAt              *int64                                                  `json:"expires_at,omitempty"`
	Subscription           *QuoteUpdateSubscriptionQuoteForItemsSubscription       `json:"subscription,omitempty"`
	SubscriptionItems      []*QuoteUpdateSubscriptionQuoteForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove []string                                                `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList       *bool                                                   `json:"replace_items_list,omitempty"`
	Discounts              []*QuoteUpdateSubscriptionQuoteForItemsDiscount         `json:"discounts,omitempty"`
	ItemTiers              []*QuoteUpdateSubscriptionQuoteForItemsItemTier         `json:"item_tiers,omitempty"`
	BillingCycles          *int32                                                  `json:"billing_cycles,omitempty"`
	TermsToCharge          *int32                                                  `json:"terms_to_charge,omitempty"`
	ReactivateFrom         *int64                                                  `json:"reactivate_from,omitempty"`
	BillingAlignmentMode   QuoteBillingAlignmentMode                               `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                                `json:"coupon_ids,omitempty"`
	ReplaceCouponList      *bool                                                   `json:"replace_coupon_list,omitempty"`
	ChangeOption           QuoteChangeOption                                       `json:"change_option,omitempty"`
	ChangesScheduledAt     *int64                                                  `json:"changes_scheduled_at,omitempty"`
	ForceTermReset         *bool                                                   `json:"force_term_reset,omitempty"`
	Reactivate             *bool                                                   `json:"reactivate,omitempty"`
	BillingAddress         *QuoteUpdateSubscriptionQuoteForItemsBillingAddress     `json:"billing_address,omitempty"`
	ShippingAddress        *QuoteUpdateSubscriptionQuoteForItemsShippingAddress    `json:"shipping_address,omitempty"`
	Customer               *QuoteUpdateSubscriptionQuoteForItemsCustomer           `json:"customer,omitempty"`
	ContractTerm           *QuoteUpdateSubscriptionQuoteForItemsContractTerm       `json:"contract_term,omitempty"`
	NetTermDays            *int32                                                  `json:"net_term_days,omitempty"`
	Coupons                []*QuoteUpdateSubscriptionQuoteForItemsCoupon           `json:"coupons,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *QuoteUpdateSubscriptionQuoteForItemsRequest) payload() any { return r }

// input sub resource params single
type QuoteUpdateSubscriptionQuoteForItemsSubscription struct {
	Id                                string                                `json:"id"`
	SetupFee                          *int64                                `json:"setup_fee,omitempty"`
	StartDate                         *int64                                `json:"start_date,omitempty"`
	TrialEnd                          *int64                                `json:"trial_end,omitempty"`
	Coupon                            string                                `json:"coupon,omitempty"`
	AutoCollection                    QuoteSubscriptionAutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              QuoteSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params multi
type QuoteUpdateSubscriptionQuoteForItemsSubscriptionItem struct {
	ItemPriceId        string                              `json:"item_price_id"`
	Quantity           *int32                              `json:"quantity,omitempty"`
	QuantityInDecimal  string                              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64                              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string                              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32                              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64                              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32                              `json:"service_period_days,omitempty"`
	ChargeOnEvent      QuoteSubscriptionItemChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool                               `json:"charge_once,omitempty"`
	ChargeOnOption     QuoteSubscriptionItemChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           QuoteSubscriptionItemItemType       `json:"item_type,omitempty"`
	StartDate          *int64                              `json:"start_date,omitempty"`
	EndDate            *int64                              `json:"end_date,omitempty"`
	RampTierId         string                              `json:"ramp_tier_id,omitempty"`
}

// input sub resource params multi
type QuoteUpdateSubscriptionQuoteForItemsDiscount struct {
	ApplyOn       QuoteDiscountApplyOn       `json:"apply_on,omitempty"`
	DurationType  QuoteDiscountDurationType  `json:"duration_type"`
	Percentage    *float64                   `json:"percentage,omitempty"`
	Amount        *int64                     `json:"amount,omitempty"`
	Period        *int32                     `json:"period,omitempty"`
	PeriodUnit    QuoteDiscountPeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool                      `json:"included_in_mrr,omitempty"`
	ItemPriceId   string                     `json:"item_price_id,omitempty"`
	Quantity      *int32                     `json:"quantity,omitempty"`
	OperationType QuoteDiscountOperationType `json:"operation_type"`
	Id            string                     `json:"id,omitempty"`
	StartDate     *int64                     `json:"start_date,omitempty"`
	EndDate       *int64                     `json:"end_date,omitempty"`
}

// input sub resource params multi
type QuoteUpdateSubscriptionQuoteForItemsItemTier struct {
	ItemPriceId           string                   `json:"item_price_id,omitempty"`
	StartingUnit          *int32                   `json:"starting_unit,omitempty"`
	EndingUnit            *int32                   `json:"ending_unit,omitempty"`
	Price                 *int64                   `json:"price,omitempty"`
	StartingUnitInDecimal string                   `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                   `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                   `json:"price_in_decimal,omitempty"`
	PricingType           QuoteItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                   `json:"package_size,omitempty"`
	RampTierId            string                   `json:"ramp_tier_id,omitempty"`
}

// input sub resource params single
type QuoteUpdateSubscriptionQuoteForItemsBillingAddress struct {
	FirstName        string                              `json:"first_name,omitempty"`
	LastName         string                              `json:"last_name,omitempty"`
	Email            string                              `json:"email,omitempty"`
	Company          string                              `json:"company,omitempty"`
	Phone            string                              `json:"phone,omitempty"`
	Line1            string                              `json:"line1,omitempty"`
	Line2            string                              `json:"line2,omitempty"`
	Line3            string                              `json:"line3,omitempty"`
	City             string                              `json:"city,omitempty"`
	StateCode        string                              `json:"state_code,omitempty"`
	State            string                              `json:"state,omitempty"`
	Zip              string                              `json:"zip,omitempty"`
	Country          string                              `json:"country,omitempty"`
	ValidationStatus QuoteBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteUpdateSubscriptionQuoteForItemsShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteUpdateSubscriptionQuoteForItemsCustomer struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}

// input sub resource params single
type QuoteUpdateSubscriptionQuoteForItemsContractTerm struct {
	ActionAtTermEnd          QuoteContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params multi
type QuoteUpdateSubscriptionQuoteForItemsCoupon struct {
	Id        string `json:"id,omitempty"`
	StartDate *int64 `json:"start_date,omitempty"`
	EndDate   *int64 `json:"end_date,omitempty"`
}
type QuoteEditUpdateSubscriptionQuoteForItemsRequest struct {
	Notes                  string                                                      `json:"notes,omitempty"`
	ExpiresAt              *int64                                                      `json:"expires_at,omitempty"`
	SubscriptionItems      []*QuoteEditUpdateSubscriptionQuoteForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove []string                                                    `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList       *bool                                                       `json:"replace_items_list,omitempty"`
	Subscription           *QuoteEditUpdateSubscriptionQuoteForItemsSubscription       `json:"subscription,omitempty"`
	Discounts              []*QuoteEditUpdateSubscriptionQuoteForItemsDiscount         `json:"discounts,omitempty"`
	ItemTiers              []*QuoteEditUpdateSubscriptionQuoteForItemsItemTier         `json:"item_tiers,omitempty"`
	BillingCycles          *int32                                                      `json:"billing_cycles,omitempty"`
	TermsToCharge          *int32                                                      `json:"terms_to_charge,omitempty"`
	ReactivateFrom         *int64                                                      `json:"reactivate_from,omitempty"`
	BillingAlignmentMode   QuoteBillingAlignmentMode                                   `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                                    `json:"coupon_ids,omitempty"`
	ReplaceCouponList      *bool                                                       `json:"replace_coupon_list,omitempty"`
	ChangeOption           QuoteChangeOption                                           `json:"change_option,omitempty"`
	ChangesScheduledAt     *int64                                                      `json:"changes_scheduled_at,omitempty"`
	ForceTermReset         *bool                                                       `json:"force_term_reset,omitempty"`
	Reactivate             *bool                                                       `json:"reactivate,omitempty"`
	BillingAddress         *QuoteEditUpdateSubscriptionQuoteForItemsBillingAddress     `json:"billing_address,omitempty"`
	ShippingAddress        *QuoteEditUpdateSubscriptionQuoteForItemsShippingAddress    `json:"shipping_address,omitempty"`
	Customer               *QuoteEditUpdateSubscriptionQuoteForItemsCustomer           `json:"customer,omitempty"`
	ContractTerm           *QuoteEditUpdateSubscriptionQuoteForItemsContractTerm       `json:"contract_term,omitempty"`
	NetTermDays            *int32                                                      `json:"net_term_days,omitempty"`
	Coupons                []*QuoteEditUpdateSubscriptionQuoteForItemsCoupon           `json:"coupons,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *QuoteEditUpdateSubscriptionQuoteForItemsRequest) payload() any { return r }

// input sub resource params multi
type QuoteEditUpdateSubscriptionQuoteForItemsSubscriptionItem struct {
	ItemPriceId        string                              `json:"item_price_id"`
	Quantity           *int32                              `json:"quantity,omitempty"`
	QuantityInDecimal  string                              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64                              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string                              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32                              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64                              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32                              `json:"service_period_days,omitempty"`
	ChargeOnEvent      QuoteSubscriptionItemChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool                               `json:"charge_once,omitempty"`
	ChargeOnOption     QuoteSubscriptionItemChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           QuoteSubscriptionItemItemType       `json:"item_type,omitempty"`
	StartDate          *int64                              `json:"start_date,omitempty"`
	EndDate            *int64                              `json:"end_date,omitempty"`
	RampTierId         string                              `json:"ramp_tier_id,omitempty"`
}

// input sub resource params single
type QuoteEditUpdateSubscriptionQuoteForItemsSubscription struct {
	SetupFee                          *int64                                `json:"setup_fee,omitempty"`
	StartDate                         *int64                                `json:"start_date,omitempty"`
	TrialEnd                          *int64                                `json:"trial_end,omitempty"`
	Coupon                            string                                `json:"coupon,omitempty"`
	AutoCollection                    QuoteSubscriptionAutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              QuoteSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params multi
type QuoteEditUpdateSubscriptionQuoteForItemsDiscount struct {
	ApplyOn       QuoteDiscountApplyOn       `json:"apply_on,omitempty"`
	DurationType  QuoteDiscountDurationType  `json:"duration_type"`
	Percentage    *float64                   `json:"percentage,omitempty"`
	Amount        *int64                     `json:"amount,omitempty"`
	Period        *int32                     `json:"period,omitempty"`
	PeriodUnit    QuoteDiscountPeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool                      `json:"included_in_mrr,omitempty"`
	ItemPriceId   string                     `json:"item_price_id,omitempty"`
	Quantity      *int32                     `json:"quantity,omitempty"`
	OperationType QuoteDiscountOperationType `json:"operation_type"`
	Id            string                     `json:"id,omitempty"`
	StartDate     *int64                     `json:"start_date,omitempty"`
	EndDate       *int64                     `json:"end_date,omitempty"`
}

// input sub resource params multi
type QuoteEditUpdateSubscriptionQuoteForItemsItemTier struct {
	ItemPriceId           string                   `json:"item_price_id,omitempty"`
	StartingUnit          *int32                   `json:"starting_unit,omitempty"`
	EndingUnit            *int32                   `json:"ending_unit,omitempty"`
	Price                 *int64                   `json:"price,omitempty"`
	StartingUnitInDecimal string                   `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                   `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                   `json:"price_in_decimal,omitempty"`
	PricingType           QuoteItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                   `json:"package_size,omitempty"`
	RampTierId            string                   `json:"ramp_tier_id,omitempty"`
}

// input sub resource params single
type QuoteEditUpdateSubscriptionQuoteForItemsBillingAddress struct {
	FirstName        string                              `json:"first_name,omitempty"`
	LastName         string                              `json:"last_name,omitempty"`
	Email            string                              `json:"email,omitempty"`
	Company          string                              `json:"company,omitempty"`
	Phone            string                              `json:"phone,omitempty"`
	Line1            string                              `json:"line1,omitempty"`
	Line2            string                              `json:"line2,omitempty"`
	Line3            string                              `json:"line3,omitempty"`
	City             string                              `json:"city,omitempty"`
	StateCode        string                              `json:"state_code,omitempty"`
	State            string                              `json:"state,omitempty"`
	Zip              string                              `json:"zip,omitempty"`
	Country          string                              `json:"country,omitempty"`
	ValidationStatus QuoteBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteEditUpdateSubscriptionQuoteForItemsShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteEditUpdateSubscriptionQuoteForItemsCustomer struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}

// input sub resource params single
type QuoteEditUpdateSubscriptionQuoteForItemsContractTerm struct {
	ActionAtTermEnd          QuoteContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params multi
type QuoteEditUpdateSubscriptionQuoteForItemsCoupon struct {
	Id        string `json:"id,omitempty"`
	StartDate *int64 `json:"start_date,omitempty"`
	EndDate   *int64 `json:"end_date,omitempty"`
}
type QuoteCreateForChargeItemsAndChargesRequest struct {
	Name               string                                                  `json:"name,omitempty"`
	CustomerId         string                                                  `json:"customer_id"`
	PoNumber           string                                                  `json:"po_number,omitempty"`
	Notes              string                                                  `json:"notes,omitempty"`
	ExpiresAt          *int64                                                  `json:"expires_at,omitempty"`
	CurrencyCode       string                                                  `json:"currency_code,omitempty"`
	ItemPrices         []*QuoteCreateForChargeItemsAndChargesItemPrice         `json:"item_prices,omitempty"`
	ItemTiers          []*QuoteCreateForChargeItemsAndChargesItemTier          `json:"item_tiers,omitempty"`
	Charges            []*QuoteCreateForChargeItemsAndChargesCharge            `json:"charges,omitempty"`
	Coupon             string                                                  `json:"coupon,omitempty"`
	CouponIds          []string                                                `json:"coupon_ids,omitempty"`
	BillingAddress     *QuoteCreateForChargeItemsAndChargesBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress    *QuoteCreateForChargeItemsAndChargesShippingAddress     `json:"shipping_address,omitempty"`
	Discounts          []*QuoteCreateForChargeItemsAndChargesDiscount          `json:"discounts,omitempty"`
	TaxProvidersFields []*QuoteCreateForChargeItemsAndChargesTaxProvidersField `json:"tax_providers_fields,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *QuoteCreateForChargeItemsAndChargesRequest) payload() any { return r }

// input sub resource params multi
type QuoteCreateForChargeItemsAndChargesItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}

// input sub resource params multi
type QuoteCreateForChargeItemsAndChargesItemTier struct {
	ItemPriceId           string                   `json:"item_price_id,omitempty"`
	StartingUnit          *int32                   `json:"starting_unit,omitempty"`
	EndingUnit            *int32                   `json:"ending_unit,omitempty"`
	Price                 *int64                   `json:"price,omitempty"`
	StartingUnitInDecimal string                   `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                   `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                   `json:"price_in_decimal,omitempty"`
	PricingType           QuoteItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                   `json:"package_size,omitempty"`
}

// input sub resource params multi
type QuoteCreateForChargeItemsAndChargesCharge struct {
	Amount                 *int64                     `json:"amount,omitempty"`
	AmountInDecimal        string                     `json:"amount_in_decimal,omitempty"`
	Description            string                     `json:"description,omitempty"`
	AvalaraSaleType        QuoteChargeAvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                     `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                     `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32                     `json:"service_period,omitempty"`
}

// input sub resource params single
type QuoteCreateForChargeItemsAndChargesBillingAddress struct {
	FirstName        string                              `json:"first_name,omitempty"`
	LastName         string                              `json:"last_name,omitempty"`
	Email            string                              `json:"email,omitempty"`
	Company          string                              `json:"company,omitempty"`
	Phone            string                              `json:"phone,omitempty"`
	Line1            string                              `json:"line1,omitempty"`
	Line2            string                              `json:"line2,omitempty"`
	Line3            string                              `json:"line3,omitempty"`
	City             string                              `json:"city,omitempty"`
	StateCode        string                              `json:"state_code,omitempty"`
	State            string                              `json:"state,omitempty"`
	Zip              string                              `json:"zip,omitempty"`
	Country          string                              `json:"country,omitempty"`
	ValidationStatus QuoteBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteCreateForChargeItemsAndChargesShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type QuoteCreateForChargeItemsAndChargesDiscount struct {
	Percentage  *float64             `json:"percentage,omitempty"`
	Quantity    *int32               `json:"quantity,omitempty"`
	Amount      *int64               `json:"amount,omitempty"`
	ApplyOn     QuoteDiscountApplyOn `json:"apply_on"`
	ItemPriceId string               `json:"item_price_id,omitempty"`
}

// input sub resource params multi
type QuoteCreateForChargeItemsAndChargesTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type QuoteEditForChargeItemsAndChargesRequest struct {
	PoNumber           string                                                `json:"po_number,omitempty"`
	Notes              string                                                `json:"notes,omitempty"`
	ExpiresAt          *int64                                                `json:"expires_at,omitempty"`
	CurrencyCode       string                                                `json:"currency_code,omitempty"`
	ItemPrices         []*QuoteEditForChargeItemsAndChargesItemPrice         `json:"item_prices,omitempty"`
	ItemTiers          []*QuoteEditForChargeItemsAndChargesItemTier          `json:"item_tiers,omitempty"`
	Charges            []*QuoteEditForChargeItemsAndChargesCharge            `json:"charges,omitempty"`
	Coupon             string                                                `json:"coupon,omitempty"`
	CouponIds          []string                                              `json:"coupon_ids,omitempty"`
	BillingAddress     *QuoteEditForChargeItemsAndChargesBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress    *QuoteEditForChargeItemsAndChargesShippingAddress     `json:"shipping_address,omitempty"`
	Discounts          []*QuoteEditForChargeItemsAndChargesDiscount          `json:"discounts,omitempty"`
	TaxProvidersFields []*QuoteEditForChargeItemsAndChargesTaxProvidersField `json:"tax_providers_fields,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *QuoteEditForChargeItemsAndChargesRequest) payload() any { return r }

// input sub resource params multi
type QuoteEditForChargeItemsAndChargesItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}

// input sub resource params multi
type QuoteEditForChargeItemsAndChargesItemTier struct {
	ItemPriceId           string                   `json:"item_price_id,omitempty"`
	StartingUnit          *int32                   `json:"starting_unit,omitempty"`
	EndingUnit            *int32                   `json:"ending_unit,omitempty"`
	Price                 *int64                   `json:"price,omitempty"`
	StartingUnitInDecimal string                   `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                   `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                   `json:"price_in_decimal,omitempty"`
	PricingType           QuoteItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                   `json:"package_size,omitempty"`
}

// input sub resource params multi
type QuoteEditForChargeItemsAndChargesCharge struct {
	Amount                 *int64                     `json:"amount,omitempty"`
	AmountInDecimal        string                     `json:"amount_in_decimal,omitempty"`
	Description            string                     `json:"description,omitempty"`
	AvalaraSaleType        QuoteChargeAvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                     `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                     `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32                     `json:"service_period,omitempty"`
}

// input sub resource params single
type QuoteEditForChargeItemsAndChargesBillingAddress struct {
	FirstName        string                              `json:"first_name,omitempty"`
	LastName         string                              `json:"last_name,omitempty"`
	Email            string                              `json:"email,omitempty"`
	Company          string                              `json:"company,omitempty"`
	Phone            string                              `json:"phone,omitempty"`
	Line1            string                              `json:"line1,omitempty"`
	Line2            string                              `json:"line2,omitempty"`
	Line3            string                              `json:"line3,omitempty"`
	City             string                              `json:"city,omitempty"`
	StateCode        string                              `json:"state_code,omitempty"`
	State            string                              `json:"state,omitempty"`
	Zip              string                              `json:"zip,omitempty"`
	Country          string                              `json:"country,omitempty"`
	ValidationStatus QuoteBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type QuoteEditForChargeItemsAndChargesShippingAddress struct {
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Email            string                               `json:"email,omitempty"`
	Company          string                               `json:"company,omitempty"`
	Phone            string                               `json:"phone,omitempty"`
	Line1            string                               `json:"line1,omitempty"`
	Line2            string                               `json:"line2,omitempty"`
	Line3            string                               `json:"line3,omitempty"`
	City             string                               `json:"city,omitempty"`
	StateCode        string                               `json:"state_code,omitempty"`
	State            string                               `json:"state,omitempty"`
	Zip              string                               `json:"zip,omitempty"`
	Country          string                               `json:"country,omitempty"`
	ValidationStatus QuoteShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type QuoteEditForChargeItemsAndChargesDiscount struct {
	Percentage  *float64             `json:"percentage,omitempty"`
	Quantity    *int32               `json:"quantity,omitempty"`
	Amount      *int64               `json:"amount,omitempty"`
	ApplyOn     QuoteDiscountApplyOn `json:"apply_on"`
	ItemPriceId string               `json:"item_price_id,omitempty"`
}

// input sub resource params multi
type QuoteEditForChargeItemsAndChargesTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type QuoteListRequest struct {
	Limit          *int32           `json:"limit,omitempty"`
	Offset         string           `json:"offset,omitempty"`
	IncludeDeleted *bool            `json:"include_deleted,omitempty"`
	Id             *StringFilter    `json:"id,omitempty"`
	CustomerId     *StringFilter    `json:"customer_id,omitempty"`
	SubscriptionId *StringFilter    `json:"subscription_id,omitempty"`
	Status         *EnumFilter      `json:"status,omitempty"`
	Date           *TimestampFilter `json:"date,omitempty"`
	UpdatedAt      *TimestampFilter `json:"updated_at,omitempty"`
	SortBy         *SortFilter      `json:"sort_by,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *QuoteListRequest) payload() any { return r }

type QuoteQuoteLineGroupsForQuoteRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *QuoteQuoteLineGroupsForQuoteRequest) payload() any { return r }

type QuoteConvertRequest struct {
	Subscription          *QuoteConvertSubscription `json:"subscription,omitempty"`
	InvoiceDate           *int64                    `json:"invoice_date,omitempty"`
	InvoiceImmediately    *bool                     `json:"invoice_immediately,omitempty"`
	CreatePendingInvoices *bool                     `json:"create_pending_invoices,omitempty"`
	FirstInvoicePending   *bool                     `json:"first_invoice_pending,omitempty"`
	apiRequest            `json:"-" form:"-"`
}

func (r *QuoteConvertRequest) payload() any { return r }

// input sub resource params single
type QuoteConvertSubscription struct {
	Id                string                          `json:"id,omitempty"`
	AutoCollection    QuoteSubscriptionAutoCollection `json:"auto_collection,omitempty"`
	PoNumber          string                          `json:"po_number,omitempty"`
	AutoCloseInvoices *bool                           `json:"auto_close_invoices,omitempty"`
}
type QuoteUpdateStatusRequest struct {
	Status     QuoteStatus `json:"status"`
	Comment    string      `json:"comment,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *QuoteUpdateStatusRequest) payload() any { return r }

type QuoteExtendExpiryDateRequest struct {
	ValidTill  *int64 `json:"valid_till"`
	apiRequest `json:"-" form:"-"`
}

func (r *QuoteExtendExpiryDateRequest) payload() any { return r }

type QuoteDeleteRequest struct {
	Comment    string `json:"comment,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *QuoteDeleteRequest) payload() any { return r }

type QuotePdfRequest struct {
	ConsolidatedView *bool                `json:"consolidated_view,omitempty"`
	DispositionType  QuoteDispositionType `json:"disposition_type,omitempty"`
	apiRequest       `json:"-" form:"-"`
}

func (r *QuotePdfRequest) payload() any { return r }

// operation response
type QuoteRetrieveResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedCharge       QuotedCharge       `json:"quoted_charge,omitempty"`
	QuotedRamp         QuotedRamp         `json:"quoted_ramp,omitempty"`
	apiResponse
}

// operation response
type QuoteCreateSubForCustomerQuoteResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	apiResponse
}

// operation response
type QuoteEditCreateSubForCustomerQuoteResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	apiResponse
}

// operation response
type QuoteUpdateSubscriptionQuoteResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	apiResponse
}

// operation response
type QuoteEditUpdateSubscriptionQuoteResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	apiResponse
}

// operation response
type QuoteCreateForOnetimeChargesResponse struct {
	Quote        *Quote       `json:"quote,omitempty"`
	QuotedCharge QuotedCharge `json:"quoted_charge,omitempty"`
	apiResponse
}

// operation response
type QuoteEditOneTimeQuoteResponse struct {
	Quote        *Quote       `json:"quote,omitempty"`
	QuotedCharge QuotedCharge `json:"quoted_charge,omitempty"`
	apiResponse
}

// operation response
type QuoteCreateSubItemsForCustomerQuoteResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedRamp         QuotedRamp         `json:"quoted_ramp,omitempty"`
	apiResponse
}

// operation response
type QuoteEditCreateSubCustomerQuoteForItemsResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedRamp         QuotedRamp         `json:"quoted_ramp,omitempty"`
	apiResponse
}

// operation response
type QuoteUpdateSubscriptionQuoteForItemsResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedRamp         QuotedRamp         `json:"quoted_ramp,omitempty"`
	apiResponse
}

// operation response
type QuoteEditUpdateSubscriptionQuoteForItemsResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedRamp         QuotedRamp         `json:"quoted_ramp,omitempty"`
	apiResponse
}

// operation response
type QuoteCreateForChargeItemsAndChargesResponse struct {
	Quote        *Quote       `json:"quote,omitempty"`
	QuotedCharge QuotedCharge `json:"quoted_charge,omitempty"`
	apiResponse
}

// operation response
type QuoteEditForChargeItemsAndChargesResponse struct {
	Quote        *Quote       `json:"quote,omitempty"`
	QuotedCharge QuotedCharge `json:"quoted_charge,omitempty"`
	apiResponse
}

// operation sub response
type QuoteListQuoteResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedRamp         QuotedRamp         `json:"quoted_ramp,omitempty"`
}

// operation response
type QuoteListResponse struct {
	List       []*QuoteListQuoteResponse `json:"list,omitempty"`
	NextOffset string                    `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type QuoteQuoteLineGroupsForQuoteQuoteResponse struct {
	QuoteLineGroup QuoteLineGroup `json:"quote_line_group,omitempty"`
}

// operation response
type QuoteQuoteLineGroupsForQuoteResponse struct {
	List       []*QuoteQuoteLineGroupsForQuoteQuoteResponse `json:"list,omitempty"`
	NextOffset string                                       `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type QuoteConvertResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	Customer           Customer           `json:"customer,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedCharge       QuotedCharge       `json:"quoted_charge,omitempty"`
	QuotedRamp         QuotedRamp         `json:"quoted_ramp,omitempty"`
	Subscription       Subscription       `json:"subscription,omitempty"`
	Invoice            Invoice            `json:"invoice,omitempty"`
	CreditNote         CreditNote         `json:"credit_note,omitempty"`
	UnbilledCharges    []UnbilledCharge   `json:"unbilled_charges,omitempty"`
	apiResponse
}

// operation response
type QuoteUpdateStatusResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedCharge       QuotedCharge       `json:"quoted_charge,omitempty"`
	QuotedRamp         QuotedRamp         `json:"quoted_ramp,omitempty"`
	apiResponse
}

// operation response
type QuoteExtendExpiryDateResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedCharge       QuotedCharge       `json:"quoted_charge,omitempty"`
	QuotedRamp         QuotedRamp         `json:"quoted_ramp,omitempty"`
	apiResponse
}

// operation response
type QuoteDeleteResponse struct {
	Quote              *Quote             `json:"quote,omitempty"`
	QuotedSubscription QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedCharge       QuotedCharge       `json:"quoted_charge,omitempty"`
	QuotedRamp         QuotedRamp         `json:"quoted_ramp,omitempty"`
	apiResponse
}

// operation response
type QuotePdfResponse struct {
	Download Download `json:"download,omitempty"`
	apiResponse
}
